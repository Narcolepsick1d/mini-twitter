package cmd

import (
	"Narcolepsick1d/mini-twitter/internal/config"
	"Narcolepsick1d/mini-twitter/internal/database"
	"Narcolepsick1d/mini-twitter/internal/migrations"
	"database/sql"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func newMigrateCmd(cfg *config.Config) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "run migrations",
		Long:  "Run migrations. Specify action to run up/down migrations.",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println("Run any of sub commands") },
	}

	db := func(cmd *cobra.Command) *sql.DB {
		db, err := database.NewSQLDB(cfg)
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(-1)
		}
		return db
	}

	migrateCmd.AddCommand(
		&cobra.Command{
			Use:   "up",
			Short: "run migrations up",
			Long:  "Run migrations UP to end.",
			Run: func(cmd *cobra.Command, args []string) {
				err := migrations.Up(db(cmd))
				if err != nil {
					cmd.PrintErr(err)
				}
			},
		},
		&cobra.Command{
			Use:   "down",
			Short: "run migrations down",
			Long:  "Run migrations DOWN to the start.",
			Run: func(cmd *cobra.Command, args []string) {
				err := migrations.Down(db(cmd))
				if err != nil {
					cmd.PrintErr(err)
				}
			},
		},
		&cobra.Command{
			Use:   "steps N",
			Short: "run migrations for given steps",
			Long:  "Run migrations {N} times. Steps can be positive (for UP migrations) or negative (for DOWN migrations)",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				steps, err := strconv.Atoi(args[0])
				if err != nil {
					cmd.PrintErr(err)
					return
				}

				err = migrations.Steps(db(cmd), steps)
				if err != nil {
					cmd.PrintErr(err)
				}
			},
		},
	)

	return migrateCmd
}
