package cmd

import (
	"database/sql"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/spf13/cobra"
)

func MigrateCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			err := database.Migrate(db)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}

