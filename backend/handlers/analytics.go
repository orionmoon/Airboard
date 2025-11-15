package handlers

import (
	"net/http"
	"time"

	"airboard/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnalyticsHandler struct {
	db *gorm.DB
}

func NewAnalyticsHandler(db *gorm.DB) *AnalyticsHandler {
	return &AnalyticsHandler{db: db}
}

// TrackClick enregistre un clic sur une application
func (h *AnalyticsHandler) TrackClick(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "Unauthorized",
			Message: "Utilisateur non authentifié",
			Code:    http.StatusUnauthorized,
		})
		return
	}

	var requestData struct {
		ApplicationID uint `json:"application_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Créer l'enregistrement du clic
	click := models.ApplicationClick{
		UserID:        userID.(uint),
		ApplicationID: requestData.ApplicationID,
		ClickedAt:     time.Now(),
	}

	if err := h.db.Create(&click).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de l'enregistrement du clic",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Clic enregistré avec succès",
	})
}

// GetDashboard retourne les statistiques complètes pour le dashboard analytics
func (h *AnalyticsHandler) GetDashboard(c *gin.Context) {
	var dashboard models.AnalyticsDashboard

	// Total des clics
	h.db.Model(&models.ApplicationClick{}).Count(&dashboard.TotalClicks)

	// Nombre d'utilisateurs uniques ayant cliqué
	h.db.Model(&models.ApplicationClick{}).
		Distinct("user_id").
		Count(&dashboard.TotalUniqueUsers)

	// Top 10 applications les plus cliquées
	topApps := []models.ApplicationStats{}
	h.db.Table("application_clicks").
		Select(`
			applications.id as application_id,
			applications.name as application_name,
			applications.icon,
			applications.color,
			COUNT(application_clicks.id) as click_count,
			COUNT(DISTINCT application_clicks.user_id) as unique_users
		`).
		Joins("JOIN applications ON applications.id = application_clicks.application_id").
		Where("applications.deleted_at IS NULL").
		Group("applications.id, applications.name, applications.icon, applications.color").
		Order("click_count DESC").
		Limit(10).
		Scan(&topApps)
	dashboard.TopApplications = topApps

	// Top 10 utilisateurs les plus actifs
	topUsers := []models.UserStats{}
	h.db.Table("application_clicks").
		Select(`
			users.id as user_id,
			users.username,
			users.first_name,
			users.last_name,
			COUNT(application_clicks.id) as click_count,
			COUNT(DISTINCT application_clicks.application_id) as unique_apps,
			MAX(application_clicks.clicked_at) as last_activity
		`).
		Joins("JOIN users ON users.id = application_clicks.user_id").
		Where("users.deleted_at IS NULL").
		Group("users.id, users.username, users.first_name, users.last_name").
		Order("click_count DESC").
		Limit(10).
		Scan(&topUsers)
	dashboard.TopUsers = topUsers

	// Activité quotidienne des 30 derniers jours
	dailyStats := []models.DailyStats{}
	h.db.Table("application_clicks").
		Select(`
			DATE(clicked_at) as date,
			COUNT(*) as click_count,
			COUNT(DISTINCT user_id) as unique_users
		`).
		Where("clicked_at >= ?", time.Now().AddDate(0, 0, -30)).
		Group("DATE(clicked_at)").
		Order("date ASC").
		Scan(&dailyStats)
	dashboard.DailyActivity = dailyStats

	// Clics des 7 derniers jours
	h.db.Model(&models.ApplicationClick{}).
		Where("clicked_at >= ?", time.Now().AddDate(0, 0, -7)).
		Count(&dashboard.ClicksLast7Days)

	// Clics des 30 derniers jours
	h.db.Model(&models.ApplicationClick{}).
		Where("clicked_at >= ?", time.Now().AddDate(0, 0, -30)).
		Count(&dashboard.ClicksLast30Days)

	c.JSON(http.StatusOK, dashboard)
}

// GetApplicationStats retourne les statistiques détaillées d'une application
func (h *AnalyticsHandler) GetApplicationStats(c *gin.Context) {
	appID := c.Param("id")

	var stats struct {
		TotalClicks     int64                 `json:"total_clicks"`
		UniqueUsers     int64                 `json:"unique_users"`
		DailyActivity   []models.DailyStats   `json:"daily_activity"`
		TopUsers        []models.UserStats    `json:"top_users"`
	}

	// Total des clics pour cette application
	h.db.Model(&models.ApplicationClick{}).
		Where("application_id = ?", appID).
		Count(&stats.TotalClicks)

	// Utilisateurs uniques
	h.db.Model(&models.ApplicationClick{}).
		Where("application_id = ?", appID).
		Distinct("user_id").
		Count(&stats.UniqueUsers)

	// Activité quotidienne des 30 derniers jours
	dailyStats := []models.DailyStats{}
	h.db.Table("application_clicks").
		Select(`
			DATE(clicked_at) as date,
			COUNT(*) as click_count,
			COUNT(DISTINCT user_id) as unique_users
		`).
		Where("application_id = ? AND clicked_at >= ?", appID, time.Now().AddDate(0, 0, -30)).
		Group("DATE(clicked_at)").
		Order("date ASC").
		Scan(&dailyStats)
	stats.DailyActivity = dailyStats

	// Top utilisateurs pour cette application
	topUsers := []models.UserStats{}
	h.db.Table("application_clicks").
		Select(`
			users.id as user_id,
			users.username,
			users.first_name,
			users.last_name,
			COUNT(application_clicks.id) as click_count,
			MAX(application_clicks.clicked_at) as last_activity
		`).
		Joins("JOIN users ON users.id = application_clicks.user_id").
		Where("application_clicks.application_id = ? AND users.deleted_at IS NULL", appID).
		Group("users.id, users.username, users.first_name, users.last_name").
		Order("click_count DESC").
		Limit(10).
		Scan(&topUsers)
	stats.TopUsers = topUsers

	c.JSON(http.StatusOK, stats)
}

// GetUserStats retourne les statistiques détaillées d'un utilisateur
func (h *AnalyticsHandler) GetUserStats(c *gin.Context) {
	userID := c.Param("id")

	var stats struct {
		TotalClicks      int64                    `json:"total_clicks"`
		UniqueApps       int64                    `json:"unique_apps"`
		DailyActivity    []models.DailyStats      `json:"daily_activity"`
		TopApplications  []models.ApplicationStats `json:"top_applications"`
	}

	// Total des clics pour cet utilisateur
	h.db.Model(&models.ApplicationClick{}).
		Where("user_id = ?", userID).
		Count(&stats.TotalClicks)

	// Applications uniques
	h.db.Model(&models.ApplicationClick{}).
		Where("user_id = ?", userID).
		Distinct("application_id").
		Count(&stats.UniqueApps)

	// Activité quotidienne des 30 derniers jours
	dailyStats := []models.DailyStats{}
	h.db.Table("application_clicks").
		Select(`
			DATE(clicked_at) as date,
			COUNT(*) as click_count
		`).
		Where("user_id = ? AND clicked_at >= ?", userID, time.Now().AddDate(0, 0, -30)).
		Group("DATE(clicked_at)").
		Order("date ASC").
		Scan(&dailyStats)
	stats.DailyActivity = dailyStats

	// Top applications pour cet utilisateur
	topApps := []models.ApplicationStats{}
	h.db.Table("application_clicks").
		Select(`
			applications.id as application_id,
			applications.name as application_name,
			applications.icon,
			applications.color,
			COUNT(application_clicks.id) as click_count
		`).
		Joins("JOIN applications ON applications.id = application_clicks.application_id").
		Where("application_clicks.user_id = ? AND applications.deleted_at IS NULL", userID).
		Group("applications.id, applications.name, applications.icon, applications.color").
		Order("click_count DESC").
		Limit(10).
		Scan(&topApps)
	stats.TopApplications = topApps

	c.JSON(http.StatusOK, stats)
}
