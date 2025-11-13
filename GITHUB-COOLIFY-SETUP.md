# 🚀 Guide : GitHub → Coolify (Auto-installation)

## ✅ OUI, c'est la méthode RECOMMANDÉE !

En publiant sur GitHub et en donnant le lien à Coolify, vous obtenez :
- ✅ **Déploiement automatique** : Coolify clone et build automatiquement
- ✅ **Auto-deploy** : Chaque `git push` redéploie automatiquement
- ✅ **Versioning** : Historique complet de vos modifications
- ✅ **Collaboration** : Facile de travailler en équipe
- ✅ **Rollback facile** : Revenir à une version précédente en 1 clic
- ✅ **PostgreSQL inclus** : Pas de service séparé à créer !

---

## 📋 Étape 1 : Publier sur GitHub

### A. Créer un repository sur GitHub

1. Allez sur **https://github.com/new**
2. Configuration :
   ```
   Repository name: airboard
   Description: Modern application portal with Vue.js and Go
   Public OU Private (les deux fonctionnent avec Coolify)
   ✅ Add README (déjà présent dans le projet)
   ✅ Add .gitignore (déjà présent)
   ✅ Choose a license: MIT (déjà présent)
   ```
3. Cliquez sur **Create repository**

### B. Initialiser Git localement

```bash
cd /home/abdelaziz/Desktop/Temporary/Airboard/airboard-airboard

# Initialiser Git
git init

# Ajouter tous les fichiers
git add .

# Premier commit
git commit -m "Initial commit - Airboard v1.0.0-coolify-ready

Features:
- Application portal with groups
- User management and permissions
- JWT authentication
- Multi-language support (FR/EN/ES/AR)
- Dark mode
- Database reset functionality
- Docker ready with PostgreSQL included
- Coolify compatible with Docker Compose
"

# Ajouter le remote GitHub (remplacer YOUR_USERNAME)
git remote add origin https://github.com/YOUR_USERNAME/airboard.git

# Push vers GitHub
git branch -M main
git push -u origin main
```

### C. Pour un repository privé (optionnel)

Si votre repository est privé, Coolify a besoin d'un accès :

**Méthode 1 : Deploy Key (recommandé)**
1. Dans Coolify → Settings → SSH Keys
2. Copier la clé publique
3. GitHub → Settings → Deploy keys → Add deploy key
4. Coller la clé et activer "Allow write access"

**Méthode 2 : Personal Access Token**
1. GitHub → Settings → Developer settings → Personal access tokens
2. Generate new token avec scope `repo`
3. Copier le token
4. Coolify → Utiliser HTTPS avec token

---

## 📋 Étape 2 : Déployer sur Coolify avec GitHub

### Configuration du service

1. **New Resource** → **Service**
2. **Source Type** : **Git Repository**

#### Configuration Git
```
Git Provider: GitHub
Repository URL: https://github.com/YOUR_USERNAME/airboard.git
Branch: main
```

Si repository privé :
- Ajouter Deploy Key OU
- Utiliser Personal Access Token

#### Configuration Build
```
Build Pack: Docker Compose
Docker Compose File: docker-simple/docker-compose.yaml
```

#### Configuration Port
```
Port: 80
```

#### Health Check
```
Path: /health
Method: GET
Port: 80
Interval: 30s
```

### Variables d'environnement (OPTIONNEL)

**Les valeurs par défaut fonctionnent directement !**

Pour la production, ajoutez dans **Environment Variables** :

```bash
# Mot de passe base de données (défaut: airboard123)
DB_PASSWORD=votre-mot-de-passe-super-securise

# Secret JWT (défaut: change-this-secret-in-production)
JWT_SECRET=votre-secret-jwt-tres-long-et-securise

# Mode (défaut: release)
GIN_MODE=release
```

**💡 Générer un JWT_SECRET sécurisé** :
```bash
openssl rand -base64 32
```

### Configuration avancée (optionnel)

**Auto Deploy** :
- Activez **Auto Deploy** pour déployer automatiquement à chaque `git push`

