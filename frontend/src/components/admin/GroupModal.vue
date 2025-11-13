<template>
  <div v-if="show" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- Background overlay -->
      <div 
        class="fixed inset-0 bg-gray-900 bg-opacity-50 transition-opacity"
        @click="closeModal"
      ></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-xl text-left overflow-hidden shadow-2xl transform transition-all sm:my-8 sm:align-middle sm:max-w-3xl sm:w-full">
        <form @submit.prevent="handleSubmit">
          <!-- Header -->
          <div class="bg-white dark:bg-gray-800 px-6 pt-6 pb-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-3">
                <div class="h-10 w-10 rounded-lg bg-gray-100 dark:bg-gray-700 flex items-center justify-center">
                  <Icon icon="mdi:account-group" class="h-5 w-5 text-gray-600 dark:text-gray-400" />
                </div>
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    {{ isEdit ? 'Modifier le groupe' : 'Nouveau groupe d\'utilisateurs' }}
                  </h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ isEdit ? 'Modifiez les informations du groupe' : 'Créez un groupe pour organiser les permissions' }}
                  </p>
                </div>
              </div>
              <button 
                type="button" 
                @click="closeModal"
                class="btn-ghost"
              >
                <Icon icon="mdi:close" class="h-5 w-5" />
              </button>
            </div>
          </div>

          <!-- Form content -->
          <div class="px-6 py-4 space-y-6">
            <!-- Section: Informations de base -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:information-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Informations générales
              </h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <!-- Nom -->
                <div class="form-group">
                  <label for="name" class="form-label form-label-required">
                    Nom du groupe
                  </label>
                  <input
                    id="name"
                    v-model="form.name"
                    type="text"
                    required
                    class="form-input"
                    placeholder="Ex: Développeurs"
                  />
                  <p v-if="errors.name" class="form-error">{{ errors.name }}</p>
                </div>

                <!-- Couleur -->
                <div class="form-group">
                  <label for="color" class="form-label">
                    Couleur d'identification
                  </label>
                  <div class="flex items-center space-x-3">
                    <input
                      id="color"
                      v-model="form.color"
                      type="color"
                      class="h-10 w-16 rounded-lg border border-gray-300 dark:border-gray-600 cursor-pointer"
                    />
                    <input
                      v-model="form.color"
                      type="text"
                      class="form-input flex-1"
                      placeholder="#6B7280"
                    />
                  </div>
                </div>
              </div>

              <!-- Description -->
              <div class="form-group">
                <label for="description" class="form-label">
                  Description
                </label>
                <textarea
                  id="description"
                  v-model="form.description"
                  rows="2"
                  class="form-textarea"
                  placeholder="Description du groupe d'utilisateurs..."
                ></textarea>
              </div>
            </div>

            <!-- Section: Permissions d'accès -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:shield-key-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Permissions d'accès aux applications
              </h4>
              
              <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-4">
                  Sélectionnez les groupes d'applications auxquels ce groupe d'utilisateurs aura accès :
                </p>
                
                <div v-if="appGroups.length === 0" class="text-center py-8">
                  <Icon icon="mdi:folder-open" class="h-12 w-12 text-gray-400 mx-auto mb-2" />
                  <p class="text-sm text-gray-500">Aucun groupe d'applications disponible</p>
                  <p class="text-xs text-gray-400">Créez d'abord des groupes d'applications</p>
                </div>
                
                <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-3 max-h-60 overflow-y-auto">
                  <div v-for="appGroup in appGroups" :key="appGroup.id" 
                       class="flex items-center p-3 rounded-lg border-2 transition-all duration-200"
                       :class="form.app_group_ids.includes(appGroup.id) 
                         ? 'border-gray-400 bg-white dark:bg-gray-800' 
                         : 'border-gray-200 dark:border-gray-600 hover:border-gray-300'"
                  >
                    <input
                      :id="`appgroup_${appGroup.id}`"
                      v-model="form.app_group_ids"
                      :value="appGroup.id"
                      type="checkbox"
                      class="h-4 w-4 text-gray-600 focus:ring-gray-500 border-gray-300 rounded"
                    />
                    <label :for="`appgroup_${appGroup.id}`" class="ml-3 flex-1 cursor-pointer">
                      <div class="flex items-center space-x-3">
                        <div 
                          class="h-8 w-8 rounded-lg flex items-center justify-center"
                          :style="{ backgroundColor: appGroup.color || '#6B7280' }"
                        >
                          <Icon :icon="appGroup.icon || 'mdi:folder'" class="h-4 w-4 text-white" />
                        </div>
                        <div class="flex-1">
                          <div class="text-sm font-medium text-gray-900 dark:text-white">
                            {{ appGroup.name }}
                          </div>
                          <div class="text-xs text-gray-500 dark:text-gray-400">
                            {{ appGroup.applications?.length || 0 }} applications
                          </div>
                        </div>
                      </div>
                    </label>
                  </div>
                </div>
                
                <div class="mt-4 p-3 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
                  <div class="flex items-start space-x-2">
                    <Icon icon="mdi:information" class="h-4 w-4 text-blue-600 dark:text-blue-400 mt-0.5" />
                    <div class="text-sm text-blue-800 dark:text-blue-200">
                      <strong>Permissions sélectionnées :</strong> 
                      {{ form.app_group_ids.length }} groupe(s) d'applications
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Section: Membres (en lecture seule pour édition) -->
            <div v-if="isEdit && form.users && form.users.length > 0">
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:account-multiple-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Membres du groupe ({{ form.users.length }})
              </h4>
              
              <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4">
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
                  <div v-for="user in form.users" :key="user.id" 
                       class="flex items-center space-x-3 p-2 bg-white dark:bg-gray-800 rounded-lg"
                  >
                    <div class="h-8 w-8 rounded-full bg-gray-600 flex items-center justify-center">
                      <span class="text-xs font-medium text-white">
                        {{ getUserInitials(user) }}
                      </span>
                    </div>
                    <div class="flex-1 min-w-0">
                      <div class="text-sm font-medium text-gray-900 dark:text-white truncate">
                        {{ getUserDisplayName(user) }}
                      </div>
                      <div class="text-xs text-gray-500 truncate">
                        {{ user.email }}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Section: Options -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:tune" class="h-4 w-4 mr-2 text-gray-500" />
                Options du groupe
              </h4>
              
              <div class="flex items-center space-x-3 p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                <input
                  id="is_active"
                  v-model="form.is_active"
                  type="checkbox"
                  class="h-4 w-4 text-gray-600 focus:ring-gray-500 border-gray-300 rounded"
                />
                <label for="is_active" class="text-sm text-gray-700 dark:text-gray-300">
                  Groupe actif et permissions appliquées
                </label>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="bg-gray-50 dark:bg-gray-700 px-6 py-4 flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-3 space-y-3 space-y-reverse sm:space-y-0">
            <button
              type="button"
              @click="closeModal"
              class="btn-secondary w-full sm:w-auto"
            >
              Annuler
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="btn-primary w-full sm:w-auto"
            >
              <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ isEdit ? 'Enregistrer' : 'Créer' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { adminService } from '@/services/api'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  group: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['close', 'submit'])

