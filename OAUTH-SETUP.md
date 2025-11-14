# Configuration OAuth - Google & Microsoft

Ce guide vous explique comment activer et configurer l'authentification OAuth pour Google et Microsoft dans Airboard.

## Aperçu

Airboard supporte maintenant 3 méthodes d'authentification :
1. **Login/Password classique** : Utilisateurs locaux avec mot de passe
2. **SSO via Authentik** : Authentification via Nginx Proxy Manager + Authentik
3. **OAuth Direct** : Google et Microsoft OAuth (nouvelle fonctionnalité)

## Avantages de OAuth Direct

- ✅ Les utilisateurs peuvent se connecter avec leur compte Google ou Microsoft
- ✅ Pas besoin d'Authentik pour ces providers
- ✅ Configuration simple via l'interface admin
- ✅ Les comptes sont créés automatiquement au premier login
- ✅ Peut coexister avec Authentik SSO

## Prérequis

### Configuration de l'URL publique

**IMPORTANT** : Avant de créer les applications OAuth, configurez la variable d'environnement `PUBLIC_URL` :

**Développement local** :
```bash
# Dans backend/.env
PUBLIC_URL=http://localhost:5173
```

**Production (Coolify ou autre)** :
```bash
# Dans les variables d'environnement Coolify
PUBLIC_URL=https://tools.marocpme.gov.ma
```

Cette URL sera utilisée pour générer automatiquement les **Redirect URIs** OAuth.

### Créer des applications OAuth

