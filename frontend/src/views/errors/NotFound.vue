<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center px-4">
    <div class="max-w-md w-full text-center">
      <div class="mb-8">
        <Icon icon="mdi:compass-off" class="mx-auto h-24 w-24 text-gray-400 dark:text-gray-600" />
      </div>
      
      <h1 class="text-6xl font-bold text-gray-900 dark:text-white mb-4">404</h1>
      
      <h2 class="text-2xl font-semibold text-gray-700 dark:text-gray-300 mb-4">
        Page non trouvée
      </h2>
      
      <p class="text-gray-600 dark:text-gray-400 mb-8">
        Désolé, la page que vous recherchez n'existe pas ou a été déplacée.
      </p>
      
      <div class="space-y-4">
        <router-link to="/dashboard" class="btn-primary btn-md w-full">
          <Icon icon="mdi:home" class="h-5 w-5 mr-2" />
          Retour au tableau de bord
        </router-link>
        
        <button @click="$router.back()" class="btn-secondary btn-md w-full">
          <Icon icon="mdi:arrow-left" class="h-5 w-5 mr-2" />
          Page précédente
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
</script>