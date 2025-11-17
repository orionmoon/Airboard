<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <div class="w-full">
      <!-- Header -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
            {{ $t('announcements.title') }}
          </h1>
          <p class="mt-1 text-gray-600 dark:text-gray-400">
            {{ $t('announcements.subtitle') }}
          </p>
        </div>
        <button
          @click="openCreateModal"
          class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors flex items-center gap-2"
        >
          <Icon icon="mdi:plus" class="h-5 w-5" />
          {{ $t('announcements.new') }}
        </button>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="flex justify-center items-center py-12">
        <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-blue-600" />
      </div>

      <!-- Announcements List -->
      <div v-else-if="announcements.length > 0" class="space-y-4">
        <div
          v-for="announcement in announcements"
          :key="announcement.id"
          class="card p-6"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-3 mb-2">
                <span
                  :class="getTypeBadgeClass(announcement.type)"
                  class="px-2 py-1 rounded-md text-xs font-medium"
                >
                  {{ $t(`announcements.type_${announcement.type}`) }}
                </span>
                <span
                  :class="announcement.is_active ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200' : 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300'"
                  class="px-2 py-1 rounded-md text-xs font-medium"
                >
                  {{ announcement.is_active ? $t('common.active') : $t('common.inactive') }}
                </span>
                <span v-if="announcement.priority > 0" class="text-xs text-gray-500 dark:text-gray-400">
                  {{ $t('announcements.priority') }}: {{ announcement.priority }}
                </span>
              </div>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">
                {{ announcement.title }}
              </h3>
              <p class="text-sm text-gray-600 dark:text-gray-400 whitespace-pre-wrap">
                {{ announcement.content }}
              </p>
              <div class="flex items-center gap-4 mt-3 text-xs text-gray-500 dark:text-gray-400">
                <span v-if="announcement.start_date">
                  {{ $t('announcements.startDate') }}: {{ formatDate(announcement.start_date) }}
                </span>
                <span v-if="announcement.end_date">
                  {{ $t('announcements.endDate') }}: {{ formatDate(announcement.end_date) }}
                </span>
                <span>
                  {{ $t('announcements.created') }}: {{ formatDate(announcement.created_at) }}
                </span>
              </div>
            </div>
            <div class="flex items-center gap-2 ml-4">
              <button
                @click="openEditModal(announcement)"
                class="p-2 text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/20 rounded-lg transition-colors"
                :title="$t('common.edit')"
              >
                <Icon icon="mdi:pencil" class="h-5 w-5" />
              </button>
              <button
                @click="confirmDelete(announcement)"
                class="p-2 text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
                :title="$t('common.delete')"
              >
                <Icon icon="mdi:delete" class="h-5 w-5" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-12">
        <Icon icon="mdi:bullhorn-outline" class="h-16 w-16 mx-auto text-gray-400 mb-4" />
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">
          {{ $t('announcements.emptyTitle') }}
        </h3>
        <p class="text-gray-600 dark:text-gray-400 mb-6">
          {{ $t('announcements.emptyText') }}
        </p>
        <button
          @click="openCreateModal"
          class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors"
        >
          {{ $t('announcements.create') }}
        </button>
      </div>

      <!-- Create/Edit Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4"
        @click.self="closeModal"
      >
        <div class="bg-white dark:bg-gray-800 rounded-lg max-w-2xl w-full p-6 max-h-[90vh] overflow-y-auto">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ isEditMode ? $t('announcements.edit') : $t('announcements.create') }}
            </h2>
            <button
              @click="closeModal"
              class="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-colors"
            >
              <Icon icon="mdi:close" class="h-5 w-5" />
            </button>
          </div>

          <form @submit.prevent="saveAnnouncement" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                {{ $t('announcements.titleField') }} *
              </label>
              <input
                v-model="formData.title"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                {{ $t('announcements.contentField') }}
              </label>
              <textarea
                v-model="formData.content"
                rows="4"
                class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              ></textarea>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  {{ $t('announcements.typeField') }} *
                </label>
                <select
                  v-model="formData.type"
                  required
                  class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
                >
                  <option value="info">{{ $t('announcements.type_info') }}</option>
                  <option value="warning">{{ $t('announcements.type_warning') }}</option>
                  <option value="success">{{ $t('announcements.type_success') }}</option>
                  <option value="error">{{ $t('announcements.type_error') }}</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  {{ $t('announcements.priorityField') }}
                </label>
                <input
                  v-model.number="formData.priority"
                  type="number"
                  min="0"
                  class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  {{ $t('announcements.startDateField') }}
                </label>
                <input
                  v-model="formData.start_date"
                  type="datetime-local"
                  class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  {{ $t('announcements.endDateField') }}
                </label>
                <input
                  v-model="formData.end_date"
                  type="datetime-local"
                  class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>

            <div class="flex items-center gap-2">
              <input
                v-model="formData.is_active"
                type="checkbox"
                id="is_active"
                class="w-4 h-4 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
              />
              <label for="is_active" class="text-sm font-medium text-gray-700 dark:text-gray-300">
                {{ $t('announcements.isActive') }}
              </label>
            </div>

            <div class="flex gap-3 pt-4">
              <button
                type="button"
                @click="closeModal"
                class="flex-1 px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-white rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
              >
                {{ $t('common.cancel') }}
              </button>
              <button
                type="submit"
                :disabled="isSaving"
                class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 text-white rounded-lg font-medium transition-colors"
              >
                {{ isSaving ? $t('announcements.saving') : $t('common.save') }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Delete Confirmation Modal -->
      <div
        v-if="showDeleteConfirm"
        class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4"
        @click.self="showDeleteConfirm = false"
      >
        <div class="bg-white dark:bg-gray-800 rounded-lg max-w-md w-full p-6">
          <div class="flex items-center gap-3 mb-4">
            <div class="h-12 w-12 bg-red-100 dark:bg-red-900 rounded-full flex items-center justify-center">
              <Icon icon="mdi:alert" class="h-6 w-6 text-red-600 dark:text-red-400" />
            </div>
            <h3 class="text-xl font-bold text-gray-900 dark:text-white">
              {{ $t('announcements.confirmDeleteTitle') }}
            </h3>
          </div>

          <p class="text-gray-700 dark:text-gray-300 mb-6">
            {{ $t('announcements.confirmDeleteText', { title: announcementToDelete?.title }) }}
          </p>

          <div class="flex gap-3">
            <button
              @click="showDeleteConfirm = false"
              class="flex-1 px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-white rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
            >
              {{ $t('common.cancel') }}
            </button>
            <button
              @click="deleteAnnouncement"
              :disabled="isDeleting"
              class="flex-1 px-4 py-2 bg-red-600 hover:bg-red-700 disabled:bg-red-400 text-white rounded-lg font-medium transition-colors"
            >
              {{ isDeleting ? $t('announcements.deleting') : $t('common.delete') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { announcementsService } from '@/services/api'

const announcements = ref([])
const isLoading = ref(false)
const isSaving = ref(false)
const isDeleting = ref(false)
const showModal = ref(false)
const showDeleteConfirm = ref(false)
const isEditMode = ref(false)
const announcementToDelete = ref(null)

const formData = ref({
  title: '',
  content: '',
  type: 'info',
  priority: 0,
  is_active: true,
  start_date: null,
  end_date: null
})

const loadAnnouncements = async () => {
  isLoading.value = true
  try {
    announcements.value = await announcementsService.getAllAnnouncements()
  } catch (error) {
    console.error('Error loading announcements:', error)
  } finally {
    isLoading.value = false
  }
}

const openCreateModal = () => {
  isEditMode.value = false
  formData.value = {
    title: '',
    content: '',
    type: 'info',
    priority: 0,
    is_active: true,
    start_date: null,
    end_date: null
  }
  showModal.value = true
}

const openEditModal = (announcement) => {
  isEditMode.value = true
  formData.value = {
    id: announcement.id,
    title: announcement.title,
    content: announcement.content,
    type: announcement.type,
    priority: announcement.priority,
    is_active: announcement.is_active,
    start_date: announcement.start_date ? formatDateTimeLocal(announcement.start_date) : null,
    end_date: announcement.end_date ? formatDateTimeLocal(announcement.end_date) : null
  }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const saveAnnouncement = async () => {
  isSaving.value = true
  try {
    const data = {
      ...formData.value,
      start_date: formData.value.start_date ? new Date(formData.value.start_date).toISOString() : null,
      end_date: formData.value.end_date ? new Date(formData.value.end_date).toISOString() : null
    }
    delete data.id

    if (isEditMode.value) {
      await announcementsService.updateAnnouncement(formData.value.id, data)
    } else {
      await announcementsService.createAnnouncement(data)
    }

    closeModal()
    await loadAnnouncements()
  } catch (error) {
    console.error('Error saving announcement:', error)
  } finally {
    isSaving.value = false
  }
}

const confirmDelete = (announcement) => {
  announcementToDelete.value = announcement
  showDeleteConfirm.value = true
}

const deleteAnnouncement = async () => {
  isDeleting.value = true
  try {
    await announcementsService.deleteAnnouncement(announcementToDelete.value.id)
    showDeleteConfirm.value = false
    await loadAnnouncements()
  } catch (error) {
    console.error('Error deleting announcement:', error)
  } finally {
    isDeleting.value = false
  }
}

const getTypeBadgeClass = (type) => {
  const classes = {
    info: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
    warning: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
    success: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200',
    error: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200'
  }
  return classes[type] || classes.info
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('fr-FR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const formatDateTimeLocal = (date) => {
  if (!date) return null
  const d = new Date(date)
  const offset = d.getTimezoneOffset()
  d.setMinutes(d.getMinutes() - offset)
  return d.toISOString().slice(0, 16)
}

onMounted(() => {
  loadAnnouncements()
})
</script>
