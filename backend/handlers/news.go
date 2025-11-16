package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"airboard/models"
)

type NewsHandler struct {
	db *gorm.DB
}

func NewNewsHandler(db *gorm.DB) *NewsHandler {
	return &NewsHandler{db: db}
}

// GetNews - Liste des news (accessible à tous les utilisateurs connectés)
func (h *NewsHandler) GetNews(c *gin.Context) {
	var news []models.News

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// Construction de la requête
	query := h.db.Model(&models.News{}).
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		Preload("Reactions")

	// Filtres
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if priority := c.Query("priority"); priority != "" {
		query = query.Where("priority = ?", priority)
	}

	if newsType := c.Query("type"); newsType != "" {
		query = query.Where("type = ?", newsType)
	}

	// Filtre par tags (supporte plusieurs tags séparés par des virgules)
	if tags := c.Query("tags"); tags != "" {
		tagIDs := strings.Split(tags, ",")
		if len(tagIDs) > 0 {
			query = query.Joins("JOIN news_tags ON news_tags.news_id = news.id").
				Where("news_tags.tag_id IN ?", tagIDs).
				Group("news.id").
				Having("COUNT(DISTINCT news_tags.tag_id) = ?", len(tagIDs))
		}
	} else if tagID := c.Query("tag_id"); tagID != "" {
		// Garde la compatibilité avec l'ancien paramètre tag_id
		query = query.Joins("JOIN news_tags ON news_tags.news_id = news.id").
			Where("news_tags.tag_id = ?", tagID)
	}

	// Recherche
	if search := c.Query("search"); search != "" {
		query = query.Where("title LIKE ? OR summary LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filtre published only (sauf pour admin/editor)
	userRole := c.GetString("user_role")
	if userRole != "admin" && userRole != "editor" {
		query = query.Where("is_published = ?", true).
			Where("(published_at IS NULL OR published_at <= ?)", time.Now())
	}

	// Gestion de la visibilité par groupes
	// TODO: Implémenter le filtrage par target_groups pour les users non-admin

	// Tri
	sortBy := c.DefaultQuery("sort", "published_at")
	sortOrder := c.DefaultQuery("order", "desc")
	if sortBy == "pinned" {
		query = query.Order("is_pinned DESC, published_at DESC")
	} else {
		query = query.Order(sortBy + " " + sortOrder)
	}

	// Compte total
	var total int64
	query.Count(&total)

	// Récupération avec pagination
	if err := query.Offset(offset).Limit(pageSize).Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}

	// Calcul du nombre de pages
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, models.NewsListResponse{
		News:       news,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}

// GetNewsBySlug - Récupérer une news par son slug
func (h *NewsHandler) GetNewsBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var news models.News
	if err := h.db.Preload("Author").
		Preload("Category").
		Preload("Tags").
		Preload("Reactions").
		Where("slug = ?", slug).
		First(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		}
		return
	}

	// Vérifier si publié (sauf admin/editor)
	userRole := c.GetString("user_role")
	if userRole != "admin" && userRole != "editor" {
		if !news.IsPublished {
			c.JSON(http.StatusForbidden, gin.H{"error": "News not published"})
			return
		}
	}

	c.JSON(http.StatusOK, news)
}

// CreateNews - Créer une news (admin/editor uniquement)
func (h *NewsHandler) CreateNews(c *gin.Context) {
	var req models.NewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Récupérer l'ID de l'utilisateur connecté
	userID := c.GetUint("user_id")

	news := models.News{
		Title:       req.Title,
		Summary:     req.Summary,
		Content:     req.Content,
		CoverImage:  req.CoverImage,
		Type:        req.Type,
		Priority:    req.Priority,
		IsPinned:    req.IsPinned,
		IsPublished: req.IsPublished,
		PublishedAt: req.PublishedAt,
		ExpiresAt:   req.ExpiresAt,
		CategoryID:  req.CategoryID,
		AuthorID:    userID,
	}

	// Si publié sans date, mettre la date actuelle
	if news.IsPublished && news.PublishedAt == nil {
		now := time.Now()
		news.PublishedAt = &now
	}

	// Créer la news
	if err := h.db.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
		return
	}

	// Associer les tags
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		h.db.Where("id IN ?", req.TagIDs).Find(&tags)
		h.db.Model(&news).Association("Tags").Replace(tags)
	}

	// Associer les groupes cibles
	if len(req.TargetGroupIDs) > 0 {
		var groups []models.Group
		h.db.Where("id IN ?", req.TargetGroupIDs).Find(&groups)
		h.db.Model(&news).Association("TargetGroups").Replace(groups)
	}

	// Recharger avec les relations
	h.db.Preload("Author").
		Preload("Category").
		Preload("Tags").
		Preload("TargetGroups").
		First(&news, news.ID)

	c.JSON(http.StatusCreated, news)
}

