package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func main() {
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	model := newHomePageModel()

	view := NewView()
	view.Parse("index", "view/index.html")

	controller := mux.NewRouter()
	controller = attachLogging(controller, log)

	// Choose the folder to serve
	staticDir := "/assets/"

	// Create the route
	controller.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	controller.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view.Exec("index", w, model)
	})

	if err := http.ListenAndServe(":8000", controller); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func attachLogging(controller *mux.Router, log zerolog.Logger) *mux.Router {
	controller.Use(hlog.NewHandler(log))
	controller.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.Path).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	controller.Use(hlog.RemoteAddrHandler("ip"))
	controller.Use(hlog.RefererHandler("referer"))
	controller.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	return controller
}
