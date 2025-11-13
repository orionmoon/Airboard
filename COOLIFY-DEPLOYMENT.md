# 🚀 Guide de déploiement Coolify - Airboard

Ce guide vous permet de déployer Airboard sur Coolify en quelques minutes.

---

## 📋 Prérequis

- Un serveur Coolify configuré
- Accès à un repository Git (GitHub, GitLab, etc.)
- 1 GB RAM minimum
- 2 GB espace disque

---

## 🎯 Méthode 1 : Déploiement Simple (RECOMMANDÉ)

### Étape 1 : Créer le service PostgreSQL

1. Dans Coolify, créez un nouveau service **Database**
2. Sélectionnez **PostgreSQL 15**
3. Configuration :
   ```
   Name: airboard-db
   Database: airboard
   Username: airboard
   Password: [généré automatiquement par Coolify]
   ```
4. Déployez la base de données

### Étape 2 : Créer le service Airboard

1. Créez un nouveau **Service**
2. Type de source : **Git Repository**
3. Configuration :

   **Repository** :
   ```
   URL: https://github.com/votre-username/airboard.git
   Branch: main
   ```

   **Build** :
   ```
   Build Pack: Docker
   Dockerfile Location: docker-simple/Dockerfile
   Build Context: .
   ```

   **Port** :
   ```
   Port: 80
   ```

   **Health Check** :
   ```
   Path: /health
   Method: GET
   Port: 80
   ```

### Étape 3 : Configurer les variables d'environnement

Dans les **Environment Variables** du service Airboard :

```bash
# Base de données (copier depuis le service PostgreSQL)
DB_HOST=airboard-db-container-name  # Nom du conteneur PostgreSQL
DB_PORT=5432
DB_NAME=airboard
DB_USER=airboard
DB_PASSWORD=<copier depuis service PostgreSQL>

# JWT Secret (générer un secret sécurisé)
JWT_SECRET=votre-secret-tres-long-et-securise-ici-minimum-32-caracteres

# Mode
GIN_MODE=release
```

**💡 Astuce** : Pour `DB_HOST`, utilisez le nom du service PostgreSQL que Coolify a créé.

### Étape 4 : Déployer

1. Cliquez sur **Deploy**
2. Attendez la fin du build (~2-3 minutes)
3. Coolify générera automatiquement une URL

### Étape 5 : Configurer le domaine (optionnel)

