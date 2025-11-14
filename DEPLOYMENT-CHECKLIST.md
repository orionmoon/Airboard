# Checklist de Déploiement - OAuth Configuration

## Variables d'environnement à configurer en production

### Dans Coolify (ou votre plateforme de déploiement)

```bash
# ============================================
# OBLIGATOIRE : URL publique de l'application
# ============================================
PUBLIC_URL=https://tools.marocpme.gov.ma

# ============================================
# Base de données
# ============================================
DB_HOST=postgres
DB_PORT=5432
DB_USER=airboard
DB_PASSWORD=<CHANGEZ-MOI>
DB_NAME=airboard

# ============================================
# JWT (CHANGER EN PRODUCTION!)
# ============================================
JWT_SECRET=<GÉNÉREZ-UNE-CLÉ-SÉCURISÉE-DE-32-CARACTÈRES>

# ============================================
# SSO (si utilisé avec Authentik)
# ============================================
SSO_ENABLED=true
SSO_AUTO_PROVISION=true
SSO_DEFAULT_ROLE=user
SSO_DEFAULT_GROUP=Common
SSO_ADMIN_GROUPS=airboard-admins

# ============================================
# Application
# ============================================
GIN_MODE=release
```

## Étapes de déploiement

### 1. Configurer les variables d'environnement

- [ ] Définir `PUBLIC_URL=https://tools.marocpme.gov.ma` dans Coolify
- [ ] Vérifier que `JWT_SECRET` est unique et sécurisé
- [ ] Configurer les variables de base de données

### 1.1 Problème PostgreSQL "unhealthy"

Si vous rencontrez l'erreur `dependency failed to start: container postgres is unhealthy` :

**Solution 1 : Supprimer le volume PostgreSQL corrompu**
```bash
# Dans Coolify, aller dans l'onglet Storages/Volumes
# Supprimer le volume postgres_data
# Redéployer l'application
```

**Solution 2 : Augmenter le timeout du healthcheck**
Le docker-compose.yaml a été mis à jour avec un healthcheck plus robuste :
- `start_period: 30s` - Donne 30 secondes à PostgreSQL pour démarrer
- `retries: 10` - Augmente le nombre de tentatives
- `timeout: 10s` - Augmente le timeout de chaque check

Si le problème persiste, vérifiez les logs PostgreSQL dans Coolify pour identifier la cause exacte.

### 2. Déployer l'application

- [ ] Push le code sur Git
- [ ] Coolify déploie automatiquement
- [ ] Vérifier que le backend démarre sans erreur
- [ ] Vérifier les logs : `docker logs airboard-backend`

Vous devriez voir :
```
✅ Google OAuth provider créé (désactivé par défaut) - Redirect: https://tools.marocpme.gov.ma/auth/oauth/google/callback
✅ Microsoft OAuth provider créé (désactivé par défaut) - Redirect: https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback
🚀 Serveur Airboard démarré sur le port 8080
```

### 3. Configurer OAuth dans Google Console

**Redirect URI à ajouter** :
```
https://tools.marocpme.gov.ma/auth/oauth/google/callback
```

- [ ] Créer OAuth Client ID dans Google Cloud Console
- [ ] Ajouter le Redirect URI ci-dessus
- [ ] Copier Client ID et Client Secret

### 4. Configurer OAuth dans Azure Portal

**Redirect URI à ajouter** :
```
https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback
```

- [ ] Créer App Registration dans Azure Portal
- [ ] Ajouter le Redirect URI ci-dessus
- [ ] Créer un Client Secret
- [ ] Copier Application (client) ID et Client Secret

### 5. Activer OAuth dans Airboard

- [ ] Se connecter en admin : `https://tools.marocpme.gov.ma`
- [ ] Aller dans **Administration > OAuth**
- [ ] Configurer Google :
  - ✅ Enable this OAuth provider
  - Client ID : coller depuis Google Console
  - Client Secret : coller depuis Google Console
  - Vérifier Redirect URI : doit afficher `https://tools.marocpme.gov.ma/auth/oauth/google/callback`
  - Sauvegarder
- [ ] Configurer Microsoft :
  - ✅ Enable this OAuth provider
  - Client ID : coller depuis Azure Portal
  - Client Secret : coller depuis Azure Portal
  - Vérifier Redirect URI : doit afficher `https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback`
  - Sauvegarder

### 6. Tester

- [ ] Se déconnecter
- [ ] Aller sur `https://tools.marocpme.gov.ma/auth/login`
- [ ] Vérifier que les boutons "Continue with Google" et "Continue with Microsoft" apparaissent
- [ ] Tester la connexion avec Google
- [ ] Tester la connexion avec Microsoft
- [ ] Vérifier que les utilisateurs sont créés automatiquement

### 7. Configuration Nginx Proxy Manager

Ajouter dans l'onglet **Advanced** :

```nginx
# Augmenter les timeouts pour OAuth
proxy_connect_timeout 60s;
proxy_send_timeout 60s;
proxy_read_timeout 60s;

# Support OAuth callbacks (pas de blocage)
location /auth/oauth/ {
    proxy_pass http://airboard-frontend;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
}
```

## Changement d'URL

Si vous changez l'URL publique de l'application (ex: `tools.marocpme.gov.ma` → `nouveau-domaine.com`) :

### 1. Mettre à jour la variable d'environnement

```bash
PUBLIC_URL=https://nouveau-domaine.com
```

### 2. Redéployer l'application

Coolify redéploiera et mettra à jour automatiquement les Redirect URIs dans la base de données.

### 3. Mettre à jour Google Console

- Aller dans Google Cloud Console
- Mettre à jour les Authorized Redirect URIs :
  - Ajouter : `https://nouveau-domaine.com/auth/oauth/google/callback`
  - (Optionnel) Supprimer l'ancienne : `https://tools.marocpme.gov.ma/auth/oauth/google/callback`

### 4. Mettre à jour Azure Portal

- Aller dans Azure Portal > App Registration
- Aller dans Authentication
- Mettre à jour les Redirect URIs :
  - Ajouter : `https://nouveau-domaine.com/auth/oauth/microsoft/callback`
  - (Optionnel) Supprimer l'ancienne : `https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback`

### 5. Vérifier dans Airboard Admin

- Se connecter en admin
- Aller dans Administration > OAuth
- Vérifier que les Redirect URIs affichent le nouveau domaine

## Dépannage

### Les Redirect URIs ne se mettent pas à jour

Si après redéploiement les Redirect URIs sont toujours les anciens :

```sql
-- Connectez-vous à PostgreSQL
docker exec -it airboard-postgres psql -U airboard

-- Mettre à jour manuellement
UPDATE oauth_providers
SET redirect_uri = 'https://nouveau-domaine.com/auth/oauth/google/callback'
WHERE provider_name = 'google';

UPDATE oauth_providers
SET redirect_uri = 'https://nouveau-domaine.com/auth/oauth/microsoft/callback'
WHERE provider_name = 'microsoft';
```

### Vérifier la configuration actuelle

```bash
# Logs backend
docker logs airboard-backend | grep "OAuth provider"

# Devrait afficher :
# ✅ Google OAuth provider créé - Redirect: https://tools.marocpme.gov.ma/auth/oauth/google/callback
```

---

**Documentation complète** : Voir [OAUTH-SETUP.md](OAUTH-SETUP.md)
