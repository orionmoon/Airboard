<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-container">
      <div class="modal-panel sm:max-w-2xl sm:w-full">
        <form @submit.prevent="handleSubmit">
          <!-- Header -->
          <div class="modal-header">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="modal-title">
                  {{ isEdit ? 'Edit Application' : 'New Application' }}
                </h3>
                <p class="modal-subtitle">
                  {{ isEdit ? 'Modify application settings' : 'Add a new application to the portal' }}
                </p>
              </div>
              <button
                type="button"
                @click="closeModal"
                class="btn-ghost p-2"
              >
                <Icon icon="mdi:close" class="h-5 w-5" />
              </button>
            </div>
          </div>

          <!-- Content -->
          <div class="modal-content space-y-6">
            <!-- Basic Information -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:information-outline" class="section-icon" />
                <h4 class="section-title">Basic Information</h4>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="form-group">
                  <label for="name" class="form-label form-label-required">
                    Application Name
                  </label>
                  <input
                    id="name"
                    v-model="form.name"
                    type="text"
                    required
                    class="form-input"
                    placeholder="GitLab"
                  />
                  <p v-if="errors.name" class="form-error">{{ errors.name }}</p>
                </div>

                <div class="form-group">
                  <label for="app_group_id" class="form-label form-label-required">
                    App Group
                  </label>
                  <select
                    id="app_group_id"
                    v-model="form.app_group_id"
                    required
                    class="form-select"
                  >
                    <option value="">Select a group</option>
                    <option
                      v-for="group in availableAppGroups"
                      :key="group.id"
                      :value="group.id"
                    >
                      {{ group.name }}
                    </option>
                  </select>
                  <p v-if="errors.app_group_id" class="form-error">{{ errors.app_group_id }}</p>
                  <p v-if="availableAppGroups.length === 0" class="form-help text-yellow-500">
                    No private app groups available. Create a private app group first.
                  </p>
                </div>
              </div>

              <div class="form-group">
                <label for="url" class="form-label form-label-required">
                  Application URL
                </label>
                <input
                  id="url"
                  v-model="form.url"
                  type="url"
                  required
                  class="form-input"
                  placeholder="https://gitlab.example.com"
                />
                <p v-if="errors.url" class="form-error">{{ errors.url }}</p>
                <p class="form-help">Complete URL to the application (with https://)</p>
              </div>

              <div class="form-group">
                <label for="description" class="form-label">
                  Description
                </label>
                <textarea
                  id="description"
                  v-model="form.description"
                  rows="2"
                  class="form-textarea"
                  placeholder="Application description..."
                ></textarea>
              </div>
            </div>

            <!-- Appearance -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:palette-outline" class="section-icon" />
                <h4 class="section-title">Appearance</h4>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Preview -->
                <div class="form-group">
                  <label class="form-label">Preview</label>
                  <div class="card p-4">
                    <div class="flex items-center space-x-4">
                      <div
                        class="h-12 w-12 rounded-lg flex items-center justify-center"
                        :style="{ backgroundColor: form.color }"
                      >
                        <Icon :icon="form.icon || 'mdi:application'" class="h-6 w-6 text-white" />
                      </div>
                      <div class="flex-1">
                        <div class="font-medium text-white">{{ form.name || 'Application Name' }}</div>
                        <div class="text-sm text-gray-400">{{ form.url || 'URL...' }}</div>
                      </div>
                      <Icon v-if="form.open_in_new_tab" icon="mdi:open-in-new" class="h-4 w-4 text-gray-400" />
                    </div>
                  </div>
                </div>

                <!-- Color, Icon and Order -->
                <div class="space-y-4">
                  <div class="form-group">
                    <label for="color" class="form-label">Color</label>
                    <div class="flex items-center space-x-3">
                      <input
                        id="color"
                        v-model="form.color"
                        type="color"
                        class="h-10 w-16 rounded-lg border border-gray-600 cursor-pointer bg-gray-800"
                      />
                      <input
                        v-model="form.color"
                        type="text"
                        class="form-input flex-1"
                        placeholder="#6366f1"
                      />
                    </div>
                  </div>

                  <!-- Icon Input -->
                  <IconInput
                    v-model="form.icon"
                    label="Icon"
                    placeholder="mdi:application, lucide:app-window, heroicons:squares-2x2"
                    category="application"
                    :show-suggestions="true"
                  />

                  <div class="form-group">
                    <label for="order" class="form-label">Display Order</label>
                    <input
                      id="order"
                      v-model.number="form.order"
                      type="number"
                      min="0"
                      class="form-input"
                      placeholder="0"
                    />
                  </div>
                </div>
              </div>
            </div>

            <!-- Behavior Options -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:tune" class="section-icon" />
                <h4 class="section-title">Behavior Options</h4>
              </div>

              <div class="space-y-3">
                <div class="card p-4">
                  <label class="flex items-center space-x-3 cursor-pointer">
                    <input
                      v-model="form.open_in_new_tab"
                      type="checkbox"
                      class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-600 rounded bg-gray-700"
                    />
                    <span class="text-sm text-gray-300">
                      Open in new tab
                    </span>
                  </label>
                </div>

                <div class="card p-4">
                  <label class="flex items-center space-x-3 cursor-pointer">
                    <input
                      v-model="form.is_active"
                      type="checkbox"
                      class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-600 rounded bg-gray-700"
                    />
                    <span class="text-sm text-gray-300">
                      Active and visible to users
                    </span>
                  </label>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="modal-footer">
            <button
              type="button"
              @click="closeModal"
              class="btn-secondary w-full sm:w-auto"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="loading || availableAppGroups.length === 0"
              class="btn-primary w-full sm:w-auto"
            >
              <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ isEdit ? 'Update' : 'Create' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch } from 'vue'
