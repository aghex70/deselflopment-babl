package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/spf13/cobra"
	"log"
)

func RootCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deselflopment-babl",
		Short: "Root command",
	}

	// Intialize database
	log.Println("Starting application database")
	db, err := database.NewSqlDB(*cfg.Database)
	if err != nil {
		log.Fatalf("error starting application database %+v", err.Error())
	}

	cmd.AddCommand(ServeCommand(cfg))
	cmd.AddCommand(AutoMigrateCommand(cfg))
	cmd.AddCommand(MigrateCommand(db))
	return cmd
}
