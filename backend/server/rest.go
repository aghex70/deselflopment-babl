package server

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"github.com/aghex70/deselflopment-babl/internal/handlers/calendar"
	"github.com/aghex70/deselflopment-babl/internal/handlers/entry"
	"github.com/aghex70/deselflopment-babl/internal/handlers/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	cfg             config.RestConfig
	calendarHandler calendar.Handler
	calendarService ports.CalendarServicer
	entryHandler    entry.Handler
	entryService    ports.EntryServicer
	userHandler     user.Handler
	userService     ports.UserServicer
}

func (s *RestServer) StartServer() error {
	router := gin.Default()

	// calendars
	router.POST("/calendars", s.calendarHandler.CreateCalendar)
	router.GET("/calendars/:uuid", s.calendarHandler.GetCalendar)
	router.PUT("/calendars/:uuid", s.calendarHandler.UpdateCalendar)
	router.DELETE("/calendars/:uuid", s.calendarHandler.DeleteCalendar)
	router.GET("/calendars", s.calendarHandler.List)

	// entries
	router.POST("/entries", s.entryHandler.CreateEntry)
	router.GET("/entries/:uuid", s.entryHandler.GetEntry)
	router.PUT("/entries/:uuid", s.entryHandler.UpdateEntry)
	router.DELETE("/entries/:uuid", s.entryHandler.DeleteEntry)
	router.GET("/entries", s.entryHandler.List)

	// users
	router.POST("/users", s.userHandler.CreateUser)
	router.GET("/users/:uuid", s.userHandler.GetUser)
	router.PUT("/users/:uuid", s.userHandler.UpdateUser)
	router.DELETE("/users/:uuid", s.userHandler.DeleteUser)
	router.GET("/users", s.userHandler.List)

	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}

func NewRestServer(cfg *config.RestConfig, calendarh calendar.Handler, entryh entry.Handler, userh user.Handler) *RestServer {
	return &RestServer{
		cfg:             *cfg,
		calendarHandler: calendarh,
		entryHandler:    entryh,
		userHandler:     userh,
	}
}
