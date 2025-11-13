# 🚀 Démarrage rapide Coolify - 3 minutes

## ✨ Installation en 3 étapes (TOUT inclus)

### Étape 1 : Créer le service (1 min)

1. Dans Coolify → **New Resource** → **Service**
2. Type : **Git Repository** ou **Docker Compose**

#### Option A : Depuis GitHub (recommandé)
```
Repository: https://github.com/votre-username/airboard.git
Branch: main
Type: Docker Compose
Docker Compose File: docker-simple/docker-compose.yaml
```

#### Option B : Upload direct
```
Upload le dossier du projet
Type: Docker Compose
Docker Compose File: docker-simple/docker-compose.yaml
```

### Étape 2 : Variables d'environnement (1 min)

**OPTIONNEL** - Valeurs par défaut fonctionnent directement !

Pour la production, ajoutez :
```bash
DB_PASSWORD=votre-mot-de-passe-securise
JWT_SECRET=votre-secret-jwt-genere
```

Générer JWT_SECRET :
```bash
openssl rand -base64 32
```

### Étape 3 : Déployer (1 min)

1. Cliquez sur **Deploy**
2. Attendez 2-3 minutes ⏱️
3. C'est prêt ! 🎉

## 🎯 Accéder à l'application

**URL** : Générée automatiquement par Coolify

**Connexion** :
- **Admin** : `admin@airboard.com` / `admin123`
- **User** : `user@airboard.com` / `user123`

⚠️ **Changez le mot de passe admin immédiatement !**

---

## 🔧 Ce qui est inclus automatiquement

- ✅ **PostgreSQL 15** (base de données)
- ✅ **Backend Go** (API)
- ✅ **Frontend Vue.js** (interface)
- ✅ **Nginx** (reverse proxy)
- ✅ **Données de démo** (créées automatiquement)
- ✅ **Health check** configuré
- ✅ **Volumes persistants** pour la base de données

**Aucune configuration manuelle requise !**

---

## 📋 Configuration avancée (optionnel)

### Domaine personnalisé

1. Dans le service → **Domains**
2. Ajouter : `airboard.votre-domaine.com`
3. Configurer DNS (A record → IP Coolify)
4. HTTPS automatique (Let's Encrypt)

### Auto-deploy

1. Activer **Auto Deploy** dans Coolify
2. Chaque `git push` → redéploiement automatique

### Sauvegardes

Les données PostgreSQL sont dans un volume Docker persistant :
- Volume : `postgres_data`
- Configurer les backups dans Coolify → Backups

---

## 🔐 Sécurité en production

### ⚠️ IMPORTANT - À faire après le déploiement

1. **Changer le mot de passe admin**
   - Se connecter comme admin
   - Profil → Changer le mot de passe

2. **Changer les secrets** (recommandé)
   - Ajouter `DB_PASSWORD` personnalisé
   - Ajouter `JWT_SECRET` personnalisé
   - Redéployer

3. **Activer HTTPS**
   - Ajouter un domaine personnalisé
   - Let's Encrypt automatique

---

## 🐛 Dépannage rapide

### Le service ne démarre pas

**Vérifier les logs** :
```
Service → Logs → Chercher les erreurs
```

**Erreurs communes** :
- Build en cours : Attendre la fin du build
- Port déjà utilisé : Changer le port dans Coolify

### L'application ne répond pas

**Vérifier le health check** :
```bash
curl https://votre-url.com/health
```

Réponse attendue :
```json
{"status":"ok","message":"Airboard API is running"}
```

### Cannot connect to database

**Solution** : Redémarrer le service
```
Service → Restart
```

PostgreSQL prend parfois 10-20 secondes à démarrer.

---

## 📊 Monitoring

Dans Coolify, vous pouvez voir :
- **CPU usage** (utilisation processeur)
- **Memory usage** (utilisation mémoire)
- **Logs en temps réel** (logs live)
- **Health status** (statut santé)

---

## 🔄 Mise à jour

### Via GitHub (Auto-deploy activé)
```bash
git add .
git commit -m "Update"
git push origin main
# → Coolify déploie automatiquement
```

### Manuellement
```
Service → Restart ou Rebuild
```

---

## 💡 Pourquoi c'est si simple ?

Airboard utilise **Docker Compose** qui inclut :
- PostgreSQL (base de données)
- Application complète (Frontend + Backend + Nginx)
- Configuration automatique
- Health checks
- Volumes persistants

**Tout est pré-configuré !** Aucune étape manuelle nécessaire.

---

## 📚 Documentation complète

- [COOLIFY-DEPLOYMENT.md](COOLIFY-DEPLOYMENT.md) - Guide détaillé
- [GITHUB-COOLIFY-SETUP.md](GITHUB-COOLIFY-SETUP.md) - GitHub + Auto-deploy
- [DOCKER-DEPLOYMENT.md](DOCKER-DEPLOYMENT.md) - Docker en local
- [DEPLOYMENT-READY.md](DEPLOYMENT-READY.md) - Vue d'ensemble

---

## ✅ Checklist de déploiement

- [ ] Service créé dans Coolify
- [ ] Variables d'environnement configurées (optionnel)
- [ ] Deploy lancé
- [ ] Logs vérifiés (pas d'erreurs)
- [ ] Health check OK
- [ ] Connexion admin testée
- [ ] Mot de passe admin changé
- [ ] HTTPS activé (domaine personnalisé)
- [ ] Sauvegardes configurées

---

**C'est tout !** 🎉

**Temps total** : 3 minutes + 2-3 minutes de build = **~5 minutes**

Guide complet : [COOLIFY-DEPLOYMENT.md](COOLIFY-DEPLOYMENT.md)
