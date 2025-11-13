package middleware

import (
	"net/http"
	"strings"
	"time"

	"airboard/config"
	"airboard/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	config *config.Config
}

func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{config: cfg}
}

// RequireAuth middleware pour vérifier l'authentification
func (am *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Token d'autorisation manquant",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Extraire le token (format: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Format de token invalide",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Vérifier le token
		claims, err := am.verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Token invalide ou expiré",
				Code:    http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// Stocker les informations de l'utilisateur dans le contexte
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// RequireAdmin middleware pour vérifier les droits administrateur
func (am *AuthMiddleware) RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, models.ErrorResponse{
				Error:   "Forbidden",
				Message: "Droits administrateur requis",
				Code:    http.StatusForbidden,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// GenerateToken génère un token JWT
func (am *AuthMiddleware) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * time.Duration(am.config.JWT.TokenExpirationHours)).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(am.config.JWT.Secret))
}

// GenerateRefreshToken génère un refresh token
func (am *AuthMiddleware) GenerateRefreshToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * time.Duration(am.config.JWT.RefreshExpirationDays)).Unix(),
		"iat":      time.Now().Unix(),
		"type":     "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(am.config.JWT.Secret))
}

// VerifyToken vérifie et parse un token JWT
func (am *AuthMiddleware) verifyToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(am.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Vérifier l'expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, jwt.ErrTokenExpired
		}
	}

	// Extraire les informations utilisateur
	userClaims := &models.Claims{
		UserID:   uint(claims["user_id"].(float64)),
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
		Email:    claims["email"].(string),
	}

	return userClaims, nil
}

// VerifyRefreshToken vérifie un refresh token
func (am *AuthMiddleware) VerifyRefreshToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(am.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Vérifier que c'est bien un refresh token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Vérifier l'expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, jwt.ErrTokenExpired
		}
	}

	// Extraire les informations utilisateur
	userClaims := &models.Claims{
		UserID:   uint(claims["user_id"].(float64)),
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
		Email:    claims["email"].(string),
	}

	return userClaims, nil
}