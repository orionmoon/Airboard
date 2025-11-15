package handlers

import (
	"net/http"
	"strconv"

	"airboard/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminHandler struct {
	db *gorm.DB
}

func NewAdminHandler(db *gorm.DB) *AdminHandler {
	return &AdminHandler{db: db}
}

// ============ GESTION DES GROUPES D'APPLICATIONS ============

// @Summary Lister les groupes d'applications
// @Description Récupère tous les groupes d'applications (admin uniquement)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Numéro de page" default(1)
// @Param limit query int false "Nombre d'éléments par page" default(10)
// @Success 200 {object} models.PaginatedResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Router /admin/app-groups [get]
func (h *AdminHandler) GetAppGroups(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var appGroups []models.AppGroup
	var total int64

	h.db.Model(&models.AppGroup{}).Count(&total)
	if err := h.db.Preload("Applications").
		Order("\"order\" ASC, name ASC").
		Limit(limit).Offset(offset).
		Find(&appGroups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des groupes d'applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	totalPages := int(total)/limit + 1
	if int(total)%limit == 0 && total > 0 {
		totalPages = int(total) / limit
	}

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Data:       appGroups,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	})
}

// @Summary Créer un groupe d'applications
// @Description Crée un nouveau groupe d'applications (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param appGroup body models.AppGroup true "Données du groupe d'applications"
// @Success 201 {object} models.AppGroup
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Router /admin/app-groups [post]
func (h *AdminHandler) CreateAppGroup(c *gin.Context) {
	var appGroup models.AppGroup
	if err := c.ShouldBindJSON(&appGroup); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Create(&appGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la création du groupe d'applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, appGroup)
}

// @Summary Modifier un groupe d'applications
// @Description Modifie un groupe d'applications existant (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID du groupe d'applications"
// @Param appGroup body models.AppGroup true "Données modifiées"
// @Success 200 {object} models.AppGroup
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /admin/app-groups/{id} [put]
func (h *AdminHandler) UpdateAppGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var appGroup models.AppGroup
	if err := h.db.First(&appGroup, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Groupe d'applications non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	var updateData models.AppGroup
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Model(&appGroup).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la modification",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, appGroup)
}

// @Summary Supprimer un groupe d'applications
// @Description Supprime un groupe d'applications (admin uniquement)
// @Tags Admin
// @Security BearerAuth
// @Param id path int true "ID du groupe d'applications"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /admin/app-groups/{id} [delete]
func (h *AdminHandler) DeleteAppGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Delete(&models.AppGroup{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Groupe d'applications supprimé avec succès",
	})
}

// ============ GESTION DES APPLICATIONS ============

// @Summary Lister les applications
// @Description Récupère toutes les applications (admin uniquement)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Numéro de page" default(1)
// @Param limit query int false "Nombre d'éléments par page" default(10)
// @Param app_group_id query int false "Filtrer par groupe d'applications"
// @Success 200 {object} models.PaginatedResponse
// @Router /admin/applications [get]
func (h *AdminHandler) GetApplications(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	appGroupID := c.Query("app_group_id")
	offset := (page - 1) * limit

	var applications []models.Application
	var total int64

	query := h.db.Model(&models.Application{})
	if appGroupID != "" {
		query = query.Where("app_group_id = ?", appGroupID)
	}

	query.Count(&total)
	if err := query.Preload("AppGroup").
		Order("\"order\" ASC, name ASC").
		Limit(limit).Offset(offset).
		Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	totalPages := int(total)/limit + 1
	if int(total)%limit == 0 && total > 0 {
		totalPages = int(total) / limit
	}

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Data:       applications,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	})
}

// @Summary Créer une application
// @Description Crée une nouvelle application (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param application body models.Application true "Données de l'application"
// @Success 201 {object} models.Application
// @Router /admin/applications [post]
func (h *AdminHandler) CreateApplication(c *gin.Context) {
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Create(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la création de l'application",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Charger la relation AppGroup
	h.db.Preload("AppGroup").First(&application, application.ID)

	c.JSON(http.StatusCreated, application)
}

// @Summary Modifier une application
// @Description Modifie une application existante (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de l'application"
// @Param application body models.Application true "Données modifiées"
// @Success 200 {object} models.Application
// @Router /admin/applications/{id} [put]
func (h *AdminHandler) UpdateApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var application models.Application
	if err := h.db.First(&application, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Application non trouvée",
			Code:    http.StatusNotFound,
		})
		return
	}

	var updateData models.Application
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Model(&application).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la modification",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Charger la relation AppGroup
	h.db.Preload("AppGroup").First(&application, application.ID)

	c.JSON(http.StatusOK, application)
}