// UpdateNews - Modifier une news
func (h *NewsHandler) UpdateNews(c *gin.Context) {
	slug := c.Param("id") // Le paramètre est en fait un slug

	var news models.News
	if err := h.db.Where("slug = ?", slug).First(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		}
		return
	}

	// Vérifier les permissions (admin peut tout modifier, editor peut modifier ses propres news)
	userID := c.GetUint("user_id")
	userRole := c.GetString("role") // Corriger aussi la clé du contexte
	if userRole != "admin" && news.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own news"})
		return
	}

	var req models.NewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mise à jour des champs
	news.Title = req.Title
	news.Summary = req.Summary
	news.Content = req.Content
	news.CoverImage = req.CoverImage
	news.Type = req.Type
	news.Priority = req.Priority
	news.CategoryID = req.CategoryID
	news.ExpiresAt = req.ExpiresAt

	// Seul admin peut épingler
	if userRole == "admin" {
		news.IsPinned = req.IsPinned
	}

	// Gestion de la publication
	wasPublished := news.IsPublished
	news.IsPublished = req.IsPublished
	if news.IsPublished && !wasPublished && news.PublishedAt == nil {
		now := time.Now()
		news.PublishedAt = &now
	}
	if req.PublishedAt != nil {
		news.PublishedAt = req.PublishedAt
	}

	// Sauvegarder
	if err := h.db.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	// Mettre à jour les tags
	if req.TagIDs != nil {
		var tags []models.Tag
		h.db.Where("id IN ?", req.TagIDs).Find(&tags)
		h.db.Model(&news).Association("Tags").Replace(tags)
	}

	// Mettre à jour les groupes cibles
	if req.TargetGroupIDs != nil {
		var groups []models.Group
		h.db.Where("id IN ?", req.TargetGroupIDs).Find(&groups)
		h.db.Model(&news).Association("TargetGroups").Replace(groups)
	}

	// Recharger avec les relations
	h.db.Preload("Author").
		Preload("Category").
		Preload("Tags").
		Preload("TargetGroups").
		First(&news, news.ID)

	c.JSON(http.StatusOK, news)
}

// DeleteNews - Supprimer une news (soft delete)
func (h *NewsHandler) DeleteNews(c *gin.Context) {
	id := c.Param("id")

	var news models.News
	if err := h.db.First(&news, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		}
		return
	}

	// Vérifier les permissions
	userID := c.GetUint("user_id")
	userRole := c.GetString("user_role")
	if userRole != "admin" && news.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own news"})
		return
	}

	if err := h.db.Delete(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}

// TogglePin - Épingler/désépingler une news (admin uniquement)
func (h *NewsHandler) TogglePin(c *gin.Context) {
	id := c.Param("id")

	var news models.News
	if err := h.db.First(&news, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		}
		return
	}

	news.IsPinned = !news.IsPinned

	if err := h.db.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// IncrementView - Incrémenter le compteur de vues
