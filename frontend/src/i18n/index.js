import { createI18n } from 'vue-i18n'

import fr from '@/locales/fr.json'
import en from '@/locales/en.json'
import es from '@/locales/es.json'
import ar from '@/locales/ar.json'

const DEFAULT_LOCALE_KEY = 'airboard_locale'

export function resolveInitialLocale() {
  const saved = localStorage.getItem(DEFAULT_LOCALE_KEY)
  if (saved) return saved
  const browser = navigator.language?.toLowerCase() || 'en'
  if (browser.startsWith('fr')) return 'fr'
  if (browser.startsWith('es')) return 'es'
  if (browser.startsWith('ar')) return 'ar'
  return 'en'
}

export const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: resolveInitialLocale(),
  fallbackLocale: 'en',
  messages: {
    fr,
    en,
    es,
    ar,
  },
})

export function applyDirectionForLocale(locale) {
  const html = document.documentElement
  if (locale === 'ar') {
    html.setAttribute('dir', 'rtl')
  } else {
    html.setAttribute('dir', 'ltr')
  }
}

export function setLocale(locale) {
  i18n.global.locale.value = locale
  localStorage.setItem(DEFAULT_LOCALE_KEY, locale)
  applyDirectionForLocale(locale)
}

// Apply direction on load
applyDirectionForLocale(i18n.global.locale.value)

export default i18n


