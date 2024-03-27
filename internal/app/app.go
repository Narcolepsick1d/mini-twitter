package app

import (
	"Narcolepsick1d/mini-twitter/internal/middleware"
	"fmt"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// App describes http handlers.
type App struct {
	*http.Server
}

// NewApp takes config params from environments and returns server instance ready to start listen.
// @title           Mini-twitter Swagger API
// @version         1.0
// @description     This is a Mini twitter service API.
// @BasePath  /api/v1
// ..
func NewApp(port int, restServer http.Handler) (*App, error) {
	sm := http.NewServeMux()

	sm.HandleFunc("/health-check/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sm.Handle(
		"/api/v1/",
		http.StripPrefix("/api/v1", middleware.CorsMiddleware(restServer)),
	)
	sm.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/api/v1/swagger/doc.json")))

	httpServer := http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		Handler:           sm,
		ReadHeaderTimeout: time.Second * 5,
	}

	return &App{
		Server: &httpServer,
	}, nil
}