func (h *NewsHandler) IncrementView(c *gin.Context) {
	id := c.Param("id")

	if err := h.db.Model(&models.News{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View counted"})
}

// AddReaction - Ajouter une réaction à une news
func (h *NewsHandler) AddReaction(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var req models.ReactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérifier si la news existe
	var news models.News
	if err := h.db.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	// Vérifier si l'utilisateur a déjà réagi
	var existingReaction models.NewsReaction
	err := h.db.Where("news_id = ? AND user_id = ?", id, userID).First(&existingReaction).Error

	if err == nil {
		// Déjà une réaction, la mettre à jour
		existingReaction.ReactionType = req.ReactionType
		if err := h.db.Save(&existingReaction).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reaction"})
			return
		}
		c.JSON(http.StatusOK, existingReaction)
		return
	}

	// Créer une nouvelle réaction
	newsID, _ := strconv.ParseUint(id, 10, 32)
	reaction := models.NewsReaction{
		NewsID:       uint(newsID),
		UserID:       userID,
		ReactionType: req.ReactionType,
	}

	if err := h.db.Create(&reaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add reaction"})
		return
	}

	c.JSON(http.StatusCreated, reaction)
}

// RemoveReaction - Retirer une réaction
func (h *NewsHandler) RemoveReaction(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	result := h.db.Where("news_id = ? AND user_id = ?", id, userID).Delete(&models.NewsReaction{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove reaction"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction removed"})
}

// GetReactions - Récupérer les réactions d'une news avec compteurs
func (h *NewsHandler) GetReactions(c *gin.Context) {
	id := c.Param("id")

	var reactions []models.NewsReaction
	if err := h.db.Where("news_id = ?", id).Find(&reactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reactions"})
		return
	}

	// Compter par type
	reactionCounts := make(map[string]int)
	for _, reaction := range reactions {
		reactionCounts[reaction.ReactionType]++
	}

	// Vérifier si l'utilisateur a réagi
	userID := c.GetUint("user_id")
	var userReaction *models.NewsReaction
	for _, reaction := range reactions {
		if reaction.UserID == userID {
			userReaction = &reaction
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"counts":        reactionCounts,
		"total":         len(reactions),
		"user_reaction": userReaction,
	})
}

// GetUnreadCount - Nombre de news non lues (pour le badge)
func (h *NewsHandler) GetUnreadCount(c *gin.Context) {
	// Compter les news publiées depuis la dernière visite
	// Pour MVP simple: toutes les news publiées dans les 30 derniers jours
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	var count int64
	h.db.Model(&models.News{}).
		Where("is_published = ?", true).
		Where("published_at >= ?", thirtyDaysAgo).
		Count(&count)

	// TODO: Implémenter le système de tracking de lecture (NewsRead)
	// Pour le moment, on retourne juste le compte des news récentes

	c.JSON(http.StatusOK, gin.H{
		"unread_count": count,
	})
}

// GetAnalytics - Statistiques du News Hub (admin/editor)
func (h *NewsHandler) GetAnalytics(c *gin.Context) {
	var stats models.NewsStatsResponse

	// Total news
	h.db.Model(&models.News{}).Count(&stats.TotalNews)

	// News publiées
	h.db.Model(&models.News{}).Where("is_published = ?", true).Count(&stats.PublishedNews)

	// News en brouillon
	h.db.Model(&models.News{}).Where("is_published = ?", false).Count(&stats.DraftNews)

	// Total vues
	h.db.Model(&models.News{}).Select("COALESCE(SUM(view_count), 0)").Scan(&stats.TotalViews)

	// Total réactions
	h.db.Model(&models.NewsReaction{}).Count(&stats.TotalReactions)

	// Top 5 news (par vues)
	var topNews []models.News
	h.db.Preload("Author").
		Preload("Category").
		Order("view_count DESC").
		Limit(5).
		Find(&topNews)

	// Convertir en NewsWithStats avec compteur de réactions
	stats.TopNews = make([]models.NewsWithStats, len(topNews))
	for i, news := range topNews {
		var reactionCount int64
		h.db.Model(&models.NewsReaction{}).Where("news_id = ?", news.ID).Count(&reactionCount)
		stats.TopNews[i] = models.NewsWithStats{
			News:          news,
			ReactionCount: reactionCount,
		}
	}

	// Réactions par type
	rows, err := h.db.Model(&models.NewsReaction{}).
		Select("reaction_type, COUNT(*) as count").
		Group("reaction_type").
		Rows()

	if err == nil {
		defer rows.Close()
		stats.ReactionsByType = make(map[string]int64)
		for rows.Next() {
			var reactionType string
			var count int64
			rows.Scan(&reactionType, &count)
			stats.ReactionsByType[reactionType] = count
		}
	}

	c.JSON(http.StatusOK, stats)
}
