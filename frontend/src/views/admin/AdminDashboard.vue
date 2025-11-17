<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <div class="w-full">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
          {{ $t('admin.overviewTitle') }}
        </h1>
        <p class="mt-1 text-gray-600 dark:text-gray-400">
          {{ $t('admin.overviewSubtitle') }}
        </p>
      </div>

      <!-- Statistiques -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div class="card p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="h-8 w-8 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center">
                <Icon icon="mdi:folder-multiple" class="h-5 w-5 text-blue-600 dark:text-blue-400" />
              </div>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('common.appGroups') }}</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ stats.total_app_groups || 0 }}
              </p>
            </div>
          </div>
        </div>

        <div class="card p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="h-8 w-8 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
                <Icon icon="mdi:application" class="h-5 w-5 text-green-600 dark:text-green-400" />
              </div>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('common.applications') }}</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ stats.total_applications || 0 }}
              </p>
            </div>
          </div>
        </div>

        <div class="card p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="h-8 w-8 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center">
                <Icon icon="mdi:account-multiple" class="h-5 w-5 text-purple-600 dark:text-purple-400" />
              </div>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('common.users') }}</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ stats.total_users || 0 }}
              </p>
            </div>
          </div>
        </div>

        <div class="card p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="h-8 w-8 bg-yellow-100 dark:bg-yellow-900 rounded-lg flex items-center justify-center">
                <Icon icon="mdi:account-group" class="h-5 w-5 text-yellow-600 dark:text-yellow-400" />
              </div>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('common.groups') }}</p>
              <p class="text-2xl font-bold text-gray-900 dark:text-white">
                {{ stats.total_groups || 0 }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Actions rapides -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <router-link to="/admin/app-groups" class="card-hover p-6">
          <div class="flex items-center gap-4">
            <div class="h-12 w-12 bg-blue-500 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:folder-multiple" class="h-6 w-6 text-white" />
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ $t('common.appGroups') }}</h3>
              <p class="text-sm text-gray-500 dark:text-gray-400">{{ $t('admin.manageAppGroups') }}</p>
            </div>
          </div>
        </router-link>

        <router-link to="/admin/applications" class="card-hover p-6">
          <div class="flex items-center gap-4">
            <div class="h-12 w-12 bg-green-500 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:application" class="h-6 w-6 text-white" />
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ $t('common.applications') }}</h3>
              <p class="text-sm text-gray-500 dark:text-gray-400">{{ $t('admin.manageApplications') }}</p>
            </div>
          </div>
        </router-link>

        <router-link to="/admin/users" class="card-hover p-6">
          <div class="flex items-center gap-4">
            <div class="h-12 w-12 bg-purple-500 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:account-multiple" class="h-6 w-6 text-white" />
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ $t('common.users') }}</h3>
              <p class="text-sm text-gray-500 dark:text-gray-400">{{ $t('admin.manageUsers') }}</p>
            </div>
          </div>
        </router-link>

        <router-link to="/admin/groups" class="card-hover p-6">
          <div class="flex items-center gap-4">
            <div class="h-12 w-12 bg-yellow-500 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:account-group" class="h-6 w-6 text-white" />
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ $t('common.groups') }}</h3>
              <p class="text-sm text-gray-500 dark:text-gray-400">{{ $t('admin.manageGroups') }}</p>
            </div>
          </div>
        </router-link>
      </div>

      <!-- Zone de danger -->
      <div class="card border-2 border-red-200 dark:border-red-800 p-6">
        <div class="flex items-start gap-4">
          <div class="flex-shrink-0">
            <div class="h-12 w-12 bg-red-100 dark:bg-red-900 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:database-refresh" class="h-6 w-6 text-red-600 dark:text-red-400" />
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-red-900 dark:text-red-100 mb-2">
              {{ $t('admin.dangerZone') }}
            </h3>
            <p class="text-sm text-gray-700 dark:text-gray-300 mb-4">
              {{ $t('admin.resetDatabaseTitle') }}
              <strong class="text-red-600 dark:text-red-400">{{ $t('admin.resetDatabaseWarning') }}</strong>
            </p>
            <button
              @click="showResetConfirmation = true"
              :disabled="isResetting"
              class="px-4 py-2 bg-red-600 hover:bg-red-700 disabled:bg-red-400 text-white rounded-lg font-medium transition-colors flex items-center gap-2"
            >
              <Icon icon="mdi:database-refresh" class="h-5 w-5" />
              <span v-if="!isResetting">{{ $t('admin.resetDatabaseButton') }}</span>
              <span v-else>{{ $t('admin.resetDatabaseInProgress') }}</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Modal de confirmation -->
      <div
        v-if="showResetConfirmation"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
        @click.self="showResetConfirmation = false"
      >
        <div class="bg-white dark:bg-gray-800 rounded-lg max-w-md w-full p-6">
          <div class="flex items-center gap-3 mb-4">
            <div class="h-12 w-12 bg-red-100 dark:bg-red-900 rounded-full flex items-center justify-center">
              <Icon icon="mdi:alert" class="h-6 w-6 text-red-600 dark:text-red-400" />
            </div>
            <h3 class="text-xl font-bold text-gray-900 dark:text-white">
              {{ $t('admin.confirmResetTitle') }}
            </h3>
          </div>

          <p class="text-gray-700 dark:text-gray-300 mb-6">
            {{ $t('admin.confirmResetMessage') }}
            <strong class="text-red-600 dark:text-red-400 block mt-2">
              {{ $t('admin.confirmResetWarning') }}
            </strong>
          </p>

          <div class="flex gap-3">
            <button
              @click="showResetConfirmation = false"
              class="flex-1 px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-white rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
            >
              {{ $t('admin.confirmResetCancel') }}
            </button>
            <button
              @click="handleResetDatabase"
              :disabled="isResetting"
              class="flex-1 px-4 py-2 bg-red-600 hover:bg-red-700 disabled:bg-red-400 text-white rounded-lg font-medium transition-colors"
            >
              {{ isResetting ? $t('admin.resetInProgress') : $t('admin.confirmResetButton') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { dashboardService, adminService } from '@/services/api'

const router = useRouter()

const stats = ref({
  total_app_groups: 0,
  total_applications: 0,
  total_users: 0,
  total_groups: 0
})

const showResetConfirmation = ref(false)
const isResetting = ref(false)

const loadStats = async () => {
  try {
    const dashboard = await dashboardService.getDashboard()
    if (dashboard.stats) {
      stats.value = dashboard.stats
    }
  } catch (error) {
    console.error('Erreur lors du chargement des statistiques:', error)
  }
}

const handleResetDatabase = async () => {
  isResetting.value = true
  try {
    const response = await adminService.resetDatabase()

    // Afficher un message de succès
    alert('Base de données réinitialisée avec succès!\n\n' + response.message + '\n\nVous allez être déconnecté.')

    // Fermer le modal
    showResetConfirmation.value = false

    // Déconnecter l'utilisateur et rediriger vers la page de connexion
    localStorage.removeItem('airboard_token')
    localStorage.removeItem('airboard_refresh_token')
    localStorage.removeItem('airboard_user')

    // Rediriger vers la page de connexion après 1 seconde
    setTimeout(() => {
      router.push('/auth/login')
      // Recharger la page pour que le serveur recrée les données initiales
      window.location.reload()
    }, 1000)
  } catch (error) {
    console.error('Erreur lors de la réinitialisation:', error)
    alert('Erreur lors de la réinitialisation de la base de données: ' + (error.response?.data?.message || error.message))
  } finally {
    isResetting.value = false
  }
}

onMounted(() => {
  loadStats()
})
</script>