// @Summary Supprimer une application
// @Description Supprime une application (admin uniquement)
// @Tags Admin
// @Security BearerAuth
// @Param id path int true "ID de l'application"
// @Success 200 {object} models.SuccessResponse
// @Router /admin/applications/{id} [delete]
func (h *AdminHandler) DeleteApplication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.db.Delete(&models.Application{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Application supprimée avec succès",
	})
}

// ============ GESTION DES UTILISATEURS ============

// @Summary Lister les utilisateurs
// @Description Récupère tous les utilisateurs (admin uniquement)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Router /admin/users [get]
func (h *AdminHandler) GetUsers(c *gin.Context) {
	var users []models.User
	if err := h.db.Preload("Groups").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des utilisateurs",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Masquer les mots de passe
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Créer un utilisateur
// @Description Crée un nouvel utilisateur (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body models.RegisterRequest true "Données de l'utilisateur"
// @Success 201 {object} models.User
// @Router /admin/users [post]
func (h *AdminHandler) CreateUser(c *gin.Context) {
	var createData struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
		IsActive  bool   `json:"is_active"`
		GroupIDs  []uint `json:"group_ids"`
	}

	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Vérifier si l'utilisateur existe déjà
	var existingUser models.User
	if err := h.db.Where("username = ? OR email = ?", createData.Username, createData.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error:   "Conflict",
			Message: "Nom d'utilisateur ou email déjà utilisé",
			Code:    http.StatusConflict,
		})
		return
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors du hashage du mot de passe",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	user := models.User{
		Username:  createData.Username,
		Email:     createData.Email,
		Password:  string(hashedPassword),
		FirstName: createData.FirstName,
		LastName:  createData.LastName,
		Role:      createData.Role,
		IsActive:  createData.IsActive,
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la création de l'utilisateur",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Associer les groupes si fournis
	if createData.GroupIDs != nil && len(createData.GroupIDs) > 0 {
		var groups []models.Group
		if err := h.db.Where("id IN ?", createData.GroupIDs).Find(&groups).Error; err == nil {
			h.db.Model(&user).Association("Groups").Replace(groups)
		}
	}

	user.Password = ""
	h.db.Preload("Groups").First(&user, user.ID)

	c.JSON(http.StatusCreated, user)
}

