package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"airboard/models"
)

// GetCategories - Liste toutes les catégories (accessible à tous)
func (h *NewsHandler) GetCategories(c *gin.Context) {
	var categories []models.NewsCategory

	query := h.db.Model(&models.NewsCategory{})

	// Filtre actif/inactif
	if active := c.Query("active"); active != "" {
		query = query.Where("is_active = ?", active == "true")
	}

	if err := query.Order("\"order\" ASC, name ASC").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategory - Récupérer une catégorie par ID
func (h *NewsHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.NewsCategory
	if err := h.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

// CreateCategory - Créer une catégorie (admin uniquement)
func (h *NewsHandler) CreateCategory(c *gin.Context) {
	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.NewsCategory{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
		Color:       req.Color,
		Order:       req.Order,
		IsActive:    req.IsActive,
	}

	if err := h.db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory - Modifier une catégorie (admin uniquement)
func (h *NewsHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.NewsCategory
	if err := h.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		}
		return
	}

	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = req.Name
	category.Description = req.Description
	category.Icon = req.Icon
	category.Color = req.Color
	category.Order = req.Order
	category.IsActive = req.IsActive

	if err := h.db.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory - Supprimer une catégorie (admin uniquement)
func (h *NewsHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.NewsCategory
	if err := h.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		}
		return
	}

	// Vérifier si des news utilisent cette catégorie
	var newsCount int64
	h.db.Model(&models.News{}).Where("category_id = ?", id).Count(&newsCount)
	if newsCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Cannot delete category with associated news",
			"news_count": newsCount,
		})
		return
	}

	if err := h.db.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// GetTags - Liste tous les tags (accessible à tous)
func (h *NewsHandler) GetTags(c *gin.Context) {
	var tags []models.Tag

	if err := h.db.Order("name ASC").Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetTag - Récupérer un tag par ID
func (h *NewsHandler) GetTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := h.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tag"})
		}
		return
	}

	c.JSON(http.StatusOK, tag)
}

// CreateTag - Créer un tag (admin/editor)
func (h *NewsHandler) CreateTag(c *gin.Context) {
	var req models.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{
		Name:  req.Name,
		Color: req.Color,
	}

	if err := h.db.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// UpdateTag - Modifier un tag (admin/editor)
func (h *NewsHandler) UpdateTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := h.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tag"})
		}
		return
	}

	var req models.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.Name = req.Name
	tag.Color = req.Color

	if err := h.db.Save(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// DeleteTag - Supprimer un tag (admin uniquement)
func (h *NewsHandler) DeleteTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := h.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tag"})
		}
		return
	}

	// Supprimer les associations avant de supprimer le tag
	h.db.Exec("DELETE FROM news_tags WHERE tag_id = ?", id)

	if err := h.db.Delete(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
