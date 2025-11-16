<template>
  <div class="flex items-center gap-1 bg-gray-100 dark:bg-gray-800 rounded-lg p-1">
    <button
      v-for="mode in viewModes"
      :key="mode.value"
      @click="$emit('update:modelValue', mode.value)"
      :class="[
        'flex items-center gap-2 px-3 py-2 rounded-md text-sm font-medium transition-all duration-200',
        modelValue === mode.value
          ? 'bg-white dark:bg-gray-700 text-primary-600 dark:text-primary-400 shadow-sm'
          : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
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
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

defineProps({
  modelValue: {
    type: String,
    default: 'grid'
  }
})

defineEmits(['update:modelValue'])

const viewModes = [
  {
    value: 'grid',
    label: 'Grille',
    icon: 'mdi:view-grid'
  },
  {
    value: 'list',
    label: 'Liste',
    icon: 'mdi:view-list'
  },
  {
    value: 'compact',
    label: 'Compact',
    icon: 'mdi:view-compact'
  }
]
</script>
