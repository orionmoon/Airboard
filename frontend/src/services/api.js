import axios from 'axios'

// Configuration de base d'Axios
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  }
})

// Logs pour le développement
if (import.meta.env.DEV) {
  api.interceptors.request.use(
    (config) => {
      console.log('🚀', config.method?.toUpperCase(), config.url, config.data)
      return config
    },
    (error) => {
      console.error('❌', 'Request Error:', error)
      return Promise.reject(error)
    }
  )

  api.interceptors.response.use(
    (response) => {
      console.log('✅', response.config.method?.toUpperCase(), response.config.url, response.data)
      return response
    },
    (error) => {
      console.error('❌', error.config?.method?.toUpperCase(), error.config?.url, error.response?.data || error.message)
      return Promise.reject(error)
    }
  )
}

// Configuration des intercepteurs avec authentification
export function setupInterceptors(router) {
  // Intercepteur de requête pour ajouter le token
  api.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem('airboard_token')
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      return config
    },
    (error) => Promise.reject(error)
  )

  // Intercepteur de réponse pour gérer l'authentification
  api.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config

      // Si erreur 401 et pas déjà une tentative de refresh
      if (error.response?.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true

        const refreshToken = localStorage.getItem('airboard_refresh_token')
        if (refreshToken) {
          try {
            const response = await api.post('/auth/refresh', {
              refresh_token: refreshToken
            })

            const { token, refresh_token } = response.data
            localStorage.setItem('airboard_token', token)
            localStorage.setItem('airboard_refresh_token', refresh_token)

            // Refaire la requête originale avec le nouveau token
            originalRequest.headers.Authorization = `Bearer ${token}`
            return api(originalRequest)
          } catch (refreshError) {
            // Refresh failed, redirect to login
            localStorage.removeItem('airboard_token')
            localStorage.removeItem('airboard_refresh_token')
            localStorage.removeItem('airboard_user')
            router.push('/auth/login')
            return Promise.reject(refreshError)
          }
        } else {
          // No refresh token, redirect to login
          router.push('/auth/login')
        }
      }

      return Promise.reject(error)
    }
  )
}

// Services API

// Auth Service
export const authService = {
  async login(credentials) {
    const response = await api.post('/auth/login', credentials)
    return response.data
  },

  async register(userData) {
    const response = await api.post('/auth/register', userData)
    return response.data
  },

  async refreshToken(refreshToken) {
    const response = await api.post('/auth/refresh', { refresh_token: refreshToken })
    return response.data
  },

  async getProfile() {
    const response = await api.get('/auth/profile')
    return response.data
  },

  async changePassword(oldPassword, newPassword) {
    const response = await api.post('/auth/change-password', {
      old_password: oldPassword,
      new_password: newPassword
    })
    return response.data
  },

  logout() {
    localStorage.removeItem('airboard_token')
    localStorage.removeItem('airboard_refresh_token')
    localStorage.removeItem('airboard_user')
  },

  async ssoAutoLogin() {
    const response = await api.get('/auth/sso/auto-login')
    return response.data
  }
}

// Dashboard Service
export const dashboardService = {
  async getDashboard() {
    const response = await api.get('/dashboard')
    return response.data
  }
}

