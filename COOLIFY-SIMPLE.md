# 🚀 Guide Coolify Simplifié - Déploiement Complet

## ⚡ Avantage : TOUT est automatique

Avec cette configuration, **PostgreSQL est inclus** dans le docker-compose !

✅ **Pas besoin de créer la base de données manuellement**
✅ **Un seul service à déployer**
✅ **Configuration automatique**
✅ **Données de démo créées automatiquement**

---

## 📋 Méthode 1 : Déploiement depuis GitHub (RECOMMANDÉ)

### Prérequis
- Un compte GitHub
- Votre code publié sur GitHub

### Étape 1 : Publier sur GitHub

```bash
cd /home/abdelaziz/Desktop/Temporary/Airboard/airboard-airboard

# Initialiser Git
git init
git add .
git commit -m "Initial commit - Airboard v1.0.0"

# Créer le repository sur github.com puis :
git remote add origin https://github.com/VOTRE-USERNAME/airboard.git
git branch -M main
git push -u origin main
```

### Étape 2 : Configurer Coolify

1. **New Resource** → **Service**
2. **Source Type** : **Git Repository**

**Configuration Git** :
```
Provider: GitHub
Repository: https://github.com/VOTRE-USERNAME/airboard.git
Branch: main
```

**Configuration Build** :
```
Type: Docker Compose
Docker Compose File: docker-simple/docker-compose.yaml
```

**Configuration Port** :
```
Port: 80
```

**Health Check** :
```
Path: /health
Method: GET
Interval: 30s
```

### Étape 3 : Variables d'environnement (OPTIONNEL)

Les valeurs par défaut fonctionnent, mais pour la production :

```bash
# Mot de passe base de données (défaut: airboard123)
DB_PASSWORD=votre-mot-de-passe-super-securise

# Secret JWT (défaut: change-this-secret-in-production)
JWT_SECRET=votre-secret-jwt-tres-long-et-securise

# Mode (défaut: release)
GIN_MODE=release
```

**Générer un JWT_SECRET sécurisé** :
```bash
openssl rand -base64 32
```

### Étape 4 : Déployer

1. Cliquez sur **Deploy**
2. Suivez les logs en temps réel
3. Attendez 2-3 minutes (build + démarrage)
4. Status → "Running" → **C'est prêt !** 🎉

### Étape 5 : Accéder

**URL** : Générée par Coolify (ex: `https://airboard-xyz.coolify.app`)

**Connexion** :
- **Admin** : `admin@airboard.com` / `admin123`
- **User** : `user@airboard.com` / `user123`

⚠️ **Changez immédiatement le mot de passe admin !**

---

## 📋 Méthode 2 : Déploiement Direct (sans Git)

### Étape 1 : Préparer l'archive

```bash
cd /home/abdelaziz/Desktop/Temporary/Airboard

# Créer une archive
tar -czf airboard-deploy.tar.gz \
  --exclude='node_modules' \
  --exclude='.git' \
  --exclude='dist' \
  airboard-airboard/
```

### Étape 2 : Upload sur Coolify

1. **New Resource** → **Service**
2. **Source Type** : **Docker Compose**
3. **Upload** : `airboard-deploy.tar.gz`
4. **Docker Compose File** : `docker-simple/docker-compose.yaml`
5. **Port** : 80
6. **Deploy**

### Étape 3 : Variables et déploiement

Même configuration que la Méthode 1 (étapes 3-5)

---

## 🔍 Vérification du déploiement

### 1. Vérifier les logs

Dans Coolify → Service → **Logs**

Vous devriez voir :
```
✅ PostgreSQL started
✅ Database initialized
✅ Admin user created: admin@airboard.com
✅ Demo groups created
✅ Backend started on port 8080
✅ Nginx started on port 80
🚀 Airboard is ready!
```

### 2. Tester le health check

```bash
curl https://votre-url.coolify.app/health
```

Réponse attendue :
```json
{
  "status": "ok",
  "message": "Airboard API is running"
}
```

### 3. Tester la connexion

1. Ouvrir l'URL dans le navigateur
2. Se connecter avec `admin@airboard.com` / `admin123`
3. Vous devriez voir le tableau de bord avec les groupes de démo

---

