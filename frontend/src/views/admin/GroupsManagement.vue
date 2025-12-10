<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <div class="w-full">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">{{ $t('groups.title') }}</h1>
            <p class="mt-1 text-gray-600 dark:text-gray-400">{{ $t('groups.subtitle') }}</p>
          </div>
          <button @click="openCreateModal" class="btn btn-primary">
            <Icon icon="mdi:plus" class="h-5 w-5 mr-2" />
            {{ $t('groups.newGroup') }}
          </button>
        </div>
      </div>

      <!-- Search and filters -->
      <div class="mb-6 flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <div class="relative">
            <Icon icon="mdi:magnify" class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
            <input
              v-model="searchQuery"
              type="text"
              :placeholder="$t('groups.searchPlaceholder')"
              class="form-input pl-10"
            />
          </div>
        </div>
        <select v-model="statusFilter" class="form-select w-full sm:w-auto">
          <option value="">{{ $t('groups.status_all') }}</option>
          <option value="active">{{ $t('groups.status_active') }}</option>
          <option value="inactive">{{ $t('groups.status_inactive') }}</option>
        </select>
      </div>

      <!-- Groups Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="group in filteredGroups" :key="group.id" class="card p-6 hover:shadow-lg transition-shadow">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div 
                class="h-12 w-12 rounded-lg flex items-center justify-center"
                :style="{ backgroundColor: group.color || '#8B5CF6' }"
              >
                <Icon icon="mdi:account-group" class="h-6 w-6 text-white" />
              </div>
              <div>
                <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ group.name }}</h3>
                <p class="text-sm text-gray-500 dark:text-gray-400">{{ group.users?.length || 0 }} {{ $t('groups.users') }}</p>
              </div>
            </div>
            
            <div class="flex space-x-2">
              <button @click="manageGroupAdmins(group)" class="btn-ghost btn-sm" title="Gérer les administrateurs">
                <Icon icon="mdi:shield-account" class="h-4 w-4" />
              </button>
              <button @click="editGroup(group)" class="btn-ghost btn-sm" :title="$t('groups.edit')">
                <Icon icon="mdi:pencil" class="h-4 w-4" />
              </button>
              <button @click="confirmDelete(group)" class="btn-ghost btn-sm text-red-600 hover:text-red-700" :title="$t('groups.remove')">
                <Icon icon="mdi:delete" class="h-4 w-4" />
              </button>
            </div>
          </div>
          
          <p v-if="group.description" class="text-sm text-gray-600 dark:text-gray-400 mb-4">{{ group.description }}</p>
          
          <!-- App Groups Access -->
          <div v-if="group.app_groups && group.app_groups.length > 0" class="mb-4">
            <p class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-2">{{ $t('groups.appsAccess') }}</p>
            <div class="flex flex-wrap gap-1">
              <span 
                v-for="appGroup in group.app_groups" 
                :key="appGroup.id"
                class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                :style="{ backgroundColor: appGroup.color + '20', color: appGroup.color }"
              >
                <Icon :icon="appGroup.icon" class="h-3 w-3 mr-1" />
                {{ appGroup.name }}
              </span>
            </div>
          </div>
          
          <!-- Users in group -->
          <div v-if="group.users && group.users.length > 0" class="mb-4">
            <p class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-2">{{ $t('groups.usersLabel') }}</p>
            <div class="flex -space-x-2 overflow-hidden">
              <div 
                v-for="(user, index) in group.users.slice(0, 5)" 
                :key="user.id"
                class="inline-block h-8 w-8 rounded-full bg-primary-500 flex items-center justify-center ring-2 ring-white dark:ring-gray-800"
                :title="getUserDisplayName(user)"
              >
                <span class="text-xs font-medium text-white">
                  {{ getUserInitials(user) }}
                </span>
              </div>
              <div v-if="group.users.length > 5" class="inline-block h-8 w-8 rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center ring-2 ring-white dark:ring-gray-800">
                <span class="text-xs font-medium text-gray-600 dark:text-gray-300">
                  +{{ group.users.length - 5 }}
                </span>
              </div>
            </div>
          </div>
          
          <div class="flex items-center justify-between text-sm">
            <span :class="group.is_active ? 'badge bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200' : 'badge bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200'">
              {{ group.is_active ? $t('common.active') : $t('common.inactive') }}
            </span>
            <span class="text-gray-500">{{ group.app_groups?.length || 0 }} {{ $t('groups.appsLabel') }}</span>
          </div>
        </div>

        <!-- Empty state -->
        <div v-if="filteredGroups.length === 0 && groups.length > 0" class="col-span-full text-center py-12">
          <Icon icon="mdi:magnify" class="mx-auto h-12 w-12 text-gray-400" />
          <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">{{ $t('groups.noResultsTitle') }}</h3>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">{{ $t('groups.noResultsText') }}</p>
        </div>

        <div v-else-if="groups.length === 0" class="col-span-full text-center py-12">
          <Icon icon="mdi:account-group" class="mx-auto h-12 w-12 text-gray-400" />
          <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">{{ $t('groups.emptyTitle') }}</h3>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">{{ $t('groups.emptyText') }}</p>
          <button @click="openCreateModal" class="btn btn-primary mt-4">
            <Icon icon="mdi:plus" class="h-5 w-5 mr-2" />
            {{ $t('groups.createGroup') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Group Modal -->
    <GroupModal
      :show="showModal"
      :group="selectedGroup"
      @close="closeModal"
      @submit="handleSubmit"
    />

    <!-- Group Admins Modal -->
    <div v-if="showGroupAdminsModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="closeGroupAdminsModal"></div>
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <div class="bg-white dark:bg-gray-800 px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="mb-6">
              <div class="flex items-center">
                <div class="h-10 w-10 rounded-lg bg-indigo-100 dark:bg-indigo-900 flex items-center justify-center mr-3">
                  <Icon icon="mdi:shield-account" class="h-5 w-5 text-indigo-600 dark:text-indigo-400" />
                </div>
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    Administrateurs du groupe
                  </h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ selectedGroupForAdmins?.name }}
                  </p>
                </div>
              </div>
            </div>

            <div class="mb-4">
              <p class="text-sm text-gray-600 dark:text-gray-400 mb-4">
                Sélectionnez les utilisateurs qui pourront administrer ce groupe et ses ressources.
              </p>

              <!-- Search -->
              <div class="mb-4">
                <input
                  v-model="adminSearchQuery"
                  type="text"
                  placeholder="Rechercher un utilisateur..."
                  class="form-input"
                />
              </div>

              <!-- Loading -->
              <div v-if="loadingUsers" class="flex justify-center py-8">
                <Icon icon="mdi:loading" class="h-6 w-6 animate-spin text-gray-400" />
              </div>

              <!-- Users list -->
              <div v-else class="max-h-96 overflow-y-auto border border-gray-200 dark:border-gray-700 rounded-lg">
                <div v-if="filteredUsersForAdmins.length === 0" class="p-4 text-center text-sm text-gray-500">
                  Aucun utilisateur trouvé
                </div>
                <div
                  v-for="user in filteredUsersForAdmins"
                  :key="user.id"
                  class="flex items-center p-3 hover:bg-gray-50 dark:hover:bg-gray-700 border-b border-gray-200 dark:border-gray-700 last:border-b-0"
                >
                  <input
                    :id="`admin_${user.id}`"
                    v-model="selectedAdminIds"
                    :value="user.id"
                    type="checkbox"
                    class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                  />
                  <label :for="`admin_${user.id}`" class="ml-3 flex items-center flex-1 cursor-pointer">
                    <div class="h-8 w-8 rounded-full bg-primary-500 flex items-center justify-center mr-3">
                      <span class="text-sm font-medium text-white">{{ getUserInitials(user) }}</span>
                    </div>
                    <div class="flex-1">
                      <p class="text-sm font-medium text-gray-900 dark:text-white">{{ getUserDisplayName(user) }}</p>
                      <p class="text-xs text-gray-500 dark:text-gray-400">@{{ user.username }} • {{ user.email }}</p>
                    </div>
                    <span
                      :class="{
                        'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200': user.role === 'admin',
                        'bg-indigo-100 text-indigo-800 dark:bg-indigo-900 dark:text-indigo-200': user.role === 'group_admin',
                        'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200': user.role === 'editor',
                        'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200': user.role === 'user'
                      }"
                      class="badge text-xs"
                    >
                      {{ user.role === 'admin' ? 'Admin' : user.role === 'group_admin' ? 'Group Admin' : user.role === 'editor' ? 'Editor' : 'User' }}
                    </span>
                  </label>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              @click="saveGroupAdmins"
              :disabled="savingAdmins"
              class="btn-primary w-full sm:w-auto sm:ml-3"
            >
              <Icon v-if="savingAdmins" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              Enregistrer
            </button>
            <button
              @click="closeGroupAdminsModal"
              class="btn btn-secondary w-full sm:w-auto mt-3 sm:mt-0"
            >
              Annuler
            </button>
          </div>
        </div>
      </div>
    </div>

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
                  {{ $t('groups.confirmDeleteTitle') }}
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ $t('groups.confirmDeleteText', { name: groupToDelete?.name }) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              @click="deleteGroup"
              :disabled="deleteLoading"
              class="btn-danger w-full sm:w-auto sm:ml-3"
            >
              <Icon v-if="deleteLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ $t('common.delete') }}
            </button>
            <button
              @click="closeDeleteModal"
              class="btn btn-secondary w-full sm:w-auto mt-3 sm:mt-0 sm:ml-3"
            >
              {{ $t('common.cancel') }}
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
import { useI18n } from 'vue-i18n'
import GroupModal from '@/components/admin/GroupModal.vue'

