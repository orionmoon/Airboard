import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { favoritesService } from '@/services/api'

export const useFavoritesStore = defineStore('favorites', () => {
  // State
  const favorites = ref(new Set())
  const isLoading = ref(false)
  const error = ref(null)

  // Getters
  const favoritesList = computed(() => Array.from(favorites.value))
  const favoritesCount = computed(() => favorites.value.size)

  // Actions
  const loadFavorites = async () => {
    try {
      isLoading.value = true
      error.value = null

      // Charger depuis l'API
      const data = await favoritesService.getFavorites()

      // Extraire les IDs des applications favorites
      const favoriteIds = data.map(app => app.id)
      favorites.value = new Set(favoriteIds)

      // Sauvegarder dans localStorage
      saveFavoritesToLocalStorage()
    } catch (err) {
      console.error('Error loading favorites:', err)
      error.value = err.message

      // En cas d'erreur, charger depuis localStorage
      loadFavoritesFromLocalStorage()
    } finally {
      isLoading.value = false
    }
  }

  const addFavorite = async (applicationId) => {
    try {
      isLoading.value = true
      error.value = null

      // Ajouter via l'API
      await favoritesService.addFavorite(applicationId)

      // Ajouter localement (créer un nouveau Set pour la réactivité)
      favorites.value = new Set([...favorites.value, applicationId])

      // Sauvegarder dans localStorage
      saveFavoritesToLocalStorage()

      return true
    } catch (err) {
      console.error('Error adding favorite:', err)
      error.value = err.message
      return false
    } finally {
      isLoading.value = false
    }
  }

  const removeFavorite = async (applicationId) => {
    try {
      isLoading.value = true
      error.value = null

      // Supprimer via l'API
      await favoritesService.removeFavorite(applicationId)

      // Supprimer localement (créer un nouveau Set pour la réactivité)
      const newFavorites = new Set(favorites.value)
      newFavorites.delete(applicationId)
      favorites.value = newFavorites

      // Sauvegarder dans localStorage
      saveFavoritesToLocalStorage()

      return true
    } catch (err) {
      console.error('Error removing favorite:', err)
      error.value = err.message
      return false
    } finally {
      isLoading.value = false
    }
  }

  const toggleFavorite = async (applicationId) => {
    if (isFavorite(applicationId)) {
      return await removeFavorite(applicationId)
    } else {
      return await addFavorite(applicationId)
    }
  }

  const isFavorite = (applicationId) => {
    return favorites.value.has(applicationId)
  }

  // LocalStorage helpers
  const saveFavoritesToLocalStorage = () => {
    try {
      const favoritesArray = Array.from(favorites.value)
      localStorage.setItem('airboard_favorite_apps', JSON.stringify(favoritesArray))
    } catch (err) {
      console.error('Error saving favorites to localStorage:', err)
    }
  }

  const loadFavoritesFromLocalStorage = () => {
    try {
      const saved = localStorage.getItem('airboard_favorite_apps')
      if (saved) {
        const favoritesArray = JSON.parse(saved)
        favorites.value = new Set(favoritesArray)
      }
    } catch (err) {
      console.error('Error loading favorites from localStorage:', err)
      favorites.value = new Set()
    }
  }

  const clearFavorites = () => {
    favorites.value = new Set()
    localStorage.removeItem('airboard_favorite_apps')
  }

  // Initialiser avec les données du localStorage au chargement
  loadFavoritesFromLocalStorage()

  return {
    // State
    favorites,
    isLoading,
    error,

    // Getters
    favoritesList,
    favoritesCount,

    // Actions
    loadFavorites,
    addFavorite,
    removeFavorite,
    toggleFavorite,
    isFavorite,
    clearFavorites
  }
})
