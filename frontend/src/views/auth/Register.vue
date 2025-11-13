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
        <h2 class="auth-title">Create Account</h2>
        <p class="auth-subtitle">Join Airboard today</p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="form-group">
            <label for="first_name" class="form-label">
              First Name
            </label>
            <input
              id="first_name"
              v-model="form.first_name"
              type="text"
              class="form-input"
              placeholder="John"
              :disabled="loading"
            />
          </div>

          <div class="form-group">
            <label for="last_name" class="form-label">
              Last Name
            </label>
            <input
              id="last_name"
              v-model="form.last_name"
              type="text"
              class="form-input"
              placeholder="Doe"
              :disabled="loading"
            />
          </div>
        </div>

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
            placeholder="johndoe"
            :disabled="loading"
          />
          <p v-if="errors.username" class="form-error">{{ errors.username }}</p>
        </div>

        <div class="form-group">
          <label for="email" class="form-label form-label-required">
            Email
          </label>
          <input
            id="email"
            v-model="form.email"
            type="email"
            required
            class="form-input"
            placeholder="john@example.com"
            :disabled="loading"
          />
          <p v-if="errors.email" class="form-error">{{ errors.email }}</p>
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
              minlength="6"
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
          <p class="form-help">At least 6 characters</p>
        </div>

        <div class="form-group">
          <label for="password_confirmation" class="form-label form-label-required">
            Confirm Password
          </label>
          <div class="relative">
            <input
              id="password_confirmation"
              v-model="form.password_confirmation"
              :type="showConfirmPassword ? 'text' : 'password'"
              required
              class="form-input pr-10"
              placeholder="Confirm your password"
              :disabled="loading"
            />
            <button
              type="button"
              @click="showConfirmPassword = !showConfirmPassword"
              class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-300"
            >
              <Icon :icon="showConfirmPassword ? 'mdi:eye-off' : 'mdi:eye'" class="h-5 w-5" />
            </button>
          </div>
          <p v-if="errors.password_confirmation" class="form-error">{{ errors.password_confirmation }}</p>
        </div>

        <div class="flex items-center">
          <input
            v-model="form.terms"
            type="checkbox"
            required
            class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-600 rounded bg-gray-700"
          />
          <span class="ml-2 text-sm text-gray-300">
            I agree to the 
            <a href="#" class="text-green-400 hover:text-green-300">Terms of Service</a>
            and 
            <a href="#" class="text-green-400 hover:text-green-300">Privacy Policy</a>
          </span>
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="btn btn-primary w-full"
        >
          <Icon v-if="loading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
          Create Account
        </button>

        <div class="text-center">
          <span class="text-sm text-gray-400">Already have an account? </span>
          <router-link
            to="/auth/login"
            class="text-sm text-green-400 hover:text-green-300 font-medium"
          >
            Sign in
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'

const router = useRouter()
const authStore = useAuthStore()
const appStore = useAppStore()

const loading = ref(false)
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const errors = ref({})

const form = reactive({
  first_name: '',
  last_name: '',
  username: '',
  email: '',
  password: '',
  password_confirmation: '',
  terms: false
})

const validateForm = () => {
  errors.value = {}
  
  if (!form.username.trim()) {
    errors.value.username = 'Username is required'
  } else if (form.username.length < 3) {
    errors.value.username = 'Username must be at least 3 characters'
  }
  
  if (!form.email.trim()) {
    errors.value.email = 'Email is required'
  } else if (!isValidEmail(form.email)) {
    errors.value.email = 'Invalid email format'
  }
  
  if (!form.password.trim()) {
    errors.value.password = 'Password is required'
  } else if (form.password.length < 6) {
    errors.value.password = 'Password must be at least 6 characters'
  }
  
  if (form.password !== form.password_confirmation) {
    errors.value.password_confirmation = 'Passwords do not match'
  }
  
  return Object.keys(errors.value).length === 0
}

const isValidEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

const handleSubmit = async () => {
  if (!validateForm()) return
  
  loading.value = true
  
  try {
    await authStore.register(form)
    appStore.showSuccess('Account created successfully! Please sign in.')
    router.push('/auth/login')
  } catch (error) {
    console.error('Registration error:', error)
    
    if (error.response?.status === 422) {
      const validationErrors = error.response.data.errors
      if (validationErrors) {
        errors.value = validationErrors
      } else {
        appStore.showError('Invalid input data')
      }
    } else if (error.response?.status === 409) {
      appStore.showError('Username or email already exists')
    } else {
      appStore.showError('Registration failed. Please try again.')
    }
  } finally {
    loading.value = false
  }
}
</script>