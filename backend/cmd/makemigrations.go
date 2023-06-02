package cmd

import (
	"database/sql"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/spf13/cobra"
)

func MakeMigrationsCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "makemigrations [filename]",
		Short: "Generate database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				panic("fooooooooooo")
			}
			filename := args[0]
			err := database.MakeMigrations(db, filename)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd

}

