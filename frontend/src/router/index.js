import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Import des vues
const Dashboard = () => import('@/views/dashboard/Dashboard.vue')
const Login = () => import('@/views/auth/Login.vue')
const Register = () => import('@/views/auth/Register.vue')
const OAuthCallback = () => import('@/views/auth/OAuthCallback.vue')

// Admin views
const AdminDashboard = () => import('@/views/admin/AdminDashboard.vue')
const AppGroupsManagement = () => import('@/views/admin/AppGroupsManagement.vue')
const ApplicationsManagement = () => import('@/views/admin/ApplicationsManagement.vue')
const UsersManagement = () => import('@/views/admin/UsersManagement.vue')
const GroupsManagement = () => import('@/views/admin/GroupsManagement.vue')
const Analytics = () => import('@/views/admin/Analytics.vue')
const AnnouncementsManagement = () => import('@/views/admin/AnnouncementsManagement.vue')
const NewsManagement = () => import('@/views/admin/NewsManagement.vue')
const NewsEditor = () => import('@/views/admin/NewsEditor.vue')

// Group Admin views
const GroupAdminDashboard = () => import('@/views/group-admin/GroupAdminDashboard.vue')
const GroupAdminAppGroups = () => import('@/views/group-admin/AppGroupsManagement.vue')
const GroupAdminApplications = () => import('@/views/group-admin/ApplicationsManagement.vue')
const GroupAdminNews = () => import('@/views/group-admin/NewsManagement.vue')
const GroupAdminNewsEditor = () => import('@/views/group-admin/NewsEditor.vue')

// News views (public)
const NewsCenter = () => import('@/views/NewsCenter.vue')
const NewsDetail = () => import('@/views/NewsDetail.vue')

// Error views
const NotFound = () => import('@/views/errors/NotFound.vue')
const Unauthorized = () => import('@/views/errors/Unauthorized.vue')

