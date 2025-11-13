<template>
  <div class="form-group">
    <label v-if="label" :for="id" class="form-label" :class="{ 'form-label-required': required }">
      {{ label }}
    </label>
    <div class="icon-input-container">
      <input
        :id="id"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        type="text"
        :placeholder="placeholder"
        :required="required"
        class="form-input icon-input"
      />
      <div class="icon-preview">
        <Icon 
          v-if="isValidIcon" 
          :icon="modelValue" 
          class="h-5 w-5" 
          @error="onIconError"
        />
        <Icon 
          v-else 
          icon="mdi:help-circle-outline" 
          class="h-5 w-5 text-gray-400" 
        />
      </div>
    </div>
    <div v-if="showHelp" class="form-help">
      {{ helpText }}
    </div>
    <div v-if="error" class="form-error">{{ error }}</div>
    
    <!-- Suggestions populaires -->
    <div v-if="showSuggestions && suggestions.length > 0" class="mt-2">
      <p class="text-xs text-gray-500 mb-2">Suggestions populaires :</p>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="suggestion in suggestions"
          :key="suggestion.icon"
          @click="selectSuggestion(suggestion.icon)"
          type="button"
          class="inline-flex items-center space-x-2 px-3 py-1 bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 rounded-md text-xs transition-colors"
        >
          <Icon :icon="suggestion.icon" class="h-3 w-3" />
          <span>{{ suggestion.label }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Icon } from '@iconify/vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  label: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Ex: mdi:home, lucide:settings, heroicons:user'
  },
  required: {
    type: Boolean,
    default: false
  },
  id: {
    type: String,
    default: () => `icon-input-${Math.random().toString(36).substr(2, 9)}`
  },
  error: {
    type: String,
    default: ''
  },
  showHelp: {
    type: Boolean,
    default: true
  },
  showSuggestions: {
    type: Boolean,
    default: true
  },
  category: {
    type: String,
    default: 'general'
  }
})

const emit = defineEmits(['update:modelValue'])

const isValidIcon = ref(true)

const helpText = computed(() => {
  return `Utilisez le format "collection:nom" (ex: mdi:home, lucide:user, heroicons:cog). Recherchez sur iconify.design`
})

const suggestions = computed(() => {
  const suggestionsByCategory = {
    general: [
      { icon: 'mdi:home', label: 'Accueil' },
      { icon: 'mdi:cog', label: 'Paramètres' },
      { icon: 'mdi:account', label: 'Utilisateur' },
      { icon: 'mdi:folder', label: 'Dossier' },
      { icon: 'mdi:file-document', label: 'Document' },
      { icon: 'mdi:chart-line', label: 'Graphique' }
    ],
    application: [
      { icon: 'mdi:application', label: 'Application' },
      { icon: 'mdi:web', label: 'Web' },
      { icon: 'mdi:git', label: 'Git' },
      { icon: 'mdi:database', label: 'Base de données' },
      { icon: 'mdi:server', label: 'Serveur' },
      { icon: 'mdi:cloud', label: 'Cloud' },
      { icon: 'mdi:email', label: 'Email' },
      { icon: 'mdi:chat', label: 'Chat' },
      { icon: 'mdi:calendar', label: 'Calendrier' },
      { icon: 'mdi:tools', label: 'Outils' }
    ],
    group: [
      { icon: 'mdi:folder-multiple', label: 'Groupes' },
      { icon: 'mdi:account-group', label: 'Équipe' },
      { icon: 'mdi:shield-check', label: 'Sécurité' },
      { icon: 'mdi:code-tags', label: 'Développement' },
      { icon: 'mdi:chart-pie', label: 'Analytics' },
      { icon: 'mdi:briefcase', label: 'Business' }
    ],
    social: [
      { icon: 'mdi:github', label: 'GitHub' },
      { icon: 'mdi:gitlab', label: 'GitLab' },
      { icon: 'mdi:slack', label: 'Slack' },
      { icon: 'mdi:discord', label: 'Discord' },
      { icon: 'mdi:microsoft-teams', label: 'Teams' },
      { icon: 'mdi:google', label: 'Google' }
    ]
  }
  
  return suggestionsByCategory[props.category] || suggestionsByCategory.general
})

// Vérifier si l'icône est valide
watch(() => props.modelValue, (newValue) => {
  if (!newValue) {
    isValidIcon.value = false
    return
  }
  
  // Vérification basique du format
  const iconPattern = /^[a-zA-Z0-9\-_]+:[a-zA-Z0-9\-_]+$/
  isValidIcon.value = iconPattern.test(newValue)
}, { immediate: true })

const onIconError = () => {
  isValidIcon.value = false
}

const selectSuggestion = (icon) => {
  emit('update:modelValue', icon)
}
</script>