const appStore = useAppStore()
const { t } = useI18n()

// State
const groups = ref([])
const showModal = ref(false)
const selectedGroup = ref(null)
const showDeleteModal = ref(false)
const groupToDelete = ref(null)
const deleteLoading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')

// Group Admins Modal
const showGroupAdminsModal = ref(false)
const selectedGroupForAdmins = ref(null)
const allUsers = ref([])
const selectedAdminIds = ref([])
const loadingUsers = ref(false)
const savingAdmins = ref(false)
const adminSearchQuery = ref('')

// Computed
const filteredGroups = computed(() => {
  let filtered = groups.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(group =>
      group.name.toLowerCase().includes(query) ||
      group.description?.toLowerCase().includes(query)
    )
  }

  // Filter by status
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(group => group.is_active)
  } else if (statusFilter.value === 'inactive') {
    filtered = filtered.filter(group => !group.is_active)
  }

  return Array.isArray(filtered) ? filtered.sort((a, b) => a.name.localeCompare(b.name)) : []
})

const filteredUsersForAdmins = computed(() => {
  let filtered = allUsers.value

  if (adminSearchQuery.value) {
    const query = adminSearchQuery.value.toLowerCase()
    filtered = filtered.filter(user =>
      user.username.toLowerCase().includes(query) ||
      user.email.toLowerCase().includes(query) ||
      user.first_name?.toLowerCase().includes(query) ||
      user.last_name?.toLowerCase().includes(query)
    )
  }

  return filtered
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

const loadGroups = async () => {
  try {
    appStore.setLoading(true)
    const data = await adminService.getGroups()
    groups.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Erreur lors du chargement des groupes:', error)
    groups.value = []
    appStore.showError(t('groups.loadError'))
  } finally {
    appStore.setLoading(false)
  }
}

const openCreateModal = () => {
  selectedGroup.value = null
  showModal.value = true
}

const editGroup = (group) => {
  selectedGroup.value = group
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedGroup.value = null
}

const confirmDelete = (group) => {
  groupToDelete.value = group
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  groupToDelete.value = null
}

const handleSubmit = async (formData) => {
  try {
    appStore.setLoading(true)
    
    if (formData.id) {
      // Update existing group - extract only the fields we need
      const { id, ...updateData } = formData
      await adminService.updateGroup(id, updateData)
      appStore.showSuccess(t('groups.updated'))
    } else {
      // Create new group - extract only the fields we need
      const { id, ...createData } = formData
      await adminService.createGroup(createData)
      appStore.showSuccess(t('groups.created'))
    }
    
    closeModal()
    await loadGroups()
  } catch (error) {
    console.error('Erreur lors de la sauvegarde:', error)
    appStore.showError(t('groups.saveError'))
  } finally {
    appStore.setLoading(false)
  }
}

const deleteGroup = async () => {
  if (!groupToDelete.value) return

  try {
    deleteLoading.value = true
    await adminService.deleteGroup(groupToDelete.value.id)
    appStore.showSuccess(t('groups.deleted'))
    closeDeleteModal()
    await loadGroups()
  } catch (error) {
    console.error('Erreur lors de la suppression:', error)
    appStore.showError(t('groups.deleteError'))
  } finally {
    deleteLoading.value = false
  }
}

// Group Admins Management
const manageGroupAdmins = async (group) => {
  selectedGroupForAdmins.value = group
  showGroupAdminsModal.value = true
  adminSearchQuery.value = ''

  try {
    loadingUsers.value = true

    // Load all users
    const usersData = await adminService.getUsers()
    allUsers.value = Array.isArray(usersData) ? usersData : (usersData.data || [])

    // Load current group admins
    const adminsData = await adminService.getGroupAdmins(group.id)
    const currentAdmins = Array.isArray(adminsData) ? adminsData : (adminsData.data || [])
    selectedAdminIds.value = currentAdmins.map(admin => admin.id)
  } catch (error) {
    console.error('Error loading group admins:', error)
    appStore.showError('Erreur lors du chargement des données')
  } finally {
    loadingUsers.value = false
  }
}

const closeGroupAdminsModal = () => {
  showGroupAdminsModal.value = false
  selectedGroupForAdmins.value = null
  allUsers.value = []
  selectedAdminIds.value = []
  adminSearchQuery.value = ''
}

const saveGroupAdmins = async () => {
  if (!selectedGroupForAdmins.value) return

  try {
    savingAdmins.value = true
    await adminService.assignGroupAdmins(selectedGroupForAdmins.value.id, selectedAdminIds.value)
    appStore.showSuccess('Administrateurs de groupe mis à jour avec succès')
    closeGroupAdminsModal()
  } catch (error) {
    console.error('Error saving group admins:', error)
    appStore.showError('Erreur lors de la sauvegarde')
  } finally {
    savingAdmins.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadGroups()
})
</script>