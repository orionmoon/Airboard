# Guide de déploiement Coolify - Airboard

## 🚀 Déploiement rapide

### Étape 1 : Créer un nouveau projet

1. Connectez-vous à votre instance Coolify
2. Créez un nouveau projet ou sélectionnez un projet existant

### Étape 2 : Configurer le déploiement

1. **Type de déploiement** : Git Repository
2. **Source** : GitHub
3. **Repository** : `https://github.com/orionmoon/airboard.git`
4. **Branch** : `main`
5. **Build Pack** : Docker Compose

### Étape 3 : Configuration Docker Compose

**IMPORTANT** : Sélectionnez `docker-compose.prod.yaml` au lieu de `docker-compose.yaml`

Dans les paramètres du service :
- **Docker Compose File** : `docker-compose.prod.yaml`

**Pourquoi ?**
- `docker-compose.yaml` : Pour installation locale (expose le port 80 directement)
- `docker-compose.prod.yaml` : Pour Coolify (utilise le proxy Traefik de Coolify)

### Étape 4 : Variables d'environnement

Configurez les variables suivantes dans Coolify :

#### Variables obligatoires

```env
# Base de données
DB_NAME=airboard
DB_USER=airboard
DB_PASSWORD=VotreMotDePasseSecurise123!

# JWT (IMPORTANT : utilisez une clé longue et aléatoire)
JWT_SECRET=votre-cle-secrete-tres-longue-minimum-32-caracteres-aleatoires

# URL publique de votre application
PUBLIC_URL=https://votre-domaine.com
```

#### Variables SSO (si vous utilisez Authentik)

```env
SSO_ENABLED=true
SSO_AUTO_PROVISION=true
SSO_DEFAULT_ROLE=user
SSO_DEFAULT_GROUP=Common
SSO_ADMIN_GROUPS=airboard-admins
```

#### Variables OAuth (optionnelles, configurables via l'interface admin)

Les configurations OAuth (Google, Microsoft) se font maintenant directement via l'interface d'administration après le premier déploiement.

### Étape 5 : Déployer

1. Cliquez sur **Deploy**
2. Attendez la fin du build (peut prendre 5-10 minutes la première fois)
3. Vérifiez les logs pour confirmer le succès

### Étape 6 : Configuration post-déploiement

1. **Accédez à votre application** : `https://votre-domaine.com`
2. **Connectez-vous** avec les identifiants par défaut :
   - Username : `admin`
   - Password : `admin123`
3. **Changez immédiatement le mot de passe admin**
4. **Configurez OAuth** (si nécessaire) via Administration > OAuth

## 🔧 Architecture sur Coolify

```
Internet
  ↓
Coolify Proxy (Traefik)
  ↓
Frontend Container (nginx) - expose: 80
  ↓ (proxy /api/v1/*)
Backend Container (Go API) - expose: 8080
  ↓
PostgreSQL Container - expose: 5432
```

**Important** : Le frontend n'expose PAS le port 80 directement. Coolify gère le routage via son proxy Traefik.

## 📋 Vérifications

### Vérifier que le déploiement fonctionne

```bash
# Depuis le serveur Coolify, vérifier les containers
docker ps | grep airboard

# Vous devriez voir 3 containers :
# - airboard-postgres
# - airboard-backend
# - airboard-frontend
```

### Vérifier les logs

Dans Coolify :
1. Allez dans votre service
2. Onglet **Logs**
3. Vérifiez qu'il n'y a pas d'erreurs

**Logs attendus** :
- Frontend : nginx démarré sur le port 80
- Backend : serveur Go démarré sur le port 8080
- PostgreSQL : base de données prête

### Tester la connexion

1. Accédez à `https://votre-domaine.com`
2. Vous devriez voir la page de connexion
3. Testez la connexion avec `admin` / `admin123`

## 🐛 Dépannage

### Erreur : "Port 80 already allocated"

**Cause** : Vous utilisez `docker-compose.yaml` au lieu de `docker-compose.prod.yaml`

