package server

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"github.com/aghex70/deselflopment-babl/internal/handlers/calendar"
	"github.com/aghex70/deselflopment-babl/internal/handlers/event"
	"github.com/aghex70/deselflopment-babl/internal/handlers/user"
	_ "github.com/aghex70/deselflopment-babl/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

// @title  API
// @version 1.0
// @description deselflopment-babl API server sample

// @contact.name API Support
// @contact.url https://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:12001
// @BasePath /api/v1
// @query.collection.format multi

type RestServer struct {
	cfg               config.RestConfig
	calendarHandler calendar.CalendarHandler
	calendarService ports.CalendarServicer
	eventHandler event.EventHandler
	eventService ports.EventServicer
	userHandler user.UserHandler
	userService ports.UserServicer
}

func (s *RestServer) StartServer() error {
	router := gin.Default()

	// calendars
	router.POST("/calendars", s.Handler.CreateCalendar)
	router.GET("/calendars/:uuid", s.Handler.GetCalendar)
	router.PUT("/calendars/:uuid", s.Handler.UpdateCalendar)
	router.DELETE("/calendars/:uuid", s.Handler.DeleteCalendar)
	router.GET("/calendars", s.Handler.List)

	// events
	router.POST("/events", s.Handler.CreateEvent)
	router.GET("/events/:uuid", s.Handler.GetEvent)
	router.PUT("/events/:uuid", s.Handler.UpdateEvent)
	router.DELETE("/events/:uuid", s.Handler.DeleteEvent)
	router.GET("/events", s.Handler.List)

	// users
	router.POST("/users", s.Handler.CreateUser)
	router.GET("/users/:uuid", s.Handler.GetUser)
	router.PUT("/users/:uuid", s.Handler.UpdateUser)
	router.DELETE("/users/:uuid", s.Handler.DeleteUser)
	router.GET("/users", s.Handler.List)

	// swagger documentation
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}

func NewRestServer(cfg *config.RestConfig , calendarh calendar.Handler, eventh event.Handler, userh user.Handler) *RestServer {
	return &RestServer{
		cfg:               *cfg,
		calendarHandler: calendarh,
		eventHandler: eventh,
		userHandler: userh,
		}
}