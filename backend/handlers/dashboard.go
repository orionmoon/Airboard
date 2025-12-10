package handlers

import (
	"net/http"

	"airboard/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DashboardHandler struct {
	db *gorm.DB
}

func NewDashboardHandler(db *gorm.DB) *DashboardHandler {
	return &DashboardHandler{db: db}
}

// @Summary Dashboard utilisateur
// @Description Récupère les groupes d'applications et applications accessibles à l'utilisateur
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.DashboardResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /dashboard [get]
func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "Unauthorized",
			Message: "Utilisateur non authentifié",
			Code:    http.StatusUnauthorized,
		})
		return
	}

	role, _ := c.Get("role")

	var appGroups []models.AppGroupWithApps

	if role == "admin" {
		// Les admins voient tout
		var allAppGroups []models.AppGroup
		if err := h.db.Where("is_active = ?", true).
			Order("\"order\" ASC, name ASC").
			Find(&allAppGroups).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "Internal Server Error",
				Message: "Erreur lors de la récupération des groupes d'applications",
				Code:    http.StatusInternalServerError,
			})
			return
		}

		for _, ag := range allAppGroups {
			var apps []models.Application
			h.db.Where("app_group_id = ? AND is_active = ?", ag.ID, true).
				Order("\"order\" ASC, name ASC").
				Find(&apps)

			appGroups = append(appGroups, models.AppGroupWithApps{
				AppGroup:     ag,
				Applications: apps,
			})
		}
	} else if role == "group_admin" {
		// Les group admins voient les AppGroups des groupes qu'ils administrent
		var groupAdminAppGroups []models.AppGroup
		if err := h.db.Table("app_groups").
			Joins("JOIN group_app_groups ON app_groups.id = group_app_groups.app_group_id").
			Joins("JOIN group_admins ON group_app_groups.group_id = group_admins.group_id").
			Where("group_admins.user_id = ? AND app_groups.is_active = ?", userID, true).
			Order("app_groups.\"order\" ASC, app_groups.name ASC").
			Find(&groupAdminAppGroups).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "Internal Server Error",
				Message: "Erreur lors de la récupération des groupes d'applications",
				Code:    http.StatusInternalServerError,
			})
			return
		}

		for _, ag := range groupAdminAppGroups {
			var apps []models.Application
			h.db.Where("app_group_id = ? AND is_active = ?", ag.ID, true).
				Order("\"order\" ASC, name ASC").
				Find(&apps)

			appGroups = append(appGroups, models.AppGroupWithApps{
				AppGroup:     ag,
				Applications: apps,
			})
		}
	} else {
		// Les utilisateurs normaux voient seulement leurs groupes autorisés
		var userAppGroups []models.AppGroup
		if err := h.db.Table("app_groups").
			Joins("JOIN group_app_groups ON app_groups.id = group_app_groups.app_group_id").
			Joins("JOIN user_groups ON group_app_groups.group_id = user_groups.group_id").
			Where("user_groups.user_id = ? AND app_groups.is_active = ?", userID, true).
			Order("app_groups.\"order\" ASC, app_groups.name ASC").
			Find(&userAppGroups).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "Internal Server Error",
				Message: "Erreur lors de la récupération des groupes d'applications",
				Code:    http.StatusInternalServerError,
			})
			return
		}

		for _, ag := range userAppGroups {
			var apps []models.Application
			h.db.Where("app_group_id = ? AND is_active = ?", ag.ID, true).
				Order("\"order\" ASC, name ASC").
				Find(&apps)

			appGroups = append(appGroups, models.AppGroupWithApps{
				AppGroup:     ag,
				Applications: apps,
			})
		}
	}

	// Statistiques (pour les admins seulement)
	stats := models.DashboardStats{}
	if role == "admin" {
		h.db.Model(&models.AppGroup{}).Where("is_active = ?", true).Count(&stats.TotalAppGroups)
		h.db.Model(&models.Application{}).Where("is_active = ?", true).Count(&stats.TotalApplications)
		h.db.Model(&models.User{}).Where("is_active = ?", true).Count(&stats.TotalUsers)
		h.db.Model(&models.Group{}).Where("is_active = ?", true).Count(&stats.TotalGroups)
	}

	response := models.DashboardResponse{
		AppGroups: appGroups,
		Stats:     stats,
	}

	c.JSON(http.StatusOK, response)
}