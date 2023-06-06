package server

import (
	"github.com/aghex70/deselflopment-babl/config"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"github.com/aghex70/deselflopment-babl/internal/handlers/entry"
	"github.com/aghex70/deselflopment-babl/internal/handlers/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
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
	cfg          config.RestConfig
	entryHandler entry.Handler
	entryService ports.EntryServicer
	userHandler  user.Handler
	userService  ports.UserServicer
}

func (s *RestServer) StartServer() error {
	router := gin.Default()

	// entries
	router.POST("/entries", s.entryHandler.CreateEntry)
	router.GET("/entries/:uuid", s.entryHandler.GetEntry)
	router.PUT("/entries/:uuid", s.entryHandler.UpdateEntry)
	router.DELETE("/entries/:uuid", s.entryHandler.DeleteEntry)
	router.GET("/entries", s.entryHandler.List)

	// users
	router.GET("/users/:uuid", s.userHandler.GetUser)
	router.DELETE("/users/:uuid", s.userHandler.DeleteUser)
	router.GET("/users", s.userHandler.List)

	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":" + strconv.Itoa(s.cfg.Port))
	if err != nil {
		return err
	}
	return nil
}

func NewRestServer(cfg *config.RestConfig, entryh entry.Handler, userh user.Handler) *RestServer {
	return &RestServer{
		cfg:          *cfg,
		entryHandler: entryh,
		userHandler:  userh,
	}
}
