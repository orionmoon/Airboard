package main

import (
	"log"

	"airboard/config"
	"airboard/handlers"
	"airboard/middleware"
	"airboard/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Charger la configuration
	cfg := config.LoadConfig()

	// Configuration Gin
	gin.SetMode(cfg.Server.Mode)

	// Connexion à la base de données
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Migrations
	if err := db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.AppGroup{},
		&models.Application{},
		&models.AppSettings{},
	); err != nil {
		log.Fatal("Erreur lors des migrations:", err)
	}

	// Créer les données initiales
	createInitialData(db)

	// Initialisation des middlewares
	authMiddleware := middleware.NewAuthMiddleware(cfg)

	// Initialisation des handlers
	authHandler := handlers.NewAuthHandler(db, authMiddleware)
	dashboardHandler := handlers.NewDashboardHandler(db)
	adminHandler := handlers.NewAdminHandler(db)
	settingsHandler := handlers.NewSettingsHandler(db)

	// Configuration du routeur
	router := gin.Default()

	// Middleware CORS
	router.Use(middleware.SetupCORS(cfg))

	// Routes publiques
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/refresh", authHandler.RefreshToken)
		}
	}

	// Routes protégées
	protected := api.Group("/")
	protected.Use(authMiddleware.RequireAuth())
	{
		// Profil utilisateur
		protected.GET("/auth/profile", authHandler.GetProfile)
		protected.POST("/auth/change-password", authHandler.ChangePassword)

		// Dashboard
		protected.GET("/dashboard", dashboardHandler.GetDashboard)

		// Routes admin
		admin := protected.Group("/admin")
		admin.Use(authMiddleware.RequireAdmin())
		{
			// Gestion des groupes d'applications
			admin.GET("/app-groups", adminHandler.GetAppGroups)
			admin.POST("/app-groups", adminHandler.CreateAppGroup)
			admin.PUT("/app-groups/:id", adminHandler.UpdateAppGroup)
			admin.DELETE("/app-groups/:id", adminHandler.DeleteAppGroup)

			// Gestion des applications
			admin.GET("/applications", adminHandler.GetApplications)
			admin.POST("/applications", adminHandler.CreateApplication)
			admin.PUT("/applications/:id", adminHandler.UpdateApplication)
			admin.DELETE("/applications/:id", adminHandler.DeleteApplication)

			// Gestion des utilisateurs
			admin.GET("/users", adminHandler.GetUsers)
			admin.POST("/users", adminHandler.CreateUser)
			admin.PUT("/users/:id", adminHandler.UpdateUser)
			admin.DELETE("/users/:id", adminHandler.DeleteUser)

			// Gestion des groupes d'utilisateurs
			admin.GET("/groups", adminHandler.GetGroups)
			admin.POST("/groups", adminHandler.CreateGroup)
			admin.PUT("/groups/:id", adminHandler.UpdateGroup)
			admin.DELETE("/groups/:id", adminHandler.DeleteGroup)

			// Gestion des paramètres de l'application
			admin.GET("/settings", settingsHandler.GetAppSettings)
			admin.PUT("/settings", settingsHandler.UpdateAppSettings)
			admin.POST("/settings/reset", settingsHandler.ResetAppSettings)

			// Gestion de la base de données
			admin.POST("/database/reset", adminHandler.ResetDatabase)
		}
	}

	// Route de santé
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Airboard API is running",
		})
	})

	// Documentation Swagger (optionnel)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("🚀 Serveur Airboard démarré sur le port %s", cfg.Server.Port)
	log.Printf("📊 Dashboard: http://localhost:%s/health", cfg.Server.Port)
	log.Printf("📚 Mode: %s", cfg.Server.Mode)

	// Démarrer le serveur
	router.Run(":" + cfg.Server.Port)
}

