package calendar

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/aghex70/deselflopment-babl/docs"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"net/http"
)

type Handler struct {
	calendarService ports.CalendarServicer
}

// CreateCalendar godoc
// @Summary Create a calendar
// @Description Create a calendar
// @ID create-calendar
// @Accept json
// @Produce json
// @Param calendar body domain.Calendar true "calendar data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Calendar, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /calendar [post]
func (h Handler) CreateCalendar(c *gin.Context) {
	var r ports.CreateCalendarRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nc, err := h.calendarService.Create(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"calendar": nc})
}

// UpdateCalendar godoc
// @Summary Update a calendar
// @Description Update a calendar
// @ID update-calendar
// @Accept json
// @Produce json
// @Param  id path string true "calendar ID" calendar body domain.Calendar true "calendar data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Calendar, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /calendar/{id} [put]
func (h Handler) UpdateCalendar(c *gin.Context) {
	var r ports.UpdateCalendarRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc, err := h.calendarService.Update(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"calendar": uc})
}

// GetCalendar godoc
// @Summary Get a calendar
// @Description Get a calendar by ID
// @ID get-calendar
// @Accept json
// @Produce json
// @Param id path string true "calendar ID"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Calendar, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /calendar/{id} [get]
func (h Handler) GetCalendar(c *gin.Context) {
	uuid := c.Param("uuid")
	cc, err := h.calendarService.Get(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"calendar": cc})
}

// DeleteCalendar godoc
// @Summary Delete a calendar
// @Description Delete a calendar by ID
// @ID delete-calendar
// @Accept json
// @Produce json
// @Param id path string true "calendar ID"
// @Success 204 {object} handlers.JSONOKResponse{data=domain.Calendar, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /calendar/{id} [delete]
func (h Handler) DeleteCalendar(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.calendarService.Delete(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary Retrieve all calendars
// @Description Retrieve all calendars
// @ID list-calendar
// @Accept json
// @Produce json
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Calendar, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /calendars [get]
func (h Handler) List(c *gin.Context) {
	cs, err := h.calendarService.List(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"calendars": cs})
}

func NewHandler(cs ports.CalendarServicer) Handler {
	return Handler{
		calendarService: cs,
	}
}
