package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	entryService "github.com/aghex70/deselflopment-babl/internal/core/services/entry"
	userService "github.com/aghex70/deselflopment-babl/internal/core/services/user"
	entryHandler "github.com/aghex70/deselflopment-babl/internal/handlers/entry"
	userHandler "github.com/aghex70/deselflopment-babl/internal/handlers/user"
	entryRepository "github.com/aghex70/deselflopment-babl/internal/stores/entry"
	userRepository "github.com/aghex70/deselflopment-babl/internal/stores/user"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/aghex70/deselflopment-babl/server"
	"github.com/spf13/cobra"
)

func ServeCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				panic(err)
			}
			entryR, _ := entryRepository.NewGormRepository(gdb)
			userR, _ := userRepository.NewGormRepository(gdb)

			entryS, _ := entryService.NewService(entryR)
			entryH := entryHandler.NewHandler(entryS)

			userS, _ := userService.NewService(userR)
			userH := userHandler.NewHandler(userS)

			s := server.NewRestServer(cfg.Server.Rest, entryH, userH)
			err = s.StartServer()
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
