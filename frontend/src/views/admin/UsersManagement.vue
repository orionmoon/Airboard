<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">{{ $t('users.title') }}</h1>
            <p class="mt-1 text-gray-600 dark:text-gray-400">{{ $t('users.subtitle') }}</p>
          </div>
          <button @click="openCreateModal" class="btn btn-primary">
            <Icon icon="mdi:plus" class="h-5 w-5 mr-2" />
            {{ $t('users.new') }}
          </button>
        </div>
      </div>

      <!-- View Toggle -->
      <div class="mb-6 flex items-center justify-between">
        <div class="flex items-center gap-2 bg-gray-100 dark:bg-gray-800 p-1 rounded-lg">
          <button
            @click="showDeleted = false"
            :class="[
              'px-4 py-2 rounded-md transition-colors',
              !showDeleted
                ? 'bg-white dark:bg-gray-700 text-gray-900 dark:text-white shadow-sm'
                : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white'
            ]"
          >
            <Icon icon="mdi:account-check" class="h-5 w-5 inline mr-2" />
            Utilisateurs actifs
          </button>
          <button
            @click="showDeleted = true"
            :class="[
              'px-4 py-2 rounded-md transition-colors',
              showDeleted
                ? 'bg-white dark:bg-gray-700 text-gray-900 dark:text-white shadow-sm'
                : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white'
            ]"
          >
            <Icon icon="mdi:delete" class="h-5 w-5 inline mr-2" />
            Corbeille
            <span v-if="deletedUsers.length > 0" class="ml-2 bg-red-500 text-white text-xs px-2 py-0.5 rounded-full">
              {{ deletedUsers.length }}
            </span>
          </button>
        </div>
      </div>

      <!-- Search and filters (only for active users) -->
      <div v-if="!showDeleted" class="mb-6 flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <div class="relative">
            <Icon icon="mdi:magnify" class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
            <input
              v-model="searchQuery"
              type="text"
              :placeholder="$t('users.searchPlaceholder')"
              class="form-input pl-10"
            />
          </div>
        </div>
        <select v-model="roleFilter" class="form-select w-full sm:w-auto">
          <option value="">{{ $t('users.filter_all_roles') }}</option>
          <option value="admin">{{ $t('users.filter_admins') }}</option>
          <option value="editor">{{ $t('users.filter_editors') }}</option>
          <option value="user">{{ $t('users.filter_users') }}</option>
        </select>
        <select v-model="statusFilter" class="form-select w-full sm:w-auto">
          <option value="">{{ $t('common.status_all') }}</option>
          <option value="active">{{ $t('common.status_active') }}</option>
          <option value="inactive">{{ $t('common.status_inactive') }}</option>
        </select>
      </div>

      <!-- Users Table -->
      <div class="card overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
            <thead class="bg-gray-50 dark:bg-gray-800">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Utilisateur
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Email
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Rôle
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Groupes
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Statut
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white dark:bg-gray-900 divide-y divide-gray-200 dark:divide-gray-700">
              <tr v-for="user in displayedUsers" :key="user.id" class="hover:bg-gray-50 dark:hover:bg-gray-800">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="h-10 w-10 rounded-full bg-primary-500 flex items-center justify-center">
                      <span class="text-sm font-medium text-white">{{ getUserInitials(user) }}</span>
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900 dark:text-white">{{ getUserDisplayName(user) }}</div>
                      <div class="text-sm text-gray-500">@{{ user.username }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-white">
                  {{ user.email }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
                    :class="{
                      'badge bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200': user.role === 'admin',
                      'badge bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200': user.role === 'editor',
                      'badge bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200': user.role === 'user'
                    }"
                  >
                    {{ user.role === 'admin' ? $t('users.role_admin') : user.role === 'editor' ? $t('users.role_editor') : $t('users.role_user') }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex flex-wrap gap-1">
                    <span
                      v-for="group in user.groups"
                      :key="group.id"
                      class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                      :style="{ backgroundColor: group.color + '20', color: group.color }"
                    >
                      {{ group.name }}
                    </span>
                    <span v-if="!user.groups || user.groups.length === 0" class="text-sm text-gray-500">
                      {{ $t('users.no_group') }}
                    </span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="user.is_active ? 'badge bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200' : 'badge bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200'">
                    {{ user.is_active ? $t('common.active') : $t('common.inactive') }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                  <!-- Actions pour utilisateurs actifs -->
                  <div v-if="!showDeleted" class="flex gap-2">
                    <button @click="editUser(user)" class="btn-ghost btn-sm" :title="$t('common.edit')">
                      <Icon icon="mdi:pencil" class="h-4 w-4" />
                    </button>
                    <button
                      @click="toggleUserStatus(user)"
                      class="btn-ghost btn-sm"
                      :class="user.is_active ? 'text-orange-600 hover:text-orange-700' : 'text-green-600 hover:text-green-700'"
                      :title="user.is_active ? $t('users.deactivate') : $t('users.activate')"
                    >
                      <Icon :icon="user.is_active ? 'mdi:account-off' : 'mdi:account-check'" class="h-4 w-4" />
                    </button>
                    <button
                      @click="confirmDelete(user)"
                      class="btn-ghost btn-sm text-red-600 hover:text-red-700"
                      :title="$t('common.delete')"
                    >
                      <Icon icon="mdi:delete" class="h-4 w-4" />
                    </button>
                  </div>
                  <!-- Actions pour utilisateurs supprimés -->
                  <div v-else class="flex gap-2 justify-end">
                    <button
                      @click="restoreUser(user)"
                      class="inline-flex items-center justify-center px-3 py-2 rounded-md text-sm font-medium bg-green-100 text-green-700 hover:bg-green-200 dark:bg-green-900 dark:text-green-300 dark:hover:bg-green-800 transition-colors"
                      title="Restaurer l'utilisateur"
                    >
                      <Icon icon="mdi:restore" class="h-4 w-4" />
                    </button>
                    <button
                      @click="confirmPermanentDelete(user)"
                      class="inline-flex items-center justify-center px-3 py-2 rounded-md text-sm font-medium bg-red-100 text-red-700 hover:bg-red-200 dark:bg-red-900 dark:text-red-300 dark:hover:bg-red-800 transition-colors"
                      title="Supprimer définitivement (irréversible)"
                    >
                      <Icon icon="mdi:delete-forever" class="h-4 w-4" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Empty state -->
        <div v-if="filteredUsers.length === 0 && users.length > 0" class="text-center py-12">
          <Icon icon="mdi:magnify" class="mx-auto h-12 w-12 text-gray-400" />
          <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">{{ $t('users.noResultsTitle') }}</h3>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">{{ $t('users.noResultsText') }}</p>
        </div>

        <div v-else-if="users.length === 0" class="text-center py-12">
          <Icon icon="mdi:account-multiple" class="mx-auto h-12 w-12 text-gray-400" />
          <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">{{ $t('users.emptyTitle') }}</h3>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">{{ $t('users.emptyText') }}</p>
          <button @click="openCreateModal" class="btn-primary mt-4">
            <Icon icon="mdi:plus" class="h-5 w-5 mr-2" />
            {{ $t('users.create') }}
          </button>
        </div>
      </div>
    </div>

    <!-- User Modal -->
    <UserModal
      :show="showModal"
      :user="selectedUser"
      @close="closeModal"
      @submit="handleSubmit"
    />

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="closeDeleteModal"></div>
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white dark:bg-gray-800 px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <Icon icon="mdi:alert" class="h-6 w-6 text-red-600" />
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white">
                  Supprimer l'utilisateur
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    Êtes-vous sûr de vouloir supprimer l'utilisateur "<strong>{{ userToDelete?.username }}</strong>" ?
                    Cette action est irréversible.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              @click="deleteUser"
              :disabled="deleteLoading"
              class="btn-danger w-full sm:w-auto sm:ml-3"
            >
              <Icon v-if="deleteLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              Supprimer
            </button>
            <button
              @click="closeDeleteModal"
              class="btn-secondary w-full sm:w-auto mt-3 sm:mt-0 sm:ml-3"
            >
              Annuler
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { adminService } from '@/services/api'
import { useAppStore } from '@/stores/app'
import UserModal from '@/components/admin/UserModal.vue'

const appStore = useAppStore()

// State
const users = ref([])
const deletedUsers = ref([])
const showDeleted = ref(false)
const showModal = ref(false)
const selectedUser = ref(null)
const showDeleteModal = ref(false)
const userToDelete = ref(null)
const deleteLoading = ref(false)
const searchQuery = ref('')
const roleFilter = ref('')
const statusFilter = ref('')

// Computed
const filteredUsers = computed(() => {
  let filtered = users.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(user => 
      user.username.toLowerCase().includes(query) ||
      user.email.toLowerCase().includes(query) ||
      user.first_name?.toLowerCase().includes(query) ||
      user.last_name?.toLowerCase().includes(query)
    )
  }

  // Filter by role
  if (roleFilter.value) {
    filtered = filtered.filter(user => user.role === roleFilter.value)
  }

  // Filter by status
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(user => user.is_active)
  } else if (statusFilter.value === 'inactive') {
    filtered = filtered.filter(user => !user.is_active)
  }

  return Array.isArray(filtered) ? filtered.sort((a, b) => a.username.localeCompare(b.username)) : []
})

// Computed pour afficher soit les utilisateurs actifs, soit les supprimés
const displayedUsers = computed(() => {
  return showDeleted.value ? deletedUsers.value : filteredUsers.value
})

// Methods
const getUserInitials = (user) => {
  if (user.first_name && user.last_name) {
    return `${user.first_name.charAt(0)}${user.last_name.charAt(0)}`.toUpperCase()
  }
  return user.username.charAt(0).toUpperCase()
}

const getUserDisplayName = (user) => {
  if (user.first_name && user.last_name) {
    return `${user.first_name} ${user.last_name}`
  }
  return user.username
}

const loadUsers = async () => {
  try {
    appStore.setLoading(true)
    const data = await adminService.getUsers()
    users.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Erreur lors du chargement des utilisateurs:', error)
    users.value = []
    appStore.showError('Erreur lors du chargement des utilisateurs')
  } finally {
    appStore.setLoading(false)
  }
}

const loadDeletedUsers = async () => {
  try {
    const data = await adminService.getDeletedUsers()
    deletedUsers.value = Array.isArray(data.users) ? data.users : []
  } catch (error) {
    console.error('Erreur lors du chargement des utilisateurs supprimés:', error)
    deletedUsers.value = []
  }
}

const openCreateModal = () => {
  selectedUser.value = null
  showModal.value = true
}

const editUser = (user) => {
  selectedUser.value = user
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedUser.value = null
}

const confirmDelete = (user) => {
  userToDelete.value = user
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  userToDelete.value = null
}

const handleSubmit = async (formData) => {
  try {
    appStore.setLoading(true)
    
    if (formData.id) {
      // Update existing user - only send relevant fields
      const { id, ...updateData } = formData
      await adminService.updateUser(id, updateData)
      appStore.showSuccess('Utilisateur modifié avec succès')
    } else {
      // Create new user - only send relevant fields  
      const { id, ...createData } = formData
      await adminService.createUser(createData)
      appStore.showSuccess('Utilisateur créé avec succès')
    }
    
    closeModal()
    await loadUsers()
  } catch (error) {
    console.error('Erreur lors de la sauvegarde:', error)
    appStore.showError('Erreur lors de la sauvegarde de l\'utilisateur')
  } finally {
    appStore.setLoading(false)
  }
}

const toggleUserStatus = async (user) => {
  try {
    // Only send the fields that need to be updated
    const updatedData = {
      is_active: !user.is_active
    }
    
    await adminService.updateUser(user.id, updatedData)
    appStore.showSuccess(`Utilisateur ${!user.is_active ? 'activé' : 'désactivé'} avec succès`)
    await loadUsers()
  } catch (error) {
    console.error('Erreur lors de la modification du statut:', error)
    appStore.showError('Erreur lors de la modification du statut')
  }
}

const deleteUser = async () => {
  if (!userToDelete.value || !userToDelete.value.id) {
    console.error('Aucun utilisateur sélectionné pour suppression')
    return
  }

  try {
    deleteLoading.value = true
    console.log('Suppression de l\'utilisateur:', userToDelete.value.id)
    await adminService.deleteUser(userToDelete.value.id)
    appStore.showSuccess('Utilisateur supprimé avec succès')
    closeDeleteModal()
    await loadUsers()
    await loadDeletedUsers()
  } catch (error) {
    console.error('Erreur lors de la suppression:', error)
    appStore.showError('Erreur lors de la suppression de l\'utilisateur')
  } finally {
    deleteLoading.value = false
  }
}

const restoreUser = async (user) => {
  try {
    await adminService.restoreUser(user.id)
    appStore.showSuccess(`Utilisateur ${user.username} restauré avec succès`)
    await loadUsers()
    await loadDeletedUsers()
  } catch (error) {
    console.error('Erreur lors de la restauration:', error)
    appStore.showError('Erreur lors de la restauration de l\'utilisateur')
  }
}

const confirmPermanentDelete = async (user) => {
  if (confirm(`⚠️ ATTENTION: Supprimer DÉFINITIVEMENT l'utilisateur "${user.username}" ?\n\nCette action est IRRÉVERSIBLE et supprimera toutes les données associées.`)) {
    try {
      await adminService.permanentlyDeleteUser(user.id)
      appStore.showSuccess('Utilisateur supprimé définitivement')
      await loadDeletedUsers()
    } catch (error) {
      console.error('Erreur lors de la suppression définitive:', error)
      appStore.showError('Erreur lors de la suppression définitive')
    }
  }
}

// Lifecycle
onMounted(() => {
  loadUsers()
  loadDeletedUsers()
})
</script>