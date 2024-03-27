package cmd

import (
	"Narcolepsick1d/mini-twitter/internal/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	setupLogger()

	rootCmd := &cobra.Command{
		Use:   "run-server",
		Short: "Run server",
		Long:  "Run server is a default command",
		Run:   func(cmd *cobra.Command, args []string) { runServer(cfg) },
	}

	rootCmd.AddCommand(newMigrateCmd(cfg))

	return rootCmd
}

func Execute() {
	rootCmd := createRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
