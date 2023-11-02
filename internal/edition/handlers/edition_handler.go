package handlers

import (
	"go_newsletter_api/internal/edition/model"
	"go_newsletter_api/internal/edition/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditionHandler struct {
	service service.EditionService
}

func NewEditionHandler(service service.EditionService) *EditionHandler {
	return &EditionHandler{service}
}

func CreateEditionHandler(h *EditionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var edition model.Edition
		if err := c.ShouldBindJSON(&edition); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := h.service.CreateEdition(&edition); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create edition", "details": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, edition)
	}
}

func UpdateEditionHandler(h *EditionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := parseID(c)
		var edition model.Edition
		if err := c.ShouldBindJSON(&edition); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		edition.ID = id
		if err := h.service.UpdateEdition(&edition); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update edition", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, edition)
	}
}

func GetEditionByIDHandler(h *EditionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := parseID(c)
		edition, err := h.service.GetEditionByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Edition not found"})
			return
		}

		c.JSON(http.StatusOK, edition)
	}
}

func GetEditionsByNewsletterIDHandler(h *EditionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsletterID := parsenewsletterID(c)
		page, pageSize := parsePaginationParams(c)
		editions, count, err := h.service.GetEditionsByNewsletterID(newsletterID, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve editions", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"editions": editions, "count": count})
	}
}

func DeleteEditionHandler(h *EditionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := parseID(c)
		if err := h.service.DeleteEdition(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete edition", "details": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "Edition deleted successfully"})
	}
}

func parseID(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}

func parsenewsletterID(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("newsletterID"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}

func parsePaginationParams(c *gin.Context) (page, pageSize int) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ = strconv.Atoi(pageStr)

	pageSizeStr := c.DefaultQuery("page_size", "10")
	pageSize, _ = strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	return page, pageSize
}