1. Dans **Domains**, ajoutez votre domaine personnalisé
2. Activez **HTTPS** (Let's Encrypt automatique)

---

## 🎯 Méthode 2 : Avec Docker Compose

### Prérequis

Créez un fichier `docker-compose.coolify.yaml` à la racine du projet :

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: airboard
      POSTGRES_USER: airboard
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U airboard"]
      interval: 10s
      timeout: 5s
      retries: 5

  airboard:
    build:
      context: .
      dockerfile: docker-simple/Dockerfile
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_NAME=airboard
      - DB_USER=airboard
      - DB_PASSWORD=${DB_PASSWORD}
      - JWT_SECRET=${JWT_SECRET}
      - GIN_MODE=release
    ports:
      - "80:80"
    depends_on:
      postgres:
        condition: service_healthy

volumes:
  postgres_data:
```

### Configuration Coolify

1. Type : **Docker Compose**
2. Docker Compose File : `docker-compose.coolify.yaml`
3. Variables d'environnement :
   ```bash
   DB_PASSWORD=votre-mot-de-passe-securise
   JWT_SECRET=votre-secret-jwt-tres-long
   ```

---

## 🔑 Connexion après déploiement

Une fois déployé, accédez à votre application :

**URL** : `https://votre-domaine.com` ou `https://generated-by-coolify.app`

**Comptes par défaut** :
- **Admin** : `admin@airboard.com` / `admin123`
- **User** : `user@airboard.com` / `user123`

⚠️ **IMPORTANT** : Changez le mot de passe admin immédiatement après la première connexion !

---

## 🔧 Configuration avancée

### Variables d'environnement disponibles

| Variable | Description | Défaut | Requis |
|----------|-------------|--------|--------|
| `DB_HOST` | Hôte PostgreSQL | `postgres` | ✅ |
| `DB_PORT` | Port PostgreSQL | `5432` | ✅ |
| `DB_NAME` | Nom de la base | `airboard` | ✅ |
| `DB_USER` | Utilisateur DB | `airboard` | ✅ |
| `DB_PASSWORD` | Mot de passe DB | - | ✅ |
| `JWT_SECRET` | Secret JWT | - | ✅ |
| `GIN_MODE` | Mode Go | `release` | ✅ |
| `PORT` | Port interne | `8080` | ❌ |

### Générer un JWT_SECRET sécurisé

```bash
# Sur votre machine locale
openssl rand -base64 32

# Ou avec Python
python3 -c "import secrets; print(secrets.token_urlsafe(32))"

# Ou avec Node.js
node -e "console.log(require('crypto').randomBytes(32).toString('base64'))"
```

---

## 🔍 Vérification du déploiement

### Health Check

Testez le health check :
```bash
curl https://votre-domaine.com/health
```

Réponse attendue :
```json
{
  "status": "ok",
  "message": "Airboard API is running"
}
```

### Logs

Dans Coolify, consultez les logs :
1. Allez dans votre service
2. Cliquez sur **Logs**
3. Vérifiez :
   ```
   ✅ Utilisateur admin créé: admin@airboard.com / admin123
   ✅ Groupe Développement créé avec applications de demo
   ✅ Groupe Production créé avec applications de demo
   🚀 Serveur Airboard démarré sur le port 8080
   ```

---

## 🔒 Sécurité en production

### Checklist de sécurité

- [ ] **JWT_SECRET** : Généré de manière aléatoire et long (min. 32 caractères)
- [ ] **DB_PASSWORD** : Mot de passe fort pour PostgreSQL
- [ ] **Mot de passe admin** : Changé après la première connexion
- [ ] **HTTPS** : Activé via Coolify (Let's Encrypt)
- [ ] **Sauvegardes** : Configurées dans Coolify pour PostgreSQL
- [ ] **Mises à jour** : Planifiées régulièrement

### Changer le mot de passe admin

1. Se connecter avec `admin@airboard.com` / `admin123`
2. Aller dans **Profil** (icône utilisateur)
3. Cliquer sur **Changer le mot de passe**
4. Entrer :
   - Ancien mot de passe : `admin123`
   - Nouveau mot de passe : `votre-nouveau-mot-de-passe-fort`
5. Sauvegarder

---

## 💾 Sauvegardes

### Sauvegarder la base de données

Coolify peut automatiquement sauvegarder PostgreSQL :

1. Allez dans votre service **PostgreSQL**
2. Onglet **Backups**
3. Configurez :
   - Fréquence : Quotidienne
   - Rétention : 7 jours (ou plus)
   - Destination : S3, local, etc.

### Sauvegarde manuelle

```bash
# Depuis le conteneur PostgreSQL
docker exec airboard-db pg_dump -U airboard airboard > backup.sql

# Restauration
docker exec -i airboard-db psql -U airboard airboard < backup.sql
```

---

## 🔄 Mise à jour de l'application

### Via Coolify (automatique)

1. Poussez vos modifications sur Git :
   ```bash
   git add .
   git commit -m "Update application"
   git push origin main
   ```

2. Dans Coolify :
   - Cliquez sur **Restart** ou **Rebuild**
   - Ou activez **Auto Deploy** pour déployer automatiquement à chaque push

### Réinitialisation de la base de données

Si vous devez repartir à zéro :

1. Se connecter en tant qu'admin
2. Aller dans **Administration**
3. Descendre jusqu'à **Zone de danger**
4. Cliquer sur **Réinitialiser la base de données**
5. Confirmer
6. Dans Coolify, **Restart** le service backend

---

## 🐛 Dépannage

### L'application ne démarre pas

**Vérifier les logs** :
```bash
# Dans Coolify, onglet Logs
# Rechercher des erreurs comme :
- "Erreur de connexion à la base de données"
- "JWT_SECRET not set"
```

**Solutions** :
- Vérifier que PostgreSQL est démarré
- Vérifier les variables d'environnement
- Vérifier le `DB_HOST` (nom du conteneur PostgreSQL)

### Erreur "cannot connect to database"

**Cause** : `DB_HOST` incorrect

**Solution** :
1. Trouvez le nom du conteneur PostgreSQL dans Coolify
2. Mettez à jour `DB_HOST` avec ce nom
3. Redéployez

### L'interface ne charge pas

**Vérifier** :
1. Le service est bien démarré
2. Le port 80 est exposé
3. Le domaine pointe vers le bon service

### Health check échoue

**Vérifier** :
```bash
# Tester manuellement
curl https://votre-domaine.com/health

# Si erreur, vérifier :
- Nginx est bien démarré
- Le backend répond sur le port 8080
- Supervisor gère bien les deux services
```

---

## 📊 Monitoring

### Metrics dans Coolify

Coolify fournit automatiquement :
- CPU usage
- Memory usage
- Disk usage
- Network traffic

### Logs applicatifs

Pour voir les logs détaillés :
```bash
# Dans Coolify
Service → Logs → Toggle "Follow Logs"
```

---

## 🚀 Optimisations

### Performance

1. **Cache Nginx** : Déjà configuré dans `docker-simple/nginx.conf`
2. **Compression gzip** : Déjà activée
3. **Build cache** : Coolify utilise automatiquement le cache Docker

### Scaling

Pour gérer plus de trafic :
1. Augmentez les ressources du conteneur dans Coolify
2. Activez le load balancing si disponible
3. Utilisez un CDN pour les assets statiques

---

## 📚 Ressources

- [Documentation Coolify](https://coolify.io/docs)
- [Documentation Airboard](README.md)
- [Guide Docker](DOCKER-DEPLOYMENT.md)
- [Structure du projet](PROJECT-STRUCTURE.md)

---

## ✅ Checklist de déploiement

- [ ] Repository Git créé et accessible
- [ ] Service PostgreSQL créé dans Coolify
- [ ] Variables d'environnement configurées
- [ ] JWT_SECRET généré de manière sécurisée
- [ ] Service Airboard déployé
- [ ] Health check réussi
- [ ] Connexion admin testée
- [ ] Mot de passe admin changé
- [ ] HTTPS activé
- [ ] Sauvegardes configurées
- [ ] Domaine personnalisé configuré (optionnel)

---

## 🎉 C'est tout !

Votre application Airboard est maintenant déployée sur Coolify !

**Temps de déploiement estimé** : 10-15 minutes

Pour toute question, consultez :
- Les logs dans Coolify
- Le guide de dépannage ci-dessus
- La documentation complète dans `DOCKER-DEPLOYMENT.md`

---

**Version** : 1.0.0-clean
**Dernière mise à jour** : 13 novembre 2025