func createInitialData(db *gorm.DB) {
	// Créer ou réinitialiser un utilisateur admin par défaut
	var adminUser models.User
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	if err := db.Where("username = ?", "admin").First(&adminUser).Error; err != nil {
		// Créer l'utilisateur admin
		admin := models.User{
			Username:  "admin",
			Email:     "admin@airboard.com",
			Password:  string(hashedPassword),
			FirstName: "Admin",
			LastName:  "Airboard",
			Role:      "admin",
			IsActive:  true,
		}
		db.Create(&admin)
		log.Println("✅ Utilisateur admin créé: admin@airboard.com / admin123")
	} else {
		// Réinitialiser le mot de passe si l'utilisateur existe déjà
		adminUser.Password = string(hashedPassword)
		adminUser.IsActive = true
		adminUser.Role = "admin"
		db.Save(&adminUser)
		log.Println("🔄 Mot de passe admin réinitialisé: admin@airboard.com / admin123")
	}

	// Créer un utilisateur normal par défaut
	var normalUser models.User
	if err := db.Where("username = ?", "user").First(&normalUser).Error; err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
		user := models.User{
			Username:  "user",
			Email:     "user@airboard.com",
			Password:  string(hashedPassword),
			FirstName: "User",
			LastName:  "Demo",
			Role:      "user",
			IsActive:  true,
		}
		db.Create(&user)
		log.Println("✅ Utilisateur demo créé: user@airboard.com / user123")
	}

	// Créer des groupes d'applications de démonstration
	var devGroup models.AppGroup
	if err := db.Where("name = ?", "Développement").First(&devGroup).Error; err != nil {
		devGroup = models.AppGroup{
			Name:        "Développement",
			Description: "Applications de développement",
			Color:       "#3B82F6",
			Icon:        "mdi:code-tags",
			Order:       1,
			IsActive:    true,
		}
		db.Create(&devGroup)

		// Applications de développement
		apps := []models.Application{
			{
				Name:         "GitLab",
				Description:  "Gestion de code source",
				URL:          "https://gitlab.com",
				Icon:         "mdi:gitlab",
				Color:        "#FC6D26",
				Order:        1,
				IsActive:     true,
				OpenInNewTab: true,
				AppGroupID:   devGroup.ID,
			},
			{
				Name:         "Jenkins",
				Description:  "Intégration continue",
				URL:          "https://jenkins.io",
				Icon:         "mdi:robot-industrial",
				Color:        "#D33833",
				Order:        2,
				IsActive:     true,
				OpenInNewTab: true,
				AppGroupID:   devGroup.ID,
			},
		}
		for _, app := range apps {
			db.Create(&app)
		}
		log.Println("✅ Groupe Développement créé avec applications de demo")
	}

	// Créer un groupe Production
	var prodGroup models.AppGroup
	if err := db.Where("name = ?", "Production").First(&prodGroup).Error; err != nil {
		prodGroup = models.AppGroup{
			Name:        "Production",
			Description: "Applications de production",
			Color:       "#10B981",
			Icon:        "mdi:server",
			Order:       2,
			IsActive:    true,
		}
		db.Create(&prodGroup)

		// Applications de production
		apps := []models.Application{
			{
				Name:         "Grafana",
				Description:  "Monitoring et métriques",
				URL:          "https://grafana.com",
				Icon:         "mdi:chart-line",
				Color:        "#F46800",
				Order:        1,
				IsActive:     true,
				OpenInNewTab: true,
				AppGroupID:   prodGroup.ID,
			},
			{
				Name:         "Prometheus",
				Description:  "Collecte de métriques",
				URL:          "https://prometheus.io",
				Icon:         "mdi:database-search",
				Color:        "#E6522C",
				Order:        2,
				IsActive:     true,
				OpenInNewTab: true,
				AppGroupID:   prodGroup.ID,
			},
		}
		for _, app := range apps {
			db.Create(&app)
		}
		log.Println("✅ Groupe Production créé avec applications de demo")
	}

	// Créer un groupe d'utilisateurs de démonstration
	var demoGroup models.Group
	if err := db.Where("name = ?", "Développeurs").First(&demoGroup).Error; err != nil {
		demoGroup = models.Group{
			Name:        "Développeurs",
			Description: "Équipe de développement",
			Color:       "#8B5CF6",
			IsActive:    true,
		}
		db.Create(&demoGroup)

		// Associer l'utilisateur normal au groupe
		if err := db.Where("username = ?", "user").First(&normalUser).Error; err == nil {
			db.Model(&demoGroup).Association("Users").Append(&normalUser)
		}

		// Associer le groupe aux groupes d'applications
		db.Model(&demoGroup).Association("AppGroups").Append(&devGroup)
		db.Model(&demoGroup).Association("AppGroups").Append(&prodGroup)

		log.Println("✅ Groupe d'utilisateurs Développeurs créé avec permissions")
	}
}
