package user

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/aghex70/deselflopment-babl/docs"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"net/http"
)

type Handler struct {
	userService ports.UserServicer
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body domain.User true "user data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.User, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /user [post]
func (h Handler) CreateUser(c *gin.Context) {
	var r ports.CreateUserRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nu, err := h.userService.Create(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": nu})
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @ID update-user
// @Accept json
// @Produce json
// @Param  id path string true "user ID" user body domain.User true "user data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.User, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /user/{id} [put]
func (h Handler) UpdateUser(c *gin.Context) {
	var r ports.UpdateUserRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uu, err := h.userService.Update(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": uu})
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @ID get-user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.User, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /user/{id} [get]
func (h Handler) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")
	uu, err := h.userService.Get(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": uu})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @ID delete-user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 204 {object} handlers.JSONOKResponse{data=domain.User, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /user/{id} [delete]
func (h Handler) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.userService.Delete(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary Retrieve all users
// @Description Retrieve all users
// @ID list-user
// @Accept json
// @Produce json
// @Success 200 {object} handlers.JSONOKResponse{data=domain.User, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /users [get]
func (h Handler) List(c *gin.Context) {
	us, err := h.userService.List(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": us})
}

func NewHandler(us ports.UserServicer) Handler {
	return Handler{
		userService: us,
	}
}
