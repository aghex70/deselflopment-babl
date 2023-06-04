package entry

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	entryService ports.EntryServicer
}

// CreateEntry godoc
// @Summary Create a entry
// @Description Create a entry
// @ID create-entry
// @Accept json
// @Produce json
// @Param entry body domain.Entry true "entry data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Entry, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /entry [post]
func (h Handler) CreateEntry(c *gin.Context) {
	var r ports.CreateEntryRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ne, err := h.entryService.Create(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"entry": ne})
}

// UpdateEntry godoc
// @Summary Update a entry
// @Description Update a entry
// @ID update-entry
// @Accept json
// @Produce json
// @Param  id path string true "entry ID" entry body domain.Entry true "entry data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Entry, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /entry/{id} [put]
func (h Handler) UpdateEntry(c *gin.Context) {
	var r ports.UpdateEntryRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ue, err := h.entryService.Update(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entry": ue})
}

// GetEntry godoc
// @Summary Get a entry
// @Description Get a entry by ID
// @ID get-entry
// @Accept json
// @Produce json
// @Param id path string true "entry ID"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Entry, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /entry/{id} [get]
func (h Handler) GetEntry(c *gin.Context) {
	uuid := c.Param("uuid")
	ee, err := h.entryService.Get(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entry": ee})
}

// DeleteEntry godoc
// @Summary Delete a entry
// @Description Delete a entry by ID
// @ID delete-entry
// @Accept json
// @Produce json
// @Param id path string true "entry ID"
// @Success 204 {object} handlers.JSONOKResponse{data=domain.Entry, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /entry/{id} [delete]
func (h Handler) DeleteEntry(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.entryService.Delete(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary Retrieve all entries
// @Description Retrieve all entries
// @ID list-entry
// @Accept json
// @Produce json
// @Success 200 {object} handlers.JSONOKResponse{data=domain.Entry, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /entries [get]
func (h Handler) List(c *gin.Context) {
	es, err := h.entryService.List(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entries": es})
}

func NewHandler(es ports.EntryServicer) Handler {
	return Handler{
		entryService: es,
	}
}
