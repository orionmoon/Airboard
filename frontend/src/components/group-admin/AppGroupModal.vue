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
                  {{ isEdit ? 'Modifier le groupe' : 'Nouveau groupe' }}
                </h3>
                <p class="modal-subtitle">
                  {{ isEdit ? 'Modifier les paramètres du groupe d\'applications' : 'Créer un nouveau groupe pour organiser les applications' }}
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
            <!-- Owner Group Selection -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:shield-account" class="section-icon" />
                <h4 class="section-title">Groupe propriétaire</h4>
              </div>

              <div class="form-group">
                <label for="owner_group_id" class="form-label form-label-required">
                  Groupe
                </label>
                <select
                  id="owner_group_id"
                  v-model="form.owner_group_id"
                  required
                  class="form-select"
                  :disabled="isEdit"
                >
                  <option value="">Sélectionner un groupe</option>
                  <option
                    v-for="group in managedGroups"
                    :key="group.id"
                    :value="group.id"
                  >
                    {{ group.name }}
                  </option>
                </select>
                <p class="text-xs text-gray-400 mt-1">
                  Ce groupe d'applications sera privé et accessible uniquement par les membres de ce groupe.
                  {{ isEdit ? 'Le groupe propriétaire ne peut pas être modifié.' : '' }}
                </p>
                <p v-if="errors.owner_group_id" class="form-error">{{ errors.owner_group_id }}</p>
              </div>
            </div>

            <!-- Basic Information -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:information-outline" class="section-icon" />
                <h4 class="section-title">Informations de base</h4>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
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
                    placeholder="Outils de développement"
                  />
                  <p v-if="errors.name" class="form-error">{{ errors.name }}</p>
                </div>

                <div class="form-group">
                  <label for="order" class="form-label">
                    Ordre d'affichage
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
                  placeholder="Décrivez ce groupe d'applications..."
                ></textarea>
              </div>
            </div>

            <!-- Appearance -->
            <div>
              <div class="section-header">
                <Icon icon="mdi:palette-outline" class="section-icon" />
                <h4 class="section-title">Apparence</h4>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Preview -->
                <div class="form-group">
                  <label class="form-label">Aperçu</label>
                  <div class="card p-4">
                    <div class="flex items-center space-x-3">
                      <div
                        class="h-12 w-12 rounded-lg flex items-center justify-center"
                        :style="{ backgroundColor: form.color }"
                      >
                        <Icon :icon="form.icon || 'mdi:folder'" class="h-6 w-6 text-white" />
                      </div>
                      <div>
                        <div class="font-medium text-white">{{ form.name || 'Nom du groupe' }}</div>
                        <div class="text-sm text-gray-400">{{ form.description || 'Description...' }}</div>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Color and Icon -->
                <div class="space-y-4">
                  <div class="form-group">
                    <label for="color" class="form-label">Couleur</label>
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
                    label="Icône"
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
                    Actif et visible pour les utilisateurs
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
              Annuler
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="btn-primary w-full sm:w-auto"
            >
              <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              {{ isEdit ? 'Mettre à jour' : 'Créer' }}
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
  },
  managedGroups: {
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
  color: '#10b981',
  icon: 'mdi:folder',
  order: 0,
  is_active: true,
  owner_group_id: '',
  is_private: true // Toujours privé pour les group admins
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
      is_active: newVal.is_active ?? true,
      owner_group_id: newVal.owner_group_id || '',
      is_private: true
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
      is_active: props.appGroup.is_active ?? true,
      owner_group_id: props.appGroup.owner_group_id || '',
      is_private: true
    })
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      color: '#10b981',
      icon: 'mdi:folder',
      order: 0,
      is_active: true,
      owner_group_id: props.managedGroups.length === 1 ? props.managedGroups[0].id : '',
      is_private: true
    })
  }
  errors.value = {}
}

const validateForm = () => {
  errors.value = {}

  if (!form.name.trim()) {
    errors.value.name = 'Le nom est requis'
  }

  if (!form.owner_group_id) {
    errors.value.owner_group_id = 'Le groupe propriétaire est requis'
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
