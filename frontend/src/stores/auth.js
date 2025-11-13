import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authService } from '@/services/api'

export const useAuthStore = defineStore('auth', () => {
  // État
  const user = ref(null)
  const token = ref(null)
  const refreshToken = ref(null)
  const isLoading = ref(false)

  // Getters
  const isAuthenticated = computed(() => !!token.value && !!user.value)
  const isAdmin = computed(() => user.value?.role === 'admin')
  const userInitials = computed(() => {
    if (!user.value) return 'U'
    const firstName = user.value.first_name || user.value.username
    const lastName = user.value.last_name || ''
    
    if (firstName && lastName) {
      return `${firstName.charAt(0)}${lastName.charAt(0)}`.toUpperCase()
    }
    return firstName.charAt(0).toUpperCase()
  })

  const userDisplayName = computed(() => {
    if (!user.value) return 'Utilisateur'
    
    if (user.value.first_name && user.value.last_name) {
      return `${user.value.first_name} ${user.value.last_name}`
    }
    
    return user.value.username
  })

  // Actions
  const login = async (credentials) => {
    try {
      isLoading.value = true
      const response = await authService.login(credentials)
      
      // Stocker les données dans l'ordre correct
      token.value = response.token
      refreshToken.value = response.refresh_token
      user.value = response.user

      // Persistance locale
      localStorage.setItem('airboard_token', response.token)
      localStorage.setItem('airboard_refresh_token', response.refresh_token)
      localStorage.setItem('airboard_user', JSON.stringify(response.user))

      // Force reactive update
      console.log('🔐 Login success:', { 
        isAuthenticated: isAuthenticated.value, 
        user: user.value?.username,
        token: !!token.value 
      })

      return response
    } catch (error) {
      console.error('Erreur de connexion:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  const register = async (userData) => {
    try {
      isLoading.value = true
      const response = await authService.register(userData)
      
      // Stocker les données
      user.value = response.user
      token.value = response.token
      refreshToken.value = response.refresh_token

      // Persistance locale
      localStorage.setItem('airboard_token', response.token)
      localStorage.setItem('airboard_refresh_token', response.refresh_token)
      localStorage.setItem('airboard_user', JSON.stringify(response.user))

      return response
    } catch (error) {
      console.error('Erreur d\'inscription:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  const logout = () => {
    // Nettoyer le store
    user.value = null
    token.value = null
    refreshToken.value = null

    // Nettoyer le localStorage
    authService.logout()
  }

  const loadFromStorage = () => {
    try {
      const storedToken = localStorage.getItem('airboard_token')
      const storedRefreshToken = localStorage.getItem('airboard_refresh_token')
      const storedUser = localStorage.getItem('airboard_user')

      if (storedToken && storedUser) {
        token.value = storedToken
        refreshToken.value = storedRefreshToken
        user.value = JSON.parse(storedUser)
        return true
      }
    } catch (error) {
      console.error('Erreur lors du chargement depuis le storage:', error)
      logout()
    }
    return false
  }

  const updateProfile = async () => {
    try {
      const profile = await authService.getProfile()
      user.value = profile
      localStorage.setItem('airboard_user', JSON.stringify(profile))
      return profile
    } catch (error) {
      console.error('Erreur lors de la mise à jour du profil:', error)
      throw error
    }
  }

  const setUser = (userData) => {
    user.value = userData
    localStorage.setItem('airboard_user', JSON.stringify(userData))
  }

  const updateTokens = (newToken, newRefreshToken) => {
    token.value = newToken
    refreshToken.value = newRefreshToken
    localStorage.setItem('airboard_token', newToken)
    localStorage.setItem('airboard_refresh_token', newRefreshToken)
  }

  const autoLoginSSO = async () => {
    try {
      isLoading.value = true

      // Appeler la route SSO auto-login
      const response = await authService.ssoAutoLogin()

      // Stocker les données
      token.value = response.token
      refreshToken.value = response.refresh_token
      user.value = response.user

      // Persistance locale
      localStorage.setItem('airboard_token', response.token)
      localStorage.setItem('airboard_refresh_token', response.refresh_token)
      localStorage.setItem('airboard_user', JSON.stringify(response.user))

      console.log('🔐 SSO Auto-login success:', {
        user: user.value?.username,
        email: user.value?.email,
        provider: user.value?.sso_provider
      })

      return response
    } catch (error) {
      // Si SSO échoue, ne pas afficher d'erreur (mode classique)
      console.log('ℹ️ SSO non disponible, passage en mode classique')
      return null
    } finally {
      isLoading.value = false
    }
  }

  return {
    // État
    user,
    token,
    refreshToken,
    isLoading,

    // Getters
    isAuthenticated,
    isAdmin,
    userInitials,
    userDisplayName,

    // Actions
    login,
    register,
    logout,
    loadFromStorage,
    updateProfile,
    setUser,
    updateTokens,
    autoLoginSSO,
  }
})