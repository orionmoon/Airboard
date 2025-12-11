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
                  {{ isEdit ? 'Edit App Group' : 'New App Group' }}
                </h3>
                <p class="modal-subtitle">
                  {{ isEdit ? 'Modify the application group settings' : 'Create a new group to organize applications' }}
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
                    Group Name
                  </label>
                  <input
                    id="name"
                    v-model="form.name"
                    type="text"
                    required
                    class="form-input"
                    placeholder="Development Tools"
                  />
                  <p v-if="errors.name" class="form-error">{{ errors.name }}</p>
                </div>

                <div class="form-group">
                  <label for="order" class="form-label">
                    Display Order
                  </label>
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

              <div class="form-group">
                <label for="description" class="form-label">
                  Description
                </label>
                <textarea
                  id="description"
                  v-model="form.description"
                  rows="3"
                  class="form-textarea"
                  placeholder="Describe this application group..."
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
                    <div class="flex items-center space-x-3">
                      <div 
                        class="h-12 w-12 rounded-lg flex items-center justify-center"
                        :style="{ backgroundColor: form.color }"
                      >
                        <Icon :icon="form.icon || 'mdi:folder'" class="h-6 w-6 text-white" />
                      </div>
                      <div>
                        <div class="font-medium text-gray-900 dark:text-white">{{ form.name || 'Group Name' }}</div>
                        <div class="text-sm text-gray-600 dark:text-gray-400">{{ form.description || 'Description...' }}</div>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Color and Icon -->
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
                        placeholder="#10b981"
                      />
                    </div>
                  </div>

                  <!-- Icon Input with free text -->
                  <IconInput
                    v-model="form.icon"
                    label="Icon"
                    placeholder="mdi:folder, lucide:folder, heroicons:folder"
                    category="group"
                    :show-suggestions="true"
                  />
                </div>
              </div>
            </div>

            <!-- Access Control (Admin only) -->
            <div v-if="userRole === 'admin'">
              <div class="section-header">
                <Icon icon="mdi:shield-lock-outline" class="section-icon" />
                <h4 class="section-title">Access Control</h4>
              </div>

              <div class="card p-4 space-y-4">
                <div>
                  <label class="flex items-center space-x-3 cursor-pointer">
                    <input
                      v-model="form.is_private"
                      type="checkbox"
                      class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700"
                    />
                    <div>
                      <span class="text-sm text-gray-700 dark:text-gray-300 font-medium">Private App Group</span>
                      <p class="text-xs text-gray-600 dark:text-gray-400 mt-1">
                        Private app groups can only be managed by their owner group admin
                      </p>
                    </div>
                  </label>
                </div>

                <div v-if="form.is_private" class="form-group">
                  <label for="owner_group_id" class="form-label form-label-required">
                    Owner Group
                  </label>
                  <select
                    id="owner_group_id"
                    v-model.number="form.owner_group_id"
                    required
                    class="form-select"
                  >
                    <option :value="null" disabled>Select owner group...</option>
                    <option v-for="group in groups" :key="group.id" :value="group.id">
                      {{ group.name }}
                    </option>
                  </select>
                  <p v-if="errors.owner_group_id" class="form-error">{{ errors.owner_group_id }}</p>
                  <p v-else class="text-xs text-gray-600 dark:text-gray-400 mt-1">
                    The group admin of this group will be able to manage this app group
                  </p>
                </div>
              </div>
            </div>

            <!-- Options -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:tune" class="section-icon" />
                <h4 class="section-title">Options</h4>
              </div>

              <div class="card p-4">
                <label class="flex items-center space-x-3 cursor-pointer">
                  <input
                    v-model="form.is_active"
                    type="checkbox"
                    class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700"
                  />
                  <span class="text-sm text-gray-700 dark:text-gray-300">
                    Active and visible to users
                  </span>
                </label>
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
              :disabled="loading"
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
import { useAuthStore } from '@/stores/auth'
import { adminService } from '@/services/api'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  appGroup: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['close', 'submit'])

const authStore = useAuthStore()
const userRole = computed(() => authStore.user?.role)

const loading = ref(false)
const errors = ref({})
const groups = ref([])

const form = reactive({
  name: '',
  description: '',
  color: '#10b981',
  icon: 'mdi:folder',
  order: 0,
  is_active: true,
  is_private: false,
  owner_group_id: null
})

const isEdit = computed(() => !!props.appGroup)

// Load groups for owner selection (admin only)
const loadGroups = async () => {
  if (userRole.value === 'admin') {
    try {
      const data = await adminService.getGroups()
      groups.value = data || []
    } catch (error) {
      console.error('Failed to load groups:', error)
    }
  }
}

// Reset form when modal opens/closes
watch(() => props.show, async (newVal) => {
  if (newVal) {
    await loadGroups() // Load groups first before resetting form
    resetForm()
  }
})

watch(() => props.appGroup, (newVal) => {
  if (newVal && props.show) {
    // Ensure owner_group_id is properly typed (number or null)
    const ownerGroupId = newVal.owner_group_id ? Number(newVal.owner_group_id) : null

    Object.assign(form, {
      name: newVal.name || '',
      description: newVal.description || '',
      color: newVal.color || '#10b981',
      icon: newVal.icon || 'mdi:folder',
      order: newVal.order || 0,
      is_active: newVal.is_active ?? true,
      is_private: newVal.is_private ?? false,
      owner_group_id: ownerGroupId
    })
  }
}, { immediate: true })

const resetForm = () => {
  if (props.appGroup) {
    // Ensure owner_group_id is properly typed (number or null)
    const ownerGroupId = props.appGroup.owner_group_id ? Number(props.appGroup.owner_group_id) : null

    Object.assign(form, {
      name: props.appGroup.name || '',
      description: props.appGroup.description || '',
      color: props.appGroup.color || '#10b981',
      icon: props.appGroup.icon || 'mdi:folder',
      order: props.appGroup.order || 0,
      is_active: props.appGroup.is_active ?? true,
      is_private: props.appGroup.is_private ?? false,
      owner_group_id: ownerGroupId
    })
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      color: '#10b981',
      icon: 'mdi:folder',
      order: 0,
      is_active: true,
      is_private: false,
      owner_group_id: null
    })
  }
  errors.value = {}
}

const validateForm = () => {
  errors.value = {}

  if (!form.name.trim()) {
    errors.value.name = 'Name is required'
  }

  // Validate owner group for private app groups (admin only)
  if (userRole.value === 'admin' && form.is_private && !form.owner_group_id) {
    errors.value.owner_group_id = 'Owner group is required for private app groups'
  }

  return Object.keys(errors.value).length === 0
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  try {
    const formData = { ...form }
    if (isEdit.value) {
      formData.id = props.appGroup.id
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