// @Summary Modifier un utilisateur
// @Description Modifie un utilisateur existant (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de l'utilisateur"
// @Param user body models.User true "Données modifiées"
// @Success 200 {object} models.User
// @Router /admin/users/{id} [put]
func (h *AdminHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Utilisateur non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	var updateData struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
		IsActive  *bool  `json:"is_active"`
		Password  string `json:"password,omitempty"`
		GroupIDs  []uint `json:"group_ids"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Mise à jour des champs
	if updateData.Username != "" {
		user.Username = updateData.Username
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}
	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName
	}
	if updateData.LastName != "" {
		user.LastName = updateData.LastName
	}
	if updateData.Role != "" {
		user.Role = updateData.Role
	}
	if updateData.IsActive != nil {
		user.IsActive = *updateData.IsActive
	}

	// Hash du nouveau mot de passe si fourni
	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:   "Internal Server Error",
				Message: "Erreur lors du hashage du mot de passe",
				Code:    http.StatusInternalServerError,
			})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la modification",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Mise à jour des groupes si fournis
	if updateData.GroupIDs != nil {
		var groups []models.Group
		if err := h.db.Where("id IN ?", updateData.GroupIDs).Find(&groups).Error; err == nil {
			h.db.Model(&user).Association("Groups").Replace(groups)
		}
	}

	// Masquer le mot de passe
	user.Password = ""

	// Charger les relations
	h.db.Preload("Groups").First(&user, user.ID)

	c.JSON(http.StatusOK, user)
}

// @Summary Supprimer un utilisateur
// @Description Supprime un utilisateur (admin uniquement)
// @Tags Admin
// @Security BearerAuth
// @Param id path int true "ID de l'utilisateur"
// @Success 200 {object} models.SuccessResponse
// @Router /admin/users/{id} [delete]
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Vérifier que l'utilisateur existe
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Utilisateur non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Supprimer les associations avec les groupes
	h.db.Model(&user).Association("Groups").Clear()

	// Supprimer l'utilisateur
	if err := h.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Utilisateur supprimé avec succès",
	})
}

// GetDeletedUsers récupère tous les utilisateurs supprimés (soft deleted)
func (h *AdminHandler) GetDeletedUsers(c *gin.Context) {
	var users []models.User

	// Utiliser Unscoped() pour voir les enregistrements soft deleted
	if err := h.db.Unscoped().Where("deleted_at IS NOT NULL").Preload("Groups").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des utilisateurs supprimés",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Masquer les mots de passe
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// RestoreUser restaure un utilisateur supprimé
func (h *AdminHandler) RestoreUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Trouver l'utilisateur supprimé
	var user models.User
	if err := h.db.Unscoped().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Utilisateur non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Vérifier qu'il est bien supprimé
	if user.DeletedAt.Time.IsZero() {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Cet utilisateur n'est pas supprimé",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Restaurer l'utilisateur
	user.DeletedAt = gorm.DeletedAt{}
	if err := h.db.Unscoped().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la restauration",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Utilisateur restauré avec succès",
		Data:    user,
	})
}

// PermanentlyDeleteUser supprime définitivement un utilisateur
func (h *AdminHandler) PermanentlyDeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Trouver l'utilisateur supprimé
	var user models.User
	if err := h.db.Unscoped().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Utilisateur non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Vérifier qu'il est bien supprimé
	if user.DeletedAt.Time.IsZero() {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Cet utilisateur doit d'abord être supprimé",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Supprimer définitivement les associations
	h.db.Unscoped().Model(&user).Association("Groups").Clear()

	// Supprimer définitivement
	if err := h.db.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression définitive",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Utilisateur supprimé définitivement",
	})
}

// @Summary Assigner un utilisateur à des groupes
// @Description Assigne un utilisateur à des groupes d'utilisateurs (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de l'utilisateur"
// @Param groups body []int true "IDs des groupes"
// @Success 200 {object} models.User
// @Router /admin/users/{id}/groups [put]
func (h *AdminHandler) AssignUserToGroups(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Utilisateur non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	var groupIDs []uint
	if err := c.ShouldBindJSON(&groupIDs); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "IDs de groupes invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Récupérer les groupes
	var groups []models.Group
	if err := h.db.Where("id IN ?", groupIDs).Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des groupes",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Remplacer les associations
	if err := h.db.Model(&user).Association("Groups").Replace(groups); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de l'assignation",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Recharger avec les relations
	h.db.Preload("Groups").First(&user, user.ID)
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// @Summary Lister les groupes
// @Description Récupère tous les groupes d'utilisateurs (admin uniquement)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Group
// @Router /admin/groups [get]
func (h *AdminHandler) GetGroups(c *gin.Context) {
	var groups []models.Group
	if err := h.db.Preload("Users").Preload("AppGroups").Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des groupes",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// @Summary Créer un groupe
// @Description Crée un nouveau groupe d'utilisateurs (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body models.Group true "Données du groupe"
// @Success 201 {object} models.Group
// @Router /admin/groups [post]
func (h *AdminHandler) CreateGroup(c *gin.Context) {
	var createData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		IsActive    bool   `json:"is_active"`
		AppGroupIDs []uint `json:"app_group_ids"`
	}

	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	group := models.Group{
		Name:        createData.Name,
		Description: createData.Description,
		Color:       createData.Color,
		IsActive:    createData.IsActive,
	}

	if err := h.db.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la création du groupe",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Associer les groupes d'applications si fournis
	if createData.AppGroupIDs != nil && len(createData.AppGroupIDs) > 0 {
		var appGroups []models.AppGroup
		if err := h.db.Where("id IN ?", createData.AppGroupIDs).Find(&appGroups).Error; err == nil {
			h.db.Model(&group).Association("AppGroups").Replace(appGroups)
		}
	}

	// Charger les relations
	h.db.Preload("Users").Preload("AppGroups").First(&group, group.ID)

	c.JSON(http.StatusCreated, group)
}

// @Summary Modifier un groupe
// @Description Modifie un groupe d'utilisateurs existant (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID du groupe"
// @Param group body models.Group true "Données modifiées"
// @Success 200 {object} models.Group
// @Router /admin/groups/{id} [put]
func (h *AdminHandler) UpdateGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var group models.Group
	if err := h.db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Groupe non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	var updateData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		IsActive    *bool  `json:"is_active"`
		AppGroupIDs []uint `json:"app_group_ids"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "Données invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Mise à jour des champs
	if updateData.Name != "" {
		group.Name = updateData.Name
	}
	if updateData.Description != "" {
		group.Description = updateData.Description
	}
	if updateData.Color != "" {
		group.Color = updateData.Color
	}
	if updateData.IsActive != nil {
		group.IsActive = *updateData.IsActive
	}

	if err := h.db.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la modification",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Mise à jour des groupes d'applications si fournis
	if updateData.AppGroupIDs != nil {
		var appGroups []models.AppGroup
		// Si le tableau n'est pas vide, chercher les groupes
		if len(updateData.AppGroupIDs) > 0 {
			if err := h.db.Where("id IN ?", updateData.AppGroupIDs).Find(&appGroups).Error; err == nil {
				h.db.Model(&group).Association("AppGroups").Replace(appGroups)
			}
		} else {
			// Tableau vide signifie supprimer toutes les associations
			h.db.Model(&group).Association("AppGroups").Clear()
		}
	}

	// Charger les relations
	h.db.Preload("Users").Preload("AppGroups").First(&group, group.ID)

	c.JSON(http.StatusOK, group)
}