**Solution** :
1. Dans Coolify, allez dans les paramètres du service
2. Changez **Docker Compose File** vers `docker-compose.prod.yaml`
3. Redéployez

### Erreur : "Backend connection refused"

**Cause** : Le backend ne peut pas se connecter à PostgreSQL

**Solutions** :
1. Vérifiez que les 3 containers sont démarrés : `docker ps`
2. Vérifiez les variables d'environnement (DB_HOST doit être `postgres`)
3. Attendez que PostgreSQL soit prêt (healthcheck)

### Erreur : "Cannot connect to database"

**Cause** : Variables d'environnement de base de données incorrectes

**Solutions** :
1. Vérifiez `DB_USER`, `DB_PASSWORD`, `DB_NAME` dans Coolify
2. Assurez-vous qu'ils correspondent entre tous les services
3. Redéployez après modification

### Page blanche / Erreur 404

**Cause** : Le frontend ne peut pas communiquer avec le backend

**Solutions** :
1. Vérifiez que nginx est configuré correctement
2. Vérifiez les logs du frontend : `docker logs airboard-frontend`
3. Vérifiez que le backend répond : `docker logs airboard-backend`

## 🔄 Mise à jour

Pour mettre à jour Airboard sur Coolify :

1. **Depuis Coolify** :
   - Cliquez sur **Redeploy** dans votre service
   - Coolify va récupérer les derniers commits de GitHub

2. **Forcer le rebuild** :
   - Si des modifications du Dockerfile ont été faites
   - Cochez **Force Rebuild** avant de redéployer

3. **Vérifier après la mise à jour** :
   - Consultez les logs
   - Testez la connexion
   - Vérifiez que vos données sont préservées

## 🔐 Sécurité

### Bonnes pratiques

1. **Changez les mots de passe par défaut** :
   - Admin : changez `admin123` immédiatement
   - Base de données : utilisez un mot de passe fort

2. **Utilisez des secrets forts** :
   ```bash
   # Générer un JWT_SECRET aléatoire
   openssl rand -hex 32
   ```

3. **Activez HTTPS** :
   - Coolify gère automatiquement Let's Encrypt
   - Assurez-vous que HTTPS est activé dans les paramètres du service

4. **Limitez les accès** :
   - Configurez les groupes Authentik pour limiter l'accès admin
   - Utilisez SSO au lieu des mots de passe locaux

## 📊 Monitoring

### Logs en temps réel

```bash
# Sur le serveur Coolify
docker logs -f airboard-frontend
docker logs -f airboard-backend
docker logs -f airboard-postgres
```

### Vérifier l'état des services

```bash
docker ps | grep airboard
docker stats $(docker ps --filter name=airboard --format "{{.Names}}")
```

### Base de données

```bash
# Connexion à PostgreSQL
docker exec -it airboard-postgres psql -U airboard -d airboard

# Vérifier le nombre d'utilisateurs
SELECT COUNT(*) FROM users;

# Vérifier le nombre d'applications
SELECT COUNT(*) FROM applications;
```

## 🆘 Support

En cas de problème :

1. Consultez les logs dans Coolify
2. Vérifiez la [documentation principale](README.md)
3. Consultez [DOCKER_INSTALL.md](DOCKER_INSTALL.md) pour le dépannage
4. Ouvrez une issue sur GitHub

## 🎯 Checklist de déploiement

- [ ] Projet créé dans Coolify
- [ ] Repository GitHub configuré
- [ ] `docker-compose.prod.yaml` sélectionné
- [ ] Variables d'environnement configurées
- [ ] `DB_PASSWORD` modifié (pas la valeur par défaut)
- [ ] `JWT_SECRET` généré de manière aléatoire
- [ ] `PUBLIC_URL` configuré avec votre domaine
- [ ] Déploiement réussi
- [ ] Containers démarrés (postgres, backend, frontend)
- [ ] Application accessible via HTTPS
- [ ] Connexion admin testée
- [ ] Mot de passe admin changé
- [ ] SSO configuré (si applicable)
