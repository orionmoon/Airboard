<template>
  <button
    @click="cycleZoom"
    class="flex items-center gap-2 px-3 py-2 rounded-lg bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors text-sm font-medium text-gray-700 dark:text-gray-300"
    :title="currentLevel.label"
  >
    <Icon :icon="currentLevel.icon" class="h-4 w-4" />
    <span class="text-xs">{{ modelValue }}%</span>
  </button>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({
  modelValue: {
    type: String,
    default: '100'
  }
})

const emit = defineEmits(['update:modelValue'])

const zoomLevels = computed(() => [
  { value: '75', label: t('dashboard.zoom.extraSmall'), icon: 'mdi:magnify-minus-outline' },
  { value: '85', label: t('dashboard.zoom.small'), icon: 'mdi:magnify-minus' },
  { value: '100', label: t('dashboard.zoom.normal'), icon: 'mdi:magnify' },
  { value: '115', label: t('dashboard.zoom.large'), icon: 'mdi:magnify-plus' }
])

const currentLevel = computed(() => {
  return zoomLevels.value.find(level => level.value === props.modelValue) || zoomLevels.value[2]
})

const cycleZoom = () => {
  const currentIndex = zoomLevels.value.findIndex(level => level.value === props.modelValue)
  const nextIndex = (currentIndex + 1) % zoomLevels.value.length
  emit('update:modelValue', zoomLevels.value[nextIndex].value)
}
</script>