**Custom Domain** :
- Ajoutez votre domaine dans **Domains**
- HTTPS sera automatiquement configuré (Let's Encrypt)

**Resource Limits** :
```
CPU: 1 core (2 cores recommandé)
Memory: 1 GB (512 MB minimum)
```

### Déployer

1. Cliquez sur **Deploy**
2. Suivez les logs en temps réel
3. Attendez 2-3 minutes (build + start)
4. Status devient "Running" → C'est prêt ! 🎉

---

## 📋 Étape 3 : Accéder à l'application

### URL générée par Coolify

Coolify génère automatiquement une URL :
```
https://airboard-[random].coolify.app
```

### Ajouter un domaine personnalisé (optionnel)

1. Dans votre service Coolify → **Domains**
2. Cliquez sur **Add Domain**
3. Entrez : `airboard.votre-domaine.com`
4. Configurez votre DNS :
   ```
   Type: A
   Name: airboard
   Value: [IP de votre serveur Coolify]
   ```
5. Attendez la propagation DNS (quelques minutes)
6. HTTPS sera automatiquement configuré (Let's Encrypt)

### Première connexion

URL : `https://votre-domaine.coolify.app`

**Comptes par défaut** :
- **Admin** : `admin@airboard.com` / `admin123`
- **User** : `user@airboard.com` / `user123`

⚠️ **IMPORTANT** : Changez immédiatement le mot de passe admin !

---

## 🔄 Workflow de développement

### Faire des modifications

```bash
# 1. Modifier le code localement
vim frontend/src/...

# 2. Tester localement
./airboard.sh start-dev

# 3. Commit
git add .
git commit -m "Add new feature"

# 4. Push vers GitHub
git push origin main

# 5. Coolify déploie automatiquement (si Auto Deploy activé)
# Sinon, cliquer manuellement sur "Restart" dans Coolify
```

### Auto-deploy activé

Si vous avez activé **Auto Deploy** dans Coolify :
- Chaque `git push` → Coolify détecte le changement
- Build automatique de la nouvelle version
- Déploiement sans interruption (zero-downtime)
- Logs disponibles en temps réel

---

## 🔐 Sécurité

### Secrets à ne PAS commiter

Le fichier `.gitignore` (déjà présent) exclut :

```gitignore
# Environnement
.env
.env.local
.env.*.local

# Node
node_modules/
dist/
*.log

# Go
backend/backend
backend/backend-test
backend/main

# IDE
.vscode/
.idea/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db
```

### Variables d'environnement sensibles

**❌ NE PAS faire** :
```yaml
# docker-compose.yaml
JWT_SECRET: mon-secret-en-dur  # DANGEREUX !
```

**✅ FAIRE** :
```yaml
# docker-compose.yaml
JWT_SECRET: ${JWT_SECRET:-change-this-secret-in-production}
```

Et configurer dans Coolify → Environment Variables

---

## 🐛 Dépannage

### Build échoue

**Erreur** : `Cannot find docker-compose file`

**Solution** :
```
Vérifier Docker Compose File: docker-simple/docker-compose.yaml
Build Context: . (point = racine)
```

### Cannot connect to database

**Erreur** : `dial tcp: lookup postgres: no such host`

**Cause** : PostgreSQL pas encore démarré

**Solution** :
```
Attendre 30 secondes (PostgreSQL démarre)
Redémarrer le service si nécessaire
```

### Health check échoue

**Erreur** : `Health check failed: connection refused`

**Solution** :
```
Vérifier :
1. Port exposé : 80
2. Health path : /health
3. Service bien démarré (logs)
4. Attendre que tous les services soient up (PostgreSQL + Backend + Nginx)
```

### Webhook ne fonctionne pas (Auto Deploy)

**Solution** :
```
1. GitHub → Settings → Webhooks
2. Vérifier l'URL du webhook Coolify
3. Vérifier les logs de delivery
4. Re-créer le webhook si nécessaire
```

---

## 📊 Monitoring

### Logs en temps réel

Dans Coolify :
```
Service → Logs → Toggle "Follow Logs"
```

Vous verrez :
```
✅ PostgreSQL started
✅ Database initialized
✅ Admin user created: admin@airboard.com
✅ Demo groups created
✅ Backend started on port 8080
✅ Nginx started on port 80
🚀 Airboard is ready!
```

### Metrics

Coolify affiche automatiquement :
- CPU usage
- Memory usage
- Network I/O
- Disk usage

---

## 🚀 Avantages de cette méthode

| Avantage | Description |
|----------|-------------|
| ✅ **Auto-deploy** | Push → Deploy automatique |
| ✅ **Versioning** | Historique complet Git |
| ✅ **Rollback** | Revenir à une version précédente |
| ✅ **Collaboration** | Plusieurs développeurs |
| ✅ **CI/CD** | Intégration continue |
| ✅ **Sauvegardes** | Code sauvegardé sur GitHub |
| ✅ **HTTPS auto** | Let's Encrypt gratuit |
| ✅ **Zero config** | Coolify détecte tout |
| ✅ **PostgreSQL inclus** | Pas de service séparé ! |

---

## 📝 Checklist complète

### Avant de publier sur GitHub
- [x] Nettoyer le projet
- [x] Vérifier `.gitignore`
- [x] Tester localement
- [x] Documenter (README.md)
- [x] Supprimer les secrets

### Publication GitHub
- [ ] Créer repository sur GitHub
- [ ] `git init` local
- [ ] `git add .`
- [ ] `git commit -m "Initial commit"`
- [ ] `git remote add origin`
- [ ] `git push -u origin main`

### Déploiement Coolify
- [ ] Créer service dans Coolify
- [ ] Configurer Git (URL + Branch)
- [ ] Configurer Build (Docker Compose)
- [ ] Ajouter variables d'environnement (optionnel)
- [ ] Générer JWT_SECRET (pour production)
- [ ] Deploy
- [ ] Vérifier logs
- [ ] Tester l'accès

### Après déploiement
- [ ] Changer mot de passe admin
- [ ] Configurer domaine personnalisé
- [ ] Activer Auto Deploy
- [ ] Configurer sauvegardes
- [ ] Tester la réinitialisation DB

---

## 🎁 Template de README GitHub

Voici ce que vous pouvez ajouter à votre `README.md` pour GitHub :

```markdown
## 🚀 Déploiement rapide sur Coolify

### Prérequis
- Un serveur Coolify
- 5 minutes de votre temps

### Étapes (3 clics)
1. Dans Coolify → **New Resource** → **Service** → **Git Repository**
2. Configuration :
   ```
   Repository: https://github.com/YOUR_USERNAME/airboard.git
   Type: Docker Compose
   Docker Compose File: docker-simple/docker-compose.yaml
   Port: 80
   ```
3. (Optionnel) Variables d'environnement :
   ```bash
   DB_PASSWORD=<votre-mot-de-passe>
   JWT_SECRET=<votre-secret>
   ```
4. **Deploy** → Attendez 2-3 min → C'est prêt ! 🎉

**✨ PostgreSQL est inclus automatiquement !**

Guide complet : [QUICK-START-COOLIFY.md](QUICK-START-COOLIFY.md)

### One-click deploy
[![Deploy on Coolify](https://img.shields.io/badge/Deploy%20on-Coolify-blue)](https://coolify.io)
```

---

## 🎉 Résumé

**GitHub + Coolify = Déploiement parfait !**

1. Publier sur GitHub : **5 minutes**
2. Configurer Coolify : **3 minutes**
3. **Auto-deploy** : Activé ✅
4. **HTTPS** : Automatique ✅
5. **Monitoring** : Intégré ✅
6. **PostgreSQL** : Inclus ✅

**Total** : 8 minutes pour un déploiement production-ready complet !

**Aucune étape manuelle de création de base de données !**

---

## 💡 Ce qui est inclus automatiquement

Avec le docker-compose, vous obtenez **TOUT** en un seul service :

- ✅ **PostgreSQL 15** (base de données)
- ✅ **Backend Go** (API RESTful)
- ✅ **Frontend Vue.js** (interface)
- ✅ **Nginx** (reverse proxy)
- ✅ **Données de démo** (créées au démarrage)
- ✅ **Health checks** (monitoring)
- ✅ **Volumes persistants** (données sauvegardées)

**Tout est pré-configuré !** Aucune configuration manuelle nécessaire.

---

**Documentation complète** :
- [QUICK-START-COOLIFY.md](QUICK-START-COOLIFY.md) - Quick start 3 minutes
- [COOLIFY-SIMPLE.md](COOLIFY-SIMPLE.md) - Guide détaillé simplifié
- [DEPLOYMENT-READY.md](DEPLOYMENT-READY.md) - Vue d'ensemble

---

**Version** : 1.0.0-coolify-ready
**Date** : 13 novembre 2025

---

🚀 **Prêt pour GitHub + Coolify ! Tout est automatique !** 🎉