// Admin Services
export const adminService = {
  // App Groups
  async getAppGroups(params = {}) {
    const response = await api.get('/admin/app-groups', { params })
    return response.data.data || response.data
  },

  async createAppGroup(data) {
    const response = await api.post('/admin/app-groups', data)
    return response.data
  },

  async updateAppGroup(id, data) {
    const response = await api.put(`/admin/app-groups/${id}`, data)
    return response.data
  },

  async deleteAppGroup(id) {
    const response = await api.delete(`/admin/app-groups/${id}`)
    return response.data
  },

  // Applications
  async getApplications(params = {}) {
    const response = await api.get('/admin/applications', { params })
    return response.data.data || response.data
  },

  async createApplication(data) {
    const response = await api.post('/admin/applications', data)
    return response.data
  },

  async updateApplication(id, data) {
    const response = await api.put(`/admin/applications/${id}`, data)
    return response.data
  },

  async deleteApplication(id) {
    const response = await api.delete(`/admin/applications/${id}`)
    return response.data
  },

  // Users
  async getUsers() {
    const response = await api.get('/admin/users')
    return response.data.data || response.data
  },

  async createUser(data) {
    const response = await api.post('/admin/users', data)
    return response.data
  },

  async updateUser(id, data) {
    const response = await api.put(`/admin/users/${id}`, data)
    return response.data
  },

  async deleteUser(id) {
    const response = await api.delete(`/admin/users/${id}`)
    return response.data
  },

  // Groups
  async getGroups() {
    const response = await api.get('/admin/groups')
    return response.data.data || response.data
  },

  async createGroup(data) {
    const response = await api.post('/admin/groups', data)
    return response.data
  },

  async updateGroup(id, data) {
    const response = await api.put(`/admin/groups/${id}`, data)
    return response.data
  },

  async deleteGroup(id) {
    const response = await api.delete(`/admin/groups/${id}`)
    return response.data
  },

  // App Settings
  async getAppSettings() {
    const response = await api.get('/admin/settings')
    return response.data.data || response.data
  },

  async updateAppSettings(data) {
    const response = await api.put('/admin/settings', data)
    return response.data.data || response.data
  },

  async resetAppSettings() {
    const response = await api.post('/admin/settings/reset')
    return response.data.data || response.data
  },

  // Database
  async resetDatabase() {
    const response = await api.post('/admin/database/reset')
    return response.data
  },

  // OAuth Providers
  async getOAuthProviders() {
    const response = await api.get('/admin/oauth/providers')
    return response.data
  },

  async updateOAuthProvider(id, data) {
    const response = await api.put(`/admin/oauth/providers/${id}`, data)
    return response.data
  },

  // Deleted Users Management
  async getDeletedUsers() {
    const response = await api.get('/admin/users/deleted')
    return response.data
  },

  async restoreUser(id) {
    const response = await api.post(`/admin/users/${id}/restore`)
    return response.data
  },

  async permanentlyDeleteUser(id) {
    const response = await api.delete(`/admin/users/${id}/permanent`)
    return response.data
  }
}

// OAuth Service (public)
export const oauthService = {
  async getEnabledProviders() {
    const response = await api.get('/auth/oauth/providers')
    return response.data
  },

  async initiateOAuth(providerName) {
    const response = await api.get(`/auth/oauth/${providerName}/initiate`)
    return response.data
  },

  async handleCallback(providerName, code, state) {
    const response = await api.get(`/auth/oauth/${providerName}/callback`, {
      params: { code, state }
    })
    return response.data
  }
}

// Favorites Service
export const favoritesService = {
  async getFavorites() {
    const response = await api.get('/user/favorites')
    return response.data
  },

  async addFavorite(applicationId) {
    const response = await api.post('/user/favorites', { application_id: applicationId })
    return response.data
  },

  async removeFavorite(applicationId) {
    const response = await api.delete(`/user/favorites/${applicationId}`)
    return response.data
  },

  async isFavorite(applicationId) {
    const response = await api.get(`/user/favorites/${applicationId}/check`)
    return response.data
  }
}

// Analytics Service
export const analyticsService = {
  async trackClick(applicationId) {
    try {
      const response = await api.post('/analytics/track', { application_id: applicationId })
      return response.data
    } catch (error) {
      // Ne pas bloquer l'application si le tracking échoue
      console.error('Error tracking click:', error)
      return null
    }
  },

  async getDashboard() {
    const response = await api.get('/admin/analytics/dashboard')
    return response.data
  },

  async getApplicationStats(applicationId) {
    const response = await api.get(`/admin/analytics/applications/${applicationId}`)
    return response.data
  },

  async getUserStats(userId) {
    const response = await api.get(`/admin/analytics/users/${userId}`)
    return response.data
  }
}

