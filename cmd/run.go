package cmd

import (
	"expvar"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"runtime"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.octolab.org/fn"
	"go.octolab.org/strings"

	"github.com/kamilsk/passport/pkg/config"
	"github.com/kamilsk/passport/pkg/server"
	"github.com/kamilsk/passport/pkg/server/router/chi"
	"github.com/kamilsk/passport/pkg/service"
	"github.com/kamilsk/passport/pkg/storage"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime.GOMAXPROCS(cnf.Union.CPUs)

		var (
			repo    = storage.Must(storage.Database(cnf.Union.DatabaseConfig))
			handler = chi.NewRouter(
				server.New(
					cnf.Union.ServerConfig,
					service.New(cnf.Union.ServiceConfig, repo),
				),
			)
		)

		if cnf.Union.MonitoringConfig.Enabled {
			if err := startMonitoring(cnf.Union.MonitoringConfig); err != nil {
				return err
			}
		}
		if cnf.Union.ProfilingConfig.Enabled {
			if err := startProfiler(cnf.Union.ProfilingConfig); err != nil {
				return err
			}
		}
		return startHTTPServer(cnf.Union.ServerConfig, handler)
	},
}

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("max_cpus") },
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("http_port") },
		func() error { return v.BindEnv("profiling_port") },
		func() error { return v.BindEnv("monitoring_port") },
		func() error { return v.BindEnv("grpc_port") },
		func() error { return v.BindEnv("read_timeout") },
		func() error { return v.BindEnv("read_header_timeout") },
		func() error { return v.BindEnv("write_timeout") },
		func() error { return v.BindEnv("idle_timeout") },
		func() error { return v.BindEnv("base_url") },
		func() error {
			v.SetDefault("max_cpus", defaults["max_cpus"])
			v.SetDefault("bind", defaults["bind"])
			v.SetDefault("http_port", defaults["http_port"])
			v.SetDefault("profiling_port", defaults["profiling_port"])
			v.SetDefault("monitoring_port", defaults["monitoring_port"])
			v.SetDefault("grpc_port", defaults["grpc_port"])

			bind := v.GetString("bind")
			v.SetDefault("host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("http_port"))))
			v.SetDefault("profiling_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("profiling_port"))))
			v.SetDefault("monitoring_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("monitoring_port"))))
			v.SetDefault("grpc_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("grpc_port"))))

			v.SetDefault("read_timeout", defaults["read_timeout"])
			v.SetDefault("read_header_timeout", defaults["read_header_timeout"])
			v.SetDefault("write_timeout", defaults["write_timeout"])
			v.SetDefault("idle_timeout", defaults["idle_timeout"])
			v.SetDefault("base_url", defaults["base_url"])
			return nil
		},
		func() error {
			flags := runCmd.Flags()
			flags.IntVarP(&cnf.Union.CPUs,
				"cpus", "C", v.GetInt("max_cpus"), "maximum number of CPUs that can be executing simultaneously")
			flags.StringVarP(&cnf.Union.ServerConfig.Interface,
				"host", "H", v.GetString("host"), "web server host")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadTimeout,
				"read-timeout", "", v.GetDuration("read_timeout"),
				"maximum duration for reading the entire request, including the body")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadHeaderTimeout,
				"read-header-timeout", "", v.GetDuration("read_header_timeout"),
				"amount of time allowed to read request headers")
			flags.DurationVarP(&cnf.Union.ServerConfig.WriteTimeout,
				"write-timeout", "", v.GetDuration("write_timeout"),
				"maximum duration before timing out writes of the response")
			flags.DurationVarP(&cnf.Union.ServerConfig.IdleTimeout,
				"idle-timeout", "", v.GetDuration("idle_timeout"),
				"maximum amount of time to wait for the next request when keep-alive is enabled")
			flags.BoolVarP(&cnf.Union.ProfilingConfig.Enabled,
				"with-profiling", "", false, "enable pprof on /pprof/* and /debug/pprof/")
			flags.StringVarP(&cnf.Union.ProfilingConfig.Interface,
				"profiling-host", "", v.GetString("profiling_host"), "profiling host")
			flags.BoolVarP(&cnf.Union.MonitoringConfig.Enabled,
				"with-monitoring", "", false, "enable prometheus on /monitoring and expvar on /vars")
			flags.StringVarP(&cnf.Union.MonitoringConfig.Interface,
				"monitoring-host", "", v.GetString("monitoring_host"), "monitoring host")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.StringVarP(&cnf.Union.ServerConfig.BaseURL, "base-url", "", v.GetString("base_url"),
				"hostname (and path) to the root")
			return nil
		},
	)
	db(runCmd)
}

func startHTTPServer(cnf config.ServerConfig, handler http.Handler) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	srv := &http.Server{Addr: cnf.Interface, Handler: handler,
		ReadTimeout:       cnf.ReadTimeout,
		ReadHeaderTimeout: cnf.ReadHeaderTimeout,
		WriteTimeout:      cnf.WriteTimeout,
		IdleTimeout:       cnf.IdleTimeout,
	}
	log.Println("start web server at", listener.Addr())
	return srv.Serve(listener)
}

func startMonitoring(cnf config.MonitoringConfig) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	go func() {
		mux := &http.ServeMux{}
		mux.Handle("/monitoring", promhttp.Handler())
		mux.Handle("/vars", expvar.Handler())
		log.Println("start monitoring server at", listener.Addr())
		_ = http.Serve(listener, mux) // TODO issue#97
		_ = listener.Close()
	}()
	return nil
}

func startProfiler(cnf config.ProfilingConfig) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/pprof/profile", pprof.Profile)
		mux.HandleFunc("/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/pprof/trace", pprof.Trace)
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		log.Println("start profiling server at", listener.Addr())
		_ = http.Serve(listener, mux) // TODO issue#97
		_ = listener.Close()
	}()
	return nil
}
