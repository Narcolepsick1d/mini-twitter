package cmd

import (
	"Narcolepsick1d/mini-twitter/internal/app"
	"Narcolepsick1d/mini-twitter/internal/config"
	"Narcolepsick1d/mini-twitter/internal/database"
	"Narcolepsick1d/mini-twitter/internal/rest"
	"Narcolepsick1d/mini-twitter/internal/scope"
	"Narcolepsick1d/mini-twitter/pkg/hash"
	"fmt"
	"log/slog"
)

func runServer(cfg *config.Config) {
	logger := slog.Default()

	db, err := database.NewGoquDB(cfg)
	if err != nil {
		logger.Error("unable to connect database", "error", err)
		return
	}
	hash := hash.NewSHA1Hasher(cfg.Hash.Salt)
	restServer := rest.NewHandler(rest.HandlerConfig{
		Dependencies: &scope.Dependencies{
			DB:     db,
			Hash:   hash,
			Secret: "sample secret",
		},
	})

	server, err := app.NewApp(cfg.Server.Port, restServer)
	if err != nil {
		logger.Error("unable to start server", "error", err)
		return
	}

	logger.Info("Starting server...", "address", fmt.Sprintf("http://localhost:%d", cfg.Server.Port))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