const routes = [
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: {
      requiresAuth: true,
      title: 'Tableau de bord'
    }
  },
  {
    path: '/news',
    name: 'NewsCenter',
    component: NewsCenter,
    meta: {
      requiresAuth: true,
      title: 'News Hub'
    }
  },
  {
    path: '/news/:slug',
    name: 'NewsDetail',
    component: NewsDetail,
    meta: {
      requiresAuth: true,
      title: 'Article'
    }
  },
  {
    path: '/auth/login',
    name: 'Login',
    component: Login,
    meta: { 
      requiresGuest: true,
      title: 'Connexion'
    }
  },
  {
    path: '/auth/register',
    name: 'Register',
    component: Register,
    meta: {
      requiresGuest: true,
      title: 'Inscription'
    }
  },
  {
    path: '/auth/oauth/:provider/callback',
    name: 'OAuthCallback',
    component: OAuthCallback,
    meta: {
      title: 'Authentification OAuth'
    }
  },
  {
    path: '/admin',
    name: 'Admin',
    component: AdminDashboard,
    meta: { 
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Administration'
    }
  },
  {
    path: '/admin/app-groups',
    name: 'AdminAppGroups',
    component: AppGroupsManagement,
    meta: { 
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Groupes d\'applications'
    }
  },
  {
    path: '/admin/applications',
    name: 'AdminApplications',
    component: ApplicationsManagement,
    meta: { 
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Applications'
    }
  },
  {
    path: '/admin/users',
    name: 'AdminUsers',
    component: UsersManagement,
    meta: { 
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Utilisateurs'
    }
  },
  {
    path: '/admin/groups',
    name: 'AdminGroups',
    component: GroupsManagement,
    meta: { 
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Groupes'
    }
  },
  {
    path: '/admin/settings',
    name: 'AdminSettings',
    component: () => import('@/views/admin/SettingsManagement.vue'),
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Paramètres'
    }
  },
  {
    path: '/admin/oauth',
    name: 'AdminOAuth',
    component: () => import('@/views/admin/OAuthSettings.vue'),
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      title: 'OAuth Configuration'
    }
  },
  {
    path: '/admin/analytics',
    name: 'AdminAnalytics',
    component: Analytics,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Analytics'
    }
  },
  {
    path: '/admin/announcements',
    name: 'AdminAnnouncements',
    component: AnnouncementsManagement,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      title: 'Announcements'
    }
  },
  {
    path: '/admin/news',
    name: 'AdminNews',
    component: NewsManagement,
    meta: {
      requiresAuth: true,
      requiresEditor: true,
      title: 'News Hub'
    }
  },
  {
    path: '/admin/news/:id/edit',
    name: 'AdminNewsEdit',
    component: NewsEditor,
    meta: {
      requiresAuth: true,
      requiresEditor: true,
      title: 'Edit Article'
    }
  },
  {
    path: '/admin/news/new',
    name: 'AdminNewsNew',
    component: NewsEditor,
    meta: {
      requiresAuth: true,
      requiresEditor: true,
      title: 'New Article'
    }
  },
  // Routes Group Admin
  {
    path: '/group-admin',
    name: 'GroupAdminDashboard',
    component: GroupAdminDashboard,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'Administration de Groupe'
    }
  },
  {
    path: '/group-admin/app-groups',
    name: 'GroupAdminAppGroups',
    component: GroupAdminAppGroups,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'Mes Groupes d\'Applications'
    }
  },
  {
    path: '/group-admin/applications',
    name: 'GroupAdminApplications',
    component: GroupAdminApplications,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'Mes Applications'
    }
  },
  {
    path: '/group-admin/news',
    name: 'GroupAdminNews',
    component: GroupAdminNews,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'News de Groupe'
    }
  },
  {
    path: '/group-admin/news/new',
    name: 'GroupAdminNewsNew',
    component: GroupAdminNewsEditor,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'Nouvel Article de Groupe'
    }
  },
  {
    path: '/group-admin/news/:slug/edit',
    name: 'GroupAdminNewsEdit',
    component: GroupAdminNewsEditor,
    meta: {
      requiresAuth: true,
      requiresGroupAdmin: true,
      title: 'Modifier l\'Article'
    }
  },
  {
    path: '/unauthorized',
    name: 'Unauthorized',
    component: Unauthorized,
    meta: {
      title: 'Accès non autorisé'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound,
    meta: {
      title: 'Page non trouvée'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// Guards de navigation
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // Charger l'état d'authentification depuis le localStorage SEULEMENT au premier chargement
  if (!authStore.isAuthenticated && !authStore.isLoading && !authStore.token) {
    console.log('🔄 Loading auth from storage...')
    authStore.loadFromStorage()
  }
  
  console.log('🧭 Router guard:', {
    to: to.path,
    isAuthenticated: authStore.isAuthenticated,
    hasToken: !!authStore.token,
    user: authStore.user?.username
  })

  // Définir le titre de la page
  if (to.meta.title) {
    document.title = `${to.meta.title} - Airboard`
  } else {
    document.title = 'Airboard - Portail Applicatif'
  }

  // Vérifier l'authentification
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({
      name: 'Login',
      query: { redirect: to.fullPath }
    })
    return
  }

  // Vérifier les droits admin
  if (to.meta.requiresAdmin && !authStore.isAdmin) {
    next({ name: 'Unauthorized' })
    return
  }

  // Vérifier les droits editor (admin ou editor)
  if (to.meta.requiresEditor) {
    const userRole = authStore.user?.role
    if (userRole !== 'admin' && userRole !== 'editor') {
      next({ name: 'Unauthorized' })
      return
    }
  }

  // Vérifier les droits group admin (admin ou group_admin)
  if (to.meta.requiresGroupAdmin) {
    const userRole = authStore.user?.role
    if (userRole !== 'admin' && userRole !== 'group_admin') {
      next({ name: 'Unauthorized' })
      return
    }
  }

  // Rediriger les utilisateurs connectés loin des pages d'auth
  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next({ name: 'Dashboard' })
    return
  }

  next()
})

// Gérer les erreurs de navigation
router.onError((error) => {
  console.error('Erreur de navigation:', error)
  
  // Rediriger vers la page d'erreur pour les erreurs de chargement de composant
  if (error.message.includes('Failed to fetch dynamically imported module')) {
    window.location.reload()
  }
})

export default router