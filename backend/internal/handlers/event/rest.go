package event

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/aghex70/deselflopment-babl/docs"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"net/http"
)

type Handler struct {
	eventService ports.EventServicer
}

// CreateEvent godoc
// @Summary Create a event
// @Description Create a event
// @ID create-event
// @Accept json
// @Produce json
// @Param event body domain.Event true "event data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Event, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /event [post]
func (h Handler) CreateEvent(c *gin.Context) {
	var r ports.CreateEventRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ne, err := h.eventService.Create(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"event": ne})
}

// UpdateEvent godoc
// @Summary Update a event
// @Description Update a event
// @ID update-event
// @Accept json
// @Produce json
// @Param  id path string true "event ID" event body domain.Event true "event data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Event, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /event/{id} [put]
func (h Handler) UpdateEvent(c *gin.Context) {
	var r ports.UpdateEventRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ue, err := h.eventService.Update(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": ue})
}

// GetEvent godoc
// @Summary Get a event
// @Description Get a event by ID
// @ID get-event
// @Accept json
// @Produce json
// @Param id path string true "event ID"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Event, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /event/{id} [get]
func (h Handler) GetEvent(c *gin.Context) {
	uuid := c.Param("uuid")
	ee, err := h.eventService.Get(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": ee})
}

// DeleteEvent godoc
// @Summary Delete a event
// @Description Delete a event by ID
// @ID delete-event
// @Accept json
// @Produce json
// @Param id path string true "event ID"
// @Success 204 {object} handlers.JSONOKResponse{data=domain.Event, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /event/{id} [delete]
func (h Handler) DeleteEvent(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.eventService.Delete(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary Retrieve all events
// @Description Retrieve all events
// @ID list-event
// @Accept json
// @Produce json
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Event, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /events [get]
func (h Handler) List(c *gin.Context) {
	es, err := h.eventService.List(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": es})
}

func NewHandler(es ports.EventServicer) Handler {
	return Handler{
		eventService: es,
	}
}
