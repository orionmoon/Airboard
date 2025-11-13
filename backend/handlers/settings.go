package handlers

import (
	"net/http"
	"airboard/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SettingsHandler struct {
	DB *gorm.DB
}

func NewSettingsHandler(db *gorm.DB) *SettingsHandler {
	return &SettingsHandler{DB: db}
}

// GetAppSettings récupère les paramètres de l'application
func (h *SettingsHandler) GetAppSettings(c *gin.Context) {
	var settings models.AppSettings
	
	// Récupérer les paramètres (il ne devrait y en avoir qu'un seul)
	result := h.DB.First(&settings)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Créer les paramètres par défaut s'ils n'existent pas
			settings = models.AppSettings{
				AppName:        "Airboard",
				AppIcon:        "mdi:view-dashboard",
				DashboardTitle: "Dashboard",
				WelcomeMessage: "Welcome to your application portal",
			}
			
			if err := h.DB.Create(&settings).Error; err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Error:   "database_error",
					Message: "Failed to create default settings",
					Code:    http.StatusInternalServerError,
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "database_error",
				Message: "Failed to fetch app settings",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "App settings retrieved successfully",
		Data:    settings,
	})
}

// UpdateAppSettings met à jour les paramètres de l'application
func (h *SettingsHandler) UpdateAppSettings(c *gin.Context) {
	var request models.AppSettingsRequest
	
	// Valider les données de la requête
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	
	// Récupérer les paramètres existants
	var settings models.AppSettings
	result := h.DB.First(&settings)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Créer de nouveaux paramètres
			settings = models.AppSettings{
				AppName:        request.AppName,
				AppIcon:        request.AppIcon,
				DashboardTitle: request.DashboardTitle,
				WelcomeMessage: request.WelcomeMessage,
			}
			
			if err := h.DB.Create(&settings).Error; err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Error:   "database_error",
					Message: "Failed to create app settings",
					Code:    http.StatusInternalServerError,
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "database_error",
				Message: "Failed to fetch app settings",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	} else {
		// Mettre à jour les paramètres existants
		settings.AppName = request.AppName
		settings.AppIcon = request.AppIcon
		settings.DashboardTitle = request.DashboardTitle
		settings.WelcomeMessage = request.WelcomeMessage
		
		if err := h.DB.Save(&settings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "database_error",
				Message: "Failed to update app settings",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "App settings updated successfully",
		Data:    settings,
	})
}

// ResetAppSettings remet les paramètres aux valeurs par défaut
func (h *SettingsHandler) ResetAppSettings(c *gin.Context) {
	var settings models.AppSettings
	
	// Récupérer les paramètres existants
	result := h.DB.First(&settings)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "database_error",
			Message: "Failed to fetch app settings",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	
	// Remettre aux valeurs par défaut
	settings.AppName = "Airboard"
	settings.AppIcon = "mdi:view-dashboard"
	settings.DashboardTitle = "Dashboard"
	settings.WelcomeMessage = "Welcome to your application portal"
	
	if result.Error == gorm.ErrRecordNotFound {
		// Créer de nouveaux paramètres avec les valeurs par défaut
		if err := h.DB.Create(&settings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "database_error",
				Message: "Failed to create default settings",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	} else {
		// Mettre à jour les paramètres existants
		if err := h.DB.Save(&settings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "database_error",
				Message: "Failed to reset app settings",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "App settings reset to defaults successfully",
		Data:    settings,
	})
}