## 🔐 Sécurité en production

### Checklist de sécurité

- [ ] **Mot de passe admin changé**
  - Se connecter → Profil → Changer le mot de passe

- [ ] **JWT_SECRET personnalisé**
  - Générer : `openssl rand -base64 32`
  - Ajouter dans Variables d'environnement
  - Redéployer

- [ ] **DB_PASSWORD personnalisé**
  - Choisir un mot de passe fort
  - Ajouter dans Variables d'environnement
  - Redéployer

- [ ] **HTTPS activé**
  - Ajouter domaine personnalisé
  - Let's Encrypt automatique

- [ ] **Sauvegardes configurées**
  - Coolify → Service → Backups
  - Fréquence : Quotidienne
  - Rétention : 7 jours minimum

---

## 🚀 Configuration avancée

### Auto-Deploy (GitHub uniquement)

1. Dans Coolify → Service → **Settings**
2. Activer **Auto Deploy on Push**
3. Chaque `git push` → redéploiement automatique

Coolify créé automatiquement un webhook sur GitHub.

### Domaine personnalisé

1. **Domains** → **Add Domain**
2. Entrer : `airboard.votre-domaine.com`
3. Configurer DNS :
   ```
   Type: A
   Name: airboard
   Value: [IP de votre serveur Coolify]
   ```
4. Attendre propagation DNS (quelques minutes)
5. HTTPS via Let's Encrypt : **Automatique** ✅

### Resource Limits

Configuration recommandée :
```
CPU: 1-2 cores
Memory: 1 GB (minimum 512 MB)
Disk: 5 GB
```

Dans Coolify → Service → **Resources** → Ajuster les limites

---

## 📊 Monitoring

### Metrics disponibles

Coolify affiche automatiquement :
- **CPU Usage** (utilisation processeur)
- **Memory Usage** (utilisation mémoire)
- **Disk Usage** (utilisation disque)
- **Network I/O** (trafic réseau)
- **Uptime** (temps de fonctionnement)

### Logs en temps réel

```
Service → Logs → Toggle "Follow Logs"
```

Vous verrez tous les logs de :
- PostgreSQL
- Backend (API Go)
- Nginx

---

## 🔄 Mise à jour

### Via GitHub (avec Auto-Deploy)

```bash
# Faire des modifications
vim frontend/src/...

# Tester localement
./airboard.sh start-dev

# Commit et push
git add .
git commit -m "Add new feature"
git push origin main

# → Coolify déploie automatiquement ! 🎉
```

### Manuellement

Dans Coolify :
1. **Restart** : Redémarrer sans rebuild
2. **Rebuild** : Rebuild complet + redémarrage

---

## 🐛 Dépannage

### Service ne démarre pas

**Erreur** : Build failed

**Vérifier** :
```
Logs → Chercher les erreurs de build
```

**Solutions communes** :
- Vérifier que `docker-simple/docker-compose.yaml` existe
- Vérifier que `docker-simple/Dockerfile` existe
- Vérifier les logs du build

### PostgreSQL connection failed

**Erreur** : `Cannot connect to postgres`

**Solution** :
1. Attendre 30 secondes (PostgreSQL peut prendre du temps à démarrer)
2. Vérifier les logs PostgreSQL :
   ```
   Logs → Filtrer "postgres"
   ```
3. Redémarrer le service si nécessaire

### Health check failed

**Erreur** : `Health check timeout`

**Solutions** :
1. Vérifier que le backend est démarré :
   ```bash
   curl https://votre-url/health
   ```
2. Vérifier les logs du backend
3. Augmenter le timeout du health check (60s au lieu de 30s)

### Port already in use

**Erreur** : `Port 80 already in use`

**Solution** :
1. Coolify → Service → **Port**
2. Changer pour un autre port (ex: 8080)
3. Redéployer

---

## 💾 Sauvegardes et restauration

### Configuration des sauvegardes

Dans Coolify → Service → **Backups** :

```
Frequency: Daily
Time: 02:00 AM
Retention: 7 days (ou plus)
Destination: S3 / Local / Autre
```

### Backup manuel PostgreSQL

