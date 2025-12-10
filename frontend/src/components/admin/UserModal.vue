<template>
  <div v-if="show" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- Background overlay -->
      <div 
        class="fixed inset-0 bg-gray-900 bg-opacity-50 transition-opacity"
        @click="closeModal"
      ></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-xl text-left overflow-hidden shadow-2xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
        <form @submit.prevent="handleSubmit">
          <!-- Header -->
          <div class="bg-white dark:bg-gray-800 px-6 pt-6 pb-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-3">
                <div class="h-10 w-10 rounded-lg bg-gray-100 dark:bg-gray-700 flex items-center justify-center">
                  <Icon icon="mdi:account" class="h-5 w-5 text-gray-600 dark:text-gray-400" />
                </div>
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    {{ isEdit ? 'Modifier l\'utilisateur' : 'Nouvel utilisateur' }}
                  </h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ isEdit ? 'Modifiez les informations du compte' : 'Créez un nouveau compte utilisateur' }}
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
            <!-- Section: Informations d'authentification -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:key-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Informations d'authentification
              </h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <!-- Username -->
                <div class="form-group">
                  <label for="username" class="form-label form-label-required">
                    Nom d'utilisateur
                  </label>
                  <input
                    id="username"
                    v-model="form.username"
                    type="text"
                    required
                    :disabled="isEdit"
                    class="form-input"
                    :class="{ 'bg-gray-100 dark:bg-gray-700 cursor-not-allowed': isEdit }"
                    placeholder="johndoe"
                  />
                  <p v-if="errors.username" class="form-error">{{ errors.username }}</p>
                  <p v-if="isEdit" class="form-help">Le nom d'utilisateur ne peut pas être modifié</p>
                </div>

                <!-- Email -->
                <div class="form-group">
                  <label for="email" class="form-label form-label-required">
                    Adresse email
                  </label>
                  <input
                    id="email"
                    v-model="form.email"
                    type="email"
                    required
                    class="form-input"
                    placeholder="john.doe@example.com"
                  />
                  <p v-if="errors.email" class="form-error">{{ errors.email }}</p>
                </div>
              </div>

              <!-- Password (only for new users) -->
              <div v-if="!isEdit" class="form-group">
                <label for="password" class="form-label form-label-required">
                  Mot de passe
                </label>
                <input
                  id="password"
                  v-model="form.password"
                  type="password"
                  required
                  class="form-input"
                  placeholder="••••••••"
                  minlength="6"
                />
                <p v-if="errors.password" class="form-error">{{ errors.password }}</p>
                <p class="form-help">Minimum 6 caractères</p>
              </div>
            </div>

            <!-- Section: Informations personnelles -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:account-details-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Informations personnelles
              </h4>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <!-- First Name -->
                <div class="form-group">
                  <label for="first_name" class="form-label">
                    Prénom
                  </label>
                  <input
                    id="first_name"
                    v-model="form.first_name"
                    type="text"
                    class="form-input"
                    placeholder="John"
                  />
                </div>

                <!-- Last Name -->
                <div class="form-group">
                  <label for="last_name" class="form-label">
                    Nom de famille
                  </label>
                  <input
                    id="last_name"
                    v-model="form.last_name"
                    type="text"
                    class="form-input"
                    placeholder="Doe"
                  />
                </div>
              </div>
            </div>

            <!-- Section: Permissions -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:shield-account-outline" class="h-4 w-4 mr-2 text-gray-500" />
                Permissions et groupes
              </h4>
              
              <div class="space-y-4">
                <!-- Role -->
                <div class="form-group">
                  <label for="role" class="form-label form-label-required">
                    Rôle
                  </label>
                  <select
                    id="role"
                    v-model="form.role"
                    required
                    class="form-select"
                  >
                    <option value="user">Utilisateur standard</option>
                    <option value="editor">Éditeur (gestion des actualités)</option>
                    <option value="group_admin">Administrateur de groupe</option>
                    <option value="admin">Administrateur</option>
                  </select>
                  <p v-if="errors.role" class="form-error">{{ errors.role }}</p>
                  <p class="form-help">
                    <span v-if="form.role === 'admin'">Les administrateurs ont accès à toutes les fonctionnalités</span>
                    <span v-else-if="form.role === 'group_admin'">Les administrateurs de groupe peuvent gérer les ressources des groupes dont ils sont responsables</span>
                    <span v-else-if="form.role === 'editor'">Les éditeurs peuvent créer et gérer les actualités du News Hub</span>
                    <span v-else>Les utilisateurs ont accès aux applications assignées via les groupes</span>
                  </p>
                </div>

                <!-- Groups -->
                <div class="form-group">
                  <label class="form-label">
                    Groupes d'utilisateurs
                  </label>
                  <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4 max-h-40 overflow-y-auto">
                    <div v-if="groups.length === 0" class="text-sm text-gray-500 text-center py-4">
                      Aucun groupe disponible
                    </div>
                    <div v-else class="space-y-3">
                      <div v-for="group in groups" :key="group.id" class="flex items-center">
                        <input
                          :id="`group_${group.id}`"
                          v-model="form.group_ids"
                          :value="group.id"
                          type="checkbox"
                          class="h-4 w-4 text-gray-600 focus:ring-gray-500 border-gray-300 rounded"
                        />
                        <label :for="`group_${group.id}`" class="ml-3 flex items-center space-x-2">
                          <div 
                            class="h-3 w-3 rounded-full"
                            :style="{ backgroundColor: group.color || '#6B7280' }"
                          ></div>
                          <span class="text-sm text-gray-700 dark:text-gray-300">{{ group.name }}</span>
                        </label>
                      </div>
                    </div>
                  </div>
                  <p class="form-help">Les groupes déterminent les applications accessibles</p>
                </div>
              </div>
            </div>

            <!-- Section: Options -->
            <div>
              <h4 class="text-sm font-medium text-gray-900 dark:text-white mb-4 flex items-center">
                <Icon icon="mdi:tune" class="h-4 w-4 mr-2 text-gray-500" />
                Options du compte
              </h4>
              
              <div class="flex items-center space-x-3 p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                <input
                  id="is_active"
                  v-model="form.is_active"
                  type="checkbox"
                  class="h-4 w-4 text-gray-600 focus:ring-gray-500 border-gray-300 rounded"
                />
                <label for="is_active" class="text-sm text-gray-700 dark:text-gray-300">
                  Compte actif et autorisé à se connecter
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
  user: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['close', 'submit'])

