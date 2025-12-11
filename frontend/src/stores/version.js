import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/services/api'

export const useVersionStore = defineStore('version', () => {
  // State
  const currentVersion = ref(null)
  const latestVersion = ref(null)
  const updateAvailable = ref(false)
  const releaseInfo = ref(null)
  const isCheckingUpdates = ref(false)
  const lastCheckTime = ref(null)

  // Getters
  const versionInfo = computed(() => ({
    version: currentVersion.value?.version || '1.0.0',
    buildDate: currentVersion.value?.buildDate || '',
    gitCommit: currentVersion.value?.gitCommit || '',
    changelog: currentVersion.value?.changelog || []
  }))

  const hasUpdate = computed(() => updateAvailable.value)

  const updateDetails = computed(() => {
    if (!releaseInfo.value) return null

    return {
      currentVersion: releaseInfo.value.currentVersion,
      latestVersion: releaseInfo.value.latestVersion,
      releaseDate: releaseInfo.value.releaseDate,
      releaseNotes: releaseInfo.value.releaseNotes,
      releaseURL: releaseInfo.value.releaseURL
    }
  })

  // Actions
  async function fetchVersion() {
    try {
      const response = await api.get('/version')
      currentVersion.value = response.data
      return response.data
    } catch (error) {
      console.error('Failed to fetch version:', error)
      // Fallback to default version
      currentVersion.value = {
        version: '1.0.0',
        buildDate: new Date().toISOString().split('T')[0],
        gitCommit: 'unknown',
        changelog: []
      }
      throw error
    }
  }

  async function checkForUpdates() {
    if (isCheckingUpdates.value) return

    isCheckingUpdates.value = true
    try {
      const response = await api.get('/version/check-updates')
      const data = response.data

      updateAvailable.value = data.updateAvailable || false

      if (data.updateAvailable) {
        releaseInfo.value = {
          currentVersion: data.currentVersion,
          latestVersion: data.latestVersion,
          releaseDate: data.releaseDate,
          releaseNotes: data.releaseNotes,
          releaseURL: data.releaseURL
        }
        latestVersion.value = data.latestVersion
      }

      lastCheckTime.value = new Date().toISOString()

      // Store in localStorage to persist across sessions
      localStorage.setItem('lastUpdateCheck', lastCheckTime.value)
      localStorage.setItem('updateAvailable', updateAvailable.value.toString())
      if (releaseInfo.value) {
        localStorage.setItem('releaseInfo', JSON.stringify(releaseInfo.value))
      }

      return data
    } catch (error) {
      console.error('Failed to check for updates:', error)
      // Don't throw - silently fail for update checks
      return { updateAvailable: false }
    } finally {
      isCheckingUpdates.value = false
    }
  }

  function dismissUpdate() {
    // Mark update as dismissed (but keep checking)
    localStorage.setItem('updateDismissed', latestVersion.value || '')
  }

  function shouldShowUpdateBadge() {
    if (!updateAvailable.value) return false

    const dismissedVersion = localStorage.getItem('updateDismissed')
    return dismissedVersion !== latestVersion.value
  }

  // Initialize from localStorage
  function initializeFromCache() {
    const cachedUpdateCheck = localStorage.getItem('lastUpdateCheck')
    const cachedUpdateAvailable = localStorage.getItem('updateAvailable')
    const cachedReleaseInfo = localStorage.getItem('releaseInfo')

    if (cachedUpdateCheck) {
      lastCheckTime.value = cachedUpdateCheck
    }

    if (cachedUpdateAvailable === 'true') {
      updateAvailable.value = true
    }

    if (cachedReleaseInfo) {
      try {
        releaseInfo.value = JSON.parse(cachedReleaseInfo)
        latestVersion.value = releaseInfo.value.latestVersion
      } catch (e) {
        console.error('Failed to parse cached release info:', e)
      }
    }
  }

  // Auto-check for updates periodically (every 4 hours)
  function startPeriodicUpdateCheck() {
    // Check immediately on start
    checkForUpdates()

    // Then check every 4 hours
    setInterval(() => {
      checkForUpdates()
    }, 4 * 60 * 60 * 1000) // 4 hours in milliseconds
  }

  return {
    // State
    currentVersion,
    latestVersion,
    updateAvailable,
    releaseInfo,
    isCheckingUpdates,
    lastCheckTime,

    // Getters
    versionInfo,
    hasUpdate,
    updateDetails,

    // Actions
    fetchVersion,
    checkForUpdates,
    dismissUpdate,
    shouldShowUpdateBadge,
    initializeFromCache,
    startPeriodicUpdateCheck
  }
})
