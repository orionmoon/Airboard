package models

import (
	"time"

	"gorm.io/gorm"
)

// User représente un utilisateur du système
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Username    string         `json:"username" gorm:"unique;not null"`
	Email       string         `json:"email" gorm:"unique;not null"`
	Password    string         `json:"-"` // Nullable pour les users SSO
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Role        string         `json:"role" gorm:"default:'user'"` // admin, user
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	SSOProvider string         `json:"sso_provider,omitempty"` // authentik, azure, etc.
	SSOID       string         `json:"sso_id,omitempty"`       // ID utilisateur externe
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Groups []Group `json:"groups,omitempty" gorm:"many2many:user_groups;"`
}

// Group représente un groupe d'utilisateurs
type Group struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	Color       string         `json:"color" gorm:"default:'#3B82F6'"` // Couleur pour l'affichage
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Users     []User     `json:"users,omitempty" gorm:"many2many:user_groups;"`
	AppGroups []AppGroup `json:"app_groups,omitempty" gorm:"many2many:group_app_groups;"`
}

// AppGroup représente un groupe d'applications
type AppGroup struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	Color       string         `json:"color" gorm:"default:'#10B981'"`   // Couleur pour l'affichage
	Icon        string         `json:"icon" gorm:"default:'mdi:folder'"` // Icône Iconify
	Order       int            `json:"order" gorm:"default:0"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Applications []Application `json:"applications,omitempty"`
	Groups       []Group       `json:"groups,omitempty" gorm:"many2many:group_app_groups;"`
}

// Application représente une application dans le portail
type Application struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Description  string         `json:"description"`
	URL          string         `json:"url" gorm:"not null"`
	Icon         string         `json:"icon" gorm:"default:'mdi:application'"` // Icône Iconify
	Color        string         `json:"color" gorm:"default:'#6366F1'"`        // Couleur pour l'affichage
	Order        int            `json:"order" gorm:"default:0"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	OpenInNewTab bool           `json:"open_in_new_tab" gorm:"default:true"`
	AppGroupID   uint           `json:"app_group_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	AppGroup *AppGroup `json:"app_group,omitempty"`
}

// JWT Claims structure
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

// Request/Response structures
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// Dashboard response
type DashboardResponse struct {
	AppGroups []AppGroupWithApps `json:"app_groups"`
	Stats     DashboardStats     `json:"stats"`
}

type AppGroupWithApps struct {
	AppGroup
	Applications []Application `json:"applications"`
}

type DashboardStats struct {
	TotalAppGroups    int64 `json:"total_app_groups"`
	TotalApplications int64 `json:"total_applications"`
	TotalUsers        int64 `json:"total_users"`
	TotalGroups       int64 `json:"total_groups"`
}

// AppSettings représente les paramètres de configuration de l'application
type AppSettings struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	AppName        string    `json:"app_name" gorm:"default:'Airboard'"`
	AppIcon        string    `json:"app_icon" gorm:"default:'mdi:view-dashboard'"`
	DashboardTitle string    `json:"dashboard_title" gorm:"default:'Dashboard'"`
	WelcomeMessage string    `json:"welcome_message" gorm:"default:'Welcome to your application portal'"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// AppSettingsRequest pour les requêtes de mise à jour
type AppSettingsRequest struct {
	AppName        string `json:"app_name" binding:"required,min=1"`
	AppIcon        string `json:"app_icon" binding:"required"`
	DashboardTitle string `json:"dashboard_title" binding:"required,min=1"`
	WelcomeMessage string `json:"welcome_message" binding:"required,min=1"`
}

// ChangePasswordRequest pour les changements de mot de passe
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// OAuthProvider représente un fournisseur OAuth (Google, Microsoft, etc.)
type OAuthProvider struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	ProviderName string    `json:"provider_name" gorm:"unique;not null"` // google, microsoft
	DisplayName  string    `json:"display_name" gorm:"not null"`         // "Google", "Microsoft"
	Icon         string    `json:"icon" gorm:"default:'mdi:login'"`      // Icône Iconify
	IsEnabled    bool      `json:"is_enabled" gorm:"default:false"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"-"` // Ne jamais exposer dans le JSON
	RedirectURI  string    `json:"redirect_uri"`
	AuthURL      string    `json:"auth_url"`      // URL d'autorisation OAuth
	TokenURL     string    `json:"token_url"`     // URL d'échange de token
	UserInfoURL  string    `json:"user_info_url"` // URL pour récupérer les infos utilisateur
	Scopes       string    `json:"scopes"`        // Scopes OAuth séparés par des espaces
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// OAuthProviderRequest pour les requêtes de mise à jour
type OAuthProviderRequest struct {
	ProviderName string `json:"provider_name" binding:"required"`
	DisplayName  string `json:"display_name" binding:"required"`
	Icon         string `json:"icon"`
	IsEnabled    bool   `json:"is_enabled"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	AuthURL      string `json:"auth_url"`
	TokenURL     string `json:"token_url"`
	UserInfoURL  string `json:"user_info_url"`
	Scopes       string `json:"scopes"`
}

// OAuthProviderPublic pour l'affichage public (sans secrets)
type OAuthProviderPublic struct {
	ID           uint   `json:"id"`
	ProviderName string `json:"provider_name"`
	DisplayName  string `json:"display_name"`
	Icon         string `json:"icon"`
	IsEnabled    bool   `json:"is_enabled"`
	RedirectURI  string `json:"redirect_uri"`
}
