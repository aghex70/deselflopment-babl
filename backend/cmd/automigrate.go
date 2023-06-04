package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/spf13/cobra"
)

func AutoMigrateCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "automigrate",
		Short: "Auto migrate",
		Run: func(cmd *cobra.Command, args []string) {
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				panic(err)
			}
			err = database.AutoMigrateWithTimestamp(gdb)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd

}
