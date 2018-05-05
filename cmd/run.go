package cmd

import (
	"expvar"
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	"github.com/kamilsk/passport/dao"
	"github.com/kamilsk/passport/server"
	"github.com/kamilsk/passport/server/router/chi"
	"github.com/kamilsk/passport/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime.GOMAXPROCS(asInt(cmd.Flag("cpus").Value))
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()

		if asBool(cmd.Flag("with-profiler").Value) {
			go startProfiler()
		}
		if asBool(cmd.Flag("with-monitoring").Value) {
			go startMonitoring()
		}

		handler := chi.NewRouter(
			server.New(
				cmd.Flag("base-url").Value.String(),
				service.New(
					dao.Must(dao.Connection(dsn(cmd))),
				),
			),
		)
		srv := &http.Server{Addr: addr, Handler: handler,
			ReadTimeout:       asDuration(cmd.Flag("read-timeout").Value),
			ReadHeaderTimeout: asDuration(cmd.Flag("read-header-timeout").Value),
			WriteTimeout:      asDuration(cmd.Flag("write-timeout").Value),
			IdleTimeout:       asDuration(cmd.Flag("idle-timeout").Value)}
		log.Println("starting server at", addr)
		return srv.ListenAndServe()
	},
}

func init() {
	v := viper.New()
	must(
		func() error { return v.BindEnv("max_cpus") },
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("port") },
		func() error { return v.BindEnv("read_timeout") },
		func() error { return v.BindEnv("read_header_timeout") },
		func() error { return v.BindEnv("write_timeout") },
		func() error { return v.BindEnv("idle_timeout") },
		func() error { return v.BindEnv("base_url") },
	)
	{
		v.SetDefault("max_cpus", 1)
		v.SetDefault("bind", "127.0.0.1")
		v.SetDefault("port", 80)
		v.SetDefault("read_timeout", time.Duration(0))
		v.SetDefault("read_header_timeout", time.Duration(0))
		v.SetDefault("write_timeout", time.Duration(0))
		v.SetDefault("idle_timeout", time.Duration(0))
		v.SetDefault("base_url", "http://localhost/")
	}
	{
		runCmd.Flags().Int("cpus", v.GetInt("max_cpus"),
			"maximum number of CPUs that can be executing simultaneously")
		runCmd.Flags().String("bind", v.GetString("bind"),
			"interface to which the server will bind")
		runCmd.Flags().Int("port", v.GetInt("port"),
			"port on which the server will listen")
		runCmd.Flags().Duration("read-timeout", v.GetDuration("read_timeout"),
			"maximum duration for reading the entire request, including the body")
		runCmd.Flags().Duration("read-header-timeout", v.GetDuration("read_header_timeout"),
			"amount of time allowed to read request headers")
		runCmd.Flags().Duration("write-timeout", v.GetDuration("write_timeout"),
			"maximum duration before timing out writes of the response")
		runCmd.Flags().Duration("idle-timeout", v.GetDuration("idle_timeout"),
			"maximum amount of time to wait for the next request when keep-alive is enabled")
		runCmd.Flags().Bool("with-profiler", false,
			"enable pprof on /pprof")
		runCmd.Flags().Bool("with-monitoring", false,
			"enable expvar on /vars")
		runCmd.Flags().String("base-url", v.GetString("base_url"),
			"hostname (and path) to the root")
	}
	db(runCmd)
}

func startProfiler() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/pprof/profile", pprof.Profile)
	mux.HandleFunc("/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/", pprof.Index) // net/http/pprof.handler.ServeHTTP specificity
	_ = http.ListenAndServe(":8090", mux)
}

func startMonitoring() {
	mux := &http.ServeMux{}
	expvar.Handler()
	mux.Handle("/vars", expvar.Handler())
	_ = http.ListenAndServe(":8091", mux)
}
