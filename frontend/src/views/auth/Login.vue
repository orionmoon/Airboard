<template>
  <div class="auth-container">
    <div class="auth-card">
      <!-- Header -->
      <div class="auth-header">
        <div class="flex items-center justify-center mb-4">
          <div class="h-12 w-12 bg-white rounded-lg flex items-center justify-center">
            <Icon icon="mdi:view-dashboard" class="h-6 w-6 text-black" />
          </div>
        </div>
        <h2 class="auth-title">Welcome to Airboard</h2>
        <p class="auth-subtitle">Sign in to your account</p>
        
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div class="form-group">
          <label for="username" class="form-label form-label-required">
            Username
          </label>
          <input
            id="username"
            v-model="form.username"
            type="text"
            required
            class="form-input"
            placeholder="Enter your username"
            :disabled="loading"
          />
          <p v-if="errors.username" class="form-error">{{ errors.username }}</p>
        </div>

        <div class="form-group">
          <label for="password" class="form-label form-label-required">
            Password
          </label>
          <div class="relative">
            <input
              id="password"
              v-model="form.password"
              :type="showPassword ? 'text' : 'password'"
              required
              class="form-input pr-10"
              placeholder="Enter your password"
              :disabled="loading"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-300"
            >
              <Icon :icon="showPassword ? 'mdi:eye-off' : 'mdi:eye'" class="h-5 w-5" />
            </button>
          </div>
          <p v-if="errors.password" class="form-error">{{ errors.password }}</p>
        </div>

        <div class="flex items-center">
          <label class="flex items-center">
            <input
              v-model="form.remember"
              type="checkbox"
              class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700"
            />
            <span class="ml-2 text-sm text-gray-600 dark:text-gray-300">Remember me</span>
          </label>
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="btn btn-primary w-full"
        >
          <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
          Sign in
        </button>

        <div class="text-center">
          <span class="text-sm text-gray-400">Don't have an account? </span>
          <router-link
            to="/auth/register"
            class="text-sm text-green-400 hover:text-green-300 font-medium"
          >
            Sign up
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'

const router = useRouter()
const authStore = useAuthStore()
const appStore = useAppStore()

const loading = ref(false)
const showPassword = ref(false)
const errors = ref({})

const form = reactive({
  username: '',
  password: '',
  remember: false
})

const validateForm = () => {
  errors.value = {}
  
  if (!form.username.trim()) {
    errors.value.username = 'Username is required'
  }
  
  if (!form.password.trim()) {
    errors.value.password = 'Password is required'
  }
  
  return Object.keys(errors.value).length === 0
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  
  try {
    const response = await authStore.login(form)
    console.log('✅ Login successful, redirecting...')
    
    appStore.showSuccess('Welcome back!')
    
    // Force navigation with nextTick to ensure state is updated
    await nextTick()
    
    // Check for redirect parameter
    const redirectPath = router.currentRoute.value.query.redirect || '/dashboard'
    
    // Force replace instead of push to avoid back button issues
    await router.replace(redirectPath)
    
    console.log('🎯 Navigation completed to:', redirectPath)
    
  } catch (error) {
    console.error('Login error:', error)
    
    if (error.response?.status === 401) {
      appStore.showError('Invalid username or password')
    } else if (error.response?.status === 422) {
      const validationErrors = error.response.data.errors
      if (validationErrors) {
        errors.value = validationErrors
      } else {
        appStore.showError('Invalid input data')
      }
    } else {
      appStore.showError('Login failed. Please try again.')
    }
  } finally {
    loading.value = false
  }
}
</script>