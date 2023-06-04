package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/aghex70/deselflopment-babl/server"
	"github.com/spf13/cobra"
	calendarService "github.com/aghex70/deselflopment-babl/internal/core/services/calendar"
	calendarHandler "github.com/aghex70/deselflopment-babl/internal/handlers/calendar"
	calendarRepository "github.com/aghex70/deselflopment-babl/internal/stores/calendar"
	entryService "github.com/aghex70/deselflopment-babl/internal/core/services/entry"
	entryHandler "github.com/aghex70/deselflopment-babl/internal/handlers/entry"
	entryRepository "github.com/aghex70/deselflopment-babl/internal/stores/entry"
	userService "github.com/aghex70/deselflopment-babl/internal/core/services/user"
	userHandler "github.com/aghex70/deselflopment-babl/internal/handlers/user"
	userRepository "github.com/aghex70/deselflopment-babl/internal/stores/user"
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
			calendarR, _ := calendarRepository.NewGormRepository(gdb)
			entryR, _ := entryRepository.NewGormRepository(gdb)
			userR, _ := userRepository.NewGormRepository(gdb)

			calendarS, _ := calendarService.NewService(calendarR)
			calendarH := calendarHandler.NewHandler(calendarS)

			entryS, _ := entryService.NewService(entryR)
			entryH := entryHandler.NewHandler(entryS)

			userS, _ := userService.NewService(userR)
			userH := userHandler.NewHandler(userS)

			s := server.NewRestServer(cfg.Server.Rest, calendarH, entryH, userH)
			err = s.StartServer()
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
