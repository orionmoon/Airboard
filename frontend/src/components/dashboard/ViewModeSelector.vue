<template>
  <div class="flex items-center gap-1 bg-gray-100 dark:bg-gray-800 rounded-lg p-1">
    <button
      v-for="mode in viewModes"
      :key="mode.value"
      @click="$emit('update:modelValue', mode.value)"
      :class="[
        'flex items-center gap-2 px-3 py-2 rounded-md transition-all duration-200',
        'text-sm font-medium',
        modelValue === mode.value
          ? 'bg-white dark:bg-gray-700 text-gray-900 dark:text-white shadow-sm'
          : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white'
      ]"
      :title="mode.label"
    >
      <Icon :icon="mode.icon" class="h-4 w-4" />
      <span class="hidden sm:inline">{{ mode.label }}</span>
    </button>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

defineProps({
  modelValue: {
    type: String,
    default: 'default'
  }
})

defineEmits(['update:modelValue'])

const viewModes = computed(() => [
  { value: 'grid', label: t('dashboard.viewModes.grid'), icon: 'mdi:view-grid' },
  { value: 'default', label: t('dashboard.viewModes.default'), icon: 'mdi:grid-large' },
  { value: 'compact', label: t('dashboard.viewModes.compact'), icon: 'mdi:view-compact' }
])
</script>
