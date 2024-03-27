package cmd

import (
	"log/slog"
	"os"
)

func setupLogger() {
	switch os.Getenv("ENV") {
	case "local":

	default:
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	}
}