import { Icon } from '@iconify/vue'
import IconInput from '@/components/ui/IconInput.vue'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  application: {
    type: Object,
    default: null
  },
  appGroups: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close', 'submit'])

const loading = ref(false)
const errors = ref({})

const form = reactive({
  name: '',
  description: '',
  url: '',
  app_group_id: '',
  color: '#6366f1',
  icon: 'mdi:application',
  order: 0,
  open_in_new_tab: true,
  is_active: true
})

const isEdit = computed(() => !!props.application)

// Filter to show only private app groups for group admins
const availableAppGroups = computed(() => {
  return props.appGroups.filter(group => group.is_private === true)
})

// Reset form when modal opens/closes
watch(() => props.show, (newVal) => {
  if (newVal) {
    resetForm()
  }
})

watch(() => props.application, (newVal) => {
  if (newVal) {
    Object.assign(form, {
      name: newVal.name || '',
      description: newVal.description || '',
      url: newVal.url || '',
      app_group_id: newVal.app_group_id || '',
      color: newVal.color || '#6366f1',
      icon: newVal.icon || 'mdi:application',
      order: newVal.order || 0,
      open_in_new_tab: newVal.open_in_new_tab ?? true,
      is_active: newVal.is_active ?? true
    })
  }
}, { immediate: true })

const resetForm = () => {
  if (props.application) {
    Object.assign(form, {
      name: props.application.name || '',
      description: props.application.description || '',
      url: props.application.url || '',
      app_group_id: props.application.app_group_id || '',
      color: props.application.color || '#6366f1',
      icon: props.application.icon || 'mdi:application',
      order: props.application.order || 0,
      open_in_new_tab: props.application.open_in_new_tab ?? true,
      is_active: props.application.is_active ?? true
    })
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      url: '',
      app_group_id: '',
      color: '#6366f1',
      icon: 'mdi:application',
      order: 0,
      open_in_new_tab: true,
      is_active: true
    })
  }
  errors.value = {}
}

const validateForm = () => {
  errors.value = {}

  if (!form.name.trim()) {
    errors.value.name = 'Name is required'
  }

  if (!form.url.trim()) {
    errors.value.url = 'URL is required'
  } else if (!isValidUrl(form.url)) {
    errors.value.url = 'Invalid URL format'
  }

  if (!form.app_group_id) {
    errors.value.app_group_id = 'App group is required'
  }

  return Object.keys(errors.value).length === 0
}

const isValidUrl = (url) => {
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

const handleSubmit = async () => {
  if (!validateForm()) return

  loading.value = true
  try {
    const formData = { ...form }
    if (isEdit.value) {
      formData.id = props.application.id
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