Vous devez créer des applications OAuth chez :
- **Google** : [Google Cloud Console](https://console.cloud.google.com/)
- **Microsoft** : [Azure Portal](https://portal.azure.com/)

---

## Configuration Google OAuth

### Étape 1 : Créer un projet Google Cloud

1. Allez sur [Google Cloud Console](https://console.cloud.google.com/)
2. Créez un nouveau projet ou sélectionnez un projet existant
3. Activez l'**Google+ API** (dans APIs & Services)

### Étape 2 : Créer les credentials OAuth

1. Allez dans **APIs & Services** > **Credentials**
2. Cliquez sur **Create Credentials** > **OAuth client ID**
3. Si demandé, configurez d'abord l'écran de consentement OAuth :
   - Type : **External** (ou Internal si G Suite)
   - Nom de l'application : `Airboard`
   - Email de support : votre email
   - Scopes : `email`, `profile`, `openid`
   - Domaines autorisés : `tools.marocpme.gov.ma`

4. Créez le **OAuth Client ID** :
   - Type d'application : **Web application**
   - Nom : `Airboard Production` (ou `Airboard Dev` pour le dev)
   - **Authorized redirect URIs** :
     - **Production** : `https://tools.marocpme.gov.ma/auth/oauth/google/callback`
     - **Dev local** : `http://localhost:5173/auth/oauth/google/callback`

     > 💡 **Astuce** : Vous pouvez ajouter les deux URIs dans Google Console pour tester en local et en prod avec le même Client ID

5. **Copiez** :
   - Client ID (format : `123456789-xxxxx.apps.googleusercontent.com`)
   - Client Secret

### Étape 3 : Configurer dans Airboard

1. Connectez-vous en tant qu'**admin**
2. Allez dans **Administration** > **OAuth**
3. Cliquez sur **Edit** pour Google
4. Remplissez :
   - ✅ **Enable this OAuth provider**
   - **Client ID** : collez le Client ID
   - **Client Secret** : collez le Client Secret
   - **Redirect URI** : (déjà rempli, cliquez sur copier pour vérifier)
5. Cliquez sur **Save**

---

## Configuration Microsoft OAuth

### Étape 1 : Créer une App Registration dans Azure

1. Allez sur [Azure Portal](https://portal.azure.com/)
2. Recherchez **Azure Active Directory** (ou **Microsoft Entra ID**)
3. Allez dans **App registrations** > **New registration**

### Étape 2 : Configurer l'application

1. **Nom** : `Airboard`
2. **Supported account types** :
   - Choisissez selon votre besoin :
     - **Single tenant** : Seulement votre organisation
     - **Multitenant** : Toute organisation Microsoft
     - **Personal + Work accounts** : Comptes personnels + professionnels
3. **Redirect URI** :
   - Type : **Web**
   - **Production** : `https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback`
   - **Dev local** : `http://localhost:5173/auth/oauth/microsoft/callback`

   > 💡 Vous pouvez ajouter plusieurs Redirect URIs dans Azure après la création
4. Cliquez sur **Register**

### Étape 3 : Créer un Client Secret

1. Dans votre app, allez dans **Certificates & secrets**
2. Cliquez sur **New client secret**
3. Description : `Airboard Production`
4. Expiration : Choisissez selon votre politique (24 mois recommandé)
5. **Copiez immédiatement la valeur du secret** (vous ne pourrez plus le voir après)

### Étape 4 : Noter l'Application ID

1. Retournez sur **Overview**
2. **Copiez** :
   - **Application (client) ID** (format : `a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6`)
   - **Directory (tenant) ID** (si single tenant)

### Étape 5 : Configurer les permissions API

1. Allez dans **API permissions**
2. Vérifiez que vous avez au minimum :
   - `User.Read` (Microsoft Graph)
   - `openid`
   - `email`
   - `profile`
3. Cliquez sur **Grant admin consent** si nécessaire

### Étape 6 : Configurer dans Airboard

1. Connectez-vous en tant qu'**admin**
2. Allez dans **Administration** > **OAuth**
3. Cliquez sur **Edit** pour Microsoft
4. Remplissez :
   - ✅ **Enable this OAuth provider**
   - **Client ID** : collez l'Application (client) ID
   - **Client Secret** : collez le secret créé
   - **Redirect URI** : (déjà rempli)
5. **Advanced Settings** :

   **IMPORTANT - Configuration Single Tenant vs Multi-tenant** :

   **Si vous avez choisi "Single tenant"** (recommandé pour une organisation) :
   - Remplacez `/common/` par `/[YOUR_TENANT_ID]/` dans les URLs
   - Exemple avec le Tenant ID `abc123-...` :
     ```
     https://login.microsoftonline.com/abc123-def4-5678-90ab-cdef12345678/oauth2/v2.0/authorize
     https://login.microsoftonline.com/abc123-def4-5678-90ab-cdef12345678/oauth2/v2.0/token
     ```
   - Pour trouver votre Tenant ID : Azure Portal > Azure AD > Overview > Tenant ID

   **Si vous avez choisi "Multitenant" ou "Personal + Work accounts"** :
   - Gardez `/common/` dans les URLs (pas de modification nécessaire)

6. Cliquez sur **Save**

---

## Vérification

### Tester l'authentification

1. Déconnectez-vous d'Airboard
2. Allez sur la page de login : `https://tools.marocpme.gov.ma/auth/login`
3. Vous devriez voir :
   - Un bouton **"Continue with Google"** (si activé)
   - Un bouton **"Continue with Microsoft"** (si activé)
   - Un séparateur "Or continue with credentials"
   - Le formulaire classique login/password

4. Cliquez sur un bouton OAuth :
   - Vous êtes redirigé vers Google/Microsoft
   - Autorisez l'accès
   - Vous revenez sur Airboard connecté
   - Un compte utilisateur est créé automatiquement

### Vérifier les utilisateurs

1. Connectez-vous en tant qu'admin
2. Allez dans **Administration** > **Users**
3. Les utilisateurs OAuth ont :
   - **SSO Provider** : `google` ou `microsoft`
   - **SSO ID** : leur ID externe
   - Pas de mot de passe (colonne vide)

---

## Coexistence avec Authentik

Airboard peut gérer les 3 méthodes simultanément :

| Méthode | URL | Usage |
|---------|-----|-------|
| **Login classique** | `/auth/login` | Utilisateurs locaux avec mot de passe |
| **Google OAuth** | `/auth/login` → Bouton Google | Comptes Google (@gmail.com, Workspace) |
| **Microsoft OAuth** | `/auth/login` → Bouton Microsoft | Comptes Microsoft (365, Outlook, Azure AD) |
| **Authentik SSO** | Headers via Nginx | LDAP, autres SSO via Authentik |

**Configuration recommandée** :
- Gardez Authentik pour LDAP et autres sources
- Activez Google/Microsoft pour les utilisateurs externes
- Gardez login/password pour les admins (fallback)

---

## Dépannage

### Les boutons OAuth n'apparaissent pas

1. Vérifiez que le provider est **activé** dans `/admin/oauth`
2. Vérifiez les logs backend : `docker logs airboard-backend`
3. Vérifiez que le **Client ID** est bien rempli

### Erreur "redirect_uri_mismatch"

1. Vérifiez que le **Redirect URI** dans Google/Azure est **exactement** :
   ```
   https://tools.marocpme.gov.ma/auth/oauth/google/callback
   ```
   ou
   ```
   https://tools.marocpme.gov.ma/auth/oauth/microsoft/callback
   ```
2. Pas d'espace, pas de `/` à la fin
3. HTTPS obligatoire (pas HTTP)

### Erreur Microsoft "AADSTS700016: Application not found in directory"

Cette erreur signifie que vous utilisez un **Single Tenant** mais que les URLs utilisent `/common/` :

**Symptôme** : Message d'erreur Azure indiquant que l'application n'a pas été trouvée dans le répertoire.

**Solution** :
1. Allez dans **Administration > OAuth**
2. Éditez le provider Microsoft
3. Dans **Advanced Settings**, remplacez `/common/` par votre **Tenant ID** :
   ```
   https://login.microsoftonline.com/[VOTRE_TENANT_ID]/oauth2/v2.0/authorize
   https://login.microsoftonline.com/[VOTRE_TENANT_ID]/oauth2/v2.0/token
   ```
4. Pour trouver votre Tenant ID :
   - Azure Portal > Azure Active Directory > Overview
   - Copiez le **Tenant ID** (format UUID : `abc12345-...`)

**Alternative** : Changez votre App Registration en **Multitenant** dans Azure Portal si vous voulez accepter tous les comptes Microsoft.

### Erreur après le callback

1. Vérifiez les logs : `docker logs airboard-backend`
2. Vérifiez que le **Client Secret** est correct
3. Vérifiez que les URLs (auth_url, token_url, user_info_url) sont correctes dans Advanced Settings

### L'utilisateur n'a pas les bons droits

Par défaut, les utilisateurs OAuth sont créés avec :
- **Role** : `user`
- **Group** : `common`

Pour donner le rôle admin :
1. Allez dans **Administration** > **Users**
2. Trouvez l'utilisateur OAuth
3. Éditez et changez le rôle en `admin`

---

## Sécurité

### Best Practices

1. ✅ **Utilisez HTTPS** (obligatoire pour OAuth)
2. ✅ **Gardez les Client Secrets secrets** (jamais dans le code)
3. ✅ **Limitez les Redirect URIs** (une seule URL de production)
4. ✅ **Renouvelez les secrets périodiquement** (tous les 12-24 mois)
5. ✅ **Activez MFA** sur vos comptes Google/Azure
6. ✅ **Surveillez les logs** d'authentification

### Secrets Storage

Les secrets OAuth sont stockés dans PostgreSQL :
- Champs `client_secret` dans la table `oauth_providers`
- Non exposés dans les API JSON (champ `json:"-"`)
- Accessibles uniquement aux admins via l'interface

---

## FAQ

### Puis-je désactiver un provider ?

Oui, décochez **"Enable this OAuth provider"** et sauvegardez. Le bouton disparaîtra de la page de login.

### Les utilisateurs existants peuvent-ils lier leur compte OAuth ?

Actuellement non. Si un utilisateur a le même email, deux comptes séparés seront créés. Fonctionnalité à venir.

### Puis-je forcer OAuth uniquement ?

Oui, vous pouvez masquer le formulaire classique en modifiant [Login.vue](frontend/src/views/auth/Login.vue), mais gardez un accès admin de secours.

### Puis-je ajouter d'autres providers (GitHub, GitLab) ?

Oui, mais il faudra modifier le backend Go pour ajouter le parsing des informations utilisateur spécifiques.

---

## Support

Pour toute question ou problème :
1. Vérifiez les logs : `docker logs airboard-backend`
2. Consultez la documentation officielle :
   - [Google OAuth](https://developers.google.com/identity/protocols/oauth2)
   - [Microsoft OAuth](https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-auth-code-flow)
3. Ouvrez une issue sur le repository GitHub

---

**Fait avec ❤️ pour tools.marocpme.gov.ma**
