<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ $t('oauth.title') }}</h1>
          <p class="page-subtitle">{{ $t('oauth.subtitle') }}</p>
        </div>
      </div>
    </div>

    <!-- OAuth Providers List -->
    <div class="max-w-4xl space-y-6">
      <div v-for="provider in providers" :key="provider.id" class="card">
        <div class="flex items-start justify-between mb-6">
          <div class="flex items-center gap-4">
            <div class="h-12 w-12 rounded-lg flex items-center justify-center"
                 :class="provider.is_enabled ? 'bg-green-500' : 'bg-gray-700'">
              <Icon :icon="provider.icon" class="h-6 w-6 text-white" />
            </div>
            <div>
              <h3 class="text-lg font-semibold text-white">{{ provider.display_name }}</h3>
              <p class="text-sm text-gray-400">{{ provider.provider_name }}</p>
            </div>
          </div>
          <div class="flex items-center gap-3">
            <span :class="provider.is_enabled ? 'badge-success' : 'badge-secondary'">
              {{ provider.is_enabled ? $t('common.enabled') : $t('common.disabled') }}
            </span>
            <button @click="editProvider(provider)" class="btn btn-secondary btn-sm">
              <Icon icon="mdi:pencil" class="h-4 w-4 mr-2" />
              {{ $t('common.edit') }}
            </button>
          </div>
        </div>

        <!-- Provider Info (read-only preview) -->
        <div v-if="provider.is_enabled" class="grid grid-cols-1 md:grid-cols-2 gap-4 pt-4 border-t border-gray-700">
          <div>
            <label class="block text-xs text-gray-400 mb-1">{{ $t('oauth.clientId') }}</label>
            <p class="text-sm text-gray-300 font-mono">{{ provider.client_id || $t('common.notConfigured') }}</p>
          </div>
          <div>
            <label class="block text-xs text-gray-400 mb-1">{{ $t('oauth.redirectUri') }}</label>
            <p class="text-sm text-gray-300 font-mono truncate">{{ provider.redirect_uri }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Provider Modal -->
    <div v-if="showEditModal && selectedProvider" class="modal-overlay">
      <div class="modal-container">
        <div class="modal-panel sm:max-w-2xl sm:w-full">
          <div class="modal-header">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0">
                <Icon :icon="selectedProvider.icon" class="h-6 w-6 text-blue-500" />
              </div>
              <div>
                <h3 class="modal-title">{{ $t('oauth.configure') }} {{ selectedProvider.display_name }}</h3>
                <p class="modal-subtitle">{{ $t('oauth.configureSubtitle') }}</p>
              </div>
            </div>
          </div>

          <form @submit.prevent="handleSave" class="modal-content">
            <div class="space-y-6">
              <!-- Enable/Disable Toggle -->
              <div class="form-group">
                <label class="flex items-center gap-3 cursor-pointer">
                  <input
                    v-model="editForm.is_enabled"
                    type="checkbox"
                    class="h-5 w-5 text-green-600 focus:ring-green-500 border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700"
                  />
                  <span class="text-white font-medium">{{ $t('oauth.enableProvider') }}</span>
                </label>
                <p class="form-help mt-2">{{ $t('oauth.enableProviderHelp') }}</p>
              </div>

              <!-- Client ID -->
              <div class="form-group">
                <label for="client_id" class="form-label form-label-required">{{ $t('oauth.clientId') }}</label>
                <input
                  id="client_id"
                  v-model="editForm.client_id"
                  type="text"
                  required
                  class="form-input font-mono"
                  :placeholder="getClientIdPlaceholder(selectedProvider.provider_name)"
                />
                <p class="form-help">{{ $t('oauth.clientIdHelp') }}</p>
              </div>

              <!-- Client Secret -->
              <div class="form-group">
                <label for="client_secret" class="form-label">{{ $t('oauth.clientSecret') }}</label>
                <div class="relative">
                  <input
                    id="client_secret"
                    v-model="editForm.client_secret"
                    :type="showSecret ? 'text' : 'password'"
                    class="form-input font-mono pr-10"
                    :placeholder="$t('oauth.clientSecretPlaceholder')"
                  />
                  <button
                    type="button"
                    @click="showSecret = !showSecret"
                    class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-300"
                  >
                    <Icon :icon="showSecret ? 'mdi:eye-off' : 'mdi:eye'" class="h-5 w-5" />
                  </button>
                </div>
                <p class="form-help">{{ $t('oauth.clientSecretHelp') }}</p>
              </div>

              <!-- Redirect URI (read-only) -->
              <div class="form-group">
                <label for="redirect_uri" class="form-label">{{ $t('oauth.redirectUri') }}</label>
                <div class="flex gap-2">
                  <input
                    id="redirect_uri"
                    v-model="editForm.redirect_uri"
                    type="text"
                    readonly
                    class="form-input font-mono bg-gray-700 cursor-not-allowed"
                  />
                  <button
                    type="button"
                    @click="copyToClipboard(editForm.redirect_uri)"
                    class="btn btn-secondary"
                  >
                    <Icon icon="mdi:content-copy" class="h-4 w-4" />
                  </button>
                </div>
                <p class="form-help">{{ $t('oauth.redirectUriHelp') }}</p>
              </div>

              <!-- Advanced Settings (collapsible) -->
              <div class="border-t border-gray-700 pt-4">
                <button
                  type="button"
                  @click="showAdvanced = !showAdvanced"
                  class="flex items-center gap-2 text-sm text-gray-400 hover:text-white transition-colors"
                >
                  <Icon :icon="showAdvanced ? 'mdi:chevron-down' : 'mdi:chevron-right'" class="h-5 w-5" />
                  {{ $t('oauth.advancedSettings') }}
                </button>

                <div v-if="showAdvanced" class="mt-4 space-y-4">
                  <div class="form-group">
                    <label for="auth_url" class="form-label">{{ $t('oauth.authUrl') }}</label>
                    <input
                      id="auth_url"
                      v-model="editForm.auth_url"
                      type="url"
                      class="form-input font-mono text-sm"
                    />
                  </div>

                  <div class="form-group">
                    <label for="token_url" class="form-label">{{ $t('oauth.tokenUrl') }}</label>
                    <input
                      id="token_url"
                      v-model="editForm.token_url"
                      type="url"
                      class="form-input font-mono text-sm"
                    />
                  </div>

                  <div class="form-group">
                    <label for="user_info_url" class="form-label">{{ $t('oauth.userInfoUrl') }}</label>
                    <input
                      id="user_info_url"
                      v-model="editForm.user_info_url"
                      type="url"
                      class="form-input font-mono text-sm"
                    />
                  </div>

                  <div class="form-group">
                    <label for="scopes" class="form-label">{{ $t('oauth.scopes') }}</label>
                    <input
                      id="scopes"
                      v-model="editForm.scopes"
                      type="text"
                      class="form-input font-mono text-sm"
                      placeholder="openid email profile"
                    />
                  </div>
                </div>
              </div>

              <!-- Setup Instructions -->
              <div class="bg-blue-50 dark:bg-blue-900/20 border-2 border-blue-300 dark:border-blue-500/30 rounded-lg p-4">
                <div class="flex items-start gap-3">
                  <Icon icon="mdi:information" class="h-5 w-5 text-blue-600 dark:text-blue-400 mt-0.5 flex-shrink-0" />
                  <div class="text-sm text-gray-700 dark:text-gray-300">
                    <p class="font-medium text-gray-900 dark:text-white mb-2">{{ $t('oauth.setupInstructions') }}</p>
                    <div v-if="selectedProvider.provider_name === 'google'">
                      <p class="mb-2">{{ $t('oauth.googleSetup1') }}</p>
                      <ol class="list-decimal list-inside space-y-1 text-xs">
                        <li>{{ $t('oauth.googleSetup2') }}</li>
                        <li>{{ $t('oauth.googleSetup3') }}</li>
                        <li>{{ $t('oauth.googleSetup4') }}</li>
                        <li>{{ $t('oauth.googleSetup5') }}</li>
                      </ol>
                    </div>
                    <div v-else-if="selectedProvider.provider_name === 'microsoft'">
                      <p class="mb-2">{{ $t('oauth.microsoftSetup1') }}</p>
                      <ol class="list-decimal list-inside space-y-1 text-xs">
                        <li>{{ $t('oauth.microsoftSetup2') }}</li>
                        <li>{{ $t('oauth.microsoftSetup3') }}</li>
                        <li>{{ $t('oauth.microsoftSetup4') }}</li>
                        <li>{{ $t('oauth.microsoftSetup5') }}</li>
                      </ol>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="modal-footer">
              <button
                type="button"
                @click="closeEditModal"
                class="btn btn-secondary w-full sm:w-auto"
              >
                {{ $t('common.cancel') }}
              </button>
              <button
                type="submit"
                :disabled="loading"
                class="btn btn-primary w-full sm:w-auto"
              >
                <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
                {{ $t('common.save') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { adminService } from '@/services/api'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()

// State
const loading = ref(false)
const providers = ref([])
const showEditModal = ref(false)
const selectedProvider = ref(null)
const showSecret = ref(false)
const showAdvanced = ref(false)

const editForm = reactive({
  provider_name: '',
  display_name: '',
  icon: '',
  is_enabled: false,
  client_id: '',
  client_secret: '',
  redirect_uri: '',
  auth_url: '',
  token_url: '',
  user_info_url: '',
  scopes: ''
})

// Methods
const loadProviders = async () => {
  try {
    appStore.setLoading(true)
    const data = await adminService.getOAuthProviders()
    providers.value = data.providers || []
  } catch (error) {
    console.error('Error loading OAuth providers:', error)
    appStore.showError('Failed to load OAuth providers')
  } finally {
    appStore.setLoading(false)
  }
}

const editProvider = (provider) => {
  selectedProvider.value = provider
  Object.assign(editForm, {
    provider_name: provider.provider_name,
    display_name: provider.display_name,
    icon: provider.icon,
    is_enabled: provider.is_enabled,
    client_id: provider.client_id || '',
    client_secret: '',
    redirect_uri: provider.redirect_uri,
    auth_url: provider.auth_url,
    token_url: provider.token_url,
    user_info_url: provider.user_info_url,
    scopes: provider.scopes
  })
  showEditModal.value = true
  showSecret.value = false
  showAdvanced.value = false
}

const handleSave = async () => {
  loading.value = true
  try {
    await adminService.updateOAuthProvider(selectedProvider.value.id, editForm)
    appStore.showSuccess('OAuth provider updated successfully')
    await loadProviders()
    closeEditModal()
  } catch (error) {
    console.error('Error updating OAuth provider:', error)
    appStore.showError('Failed to update OAuth provider')
  } finally {
    loading.value = false
  }
}

const closeEditModal = () => {
  showEditModal.value = false
  selectedProvider.value = null
  showSecret.value = false
  showAdvanced.value = false
}

const copyToClipboard = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    appStore.showSuccess('Copied to clipboard')
  } catch (error) {
    appStore.showError('Failed to copy to clipboard')
  }
}

const getClientIdPlaceholder = (providerName) => {
  switch (providerName) {
    case 'google':
      return '123456789-abc123.apps.googleusercontent.com'
    case 'microsoft':
      return 'a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6'
    default:
      return 'Enter your client ID'
  }
}

// Lifecycle
onMounted(() => {
  loadProviders()
})
</script>