const loading = ref(false)
const errors = ref({})
const appGroups = ref([])

const form = reactive({
  name: '',
  description: '',
  color: '#6B7280',
  app_group_ids: [],
  users: [],
  is_active: true
})

const isEdit = computed(() => !!props.group)

// Load app groups when component mounts
onMounted(async () => {
  try {
    const data = await adminService.getAppGroups()
    appGroups.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Erreur lors du chargement des groupes d\'applications:', error)
    appGroups.value = []
  }
})

// Reset form when modal opens/closes
watch(() => props.show, (newVal) => {
  if (newVal) {
    resetForm()
  }
})

watch(() => props.group, (newVal) => {
  if (newVal) {
    Object.assign(form, {
      name: newVal.name || '',
      description: newVal.description || '',
      color: newVal.color || '#6B7280',
      app_group_ids: newVal.app_groups?.map(ag => ag.id) || [],
      users: newVal.users || [],
      is_active: newVal.is_active ?? true
    })
  }
}, { immediate: true })

const resetForm = () => {
  if (props.group) {
    Object.assign(form, {
      name: props.group.name || '',
      description: props.group.description || '',
      color: props.group.color || '#6B7280',
      app_group_ids: props.group.app_groups?.map(ag => ag.id) || [],
      users: props.group.users || [],
      is_active: props.group.is_active ?? true
    })
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      color: '#6B7280',
      app_group_ids: [],
      users: [],
      is_active: true
    })
  }
  errors.value = {}
}

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

const validateForm = () => {
  errors.value = {}
  
  if (!form.name.trim()) {
    errors.value.name = 'Le nom est requis'
  }
  
  return Object.keys(errors.value).length === 0
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  try {
    const formData = { ...form }
    if (isEdit.value) {
      formData.id = props.group.id
    }
    
    // Remove users from form data as it's read-only
    delete formData.users
    
    emit('submit', formData)
  } finally {
    loading.value = false
  }
}

const closeModal = () => {
  resetForm()
  emit('close')
}
</script>