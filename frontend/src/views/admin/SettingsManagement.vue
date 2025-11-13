<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ $t('settings.title') }}</h1>
          <p class="page-subtitle">{{ $t('settings.subtitle') }}</p>
        </div>
        <div class="flex gap-3">
          <button @click="showPasswordModal = true" class="btn btn-secondary">
            <Icon icon="mdi:key" class="h-4 w-4 mr-2" />
            {{ $t('settings.changePassword') }}
          </button>
          <button @click="resetSettings" class="btn btn-secondary">
            <Icon icon="mdi:restore" class="h-4 w-4 mr-2" />
            {{ $t('settings.resetToDefault') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Settings Form -->
    <div class="max-w-4xl">
      <div class="card">
        <form @submit.prevent="handleSubmit" class="space-y-8">
          <!-- Preview Section -->
          <div>
            <div class="section-header">
              <Icon icon="mdi:eye-outline" class="section-icon" />
              <h4 class="section-title">{{ $t('settings.livePreview') }}</h4>
            </div>
            
            <div class="bg-gray-700 rounded-lg p-6 border-2 border-dashed border-gray-600">
              <!-- Header Preview -->
              <div class="flex items-center gap-3 mb-4">
                <div class="h-10 w-10 bg-white rounded-lg flex items-center justify-center">
                  <Icon :icon="form.app_icon || 'mdi:view-dashboard'" class="h-6 w-6 text-black" />
                </div>
                <h1 class="text-xl font-bold text-white">{{ form.app_name || 'Airboard' }}</h1>
              </div>
              
              <!-- Dashboard Preview -->
              <div class="bg-gray-800 rounded-lg p-4">
                <h2 class="text-lg font-semibold text-white mb-2">{{ form.dashboard_title || $t('settings.previewDashboardTitle') }}</h2>
                <p class="text-gray-300 text-sm mb-4">{{ form.welcome_message || $t('settings.previewWelcome') }}</p>
                
                <!-- Sample App Cards -->
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                  <div class="bg-gray-700 rounded-lg p-4 hover:bg-gray-600 transition-colors">
              <div class="flex items-center gap-3 mb-2">
                      <div class="h-8 w-8 bg-blue-500 rounded-lg flex items-center justify-center">
                        <Icon icon="mdi:git" class="h-4 w-4 text-white" />
                      </div>
                      <div>
                        <h3 class="font-medium text-white text-sm">GitLab</h3>
                      </div>
                    </div>
                    <p class="text-xs text-gray-400">Git repository management</p>
                  </div>
                  
                  <div class="bg-gray-700 rounded-lg p-4 hover:bg-gray-600 transition-colors">
                    <div class="flex items-center gap-3 mb-2">
                      <div class="h-8 w-8 bg-green-500 rounded-lg flex items-center justify-center">
                        <Icon icon="mdi:email" class="h-4 w-4 text-white" />
                      </div>
                      <div>
                        <h3 class="font-medium text-white text-sm">Mail</h3>
                      </div>
                    </div>
                    <p class="text-xs text-gray-400">Email management</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- App Information -->
          <div>
            <div class="section-header">
              <Icon icon="mdi:information-outline" class="section-icon" />
              <h4 class="section-title">{{ $t('settings.appInformation') }}</h4>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="form-group">
                <label for="app_name" class="form-label form-label-required">{{ $t('settings.appName') }}</label>
                <input
                  id="app_name"
                  v-model="form.app_name"
                  type="text"
                  required
                  class="form-input"
                  :placeholder="$t('settings.appNamePlaceholder')"
                />
                <p v-if="errors.app_name" class="form-error">{{ errors.app_name }}</p>
                <p class="form-help">This name appears in the header and browser title</p>
              </div>

              <IconInput
                v-model="form.app_icon"
                label="{{ $t('settings.appIcon') }}"
                placeholder="mdi:view-dashboard, lucide:layout-dashboard"
                category="general"
                :show-suggestions="true"
                required
              />
            </div>
          </div>

          <!-- Dashboard Customization -->
          <div>
            <div class="section-header">
              <Icon icon="mdi:view-dashboard-outline" class="section-icon" />
              <h4 class="section-title">{{ $t('settings.dashboardCustomization') }}</h4>
            </div>
            
            <div class="space-y-4">
              <div class="form-group">
                <label for="dashboard_title" class="form-label form-label-required">{{ $t('settings.dashboardTitle') }}</label>
                <input
                  id="dashboard_title"
                  v-model="form.dashboard_title"
                  type="text"
                  required
                  class="form-input"
                  :placeholder="$t('settings.dashboardTitlePlaceholder')"
                />
                <p v-if="errors.dashboard_title" class="form-error">{{ errors.dashboard_title }}</p>
                <p class="form-help">Main title displayed on the dashboard page</p>
              </div>

              <div class="form-group">
                <label for="welcome_message" class="form-label form-label-required">{{ $t('settings.welcomeMessage') }}</label>
                <textarea
                  id="welcome_message"
                  v-model="form.welcome_message"
                  rows="3"
                  required
                  class="form-textarea"
                  :placeholder="$t('settings.welcomeMessagePlaceholder')"
                ></textarea>
                <p v-if="errors.welcome_message" class="form-error">{{ errors.welcome_message }}</p>
                <p class="form-help">Friendly message displayed below the dashboard title</p>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex justify-end space-x-4 pt-6 border-t border-gray-700">
            <button
              type="button"
              @click="loadSettings"
              class="btn btn-secondary"
            >
              <Icon icon="mdi:refresh" class="h-4 w-4 mr-2" />
              {{ $t('settings.reload') }}
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="btn btn-primary"
            >
              <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ $t('settings.save') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Password Change Modal -->
    <div v-if="showPasswordModal" class="modal-overlay">
      <div class="modal-container">
        <div class="modal-panel sm:max-w-md sm:w-full">
          <div class="modal-header">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0">
                <Icon icon="mdi:key" class="h-6 w-6 text-green-500" />
              </div>
              <div>
                <h3 class="modal-title">{{ $t('settings.changePassword') }}</h3>
                <p class="modal-subtitle">
                  {{ $t('settings.changePasswordSubtitle') }}
                </p>
              </div>
            </div>
          </div>
          
          <form @submit.prevent="handlePasswordChange" class="modal-content">
            <div class="space-y-4">
              <div class="form-group">
                <label for="old_password" class="form-label form-label-required">{{ $t('settings.currentPassword') }}</label>
                <input
                  id="old_password"
                  v-model="passwordForm.oldPassword"
                  type="password"
                  required
                  class="form-input"
                  :placeholder="$t('settings.currentPasswordPlaceholder')"
                />
                <p v-if="passwordErrors.oldPassword" class="form-error">{{ passwordErrors.oldPassword }}</p>
              </div>

              <div class="form-group">
                <label for="new_password" class="form-label form-label-required">{{ $t('settings.newPassword') }}</label>
                <input
                  id="new_password"
                  v-model="passwordForm.newPassword"
                  type="password"
                  required
                  class="form-input"
                  :placeholder="$t('settings.newPasswordPlaceholder')"
                  minlength="6"
                />
                <p v-if="passwordErrors.newPassword" class="form-error">{{ passwordErrors.newPassword }}</p>
                <p class="form-help">Minimum 6 characters</p>
              </div>

              <div class="form-group">
                <label for="confirm_password" class="form-label form-label-required">{{ $t('settings.confirmNewPassword') }}</label>
                <input
                  id="confirm_password"
                  v-model="passwordForm.confirmPassword"
                  type="password"
                  required
                  class="form-input"
                  :placeholder="$t('settings.confirmNewPasswordPlaceholder')"
                  minlength="6"
                />
                <p v-if="passwordErrors.confirmPassword" class="form-error">{{ passwordErrors.confirmPassword }}</p>
              </div>
            </div>

            <div class="modal-footer">
              <button
                type="button"
                @click="closePasswordModal"
                class="btn btn-secondary w-full sm:w-auto"
              >
                {{ $t('common.cancel') }}
              </button>
              <button
                type="submit"
                :disabled="passwordLoading"
                class="btn btn-primary w-full sm:w-auto"
              >
                <Icon v-if="passwordLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
                {{ $t('settings.changePassword') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Reset Confirmation Modal -->
    <div v-if="showResetModal" class="modal-overlay">
      <div class="modal-container">
        <div class="modal-panel sm:max-w-lg sm:w-full">
          <div class="modal-header">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0">
                <Icon icon="mdi:restore" class="h-6 w-6 text-yellow-500" />
              </div>
              <div>
                <h3 class="modal-title">{{ $t('settings.resetTitle') }}</h3>
                <p class="modal-subtitle">
                  {{ $t('settings.resetSubtitle') }}
                </p>
              </div>
            </div>
          </div>
          
          <div class="modal-footer">
            <button @click="closeResetModal" class="btn btn-secondary w-full sm:w-auto">
              {{ $t('common.cancel') }}
            </button>
            <button
              @click="confirmReset"
              :disabled="resetLoading"
              class="btn btn-danger w-full sm:w-auto"
            >
              <Icon v-if="resetLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ $t('settings.resetButton') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { adminService } from '@/services/api'
import { authService } from '@/services/api'
import { useAppStore } from '@/stores/app'
import IconInput from '@/components/ui/IconInput.vue'

const appStore = useAppStore()

// State
const loading = ref(false)
const resetLoading = ref(false)
const showResetModal = ref(false)
const showPasswordModal = ref(false)
const passwordLoading = ref(false)
const passwordErrors = ref({})
const errors = ref({})

const form = reactive({
  app_name: '',
  app_icon: '',
  dashboard_title: '',
  welcome_message: ''
})

const passwordForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

// Methods
const loadSettings = async () => {
  try {
    appStore.setLoading(true)
    const data = await adminService.getAppSettings()
    console.log('Settings loaded:', data)
    
    Object.assign(form, {
      app_name: data.app_name || 'Airboard',
      app_icon: data.app_icon || 'mdi:view-dashboard',
      dashboard_title: data.dashboard_title || 'Dashboard',
      welcome_message: data.welcome_message || 'Welcome to your application portal'
    })
  } catch (error) {
    console.error('Error loading settings:', error)
    appStore.showError('Failed to load settings')
  } finally {
    appStore.setLoading(false)
  }
}

const validateForm = () => {
  errors.value = {}
  
  if (!form.app_name.trim()) {
    errors.value.app_name = 'Application name is required'
  }
  
  if (!form.app_icon.trim()) {
    errors.value.app_icon = 'Application icon is required'
  }
  
  if (!form.dashboard_title.trim()) {
    errors.value.dashboard_title = 'Dashboard title is required'
  }
  
  if (!form.welcome_message.trim()) {
    errors.value.welcome_message = 'Welcome message is required'
  }
  
  return Object.keys(errors.value).length === 0
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  try {
    await adminService.updateAppSettings(form)
    appStore.showSuccess('Settings updated successfully')
    // 🔧 FIX: Recharger les settings et mettre à jour le store global
    await loadSettings()
    appStore.invalidateSettings()
    // Optionnel: rafraîchir les settings dans le store global
    try {
      await appStore.refreshAppSettings()
    } catch (e) {
      console.warn('Could not refresh global settings:', e)
    }
  } catch (error) {
    console.error('Error updating settings:', error)
    appStore.showError('Failed to update settings')
  } finally {
    loading.value = false
  }
}

const resetSettings = () => {
  showResetModal.value = true
}

const closeResetModal = () => {
  showResetModal.value = false
}

const confirmReset = async () => {
  resetLoading.value = true
  try {
    await adminService.resetAppSettings()
    // 🔧 FIX: Recharger les settings et mettre à jour le store global
    await loadSettings()
    appStore.invalidateSettings()
    // Optionnel: rafraîchir les settings dans le store global
    try {
      await appStore.refreshAppSettings()
    } catch (e) {
      console.warn('Could not refresh global settings:', e)
    }
    appStore.showSuccess('Settings reset to defaults successfully')
    closeResetModal()
  } catch (error) {
    console.error('Error resetting settings:', error)
    appStore.showError('Failed to reset settings')
  } finally {
    resetLoading.value = false
  }
}

const validatePasswordForm = () => {
  passwordErrors.value = {}
  
  if (!passwordForm.oldPassword.trim()) {
    passwordErrors.value.oldPassword = 'Current password is required'
  }
  
  if (!passwordForm.newPassword.trim()) {
    passwordErrors.value.newPassword = 'New password is required'
  } else if (passwordForm.newPassword.length < 6) {
    passwordErrors.value.newPassword = 'Password must be at least 6 characters'
  }
  
  if (!passwordForm.confirmPassword.trim()) {
    passwordErrors.value.confirmPassword = 'Please confirm your password'
  } else if (passwordForm.newPassword !== passwordForm.confirmPassword) {
    passwordErrors.value.confirmPassword = 'Passwords do not match'
  }
  
  return Object.keys(passwordErrors.value).length === 0
}

const handlePasswordChange = async () => {
  if (!validatePasswordForm()) return
  
  passwordLoading.value = true
  try {
    await authService.changePassword(passwordForm.oldPassword, passwordForm.newPassword)
    appStore.showSuccess('Password changed successfully')
    closePasswordModal()
  } catch (error) {
    console.error('Error changing password:', error)
    if (error.response?.status === 401) {
      passwordErrors.value.oldPassword = 'Incorrect current password'
    } else {
      appStore.showError('Failed to change password')
    }
  } finally {
    passwordLoading.value = false
  }
}

const closePasswordModal = () => {
  showPasswordModal.value = false
  Object.assign(passwordForm, {
    oldPassword: '',
    newPassword: '',
    confirmPassword: ''
  })
  passwordErrors.value = {}
}

// Lifecycle
onMounted(() => {
  loadSettings()
})
</script>