// Announcements Service
export const announcementsService = {
  // Public - Get active announcements for users
  async getActiveAnnouncements() {
    const response = await api.get('/announcements')
    return response.data
  },

  // Admin - Get all announcements
  async getAllAnnouncements() {
    const response = await api.get('/admin/announcements')
    return response.data
  },

  // Admin - Get single announcement
  async getAnnouncement(id) {
    const response = await api.get(`/admin/announcements/${id}`)
    return response.data
  },

  // Admin - Create announcement
  async createAnnouncement(data) {
    const response = await api.post('/admin/announcements', data)
    return response.data
  },

  // Admin - Update announcement
  async updateAnnouncement(id, data) {
    const response = await api.put(`/admin/announcements/${id}`, data)
    return response.data
  },

  // Admin - Delete announcement
  async deleteAnnouncement(id) {
    const response = await api.delete(`/admin/announcements/${id}`)
    return response.data
  }
}

// News Hub Service
export const newsService = {
  // User - Get news with filters and pagination
  async getNews(params = {}) {
    const response = await api.get('/news', { params })
    return response.data
  },

  // User - Get news by slug
  async getNewsBySlug(slug) {
    const response = await api.get(`/news/article/${slug}`)
    return response.data
  },

  // User - Increment view count
  async incrementView(id) {
    const response = await api.post(`/news/${id}/view`)
    return response.data
  },

  // User - Get reactions for a news
  async getReactions(id) {
    const response = await api.get(`/news/${id}/reactions`)
    return response.data
  },

  // User - Add reaction
  async addReaction(id, reactionType) {
    const response = await api.post(`/news/${id}/react`, { reaction_type: reactionType })
    return response.data
  },

  // User - Remove reaction
  async removeReaction(id) {
    const response = await api.delete(`/news/${id}/react`)
    return response.data
  },

  // User - Get unread count
  async getUnreadCount() {
    const response = await api.get('/news/unread/count')
    return response.data
  },

  // User - Get categories
  async getCategories() {
    const response = await api.get('/news/categories')
    return response.data
  },

  // User - Get tags
  async getTags() {
    const response = await api.get('/news/tags')
    return response.data
  },

  // User - Get unread count
  async getUnreadCount() {
    const response = await api.get('/news/unread/count')
    return response.data
  },

  // Editor - Create news
  async createNews(data) {
    const response = await api.post('/editor/news', data)
    return response.data
  },

  // Editor - Update news
  async updateNews(id, data) {
    const response = await api.put(`/editor/news/${id}`, data)
    return response.data
  },

  // Editor - Delete news
  async deleteNews(id) {
    const response = await api.delete(`/editor/news/${id}`)
    return response.data
  },

  // Editor - Create tag
  async createTag(data) {
    const response = await api.post('/editor/news/tags', data)
    return response.data
  },

  // Editor - Update tag
  async updateTag(id, data) {
    const response = await api.put(`/editor/news/tags/${id}`, data)
    return response.data
  },

  // Editor - Delete tag
  async deleteTag(id) {
    const response = await api.delete(`/editor/news/tags/${id}`)
    return response.data
  },

  // Admin - Create category
  async createCategory(data) {
    const response = await api.post('/admin/news/categories', data)
    return response.data
  },

  // Admin - Update category
  async updateCategory(id, data) {
    const response = await api.put(`/admin/news/categories/${id}`, data)
    return response.data
  },

  // Admin - Delete category
  async deleteCategory(id) {
    const response = await api.delete(`/admin/news/categories/${id}`)
    return response.data
  },

  // Admin - Toggle pin
  async togglePin(id) {
    const response = await api.post(`/admin/news/${id}/pin`)
    return response.data
  },

  // Admin - Get analytics
  async getAnalytics() {
    const response = await api.get('/admin/news/analytics')
    return response.data
  }
}

export default api