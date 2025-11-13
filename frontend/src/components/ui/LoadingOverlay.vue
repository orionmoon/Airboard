<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 loading-overlay">
    <div class="bg-white dark:bg-gray-800 rounded-lg p-6 shadow-xl max-w-sm w-full mx-4">
      <div class="flex flex-col items-center space-y-4">
        <!-- Spinner -->
        <div class="loading-spinner h-8 w-8 text-primary-600"></div>
        
        <!-- Message -->
        <div class="text-center">
          <h3 class="text-lg font-medium text-gray-900 dark:text-gray-100">
            {{ title }}
          </h3>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
            {{ message }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    default: 'Chargement en cours...'
  },
  message: {
    type: String,
    default: 'Veuillez patienter pendant que nous traitons votre demande.'
  }
})
</script>