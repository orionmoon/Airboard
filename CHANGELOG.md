# 📋 Changelog - Airboard

Toutes les modifications notables du projet Airboard sont documentées ici.

---

## [1.0.0-clean] - 2025-11-13

### ✨ Nouvelles fonctionnalités

#### 🗄️ Réinitialisation de la base de données
- **Backend** : Endpoint API `POST /admin/database/reset`
  - Suppression complète de toutes les données
  - Réinitialisation des séquences PostgreSQL
  - Accessible uniquement aux administrateurs
  - Fichier : `backend/handlers/admin.go` (lignes 994-1091)

- **Frontend** : Interface de réinitialisation dans Admin Dashboard
  - Section "Zone de danger" avec design rouge distinctif
  - Modal de confirmation avec double vérification
  - Gestion des états de chargement et d'erreur
  - Déconnexion automatique après réinitialisation
  - Fichier : `frontend/src/views/admin/AdminDashboard.vue`

- **API Service** : Méthode `resetDatabase()`
  - Fichier : `frontend/src/services/api.js` (lignes 244-247)

### 🐳 Améliorations Docker

- Ajout de `.dockerignore` pour backend et frontend
- Documentation complète du déploiement Docker
- Guide de démarrage rapide
- Optimisation des builds Docker

### 📚 Documentation

- **Nouveau** : `DOCKER-DEPLOYMENT.md` - Guide complet (déploiement, configuration, dépannage)
- **Nouveau** : `QUICK-START-DOCKER.md` - Démarrage rapide en 3 étapes
- **Nouveau** : `PROJECT-STRUCTURE.md` - Structure détaillée du projet
- **Nouveau** : `CHANGELOG.md` - Ce fichier

### 🛠️ Scripts et outils

- **Nouveau** : `clean-project.sh` - Nettoyage complet du projet
  - Supprime node_modules, binaires, fichiers temporaires
  - Réduit la taille de ~122 MB à ~824 KB

- **Nouveau** : `create-docker-release.sh` - Création de releases distribuables
  - Génère une archive tar.gz prête à déployer
  - Inclut les instructions d'installation
  - Taille : ~107 KB

### 🔧 Optimisations

- Nettoyage complet du code
- Suppression des fichiers IDE et temporaires
- Optimisation de la taille du projet
- Amélioration des Dockerfiles

### 🔒 Sécurité

- Validation stricte des permissions pour la réinitialisation DB
- Double confirmation avant toute action destructive
- Messages d'avertissement clairs
- Déconnexion forcée après réinitialisation

---

## [1.0.0] - 2025-10-29 (Version initiale)

### ✨ Fonctionnalités principales

#### 🔐 Authentification et autorisation
- Authentification JWT avec refresh tokens
- Rôles utilisateur (admin, user)
- Middleware de protection des routes
- Gestion de session sécurisée

#### 👥 Gestion des utilisateurs
- CRUD complet des utilisateurs
- Profils utilisateur
- Changement de mot de passe
- Activation/désactivation de comptes

#### 📱 Portail d'applications
- Groupes d'applications personnalisables
- Applications avec icônes Iconify
- Ordre personnalisé
- Activation/désactivation
- Ouverture dans nouvel onglet

#### 🎨 Interface utilisateur
- Design moderne avec TailwindCSS
- Mode sombre
- Interface responsive
- Composants réutilisables
- Animations fluides

#### 🌐 Internationalisation
- Support multilingue (FR, EN, ES, AR)
- Changement de langue dynamique
- Traductions complètes

#### 👥 Gestion des groupes
- Groupes d'utilisateurs
- Permissions par groupe
- Association utilisateur-groupe
- Association groupe-application

#### ⚙️ Paramètres d'application
- Nom de l'application personnalisable
- Icône d'application
- Titre du dashboard
- Message de bienvenue
- Réinitialisation des paramètres

#### 🐳 Docker
- Docker Compose pour déploiement complet
- PostgreSQL 15
- Backend Go optimisé
- Frontend Nginx avec compression
- Healthchecks
- Volumes persistants

#### 📊 Dashboard
- Statistiques en temps réel
- Vue d'ensemble des applications
- Accès rapide aux fonctionnalités

### 🏗️ Architecture technique

**Backend**
- Go 1.21+
- Framework Gin
- ORM GORM
- PostgreSQL 15
- JWT pour l'authentification
- Middleware CORS

**Frontend**
- Vue.js 3
- Vite
- TailwindCSS
- Vue Router
- Pinia (state management)
- Axios
- Vue i18n
- Iconify

**Infrastructure**
- Docker & Docker Compose
- Nginx
- PostgreSQL
- Multi-stage builds
- Healthchecks

### 📝 Scripts

- `airboard.sh` - Script principal de gestion
  - start-docker : Démarrage avec Docker
  - start-dev : Mode développement
  - start-local : Mode local
  - stop : Arrêt des services
  - restart : Redémarrage
  - logs : Affichage des logs
  - status : État des services
  - clean : Nettoyage

---

## 🎯 Prochaines versions

### [1.1.0] - Planifié

- [ ] Import/Export de données
- [ ] Sauvegarde automatique de la base de données
- [ ] Logs d'audit des actions admin
- [ ] API de statistiques avancées
- [ ] Support de thèmes personnalisés
- [ ] Gestion des permissions granulaires

### [1.2.0] - Planifié

- [ ] SSO (Single Sign-On)
- [ ] Authentification à deux facteurs (2FA)
- [ ] API REST complète documentée (Swagger)
- [ ] Tests automatisés (frontend + backend)
- [ ] CI/CD avec GitHub Actions

### [2.0.0] - Vision

- [ ] Architecture microservices
- [ ] Support de Kubernetes
- [ ] Monitoring et alerting intégrés
- [ ] Dashboard de métriques avancé
- [ ] Plugin system
- [ ] API GraphQL

---

## 📖 Légende

- ✨ Nouvelle fonctionnalité
- 🔧 Amélioration
- 🐛 Correction de bug
- 🔒 Sécurité
- 📚 Documentation
- 🗑️ Suppression
- ⚡ Performance
- 🎨 UI/UX

---

**Maintenu par** : L'équipe Airboard
**Licence** : MIT
