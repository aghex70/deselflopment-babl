package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/spf13/cobra"
)

func RootCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deselflopment-babl",
		Short: "Root command",
	}

	// Intialize database
	db, err := database.NewSqlDB(*cfg.Database)
	if err != nil {
		panic(err)
	}

	cmd.AddCommand(ServeCommand(cfg))
	cmd.AddCommand(MakeMigrationsCommand(db))
	cmd.AddCommand(MigrateCommand(db))
	return cmd
}

