# Système de Versioning Airboard

Ce document explique comment fonctionne le système de versioning d'Airboard et comment publier de nouvelles versions.

## 📋 Vue d'ensemble

Airboard dispose d'un système de versioning automatique qui :
- Affiche la version actuelle dans la sidebar sous le nom "Airboard"
- Vérifie automatiquement les nouvelles versions sur GitHub
- Affiche un badge rouge "Mise à jour" lorsqu'une nouvelle version est disponible
- Permet aux utilisateurs de voir les notes de version via un modal élégant

## 🎯 Fonctionnalités

### 1. Affichage de la version
- **Emplacement** : Sidebar, juste sous le nom "Airboard"
- **Format** : `v1.0.0`
- **Couleur** : Gris discret pour ne pas perturber l'interface

### 2. Badge de mise à jour
- **Apparence** : Badge rouge pulsant avec icône et texte "Mise à jour"
- **Déclenchement** : Apparaît automatiquement quand une nouvelle version GitHub est détectée
- **Action** : Clic pour ouvrir le modal avec les détails de la release

### 3. Modal de mise à jour
- **Contenu** :
  - Comparaison visuelle des versions (actuelle → nouvelle)
  - Date de publication
  - Notes de version complètes
  - Lien direct vers la release GitHub
- **Actions** :
  - "Me le rappeler plus tard" : Cache le badge jusqu'à la prochaine version
  - "Voir sur GitHub" : Ouvre la page de release

### 4. Vérification automatique
- **Fréquence** : Toutes les 4 heures
- **Première vérification** : Au chargement de l'application
- **Cache** : Les résultats sont mis en cache dans localStorage

## 🚀 Publication d'une nouvelle version

### Méthode automatique (recommandée)

Utilisez le script `update-version.sh` :

```bash
# Syntaxe
./update-version.sh <version> [message]

# Exemple
./update-version.sh 1.1.0 "Ajout du système de versioning et nouvelles fonctionnalités"
```

Le script va :
1. ✅ Mettre à jour `version.json` avec la nouvelle version
2. ✅ Copier le fichier dans `backend/version.json`
3. ✅ Mettre à jour `frontend/package.json`
4. ✅ Générer automatiquement le changelog depuis les commits Git
5. ✅ Afficher les étapes suivantes pour finaliser la release

### Étapes complètes de publication

1. **Préparer la version**
   ```bash
   ./update-version.sh 1.1.0 "Description de la version"
   ```

2. **Vérifier les changements**
   ```bash
   git diff version.json backend/version.json frontend/package.json
   ```

3. **Committer les changements**
   ```bash
   git add version.json backend/version.json frontend/package.json
   git commit -m "chore: bump version to 1.1.0"
   ```

4. **Créer un tag Git**
   ```bash
   git tag -a v1.1.0 -m "Release version 1.1.0"
   ```

5. **Pousser sur GitHub**
   ```bash
   git push origin main
   git push origin v1.1.0
   ```

6. **Créer une Release sur GitHub**
   - Allez sur `https://github.com/VOTRE_USERNAME/VOTRE_REPO/releases/new`
   - Sélectionnez le tag `v1.1.0`
   - Titre : `Version 1.1.0`
   - Description : Ajoutez les notes de version (markdown supporté)
   - Cliquez sur "Publish release"

### Méthode manuelle

Si vous préférez faire manuellement :

1. **Éditer `version.json`** et `backend/version.json`
   ```json
   {
     "version": "1.1.0",
     "buildDate": "2025-12-11",
     "gitCommit": "abc1234",
     "changelog": [
       {
         "version": "1.1.0",
         "date": "2025-12-11",
         "changes": [
           "feat: nouvelle fonctionnalité X",
           "fix: correction du bug Y",
           "chore: mise à jour des dépendances"
         ]
       }
     ]
   }
   ```

2. **Mettre à jour `frontend/package.json`**
   ```json
   {
     "version": "1.1.0",
     ...
   }
   ```

3. Suivre les étapes 2-6 ci-dessus

## ⚙️ Configuration

### Variables d'environnement

Pour activer la vérification automatique des mises à jour, configurez ces variables dans `.env` :

```bash
# Configuration GitHub pour la vérification de mises à jour
GITHUB_REPO=username/airboard          # Format: "owner/repo"
GITHUB_TOKEN=ghp_xxxxxxxxxxxxx         # Optionnel : Personal Access Token GitHub
```

**Important** :
- Si `GITHUB_REPO` n'est pas défini, la vérification de mise à jour sera désactivée
- Le `GITHUB_TOKEN` est optionnel mais recommandé pour éviter les limites de taux de l'API GitHub (60 requêtes/heure sans token, 5000 avec token)

### Créer un GitHub Personal Access Token

1. Allez sur GitHub : Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Cliquez sur "Generate new token"
3. Donnez un nom : "Airboard Update Checker"
4. **Permissions requises** : Aucune permission spéciale n'est nécessaire (lecture publique)
5. Cliquez sur "Generate token"
6. Copiez le token et ajoutez-le dans `.env` : `GITHUB_TOKEN=ghp_...`

