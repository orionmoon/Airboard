package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireEditor vérifie que l'utilisateur a le rôle editor ou admin
func (m *AuthMiddleware) RequireEditor() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok || (roleStr != "editor" && roleStr != "admin") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Editor or Admin role required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireEditorOrAdmin - Alias pour RequireEditor (même comportement)
func (m *AuthMiddleware) RequireEditorOrAdmin() gin.HandlerFunc {
	return m.RequireEditor()
}