// @Summary Supprimer un groupe
// @Description Supprime un groupe d'utilisateurs (admin uniquement)
// @Tags Admin
// @Security BearerAuth
// @Param id path int true "ID du groupe"
// @Success 200 {object} models.SuccessResponse
// @Router /admin/groups/{id} [delete]
func (h *AdminHandler) DeleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var group models.Group
	if err := h.db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Groupe non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Supprimer les associations
	h.db.Model(&group).Association("Users").Clear()
	h.db.Model(&group).Association("AppGroups").Clear()

	// Supprimer le groupe
	if err := h.db.Delete(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Groupe supprimé avec succès",
	})
}

// @Summary Assigner des groupes d'applications à un groupe d'utilisateurs
// @Description Définit les permissions d'un groupe d'utilisateurs (admin uniquement)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID du groupe d'utilisateurs"
// @Param app_groups body []int true "IDs des groupes d'applications"
// @Success 200 {object} models.Group
// @Router /admin/groups/{id}/app-groups [put]
func (h *AdminHandler) AssignGroupToAppGroups(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "ID invalide",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var group models.Group
	if err := h.db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Not Found",
			Message: "Groupe non trouvé",
			Code:    http.StatusNotFound,
		})
		return
	}

	var appGroupIDs []uint
	if err := c.ShouldBindJSON(&appGroupIDs); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Bad Request",
			Message: "IDs de groupes d'applications invalides",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Récupérer les groupes d'applications
	var appGroups []models.AppGroup
	if err := h.db.Where("id IN ?", appGroupIDs).Find(&appGroups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la récupération des groupes d'applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Remplacer les associations
	if err := h.db.Model(&group).Association("AppGroups").Replace(appGroups); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de l'assignation",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Recharger avec les relations
	h.db.Preload("Users").Preload("AppGroups").First(&group, group.ID)

	c.JSON(http.StatusOK, group)
}

// ============ GESTION DE LA BASE DE DONNÉES ============

// @Summary Réinitialiser la base de données
// @Description Supprime toutes les données et recrée les données initiales (admin uniquement - ATTENTION: opération destructive)
// @Tags Admin
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /admin/database/reset [post]
func (h *AdminHandler) ResetDatabase(c *gin.Context) {
	// Supprimer toutes les données dans l'ordre pour éviter les erreurs de contraintes

	// Nettoyer les tables de jointure (many-to-many)
	if err := h.db.Exec("DELETE FROM user_groups").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des associations utilisateurs-groupes",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if err := h.db.Exec("DELETE FROM group_app_groups").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des associations groupes-applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Supprimer les applications
	if err := h.db.Exec("DELETE FROM applications").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Supprimer les groupes d'applications
	if err := h.db.Exec("DELETE FROM app_groups").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des groupes d'applications",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Supprimer les utilisateurs
	if err := h.db.Exec("DELETE FROM users").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des utilisateurs",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Supprimer les groupes d'utilisateurs
	if err := h.db.Exec("DELETE FROM groups").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des groupes d'utilisateurs",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Supprimer les paramètres
	if err := h.db.Exec("DELETE FROM app_settings").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Erreur lors de la suppression des paramètres",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Réinitialiser les séquences pour repartir à 1
	tables := []string{"users", "groups", "app_groups", "applications", "app_settings"}
	for _, table := range tables {
		sequenceName := table + "_id_seq"
		if err := h.db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			// Ignorer les erreurs de séquence, elles peuvent ne pas exister
			continue
		}
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Base de données réinitialisée avec succès. Veuillez redémarrer le serveur pour recréer les données initiales.",
		Data: map[string]interface{}{
			"note": "Les données de démonstration seront recréées au prochain démarrage du serveur",
		},
	})
}