const loading = ref(false)
const errors = ref({})
const groups = ref([])

const form = reactive({
  username: '',
  email: '',
  password: '',
  first_name: '',
  last_name: '',
  role: 'user',
  group_ids: [],
  is_active: true
})

const isEdit = computed(() => !!props.user)

// Load groups when component mounts
onMounted(async () => {
  try {
    const data = await adminService.getGroups()
    groups.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Erreur lors du chargement des groupes:', error)
    groups.value = []
  }
})

// Reset form when modal opens/closes
watch(() => props.show, (newVal) => {
  if (newVal) {
    resetForm()
  }
})

watch(() => props.user, (newVal) => {
  if (newVal) {
    Object.assign(form, {
      username: newVal.username || '',
      email: newVal.email || '',
      password: '',
      first_name: newVal.first_name || '',
      last_name: newVal.last_name || '',
      role: newVal.role || 'user',
      group_ids: newVal.groups?.map(g => g.id) || [],
      is_active: newVal.is_active ?? true
    })
  }
}, { immediate: true })

const resetForm = () => {
  if (props.user) {
    Object.assign(form, {
      username: props.user.username || '',
      email: props.user.email || '',
      password: '',
      first_name: props.user.first_name || '',
      last_name: props.user.last_name || '',
      role: props.user.role || 'user',
      group_ids: props.user.groups?.map(g => g.id) || [],
      is_active: props.user.is_active ?? true
    })
  } else {
    Object.assign(form, {
      username: '',
      email: '',
      password: '',
      first_name: '',
      last_name: '',
      role: 'user',
      group_ids: [],
      is_active: true
    })
  }
  errors.value = {}
}

const validateForm = () => {
  errors.value = {}
  
  if (!form.username.trim()) {
    errors.value.username = 'Le nom d\'utilisateur est requis'
  } else if (form.username.length < 3) {
    errors.value.username = 'Le nom d\'utilisateur doit contenir au moins 3 caractères'
  }
  
  if (!form.email.trim()) {
    errors.value.email = 'L\'email est requis'
  } else if (!isValidEmail(form.email)) {
    errors.value.email = 'L\'email n\'est pas valide'
  }
  
  if (!isEdit.value) {
    if (!form.password.trim()) {
      errors.value.password = 'Le mot de passe est requis'
    } else if (form.password.length < 6) {
      errors.value.password = 'Le mot de passe doit contenir au moins 6 caractères'
    }
  }
  
  if (!form.role) {
    errors.value.role = 'Le rôle est requis'
  }
  
  return Object.keys(errors.value).length === 0
}

const isValidEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  try {
    const formData = { ...form }
    if (isEdit.value) {
      formData.id = props.user.id
      // Don't send password if it's empty in edit mode
      if (!formData.password) {
        delete formData.password
      }
    }
    
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