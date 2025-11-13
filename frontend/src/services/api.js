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
  }
}

export default api