## 📁 Fichiers impliqués

### Backend
- `backend/handlers/version.go` - Endpoints API pour version et vérification de mises à jour
- `backend/version.json` - Fichier de version (copié depuis la racine)
- `backend/main.go` - Routes API `/api/v1/version` et `/api/v1/version/check-updates`

### Frontend
- `frontend/src/stores/version.js` - Store Pinia pour la gestion de version
- `frontend/src/components/layout/Sidebar.vue` - Affichage version + badge + modal
- `frontend/package.json` - Version npm du frontend

### Racine
- `version.json` - Source de vérité pour la version
- `update-version.sh` - Script de mise à jour de version
- `.env.example` - Template avec config GitHub

## 🔄 Workflow de versioning

```
┌─────────────────────────────────────────────────────────────┐
│  1. Développement des fonctionnalités                       │
│     - Commits réguliers avec messages conventionnels        │
│     - feat: / fix: / chore: / docs: etc.                    │
└─────────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────────┐
│  2. Préparation de la release                               │
│     ./update-version.sh 1.1.0 "Description"                 │
│     - Génère version.json avec changelog automatique        │
└─────────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────────┐
│  3. Commit et tag                                            │
│     git commit -m "chore: bump version to 1.1.0"            │
│     git tag -a v1.1.0 -m "Release version 1.1.0"            │
└─────────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────────┐
│  4. Push sur GitHub                                          │
│     git push origin main && git push origin v1.1.0          │
└─────────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────────┐
│  5. Création de la GitHub Release                            │
│     - Via l'interface GitHub                                │
│     - Sélection du tag v1.1.0                               │
│     - Ajout des notes de version                            │
└─────────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────────┐
│  6. Notification automatique aux utilisateurs                │
│     - Badge rouge "Mise à jour" apparaît                    │
│     - Modal avec détails de la release                      │
│     - Lien direct vers GitHub                               │
└─────────────────────────────────────────────────────────────┘
```

## 🎨 Semantic Versioning

Airboard suit le schéma de versioning sémantique (SemVer) :

```
MAJOR.MINOR.PATCH
  │     │     │
  │     │     └─── Corrections de bugs (1.0.1, 1.0.2, etc.)
  │     └─────────── Nouvelles fonctionnalités rétrocompatibles (1.1.0, 1.2.0, etc.)
  └───────────────── Changements majeurs non rétrocompatibles (2.0.0, 3.0.0, etc.)
```

**Exemples** :
- `1.0.1` - Correction de bugs SSO
- `1.1.0` - Ajout du système de versioning
- `2.0.0` - Refonte complète de l'interface

## 🧪 Test du système

Pour tester le système de versioning en développement :

1. **Créer une release de test sur GitHub**
   - Tag : `v1.1.0-test`
   - Marquer comme "pre-release"

2. **Configurer `.env`**
   ```bash
   GITHUB_REPO=votre-username/airboard
   ```

3. **Forcer une vérification de mise à jour**
   - Ouvrez la console du navigateur
   - Exécutez : `await versionStore.checkForUpdates()`
   - Le badge devrait apparaître si une nouvelle version existe

4. **Vérifier le modal**
   - Cliquez sur le badge "Mise à jour"
   - Vérifiez que le modal s'affiche correctement
   - Testez le lien "Voir sur GitHub"

## 📝 Format du Changelog

Le changelog dans `version.json` utilise le format suivant :

```json
{
  "version": "1.1.0",
  "buildDate": "2025-12-11",
  "gitCommit": "f28a436",
  "changelog": [
    {
      "version": "1.1.0",
      "date": "2025-12-11",
      "changes": [
        "feat: ajout du système de versioning automatique",
        "feat: badge de notification de mise à jour",
        "fix: correction du bug d'authentification SSO",
        "chore: mise à jour des dépendances"
      ]
    }
  ]
}
```

**Conventions de messages** :
- `feat:` - Nouvelle fonctionnalité
- `fix:` - Correction de bug
- `chore:` - Tâches de maintenance
- `docs:` - Documentation
- `refactor:` - Refactorisation de code
- `perf:` - Amélioration de performance
- `test:` - Ajout de tests

## 🐛 Dépannage

### Le badge ne s'affiche pas

1. Vérifiez que `GITHUB_REPO` est configuré dans `.env`
2. Vérifiez qu'une release existe sur GitHub avec un tag supérieur à la version actuelle
3. Ouvrez la console et vérifiez les erreurs API
4. Forcez une vérification : `await versionStore.checkForUpdates()`

### Erreur "Failed to check for updates"

1. Vérifiez la connexion Internet
2. Vérifiez que le repository GitHub est public (ou ajoutez `GITHUB_TOKEN`)
3. Vérifiez le format de `GITHUB_REPO` (doit être `owner/repo`)

### La version ne se met pas à jour

1. Vérifiez que `version.json` existe dans `backend/`
2. Redémarrez le backend après modification
3. Videz le cache du navigateur (localStorage)

## 📚 Ressources

- [Semantic Versioning](https://semver.org/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Releases](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)
- [GitHub API - Releases](https://docs.github.com/en/rest/releases/releases)
