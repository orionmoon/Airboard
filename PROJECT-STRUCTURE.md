# 📁 Structure du projet Airboard (Nettoyé)

## 📊 Vue d'ensemble

Taille totale : **~824 KB** (après nettoyage)

Le projet a été nettoyé pour ne contenir que les fichiers essentiels :
- Code source backend (Go)
- Code source frontend (Vue.js)
- Configuration Docker
- Documentation
- Scripts de gestion

## 🗂️ Arborescence

```
airboard/
├── backend/                    # API Backend Go
│   ├── config/                 # Configuration
│   │   └── config.go
│   ├── handlers/               # Gestionnaires HTTP
│   │   ├── admin.go           # ✨ NOUVEAU: avec ResetDatabase()
│   │   ├── auth.go
│   │   ├── dashboard.go
│   │   └── settings.go
│   ├── middleware/             # Middlewares
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models/                 # Modèles de données
│   │   └── models.go
│   ├── Dockerfile              # Build Docker multi-stage
│   ├── .dockerignore           # Exclusions Docker
│   ├── go.mod                  # Dépendances Go
│   ├── go.sum
│   ├── init.sql                # Script d'initialisation DB
│   ├── install.sh              # Script d'installation
│   └── main.go                 # Point d'entrée (avec route /admin/database/reset)
│
├── frontend/                   # Application Vue.js
│   ├── public/                 # Fichiers statiques
│   │   └── favicon.ico
│   ├── src/
│   │   ├── assets/            # CSS, images
│   │   │   └── css/
│   │   ├── components/        # Composants Vue
│   │   │   ├── admin/
│   │   │   ├── auth/
│   │   │   └── common/
│   │   ├── i18n/              # Traductions (FR, EN, ES, AR)
│   │   ├── router/            # Vue Router
│   │   ├── services/          # Services API
│   │   │   └── api.js         # ✨ NOUVEAU: avec resetDatabase()
│   │   ├── stores/            # Pinia stores
│   │   ├── views/             # Pages Vue
│   │   │   ├── admin/
│   │   │   │   └── AdminDashboard.vue  # ✨ NOUVEAU: avec bouton reset DB
│   │   │   ├── auth/
│   │   │   └── dashboard/
│   │   ├── App.vue
│   │   └── main.js
│   ├── Dockerfile              # Build Vue.js + Nginx
│   ├── .dockerignore           # Exclusions Docker
│   ├── index.html
│   ├── nginx.conf              # Config Nginx optimisée
│   ├── package.json            # Dépendances npm
│   ├── package-lock.json
│   ├── postcss.config.js
│   ├── tailwind.config.js
│   └── vite.config.js
│
├── docker-simple/              # Configuration Docker simplifiée
│   ├── docker-compose.yaml
│   ├── Dockerfile
│   ├── docker-run.sh
│   ├── nginx.conf
│   ├── prepare-coolify.sh
│   ├── supervisor.conf
│   ├── coolify-setup.md
│   └── README.md
│
├── releases/                   # ✨ Archives de distribution
│   ├── airboard-docker-v1.0.0-clean.tar.gz  (107 KB)
│   └── INSTALLATION-1.0.0-clean.txt
│
├── airboard.sh                 # ⭐ Script principal de gestion
├── clean-project.sh            # ✨ Script de nettoyage
├── create-docker-release.sh    # ✨ Script de création de release
│
├── docker-compose.yaml         # ⭐ Orchestration Docker (recommandé)
├── docker-compose.simple.yml
├── docker-compose.coolify.yaml
│
├── .env.example                # Template de configuration
│
├── README.md                   # Documentation principale
├── DOCKER-DEPLOYMENT.md        # ✨ Guide déploiement Docker complet
├── QUICK-START-DOCKER.md       # ✨ Démarrage rapide
├── COOLIFY-CUSTOM-COMMANDS.md
├── COOLIFY-FIX.md
└── LICENSE
```

## 📝 Fichiers essentiels

### 🎯 Pour le développement

| Fichier | Description |
|---------|-------------|
| `backend/main.go` | Point d'entrée API Go |
| `backend/handlers/admin.go` | Gestionnaire admin avec réinitialisation DB |
| `frontend/src/main.js` | Point d'entrée Vue.js |
| `frontend/src/services/api.js` | Client API |
| `frontend/src/views/admin/AdminDashboard.vue` | Dashboard admin avec bouton reset |

### 🐳 Pour Docker

| Fichier | Description |
|---------|-------------|
| `docker-compose.yaml` | Configuration complète (PostgreSQL + Backend + Frontend) |
| `backend/Dockerfile` | Build Go multi-stage optimisé |
| `frontend/Dockerfile` | Build Vue.js + Nginx |
| `frontend/nginx.conf` | Configuration Nginx avec cache et compression |

### 📚 Documentation

| Fichier | Description |
|---------|-------------|
| `README.md` | Documentation principale |
| `DOCKER-DEPLOYMENT.md` | Guide complet de déploiement Docker |
| `QUICK-START-DOCKER.md` | Démarrage rapide en 3 étapes |

### 🛠️ Scripts

| Script | Usage |
|--------|-------|
| `airboard.sh` | Gestion principale (start-docker, stop, logs, etc.) |
| `clean-project.sh` | Nettoyage du projet |
| `create-docker-release.sh` | Création d'une release distribuable |

## ✨ Nouveautés ajoutées

### 1. Réinitialisation de la base de données

**Backend** :
- `handlers/admin.go` : Fonction `ResetDatabase()` (lignes 1003-1091)
- `main.go` : Route `POST /admin/database/reset` (ligne 115)

**Frontend** :
- `services/api.js` : Méthode `resetDatabase()` (lignes 244-247)
- `views/admin/AdminDashboard.vue` :
  - Section "Zone de danger" (lignes 132-159)
  - Modal de confirmation (lignes 161-200)
  - Logique de réinitialisation (lignes 234-262)

### 2. Optimisations Docker

- `.dockerignore` pour backend et frontend
- Documentation complète du déploiement
- Scripts de création de release
- Script de nettoyage du projet

## 🚀 Commandes principales

```bash
# Développement
cd backend && go run main.go        # Backend
cd frontend && npm install && npm run dev  # Frontend

# Docker (recommandé)
./airboard.sh start-docker          # Démarrage complet
docker-compose up -d                # OU manuellement
docker-compose logs -f              # Voir les logs
docker-compose down                 # Arrêter

# Nettoyage
./clean-project.sh                  # Nettoyer le projet

# Release
./create-docker-release.sh 1.0.1   # Créer une release
```

## 📦 Distribution

La release `airboard-docker-v1.0.0-clean.tar.gz` contient tout le nécessaire pour déployer l'application :

```bash
# Extraire
tar -xzf airboard-docker-v1.0.0-clean.tar.gz
cd airboard

# Démarrer
./airboard.sh start-docker

# Accéder
# http://localhost:3000
```

## 🔑 Comptes par défaut

- **Admin** : `admin@airboard.com` / `admin123`
- **User** : `user@airboard.com` / `user123`

## 🔒 Sécurité en production

Avant le déploiement en production :

1. ✅ Changer `JWT_SECRET` dans `docker-compose.yaml`
2. ✅ Modifier les mots de passe PostgreSQL
3. ✅ Changer le mot de passe admin après connexion
4. ✅ Utiliser HTTPS avec un reverse proxy
5. ✅ Configurer des sauvegardes automatiques

---

**Version** : 1.0.0-clean
**Dernière mise à jour** : 13 novembre 2025
**Licence** : MIT