```bash
# Se connecter au conteneur
docker exec -it airboard-db sh

# Créer un backup
pg_dump -U airboard airboard > /tmp/backup.sql

# Copier le backup
docker cp airboard-db:/tmp/backup.sql ./backup-$(date +%Y%m%d).sql
```

### Restauration

```bash
# Copier le backup dans le conteneur
docker cp backup.sql airboard-db:/tmp/

# Restaurer
docker exec -it airboard-db sh
psql -U airboard airboard < /tmp/backup.sql
```

---

## 📈 Scaling

### Augmenter les ressources

Si votre application devient plus utilisée :

1. **CPU** : Augmenter à 2-4 cores
2. **Memory** : Augmenter à 2-4 GB
3. **Disk** : Surveiller l'utilisation PostgreSQL

### Load Balancing (Coolify Pro)

Pour gérer encore plus de trafic :
- Activer le load balancing dans Coolify Pro
- Déployer plusieurs instances
- Utiliser un CDN pour les assets statiques

---

## ✅ Checklist complète

### Avant le déploiement
- [x] Code testé localement
- [x] Docker Compose testé
- [x] Documentation lue
- [x] Variables d'environnement préparées

### Déploiement
- [ ] Repository GitHub créé (si méthode 1)
- [ ] Service Coolify créé
- [ ] Docker Compose configuré
- [ ] Variables d'environnement ajoutées
- [ ] Deploy réussi
- [ ] Logs vérifiés

### Post-déploiement
- [ ] Health check OK
- [ ] Connexion admin testée
- [ ] Mot de passe admin changé
- [ ] JWT_SECRET changé
- [ ] DB_PASSWORD changé
- [ ] HTTPS activé
- [ ] Domaine configuré (optionnel)
- [ ] Auto-deploy activé (optionnel)
- [ ] Sauvegardes configurées

---

## 🎁 Ce qui est inclus automatiquement

Avec ce docker-compose, vous obtenez **TOUT** en un seul service :

| Composant | Version | Description |
|-----------|---------|-------------|
| **PostgreSQL** | 15-alpine | Base de données |
| **Backend** | Go 1.21 | API RESTful |
| **Frontend** | Vue.js 3 | Interface utilisateur |
| **Nginx** | Alpine | Reverse proxy |
| **Supervisor** | Latest | Gestionnaire de services |
| **Données de démo** | - | Créées automatiquement |

**Taille totale** : ~80 MB (image optimisée)

**Aucune configuration manuelle requise !**

---

## 🎉 Résumé

### Ce que vous avez

- ✅ Application complète en **un seul docker-compose**
- ✅ PostgreSQL **inclus** (pas de service séparé)
- ✅ Configuration **automatique**
- ✅ Données de démo **créées automatiquement**
- ✅ Health check **configuré**
- ✅ HTTPS **automatique** (avec domaine)
- ✅ Auto-deploy **disponible** (GitHub)

### Temps de déploiement

- **Méthode GitHub** : 5 minutes + 3 minutes de build = **8 minutes**
- **Méthode Direct** : 3 minutes + 3 minutes de build = **6 minutes**

### Prochaines étapes

1. **Tester l'application** : Se connecter et explorer
2. **Changer les secrets** : Admin password, JWT_SECRET, DB_PASSWORD
3. **Configurer HTTPS** : Ajouter un domaine
4. **Activer Auto-deploy** : Pour déployer à chaque push (GitHub)
5. **Configurer les backups** : Pour sauvegarder PostgreSQL

---

## 📚 Documentation

- [QUICK-START-COOLIFY.md](QUICK-START-COOLIFY.md) - Quick start 3 minutes
- [GITHUB-COOLIFY-SETUP.md](GITHUB-COOLIFY-SETUP.md) - GitHub + Auto-deploy
- [DEPLOYMENT-READY.md](DEPLOYMENT-READY.md) - Vue d'ensemble
- [PROJECT-STRUCTURE.md](PROJECT-STRUCTURE.md) - Structure du projet
- [DOCKER-DEPLOYMENT.md](DOCKER-DEPLOYMENT.md) - Docker en local

---

**🎉 Airboard est prêt pour Coolify ! Tout est automatique ! 🚀**

**Version** : 1.0.0-coolify-ready
**Dernière mise à jour** : 13 novembre 2025
