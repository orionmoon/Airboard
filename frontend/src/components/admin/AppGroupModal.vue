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
                        <div class="font-medium text-white">{{ form.name || 'Group Name' }}</div>
                        <div class="text-sm text-gray-400">{{ form.description || 'Description...' }}</div>
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
                    class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-600 rounded bg-gray-700"
                  />
                  <span class="text-sm text-gray-300">
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

const loading = ref(false)
const errors = ref({})

const form = reactive({
  name: '',
  description: '',
  color: '#10b981',
  icon: 'mdi:folder',
  order: 0,
  is_active: true
})

const isEdit = computed(() => !!props.appGroup)

// Reset form when modal opens/closes
watch(() => props.show, (newVal) => {
  if (newVal) {
    resetForm()
  }
})

watch(() => props.appGroup, (newVal) => {
  if (newVal) {
    Object.assign(form, {
      name: newVal.name || '',
      description: newVal.description || '',
      color: newVal.color || '#10b981',
      icon: newVal.icon || 'mdi:folder',
      order: newVal.order || 0,
      is_active: newVal.is_active ?? true
    })
  }
}, { immediate: true })

const resetForm = () => {
  if (props.appGroup) {
    Object.assign(form, {
      name: props.appGroup.name || '',
      description: props.appGroup.description || '',
      color: props.appGroup.color || '#10b981',
      icon: props.appGroup.icon || 'mdi:folder',
      order: props.appGroup.order || 0,
      is_active: props.appGroup.is_active ?? true
    })
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      color: '#10b981',
      icon: 'mdi:folder',
      order: 0,
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