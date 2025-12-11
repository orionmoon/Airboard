# Guide Rapide - Système de Versioning

## 🎯 Résumé

Airboard dispose maintenant d'un système de versioning automatique qui :
- ✅ Affiche la version actuelle dans la sidebar (sous "Airboard")
- ✅ Vérifie automatiquement les nouvelles versions sur GitHub
- ✅ Affiche un **badge rouge pulsant** "Mise à jour" quand une nouvelle version est disponible
- ✅ Permet de voir les notes de version via un modal élégant

## 🚀 Démarrage Rapide

### 1. Configuration initiale (une seule fois)

Ajoutez dans votre fichier `.env` :

```bash
# Configuration GitHub pour vérification de mises à jour
GITHUB_REPO=votre-username/airboard    # Remplacez par votre repo GitHub
GITHUB_TOKEN=ghp_xxxxx                 # Optionnel mais recommandé
```

**Pour créer un GitHub Token** :
1. GitHub → Settings → Developer settings → Personal access tokens
2. Generate new token (classic)
3. Aucune permission spéciale requise (lecture publique)
4. Copiez le token dans `.env`

### 2. Publier une nouvelle version

**Méthode simple (recommandée)** :

```bash
# Syntaxe
./update-version.sh <version> "Description"

# Exemple
./update-version.sh 1.1.0 "Ajout du système de versioning"
```

Le script génère automatiquement :
- ✅ `version.json` avec la nouvelle version
- ✅ Changelog depuis les commits Git
- ✅ Mise à jour de `frontend/package.json`
- ✅ Instructions pour créer le tag Git et la release GitHub

**Puis suivez les instructions affichées** :

```bash
# 1. Vérifier les changements
git diff version.json

# 2. Committer
git add version.json backend/version.json frontend/package.json
git commit -m "chore: bump version to 1.1.0"

# 3. Créer le tag
git tag -a v1.1.0 -m "Release version 1.1.0"

# 4. Pousser sur GitHub
git push origin main
git push origin v1.1.0

# 5. Créer une Release sur GitHub
# Allez sur https://github.com/YOUR_USERNAME/airboard/releases/new
# Sélectionnez le tag v1.1.0 et publiez
```

### 3. Résultat

Une fois la release publiée sur GitHub :
- 🔴 Les utilisateurs verront automatiquement un **badge rouge "Mise à jour"** sous "Airboard"
- 💬 En cliquant, ils verront un modal avec :
  - Comparaison des versions (actuelle → nouvelle)
  - Date de publication
  - Notes de version complètes
  - Lien direct vers GitHub

## 📸 Aperçu

### Sidebar avec version
```
┌─────────────────────────┐
│  📱 Airboard            │
│  v1.0.0  🔴 Mise à jour │  ← Badge rouge pulsant
├─────────────────────────┤
│  📊 Dashboard           │
│  📰 News Hub            │
│  ...                    │
└─────────────────────────┘
```

### Modal de mise à jour
```
╔═══════════════════════════════════════╗
║  🔄 Nouvelle version disponible       ║
╠═══════════════════════════════════════╣
║                                       ║
║    v1.0.0  ──→  v1.1.0               ║
║                                       ║
║  Publiée le 11 décembre 2025         ║
║                                       ║
║  📝 Notes de version:                 ║
║  - feat: système de versioning        ║
║  - feat: badge de notification        ║
║  - fix: correction SSO                ║
║                                       ║
║  [Me le rappeler plus tard]           ║
║  [Fermer]  [🔗 Voir sur GitHub]      ║
╚═══════════════════════════════════════╝
```

## 🔧 Vérification automatique

Le système vérifie automatiquement les mises à jour :
- ⏰ Toutes les 4 heures
- 🚀 Au démarrage de l'application
- 💾 Les résultats sont mis en cache dans le navigateur

## 📋 Schéma de versioning

Airboard utilise le **Semantic Versioning** (SemVer) :

```
MAJOR.MINOR.PATCH
  │     │     │
  │     │     └── Corrections de bugs (1.0.1)
  │     └──────── Nouvelles fonctionnalités (1.1.0)
  └────────────── Changements majeurs (2.0.0)
```

**Exemples** :
- `1.0.1` → Correction d'un bug
- `1.1.0` → Nouvelle fonctionnalité
- `2.0.0` → Refonte majeure

## 🧪 Test en développement

Pour tester le système localement :

1. **Créer une release de test** sur votre repo GitHub (v1.0.1-test)

2. **Ouvrir la console du navigateur** et exécuter :
   ```javascript
   // Forcer une vérification
   const versionStore = useVersionStore()
   await versionStore.checkForUpdates()
   ```

3. **Le badge devrait apparaître** si la version GitHub > version locale

## 📚 Documentation complète

Pour plus de détails, consultez [VERSIONING.md](./VERSIONING.md)

## ❓ Questions fréquentes

### Le badge ne s'affiche pas ?
1. Vérifiez que `GITHUB_REPO` est configuré dans `.env`
2. Vérifiez qu'une release existe sur GitHub
3. Videz le cache du navigateur et rechargez

### Comment désactiver la vérification ?
Laissez `GITHUB_REPO` vide dans `.env`

### Faut-il un token GitHub ?
Non, mais c'est recommandé pour éviter les limites de taux (60 req/h → 5000 req/h avec token)

### Les utilisateurs doivent-ils faire quelque chose ?
Non ! La mise à jour se fait côté serveur. Les utilisateurs sont juste notifiés qu'une nouvelle version est disponible.

---

**Créé le** : 2025-12-11
**Version** : 1.0.0
