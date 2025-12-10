package middleware

import (
	"net/http"

	"airboard/models"

	"github.com/gin-gonic/gin"
)

// RequireGroupAdmin vérifie que l'utilisateur est group_admin OU admin
func (am *AuthMiddleware) RequireGroupAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Authentication requise",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok || (roleStr != "group_admin" && roleStr != "admin") {
			c.JSON(http.StatusForbidden, models.ErrorResponse{
				Error:   "Forbidden",
				Message: "Rôle Group Admin ou Admin requis",
				Code:    http.StatusForbidden,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CanManageGroup vérifie si l'utilisateur peut gérer un groupe spécifique
func CanManageGroup(c *gin.Context, groupID uint) bool {
	role := c.GetString("role")
	if role == "admin" {
		return true // Admin global peut tout gérer
	}

	if role != "group_admin" {
		return false
	}

	managedGroupIDsInterface, exists := c.Get("managed_group_ids")
	if !exists {
		return false
	}

	managedGroupIDs, ok := managedGroupIDsInterface.([]uint)
	if !ok {
		return false
	}

	for _, id := range managedGroupIDs {
		if id == groupID {
			return true
		}
	}

	return false
}

// GetManagedGroupIDs retourne les IDs des groupes gérés (helper)
func GetManagedGroupIDs(c *gin.Context) []uint {
	managedGroupIDsInterface, exists := c.Get("managed_group_ids")
	if !exists {
		return []uint{}
	}

	managedGroupIDs, ok := managedGroupIDsInterface.([]uint)
	if !ok {
		return []uint{}
	}

	return managedGroupIDs
}

// CanManageAppGroup vérifie si un group_admin peut gérer un AppGroup spécifique
// Un group_admin peut administrer un AppGroup uniquement si :
// - L'AppGroup est privé (IsPrivate = true)
// - ET l'OwnerGroupID correspond à l'un des groupes qu'il administre
// Les AppGroups publics (IsPrivate = false) ne peuvent être administrés que par l'admin global
func CanManageAppGroupWithDB(c *gin.Context, appGroupID uint, db interface{}) bool {
	role := c.GetString("role")
	if role == "admin" {
		return true // Admin global peut tout gérer
	}

	if role != "group_admin" {
		return false
	}

	managedGroupIDs := GetManagedGroupIDs(c)
	if len(managedGroupIDs) == 0 {
		return false
	}

	// Récupérer l'AppGroup pour vérifier IsPrivate et OwnerGroupID
	type AppGroupCheck struct {
		IsPrivate    bool
		OwnerGroupID *uint
	}
	var appGroup AppGroupCheck

	if dbConn, ok := db.(interface{ Raw(string, ...interface{}) interface{ Scan(interface{}) error } }); ok {
		err := dbConn.Raw(`
			SELECT is_private, owner_group_id
			FROM app_groups
			WHERE id = ? AND deleted_at IS NULL
		`, appGroupID).Scan(&appGroup)

		if err != nil {
			return false
		}
	} else {
		return false
	}

	// Vérifier si l'AppGroup est privé
	if !appGroup.IsPrivate {
		return false // Seul l'admin global peut gérer les AppGroups publics
	}

	// Vérifier si l'OwnerGroupID correspond à l'un des groupes administrés
	if appGroup.OwnerGroupID == nil {
		return false // AppGroup privé sans propriétaire ne peut pas être administré
	}

	for _, groupID := range managedGroupIDs {
		if *appGroup.OwnerGroupID == groupID {
			return true
		}
	}

	return false
}
