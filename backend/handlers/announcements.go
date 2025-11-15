package handlers

import (
	"net/http"
	"strconv"
	"time"

	"airboard/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnnouncementHandler struct {
	db *gorm.DB
}

func NewAnnouncementHandler(db *gorm.DB) *AnnouncementHandler {
	return &AnnouncementHandler{db: db}
}

// GetAllAnnouncements retourne toutes les annonces (admin)
func (h *AnnouncementHandler) GetAllAnnouncements(c *gin.Context) {
	var announcements []models.Announcement

	if err := h.db.Order("priority DESC, created_at DESC").Find(&announcements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des annonces",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, announcements)
}

// GetActiveAnnouncements retourne uniquement les annonces actives et valides (pour les users)
func (h *AnnouncementHandler) GetActiveAnnouncements(c *gin.Context) {
	var announcements []models.Announcement
	now := time.Now()

	query := h.db.Where("is_active = ?", true)

	// Filtrer par dates si elles sont définies
	query = query.Where("(start_date IS NULL OR start_date <= ?)", now)
	query = query.Where("(end_date IS NULL OR end_date >= ?)", now)

	if err := query.Order("priority DESC, created_at DESC").Find(&announcements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des annonces",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, announcements)
}

// GetAnnouncement retourne une annonce spécifique
func (h *AnnouncementHandler) GetAnnouncement(c *gin.Context) {
	id := c.Param("id")
	var announcement models.Announcement

	if err := h.db.First(&announcement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "Not Found",
				Message: "Annonce non trouvée",
				Code:    http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération de l'annonce",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

// CreateAnnouncement crée une nouvelle annonce
func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	var req models.AnnouncementRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides: " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	announcement := models.Announcement{
		Title:     req.Title,
		Content:   req.Content,
		Type:      req.Type,
		Priority:  req.Priority,
		IsActive:  req.IsActive,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := h.db.Create(&announcement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la création de l'annonce",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, announcement)
}

// UpdateAnnouncement met à jour une annonce existante
func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	id := c.Param("id")
	var announcement models.Announcement

	if err := h.db.First(&announcement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "Not Found",
				Message: "Annonce non trouvée",
				Code:    http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération de l'annonce",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	var req models.AnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides: " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Mise à jour des champs
	announcement.Title = req.Title
	announcement.Content = req.Content
	announcement.Type = req.Type
	announcement.Priority = req.Priority
	announcement.IsActive = req.IsActive
	announcement.StartDate = req.StartDate
	announcement.EndDate = req.EndDate

	if err := h.db.Save(&announcement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la mise à jour de l'annonce",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

// DeleteAnnouncement supprime une annonce (soft delete)
func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Delete(&models.Announcement{}, idInt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression de l'annonce",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Annonce supprimée avec succès",
	})
}
