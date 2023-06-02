package cmd

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/persistence/database"
	"github.com/aghex70/deselflopment-babl/server"
	"github.com/spf13/cobra"
	calendarService "github.com/aghex70/deselflopment-babl/internal/core/services/calendar"
	calendarHandler "github.com/aghex70/deselflopment-babl/internal/handlers/calendar"
	calendarRepository "github.com/aghex70/deselflopment-babl/internal/stores/calendar"
	eventService "github.com/aghex70/deselflopment-babl/internal/core/services/event"
	eventHandler "github.com/aghex70/deselflopment-babl/internal/handlers/event"
	eventRepository "github.com/aghex70/deselflopment-babl/internal/stores/event"
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
			eventR, _ := eventRepository.NewGormRepository(gdb)
			userR, _ := userRepository.NewGormRepository(gdb)

			calendarS, _ := calendarService.NewService(calendarR)
			calendarH := calendarHandler.NewHandler(calendarS)

			eventS, _ := eventService.NewService(eventR)
			eventH := eventHandler.NewHandler(eventS)

			userS, _ := userService.NewService(userR)
			userH := userHandler.NewHandler(userS)

			s := server.NewRestServer(cfg.Server.Rest, calendarH, eventH, userH)
			err = s.StartServer()
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
