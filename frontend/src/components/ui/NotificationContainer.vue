<template>
  <div class="fixed bottom-4 right-4 z-50 space-y-2">
    <TransitionGroup name="toast" tag="div">
      <div
        v-for="notification in appStore.notifications"
        :key="notification.id"
        :class="getToastClasses(notification.type)"
      >
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <Icon :icon="getIcon(notification.type)" class="h-5 w-5" />
          </div>
          <div class="ml-3 flex-1">
            <p class="text-sm font-medium">
              {{ notification.title }}
            </p>
            <p v-if="notification.message" class="mt-1 text-sm opacity-90">
              {{ notification.message }}
            </p>
          </div>
          <div class="ml-4 flex-shrink-0">
            <button
              @click="appStore.removeNotification(notification.id)"
              class="inline-flex rounded-md p-1.5 hover:bg-black hover:bg-opacity-10 focus:outline-none transition-colors"
            >
              <Icon icon="mdi:close" class="h-4 w-4" />
            </button>
          </div>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()

const getToastClasses = (type) => {
  const baseClasses = 'max-w-sm w-full p-4 rounded-lg shadow-lg pointer-events-auto'
  
  switch (type) {
    case 'success':
      return `${baseClasses} toast-success`
    case 'error':
      return `${baseClasses} toast-error`
    case 'warning':
      return `${baseClasses} toast-warning`
    case 'info':
    default:
      return `${baseClasses} toast-info`
  }
}

const getIcon = (type) => {
  switch (type) {
    case 'success':
      return 'mdi:check-circle'
    case 'error':
      return 'mdi:alert-circle'
    case 'warning':
      return 'mdi:alert'
    case 'info':
    default:
      return 'mdi:information'
  }
}
</script>

<style scoped>
.toast-enter-active {
  transition: all 0.3s ease-out;
}

.toast-leave-active {
  transition: all 0.3s ease-in;
}

.toast-enter-from {
  transform: translateX(100%);
  opacity: 0;
}

.toast-leave-to {
  transform: translateX(100%);
  opacity: 0;
}

.toast-move {
  transition: transform 0.3s ease;
}
</style>