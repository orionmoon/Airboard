<template>
  <div class="flex items-center gap-2">
    <Icon icon="mdi:sort" class="h-5 w-5 text-gray-500 dark:text-gray-400" />
    <select
      :value="modelValue"
      @change="$emit('update:modelValue', $event.target.value)"
      class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent dark:bg-gray-700 dark:text-white text-sm"
    >
      <option value="date-desc">{{ $t('news.sort.dateNewest') }}</option>
      <option value="date-asc">{{ $t('news.sort.dateOldest') }}</option>
      <option value="views-desc">{{ $t('news.sort.mostViewed') }}</option>
      <option value="reactions-desc">{{ $t('news.sort.mostReacted') }}</option>
      <option value="title-asc">{{ $t('news.sort.titleAZ') }}</option>
      <option value="title-desc">{{ $t('news.sort.titleZA') }}</option>
    </select>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'

defineProps({
  modelValue: {
    type: String,
    default: 'date-desc'
  }
})

defineEmits(['update:modelValue'])
</script>
