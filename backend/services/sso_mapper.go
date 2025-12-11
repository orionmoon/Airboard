package services

import (
	"airboard/config"
	"airboard/models"
	"log"
	"strings"

	"gorm.io/gorm"
)

// SSOMapper gère le mapping et la synchronisation des utilisateurs SSO
type SSOMapper struct {
	db     *gorm.DB
	config *config.Config
}

// NewSSOMapper crée une nouvelle instance de SSOMapper
func NewSSOMapper(db *gorm.DB, cfg *config.Config) *SSOMapper {
	return &SSOMapper{
		db:     db,
		config: cfg,
	}
}

// SSOUserInfo contient les informations d'un utilisateur provenant du SSO
type SSOUserInfo struct {
	Email     string
	Username  string
	FirstName string
	LastName  string
	Groups    []string
	SSOID     string
}

// SyncUser crée ou met à jour un utilisateur à partir des informations SSO
func (m *SSOMapper) SyncUser(info *SSOUserInfo) (*models.User, error) {
	var user models.User

	// Chercher l'utilisateur par email
	result := m.db.Preload("Groups").Preload("AdminOfGroups").Where("email = ?", info.Email).First(&user)

	isNewUser := result.Error == gorm.ErrRecordNotFound

	if isNewUser {
		// Créer un nouvel utilisateur
		log.Printf("[SSO] Création d'un nouvel utilisateur: %s (%s)", info.Username, info.Email)

		user = models.User{
			Username:    info.Username,
			Email:       info.Email,
			FirstName:   info.FirstName,
			LastName:    info.LastName,
			Password:    "", // Pas de mot de passe pour SSO
			Role:        m.determineRole(info.Groups),
			IsActive:    true,
			SSOProvider: "authentik",
			SSOID:       info.SSOID,
		}

		if err := m.db.Create(&user).Error; err != nil {
			log.Printf("[SSO] Erreur lors de la création de l'utilisateur: %v", err)
			return nil, err
		}
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		// Mettre à jour l'utilisateur existant
		log.Printf("[SSO] Mise à jour de l'utilisateur: %s (%s)", info.Username, info.Email)

		user.FirstName = info.FirstName
		user.LastName = info.LastName
		user.Role = m.determineRole(info.Groups)
		user.SSOProvider = "authentik"
		user.SSOID = info.SSOID
		user.IsActive = true

		if err := m.db.Save(&user).Error; err != nil {
			log.Printf("[SSO] Erreur lors de la mise à jour de l'utilisateur: %v", err)
			return nil, err
		}
	}

	// Synchroniser les groupes
	if err := m.syncGroups(&user, info.Groups); err != nil {
		log.Printf("[SSO] Erreur lors de la synchronisation des groupes: %v", err)
		return nil, err
	}

	// Recharger l'utilisateur avec les groupes et les groupes administrés
	if err := m.db.Preload("Groups").Preload("AdminOfGroups").First(&user, user.ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// determineRole détermine le rôle de l'utilisateur basé sur ses groupes Authentik
func (m *SSOMapper) determineRole(authentikGroups []string) string {
	// Vérifier si l'utilisateur appartient à un groupe admin
	for _, authentikGroup := range authentikGroups {
		for _, adminGroup := range m.config.SSO.AdminGroups {
			if strings.EqualFold(authentikGroup, adminGroup) {
				log.Printf("[SSO] Utilisateur attribué au rôle admin (groupe: %s)", authentikGroup)
				return "admin"
			}
		}
	}

	// Par défaut, rôle user
	return m.config.SSO.DefaultRole
}

// syncGroups synchronise les groupes de l'utilisateur avec ceux d'Authentik
func (m *SSOMapper) syncGroups(user *models.User, authentikGroups []string) error {
	var airboardGroups []models.Group

	// Mapper les groupes Authentik vers les groupes Airboard
	for _, authentikGroup := range authentikGroups {
		groupName := m.mapGroupName(authentikGroup)

		var group models.Group
		result := m.db.Where("name = ?", groupName).First(&group)

		if result.Error == gorm.ErrRecordNotFound {
			// Créer le groupe s'il n'existe pas (en mode auto-provision)
			if m.config.SSO.AutoProvision {
				log.Printf("[SSO] Création du groupe: %s (depuis %s)", groupName, authentikGroup)
				group = models.Group{
					Name:        groupName,
					Description: "Synchronisé depuis Authentik",
					IsActive:    true,
				}
				if err := m.db.Create(&group).Error; err != nil {
					log.Printf("[SSO] Erreur lors de la création du groupe %s: %v", groupName, err)
					continue
				}
			} else {
				// Si auto-provision désactivé, on ajoute au groupe par défaut
				log.Printf("[SSO] Groupe %s non trouvé, utilisation du groupe par défaut", groupName)
				if err := m.db.Where("name = ?", m.config.SSO.DefaultGroup).First(&group).Error; err != nil {
					// Si le groupe par défaut n'existe pas, le créer
					group = models.Group{
						Name:        m.config.SSO.DefaultGroup,
						Description: "Groupe par défaut",
						IsActive:    true,
					}
					if err := m.db.Create(&group).Error; err != nil {
						log.Printf("[SSO] Erreur lors de la création du groupe par défaut: %v", err)
						continue
					}
				}
			}
		} else if result.Error != nil {
			log.Printf("[SSO] Erreur lors de la recherche du groupe %s: %v", groupName, result.Error)
			continue
		}

		airboardGroups = append(airboardGroups, group)
	}

	// Si aucun groupe n'a été mappé, ajouter au groupe par défaut
	if len(airboardGroups) == 0 {
		var defaultGroup models.Group
		if err := m.db.Where("name = ?", m.config.SSO.DefaultGroup).First(&defaultGroup).Error; err != nil {
			// Créer le groupe par défaut s'il n'existe pas
			defaultGroup = models.Group{
				Name:        m.config.SSO.DefaultGroup,
				Description: "Groupe par défaut pour les utilisateurs SSO",
				IsActive:    true,
			}
			if err := m.db.Create(&defaultGroup).Error; err != nil {
				log.Printf("[SSO] Erreur lors de la création du groupe par défaut: %v", err)
				return err
			}
		}
		airboardGroups = append(airboardGroups, defaultGroup)
	}

	// Remplacer les groupes de l'utilisateur
	if err := m.db.Model(user).Association("Groups").Replace(airboardGroups); err != nil {
		return err
	}

	log.Printf("[SSO] Groupes synchronisés pour %s: %d groupe(s)", user.Email, len(airboardGroups))
	return nil
}

// mapGroupName mappe un nom de groupe Authentik vers un nom de groupe Airboard
func (m *SSOMapper) mapGroupName(authentikGroup string) string {
	// Retirer les préfixes courants (ex: "airboard-", "app-")
	groupName := strings.TrimPrefix(authentikGroup, "airboard-")
	groupName = strings.TrimPrefix(groupName, "app-")

	// Capitaliser la première lettre
	if len(groupName) > 0 {
		groupName = strings.ToUpper(groupName[:1]) + groupName[1:]
	}

	return groupName
}
