package chi

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(api interface{}, withProfiler bool) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1/tracker", func(r chi.Router) {
		r.Get("/pixel", notImplemented)
		r.Post("/fingerprint", notImplemented)
		r.Post("/heartbeat", notImplemented)
	})

	if withProfiler {
		r.Route("/debug/pprof", func(r chi.Router) {
			r.Get("/", pprof.Index)
			r.Get("/cmdline", pprof.Cmdline)
			r.Get("/profile", pprof.Profile)
			r.Get("/symbol", pprof.Symbol)
			r.Get("/trace", pprof.Trace)
		})
	}

	return r
}
