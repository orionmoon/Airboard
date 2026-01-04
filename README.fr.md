# Airboard - Portail d'Applications Moderne

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.0+-4FC08D.svg)](https://vuejs.org)
[![SSO](https://img.shields.io/badge/SSO-Authentik-orange.svg)](https://goauthentik.io)

*Un portail d'applications moderne et sécurisé avec gestion des utilisateurs, intégration SSO, et support Docker*

[Démarrage Rapide](#démarrage-rapide) • [À Propos](#à-propos) • [Fonctionnalités](#fonctionnalités) • [Installation](#installation) • [Configuration](#configuration) • [Documentation](#documentation)

---

**📖 Disponible en plusieurs langues:**

[🇬🇧 English](README.md) | [🇫🇷 Français](README.fr.md)

</div>

---

## Table des Matières

- [À Propos](#à-propos)
- [Fonctionnalités Principales](#fonctionnalités-principales)
- [Captures d'Écran](#captures-décran)
- [Démarrage Rapide](#démarrage-rapide)
- [Installation Détaillée](#installation-détaillée)
- [Configuration](#configuration)
- [Architecture](#architecture)
- [Sécurité](#sécurité)
- [API](#api)
- [Dépannage](#dépannage)
- [Contribution](#contribution)
- [License](#license)

---

## À Propos

**Airboard** est une plateforme de portail interne complète conçue pour centraliser l'accès aux applications, faciliter les communications internes, et fournir des analyses pertinentes - le tout en un seul endroit.

### Objectif Principal

Dans un environnement professionnel moderne, les organisations ont besoin de bien plus qu'un simple annuaire de liens. **Airboard** répond à ce besoin en fournissant une solution d'espace de travail numérique complète :

- **Portail d'Applications Unifié** - Tableau de bord centralisé avec accès personnalisé à tous vos outils
- **Hub d'Actualités** - Système de communication interne avec articles, tutoriels, annonces et FAQ
- **Tableau de Bord Analytique** - Suivi de l'utilisation des applications et de l'engagement des utilisateurs
- **Système d'Annonces** - Diffusion instantanée de messages importants aux équipes
- **Favoris Utilisateur** - Expérience personnalisée avec applications et contenus favoris
- **Gestion Déléguée** - Rôle Group Admin pour la gestion déléguée des groupes privés
- **Gestion Basée sur les Rôles** - Permissions granulaires pour Admins, Group Admins, Éditeurs et Utilisateurs
- **Support Multilingue** - Disponible en français, anglais, espagnol et arabe
- **Intégration SSO** - Authentification transparente avec Authentik, Microsoft 365, LDAP, etc.
- **Prêt pour Docker** - Déploiement en quelques minutes avec conteneurisation complète

### Cas d'Usage

- **Communication d'Entreprise** : Publication d'actualités, mises à jour de politiques et articles de base de connaissances
- **Services IT** : Portail d'applications + documentation technique + analyses d'utilisation
- **RH & Communications Internes** : Partage d'annonces, manuels employés, tutoriels et FAQ
- **Tableaux de Bord Départementaux** : Vues personnalisées par équipe (RH, Finance, Ventes)
- **Collaboration Projet** : Portails spécifiques aux équipes avec applications, annonces et documentation
- **Analyses & Insights** : Suivi des applications les plus utilisées et comportement utilisateur

Airboard transforme la façon dont les organisations gèrent leurs ressources numériques, communications internes et engagement utilisateur.

---

## Fonctionnalités Principales

### Authentification & Sécurité

- **Authentification Duale**
  - Login classique avec username/password
  - SSO via Authentik (Microsoft 365, Google, LDAP, etc.)
  - Auto-provisioning des utilisateurs SSO
  - Mapping automatique des rôles via groupes Authentik

- **Sécurité Renforcée (OWASP 2025)**
  - Authentification JWT avec refresh tokens
  - Hashage bcrypt configurable (coût 10-31)
  - Protection CSRF et XSS
  - Headers de sécurité (CSP, CORS)
  - Validation stricte des entrées
  - Protection contre les injections SQL
  - Limite de taux sur les endpoints sensibles

- **OAuth2 Intégré**
  - Google OAuth2
  - Microsoft OAuth2
  - Configuration via interface admin
  - Gestion des tokens sécurisée

### Gestion des Utilisateurs & Permissions

- **Système de Rôles**
  - **Admin** : Accès complet au système
  - **Group Admin** : Gestion déléguée de groupes privés et applications
  - **Editor** : Création et gestion d'articles
  - **User** : Accès aux applications assignées

- **Gestion des Utilisateurs**
  - CRUD complet des utilisateurs
  - Assignation de groupes et rôles
  - Gestion des favoris utilisateur
  - Profils utilisateur personnalisables

- **Gestion des Groupes**
  - Création et organisation de groupes
  - Groupes publics et privés
  - Assignation de Group Admins
  - Contrôle d'accès granulaire

### Système Group Admin (Délégation)

- **Gestion Déléguée**
  - Group Admins peuvent gérer leurs groupes assignés
  - Création de groupes d'applications privés
  - Gestion des applications dans leurs groupes
  - Création et modification d'articles
  - Pas d'accès aux utilisateurs ou paramètres globaux

- **Groupes d'Applications Privés**
  - Visibles uniquement aux groupes assignés
  - Ownership par groupe (OwnerGroupID)
  - Permissions vérifiées à chaque opération
  - Séparation complète des groupes publics

### Portail d'Applications

- **Organisation des Applications**
  - Groupes d'applications publics et privés
  - Icônes personnalisées (Iconify)
  - Couleurs thématiques par groupe
  - URLs et descriptions configurables
  - Tri et filtrage avancés

- **Favoris Utilisateur**
  - Épinglage d'applications favorites
  - Accès rapide depuis le tableau de bord
  - Synchronisation en temps réel

- **Analyses d'Utilisation**
  - Suivi des clics par application
  - Utilisateurs uniques
  - Graphiques d'activité sur 30 jours
  - Top applications

### Hub d'Actualités

- **Éditeur Riche (Tiptap)**
  - WYSIWYG complet
  - Code blocks avec coloration syntaxique
  - Tableaux, listes, liens
  - Alignement de texte
  - Images et médias

- **Organisation du Contenu**
  - Catégories (General, Tutorials, Announcements, FAQ)
  - Tags multiples et dynamiques
  - Slugs SEO-friendly
  - Épinglage d'articles importants

- **Engagement Utilisateur**
  - Système de réactions (👍, ❤️, 🎉)
  - Suivi des vues par utilisateur
  - Compteurs de vues uniques
  - Analytics détaillées

- **Modes d'Affichage**
  - Vue grille (cards)
  - Vue liste (compact)
  - Vue table (très compact)
  - Filtres avancés (catégorie, type, tags)
  - Recherche full-text

### Système de Sondages

- **Sondages Interactifs**
  - Sondages à choix unique ou multiples
  - Vote anonyme ou nominatif
  - Sondages planifiés (date début/fin)
  - Ciblage par groupes ou tous les utilisateurs
  - Lien vers articles ou annonces

- **Gestion des Sondages**
  - Création et édition (Admin & Group Admin)
  - Visualisation des résultats en temps réel
  - Contrôle de visibilité (toujours, après vote, après fermeture)
  - Suivi des votes et analytics

- **Participation Utilisateur**
  - Interface de vote facile
  - Visualisation des résultats selon permissions
  - Historique des votes
  - Détails des votants (si non anonyme)

- **Intégration**
  - Widget sondages sur le dashboard
  - Intégration dans les articles
  - Page dédiée aux sondages
  - Export des résultats en CSV

### Notifications In-App

- **Système de Notifications**
  - Notifications temps réel dans l'application
  - Cloche de notification avec compteur
  - Types multiples (système, actualités, annonces, événements)
  - Niveaux de priorité (normal, important, urgent)

- **Types de Notifications**
  - Notifications système (connexion, changements de rôle, accès)
  - Notifications de contenu (nouveaux articles, sondages, annonces)
  - Notifications d'engagement (réactions, commentaires, mentions)
  - Notifications personnalisées par admin

- **Expérience Utilisateur**
  - Suivi statut lu/non lu
  - Marquer comme lu/non lu
  - Tout marquer comme lu
  - Supprimer notifications
  - Redirection vers URL d'action au clic

- **Gestion**
  - Centre de notifications dédié
  - Filtres par type et statut
  - Nettoyage automatique anciennes notifications
  - Push notifications (prévu)

### Système d'Annonces

- **Bannières Dynamiques**
  - Affichage sur le tableau de bord
  - Types : Info, Warning, Error, Success
  - Période de validité (date début/fin)
  - Priorité d'affichage

- **Gestion Centralisée**
  - Création et modification via admin
  - Activation/désactivation
  - Prévisualisation

### Analytics & Reporting

- **Dashboard Analytique**
  - Statistiques globales (utilisateurs, applications, articles)
  - Graphiques d'activité quotidienne
  - Top applications et utilisateurs
  - Analytics des actualités

- **Suivi d'Utilisation**
  - Clics par application
  - Vues d'articles
  - Réactions et engagement
  - Exportation de données

### Internationalisation

- **Support Multilingue**
  - Français (fr)
  - Anglais (en)
  - Espagnol (es)
  - Arabe (ar)
  - Sélecteur de langue intégré
  - Traductions complètes

### Gestion des Versions

- **Auto-Update System**
  - Vérification automatique des nouvelles versions
  - Intégration GitHub Releases
  - Notifications de mise à jour
  - Changelog accessible
  - Configuration via GITHUB_REPO

### Interface Utilisateur

- **Design Moderne**
  - TailwindCSS responsive
  - Mode sombre
  - Animations fluides
  - Navigation intuitive

- **Composants Réutilisables**
  - Modals standardisés
  - Cartes d'applications
  - Formulaires validés
  - Loaders et notifications

---

## Captures d'Écran

### Dashboard & Applications

<div align="center">

**Tableau de Bord Principal**
![Dashboard](img/Screenshot_20251117_090240.png)

</div>

### Hub d'Actualités

<div align="center">

**Vue Grille**
![News Grid](img/Screenshot_20251117_093545.png)

**Vue Compacte**
![News Compact](img/Screenshot_20251117_093555.png)

**Vue Liste**
![News List](img/Screenshot_20251117_093524.png)

**Gestion des Actualités**
![News Management](img/Screenshot_20251117_090448.png)

**Éditeur de Texte Riche**
![Rich Editor](img/Screenshot_20251117_090503.png)

</div>

### Administration

<div align="center">

**Vue d'Ensemble**
![Overview](img/Screenshot_20251117_090338.png)

**Gestion des Groupes d'Applications**
![App Groups](img/Screenshot_20251117_090346.png)

**Gestion des Applications**
![Applications](img/Screenshot_20251117_090356.png)

**Gestion des Utilisateurs**
![Users](img/Screenshot_20251117_090403.png)

**Gestion des Groupes**
![Groups](img/Screenshot_20251117_090409.png)

**Paramètres**
![Settings](img/Screenshot_20251117_090415.png)

**OAuth**
![OAuth](img/Screenshot_20251117_090420.png)

**Analytics**
![Analytics](img/Screenshot_20251117_090429.png)

**Gestion des Annonces**
![Announcements](img/Screenshot_20251117_090439.png)

**Éditeur d'Annonces**
![Announcement Editor](img/Screenshot_20251117_090514.png)

</div>

---

## Démarrage Rapide

### Option 1: Docker Compose (Recommandé)

```bash
# Cloner le dépôt
git clone https://github.com/votre-username/airboard.git
cd airboard

# Démarrer Docker (si nécessaire)
sudo systemctl start docker

# Configurer les variables (optionnel)
cp .env.example .env
# Éditer .env si besoin

# Construire et démarrer
docker-compose build
docker-compose up -d

# Vérifier le statut
docker-compose ps
```

**Accès :** http://localhost

**Comptes par défaut :**
- **Admin** : `admin` / `admin123`
- **User** : `user` / `user123`

**IMPORTANT :** Changez les mots de passe par défaut dès la première connexion.

### Option 2: Déploiement Coolify

1. Créer un nouveau projet dans Coolify
2. Déployer depuis Git : `https://github.com/votre-username/airboard.git`
3. Build Pack : `docker-compose`
4. Fichier : `docker-compose.prod.yaml`
5. Configurer les variables d'environnement (voir [Configuration](#configuration))
6. Déployer

### Option 3: Mode Développement

```bash
# Backend
cd backend
go mod download
go run main.go

# Frontend (autre terminal)
cd frontend
npm install
npm run dev
```

**Accès :** http://localhost:3000

---

## Installation Détaillée

### Prérequis

**Pour Docker (Production):**
- Docker 20.10+
- Docker Compose 2.0+
- 2 Go RAM minimum
- 10 Go espace disque

**Pour Développement:**
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- 4 Go RAM minimum

### Installation Docker (Production)

#### 1. Installation de Docker

**Ubuntu/Debian:**
```bash
# Installer Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Ajouter l'utilisateur au groupe docker
sudo usermod -aG docker $USER

# Démarrer Docker
sudo systemctl start docker
sudo systemctl enable docker

# Vérifier
docker --version
docker-compose --version
```

**CentOS/RHEL:**
```bash
sudo yum install -y docker docker-compose
sudo systemctl start docker
sudo systemctl enable docker
```

#### 2. Cloner le Projet

```bash
git clone https://github.com/votre-username/airboard.git
cd airboard
```

#### 3. Configuration des Variables

```bash
# Copier le template
cp .env.example .env

# Éditer avec vos valeurs
nano .env
```

**Variables critiques à modifier en production :**

```env
# Base de données (OBLIGATOIRE)
DB_PASSWORD=votre_mot_de_passe_securise

# JWT (OBLIGATOIRE)
JWT_SECRET=votre-secret-jwt-tres-long-minimum-32-caracteres

# Application
PUBLIC_URL=https://votre-domaine.com

# SSO (si utilisé)
SSO_ENABLED=true
SSO_ADMIN_GROUPS=airboard-admins,votre-groupe-admin

# GitHub Updates (optionnel)
GITHUB_REPO=votre-username/airboard
GITHUB_TOKEN=ghp_votre_token_github
```

#### 4. Déploiement

```bash
# Construire les images
docker-compose build

# Démarrer les services
docker-compose up -d

# Vérifier les logs
docker-compose logs -f

# Vérifier le statut
docker-compose ps
```

#### 5. Vérification

```bash
# Tester l'accès
curl http://localhost/health

# Vérifier la base de données
docker-compose exec postgres psql -U airboard -d airboard -c "\dt"

# Logs backend
docker-compose logs backend

# Logs frontend
docker-compose logs frontend
```

### Installation Coolify (Cloud)

#### 1. Préparer le Dépôt

```bash
# Forker ou cloner
git clone https://github.com/votre-username/airboard.git
cd airboard

# Pousser vers votre repo GitHub
git remote set-url origin https://github.com/VOTRE_USERNAME/airboard.git
git push
```

#### 2. Configuration Coolify

1. **Créer un Projet**
   - Nom : Airboard
   - Environnement : Production

2. **Ajouter un Service**
   - Type : Git Repository
   - Repository : votre repo GitHub
   - Branch : main
   - Build Pack : docker-compose

3. **Sélectionner docker-compose.prod.yaml**
   - Dans Settings > Build
   - Docker Compose File : `docker-compose.prod.yaml`

4. **Configurer les Variables d'Environnement**

Variables minimum :
```env
DB_PASSWORD=mot_de_passe_securise_123
JWT_SECRET=secret-jwt-tres-long-au-moins-32-caracteres-random
PUBLIC_URL=https://airboard.votre-domaine.com
SSO_ENABLED=false
```

Variables complètes (recommandées) :
```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=airboard
DB_PASSWORD=mot_de_passe_tres_securise_123
DB_NAME=airboard

# JWT
JWT_SECRET=votre-secret-jwt-super-long-et-aleatoire-minimum-32-chars
JWT_TOKEN_EXPIRATION_HOURS=24
JWT_REFRESH_EXPIRATION_DAYS=7

# Security
BCRYPT_COST=12

# Application
GIN_MODE=release
PUBLIC_URL=https://airboard.votre-domaine.com
SIGNUP_ENABLED=true

# SSO (si Authentik utilisé)
SSO_ENABLED=true
SSO_AUTO_PROVISION=true
SSO_DEFAULT_ROLE=user
SSO_DEFAULT_GROUP=Common
SSO_ADMIN_GROUPS=airboard-admins,it-admins

# GitHub Updates (optionnel)
GITHUB_REPO=votre-username/airboard
GITHUB_TOKEN=ghp_votre_token_personnel
```

5. **Déployer**
   - Cliquer sur "Deploy"
   - Attendre la fin du build (5-10 minutes)
   - Vérifier les logs

6. **Configuration DNS**
   - Ajouter un enregistrement DNS A vers l'IP Coolify
   - Ou CNAME vers le domaine Coolify

#### 3. Vérification Post-Déploiement

```bash
# Tester la santé
curl https://airboard.votre-domaine.com/health

# Devrait retourner : {"status":"healthy"}
```

### Installation avec SSO (Authentik + Nginx Proxy Manager)

#### Architecture

```
Internet → Nginx Proxy Manager → Authentik → Coolify → Airboard
```

#### 1. Configuration Authentik

**Créer une Application :**
1. Applications > Create
2. Name : Airboard
3. Slug : airboard
4. Provider : Create new Proxy Provider

**Configurer le Provider :**
- Type : Proxy Provider
- Authorization flow : default-provider-authorization-implicit-consent
- External host : `https://airboard.votre-domaine.com`
- Internal host : `http://coolify-airboard:80`
- Mode : Forward auth (single application)

**Créer des Groupes :**
1. Directory > Groups > Create
2. Créer `airboard-admins` pour les administrateurs
3. Créer `airboard-users` pour les utilisateurs standards
4. Assigner des utilisateurs aux groupes

#### 2. Configuration Nginx Proxy Manager

**Créer un Proxy Host :**
1. Hosts > Proxy Hosts > Add Proxy Host
2. Domain Names : `airboard.votre-domaine.com`
3. Forward Hostname/IP : domaine Coolify ou IP
4. Forward Port : 80
5. Cache Assets : ON
6. Block Common Exploits : ON
7. Websockets Support : ON

**Configurer l'Authentification Authentik :**
1. Access List > Create Access List
2. Authorization > Authentik
3. Provider : Sélectionner le provider Airboard

**Configuration Avancée (Custom Nginx Config) :**
```nginx
location / {
    proxy_pass $forward_scheme://$server:$port;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-Host $host;
    proxy_set_header X-Forwarded-Port $server_port;

    # Authentik SSO Headers (CRITICAL)
    proxy_set_header X-authentik-email $authentik_email;
    proxy_set_header X-authentik-username $authentik_username;
    proxy_set_header X-authentik-groups $authentik_groups;
    proxy_set_header X-authentik-uid $authentik_uid;
}
```

#### 3. Configuration Airboard pour SSO

Dans Coolify, définir :
```env
SSO_ENABLED=true
SSO_AUTO_PROVISION=true
SSO_DEFAULT_ROLE=user
SSO_DEFAULT_GROUP=Common
SSO_ADMIN_GROUPS=airboard-admins
```

#### 4. Test du Flux SSO

1. Accéder à `https://airboard.votre-domaine.com`
2. Redirection vers Authentik
3. Login avec Microsoft 365 / Google / LDAP
4. Authentik authentifie et renvoie vers Airboard
5. Airboard auto-provisionne l'utilisateur
6. Redirection vers le dashboard

### Installation Développement

#### 1. Prérequis

```bash
# Installer Go
wget https://go.dev/dl/go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

# Installer Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
node --version
npm --version

# Installer PostgreSQL
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

#### 2. Configuration Base de Données

```bash
# Se connecter à PostgreSQL
sudo -u postgres psql

# Créer base et utilisateur
CREATE DATABASE airboard;
CREATE USER airboard WITH PASSWORD 'airboard123';
GRANT ALL PRIVILEGES ON DATABASE airboard TO airboard;
\q
```

#### 3. Configuration Backend

```bash
cd backend

# Créer .env
cat > .env << EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=airboard
DB_PASSWORD=airboard123
DB_NAME=airboard
JWT_SECRET=dev-secret-key-not-for-production
GIN_MODE=debug
SSO_ENABLED=false
EOF

# Installer dépendances
go mod download

# Lancer le serveur
go run main.go
```

Le backend démarre sur **http://localhost:8080**

#### 4. Configuration Frontend

```bash
cd frontend

# Installer dépendances
npm install

# Créer .env.local
cat > .env.local << EOF
VITE_API_URL=http://localhost:8080/api/v1
EOF

# Lancer le dev server
npm run dev
```

Le frontend démarre sur **http://localhost:3000**

#### 5. Accès & Test

- Frontend : http://localhost:3000
- Backend API : http://localhost:8080/api/v1
- Health check : http://localhost:8080/health

**Comptes de test :**
- Admin : `admin` / `admin123`
- User : `user` / `user123`

---

## Configuration

### Variables d'Environnement

Toutes les variables peuvent être configurées via :
- **Coolify** : Section Environment Variables
- **Docker Compose** : Variables shell ou fichier `.env`
- **Développement** : Fichier `.env` dans le dossier racine

#### Base de Données

| Variable | Description | Défaut | Requis |
|----------|-------------|--------|--------|
| `DB_HOST` | Hôte PostgreSQL | `postgres` | Non |
| `DB_PORT` | Port PostgreSQL | `5432` | Non |
| `DB_USER` | Utilisateur BDD | `airboard` | Non |
| `DB_PASSWORD` | Mot de passe BDD | `airboard123` | **OUI (prod)** |
| `DB_NAME` | Nom de la BDD | `airboard` | Non |

#### JWT & Sécurité

| Variable | Description | Défaut | Requis |
|----------|-------------|--------|--------|
| `JWT_SECRET` | Clé secrète JWT (min 32 chars) | Auto-généré | **OUI (prod)** |
| `JWT_TOKEN_EXPIRATION_HOURS` | Durée token accès (heures) | `24` | Non |
| `JWT_REFRESH_EXPIRATION_DAYS` | Durée refresh token (jours) | `7` | Non |
| `BCRYPT_COST` | Coût bcrypt (10-31) | `12` | Non |

**Génération sécurisée de JWT_SECRET :**
```bash
# Option 1 : OpenSSL
openssl rand -base64 64

# Option 2 : Python
python3 -c "import secrets; print(secrets.token_urlsafe(64))"

# Option 3 : Laisser vide pour auto-génération
JWT_SECRET=
```

#### Application

| Variable | Description | Défaut | Requis |
|----------|-------------|--------|--------|
| `GIN_MODE` | Mode Gin (`debug`/`release`) | `release` | Non |
| `PUBLIC_URL` | URL publique de l'app | `http://localhost` | Non |
| `SIGNUP_ENABLED` | Activer l'inscription classique | `true` | Non |

#### SSO (Authentik)

**IMPORTANT :** SSO est activé par défaut. Désactivez si vous n'utilisez pas Authentik.

| Variable | Description | Défaut | Requis |
|----------|-------------|--------|--------|
| `SSO_ENABLED` | Activer SSO | `true` | **OUI (désactiver si pas de SSO)** |
| `SSO_AUTO_PROVISION` | Auto-créer utilisateurs SSO | `true` | Non |
| `SSO_DEFAULT_ROLE` | Rôle par défaut (`user`/`admin`) | `user` | Non |
| `SSO_DEFAULT_GROUP` | Groupe par défaut | `Common` | Non |
| `SSO_ADMIN_GROUPS` | Groupes admin (séparés par virgule) | `airboard-admins` | Non |

**Headers SSO détectés :**
- `X-authentik-email` - Email utilisateur
- `X-authentik-username` - Nom d'utilisateur
- `X-authentik-groups` - Groupes (séparés par virgule)
- `X-authentik-uid` - ID unique

#### OAuth2

| Variable | Description | Configuration |
|----------|-------------|---------------|
| Google OAuth | Configuration Google | Via interface admin |
| Microsoft OAuth | Configuration Microsoft | Via interface admin |

Configuration via `/admin/oauth` dans l'interface.

#### Mises à Jour Automatiques

| Variable | Description | Exemple |
|----------|-------------|---------|
| `GITHUB_REPO` | Repo GitHub (format `owner/repo`) | `username/airboard` |
| `GITHUB_TOKEN` | Token GitHub (optionnel) | `ghp_xxxxx` |

Laisser `GITHUB_REPO` vide pour désactiver les vérifications.

#### Stockage Médias (Optionnel)

| Variable | Description | Défaut |
|----------|-------------|--------|
| `STORAGE_TYPE` | Type (`local`/`s3`/`minio`) | `local` |
| `UPLOAD_DIR` | Répertoire local | `./uploads` |
| `S3_BUCKET` | Nom bucket S3/MinIO | - |
| `S3_REGION` | Région S3 | - |
| `S3_ENDPOINT` | Endpoint MinIO | - |
| `S3_ACCESS_KEY` | Access Key | - |
| `S3_SECRET_KEY` | Secret Key | - |

### Checklist Sécurité Production

Avant de déployer en production :

- [ ] **Changer `DB_PASSWORD`** - Mot de passe fort (16+ chars)
- [ ] **Configurer `JWT_SECRET`** - Secret aléatoire (32+ chars)
- [ ] **Définir `PUBLIC_URL`** - URL HTTPS de production
- [ ] **Configurer `SSO_ENABLED`** - `true` si SSO, `false` sinon
- [ ] **Vérifier `BCRYPT_COST`** - Minimum 12, recommandé 13
- [ ] **Activer `GIN_MODE=release`** - Mode production
- [ ] **Configurer `SSO_ADMIN_GROUPS`** - Groupes réels
- [ ] **Changer mots de passe par défaut** - `admin` et `user`
- [ ] **Activer HTTPS** - Via Nginx/Caddy/Traefik
- [ ] **Configurer sauvegardes BDD** - PostgreSQL backups
- [ ] **Vérifier logs** - Aucune erreur au démarrage

**Exemple configuration production :**
```env
DB_PASSWORD=K8$mP9#nQ2@vL5*wR7&xT4!yU6
JWT_SECRET=a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6
PUBLIC_URL=https://tools.company.com
SSO_ENABLED=true
SSO_ADMIN_GROUPS=airboard-admins,it-admins,sysadmins
GIN_MODE=release
BCRYPT_COST=13
```

---

## Architecture

### Stack Technique

| Composant | Technologie | Rôle |
|-----------|-------------|------|
| **Backend** | Go 1.21 + Gin | Serveur API REST |
| **Frontend** | Vue.js 3 + Vite | Application SPA |
| **Base de Données** | PostgreSQL 15 | Persistance données |
| **ORM** | GORM | Abstraction BDD |
| **Authentification** | JWT | Auth basée tokens |
| **SSO** | Authentik | Intégration IdP |
| **Éditeur Texte** | Tiptap | WYSIWYG pour actualités |
| **Conteneurisation** | Docker + Compose | Orchestration |
| **Reverse Proxy** | Nginx | Serveur web prod |
| **Styling** | TailwindCSS | CSS utility-first |
| **Icônes** | Iconify | Bibliothèque d'icônes |
| **État** | Pinia | Gestion d'état Vue |
| **Routing** | Vue Router | Routing client |
| **i18n** | vue-i18n | Internationalisation |

### Architecture Backend (Go)

```
backend/
├── main.go                    # Point d'entrée, routes, migrations
├── config/
│   └── config.go              # Config env, SSO
├── handlers/
│   ├── auth.go                # Login, SSO, refresh
│   ├── users.go               # CRUD utilisateurs
│   ├── groups.go              # CRUD groupes
│   ├── applications.go        # CRUD applications
│   ├── group_admin.go         # Endpoints Group Admin
│   ├── analytics.go           # Analytics
│   ├── announcements.go       # Annonces
│   ├── news.go                # Actualités & réactions
│   └── news_categories.go     # Catégories & tags
├── middleware/
│   ├── auth.go                # Validation JWT
│   ├── group_admin.go         # Permissions Group Admin
│   ├── cors.go                # Headers CORS
│   ├── sso.go                 # Détection headers SSO
│   └── editor.go              # Middleware éditeur
├── models/
│   ├── models.go              # User, Group, App, Analytics
│   └── news.go                # News, Category, Tag, Reaction
└── services/
    └── sso_mapper.go          # Sync utilisateurs SSO
```

**Patterns clés :**
- **Flux SSO** : Nginx forward headers → middleware SSO → auto-provision → JWT
- **Auth** : JWT avec refresh tokens
- **Authorization** : Middlewares par rôle (`RequireAuth`, `RequireAdmin`, `RequireGroupAdmin`, `RequireEditor`)
- **Group Admin** : Gestion déléguée avec vérifications de permissions
- **BDD** : GORM avec auto-migrations
- **Routes** : Groupées sous `/api/v1/`

### Architecture Frontend (Vue.js)

```
frontend/src/
├── main.js                    # Point d'entrée
├── App.vue                    # Composant racine, SSO auto-login
├── router/
│   └── index.js               # Routes + guards
├── stores/
│   ├── auth.js                # État auth + SSO
│   ├── app.js                 # État app
│   └── favorites.js           # Favoris
├── services/
│   └── api.js                 # Tous les endpoints API
├── views/
│   ├── dashboard/             # Dashboard, Analytics
│   ├── admin/                 # Pages admin
│   ├── group-admin/           # Pages Group Admin
│   ├── auth/                  # Login, OAuth callback
│   ├── NewsCenter.vue         # Hub actualités
│   └── NewsDetail.vue         # Détail article
├── components/
│   ├── layout/                # Sidebar, Navbar
│   ├── admin/                 # Modales admin
│   ├── group-admin/           # Modales Group Admin
│   ├── news/                  # Composants actualités
│   │   ├── RichTextEditor.vue
│   │   ├── TiptapRenderer.vue
│   │   ├── NewsCard.vue
│   │   └── ViewModeSelector.vue
│   └── dashboard/             # Composants dashboard
└── locales/
    ├── fr.json                # Traductions françaises
    ├── en.json                # Traductions anglaises
    ├── es.json                # Traductions espagnoles
    └── ar.json                # Traductions arabes
```

**Patterns clés :**
- **État** : Pinia pour auth, settings, favoris
- **SSO Auto-Login** : `App.vue` vérifie session SSO au mount
- **API** : Centralisée dans `services/api.js`
- **Guards** : Vérification auth + rôles dans router
- **Édition** : Tiptap pour contenu riche

### Modèles de Données (GORM)

**Principaux modèles :**

- **User** - Utilisateurs avec rôle, SSO, groupes gérés
- **Group** - Groupes d'utilisateurs avec admins
- **AppGroup** - Catégories d'apps (public/privé)
- **Application** - Applications individuelles
- **News** - Articles avec contenu JSON (Tiptap)
- **NewsCategory** - Catégories (General, Tutorials, etc.)
- **Tag** - Tags pour articles
- **NewsReaction** - Réactions utilisateur
- **ApplicationClick** - Analytics clics
- **Announcement** - Bannières annonces

**Relations importantes :**

```
User ←→ Groups (many-to-many)
User ←→ ManagedGroups (many-to-many, Group Admins)
Group ←→ Admins (many-to-many, Group Admins)
Group ←→ AppGroups (many-to-many)
AppGroup → OwnerGroup (belongs-to, pour groupes privés)
Application → AppGroup (belongs-to)
News → Category (belongs-to)
News ←→ Tags (many-to-many)
NewsReaction → User + News (belongs-to)
```

### Flux SSO Détaillé

1. **Accès utilisateur** → `https://airboard.company.com`
2. **NPM détecte** → pas authentifié
3. **Redirection** → Authentik
4. **Auth Authentik** → Microsoft 365 / Google / LDAP
5. **NPM forward** → headers `X-authentik-*`
6. **Middleware SSO** → détecte headers
7. **Service SSO Mapper** :
   - Cherche utilisateur par email
   - Auto-provision si inexistant
   - Sync groupes Authentik → Airboard
   - Assigne rôle admin si groupe match `SSO_ADMIN_GROUPS`
8. **Génération JWT** → token + refresh token
9. **Réponse frontend** → stockage localStorage
10. **Redirection** → `/dashboard`

---

## Sécurité

### Conformité OWASP 2025

Airboard implémente les meilleures pratiques de sécurité OWASP :

#### 1. Broken Access Control
- Vérification JWT sur tous les endpoints protégés
- Middlewares de rôle (`RequireAdmin`, `RequireGroupAdmin`, `RequireEditor`)
- Validation ownership pour Group Admins
- Pas d'accès direct aux ressources d'autres groupes

#### 2. Cryptographic Failures
- Hashage bcrypt pour mots de passe (coût configurable 10-31)
- JWT signé avec secret fort (min 32 chars)
- Tokens expirables et refresh tokens
- HTTPS obligatoire en production

#### 3. Injection
- Protection SQL via GORM (prepared statements)
- Validation stricte des entrées utilisateur
- Sanitization des données Tiptap
- Pas d'exécution de commandes shell avec input utilisateur

#### 4. Insecure Design
- Architecture en couches (handlers → services → models)
- Séparation des responsabilités
- Principe du moindre privilège (Group Admins)
- Validation côté serveur systématique

#### 5. Security Misconfiguration
- Headers de sécurité (CSP, CORS, X-Frame-Options)
- Mode release en production (`GIN_MODE=release`)
- Logs sécurisés (pas de secrets)
- Configuration via variables d'environnement

#### 6. Vulnerable Components
- Dépendances mises à jour régulièrement
- Vérification CVE avec `go mod` et `npm audit`
- Images Docker officielles
- Scan sécurité automatique

#### 7. Identification Failures
- Rate limiting sur login
- Authentification multi-facteur (via SSO)
- Sessions expirables
- Refresh tokens avec rotation

#### 8. Software & Data Integrity Failures
- Vérification intégrité fichiers uploadés
- Validation types MIME
- Limites de taille fichiers
- Pas d'exécution code uploadé

#### 9. Security Logging Failures
- Logs d'authentification
- Logs d'actions admin
- Logs d'erreurs sécurité
- Monitoring SSO

#### 10. Server-Side Request Forgery
- Validation URLs externes
- Pas de requêtes vers URLs utilisateur
- Whitelist domaines autorisés

### Headers de Sécurité

**Nginx (production) :**
```nginx
# CSP
add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self'";

# Protection XSS
add_header X-XSS-Protection "1; mode=block";

# Anti-clickjacking
add_header X-Frame-Options "SAMEORIGIN";

# Type sniffing
add_header X-Content-Type-Options "nosniff";

# HTTPS only
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains";
```

### Recommandations Déploiement Sécurisé

1. **Activer HTTPS** - Obligatoire en production (Let's Encrypt, Cloudflare)
2. **Changer secrets** - `DB_PASSWORD`, `JWT_SECRET`
3. **Augmenter bcrypt** - `BCRYPT_COST=13` ou 14
4. **Restreindre accès BDD** - Pas d'exposition externe
5. **Activer firewall** - Ports 80/443 uniquement
6. **Sauvegardes régulières** - PostgreSQL dumps quotidiens
7. **Logs centralisés** - Monitoring et alertes
8. **Mises à jour** - Vérifier GitHub releases
9. **Audits sécurité** - Tests réguliers
10. **SSO recommandé** - Authentification centralisée

---

## API

### Base URL

```
http://localhost:8080/api/v1      # Développement
https://airboard.company.com/api/v1  # Production
```

### Authentification

Tous les endpoints protégés nécessitent un header `Authorization` :

```
Authorization: Bearer <jwt_token>
```

### Endpoints Principaux

#### Authentication

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| POST | `/auth/login` | Login classique | Non |
| POST | `/auth/refresh` | Rafraîchir JWT | Non |
| GET | `/auth/sso/auto-login` | Auto-login SSO | Headers |

#### Group Admin

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/group-admin/app-groups` | Groupes d'apps gérés | Group Admin |
| GET | `/group-admin/managed-groups` | Groupes gérés | Group Admin |
| GET | `/group-admin/applications` | Apps des groupes gérés | Group Admin |
| POST | `/group-admin/applications` | Créer application | Group Admin |
| PUT | `/group-admin/applications/:id` | Modifier application | Group Admin |
| DELETE | `/group-admin/applications/:id` | Supprimer application | Group Admin |
| POST | `/group-admin/news` | Créer article | Group Admin |
| PUT | `/group-admin/news/:id` | Modifier article | Group Admin |
| DELETE | `/group-admin/news/:id` | Supprimer article | Group Admin |

#### Admin - Utilisateurs

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/users` | Liste utilisateurs | Admin |
| GET | `/users/:id` | Détail utilisateur | Admin |
| POST | `/admin/users` | Créer utilisateur | Admin |
| PUT | `/admin/users/:id` | Modifier utilisateur | Admin |
| DELETE | `/admin/users/:id` | Supprimer utilisateur | Admin |

#### Admin - Groupes

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/admin/groups` | Liste groupes | Admin |
| POST | `/admin/groups` | Créer groupe | Admin |
| PUT | `/admin/groups/:id` | Modifier groupe | Admin |
| DELETE | `/admin/groups/:id` | Supprimer groupe | Admin |

#### Applications

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/dashboard` | Dashboard utilisateur | User |
| GET | `/admin/applications` | Liste applications | Admin |
| POST | `/admin/applications` | Créer application | Admin |
| PUT | `/admin/applications/:id` | Modifier application | Admin |
| DELETE | `/admin/applications/:id` | Supprimer application | Admin |

#### Favoris

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/favorites` | Liste favoris | User |
| POST | `/favorites/:id` | Ajouter favori | User |
| DELETE | `/favorites/:id` | Retirer favori | User |

#### Actualités

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/news` | Liste articles (filtres) | User |
| GET | `/news/article/:slug` | Détail article | User |
| POST | `/news/:id/view` | Incrémenter vues | User |
| GET | `/news/:id/reactions` | Réactions article | User |
| POST | `/news/:id/react` | Ajouter réaction | User |
| DELETE | `/news/:id/react` | Retirer réaction | User |
| POST | `/editor/news` | Créer article | Editor |
| PUT | `/editor/news/:id` | Modifier article | Editor |
| DELETE | `/editor/news/:id` | Supprimer article | Editor |
| POST | `/admin/news/:id/pin` | Épingler article | Admin |

#### Analytics

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/analytics/dashboard` | Dashboard analytics | Admin |
| POST | `/analytics/click` | Tracker clic app | User |
| GET | `/admin/news/analytics` | Analytics actualités | Admin |

#### Annonces

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/announcements/active` | Annonces actives | User |
| GET | `/admin/announcements` | Liste annonces | Admin |
| POST | `/admin/announcements` | Créer annonce | Admin |
| PUT | `/admin/announcements/:id` | Modifier annonce | Admin |
| DELETE | `/admin/announcements/:id` | Supprimer annonce | Admin |

#### Sondages

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/polls` | Liste sondages (avec filtres) | User |
| GET | `/polls/:id` | Détail sondage | User |
| POST | `/polls/:id/vote` | Voter sur un sondage | User |
| DELETE | `/polls/:id/vote` | Retirer vote | User |
| GET | `/polls/:id/results` | Résultats sondage | User |
| POST | `/admin/polls` | Créer sondage | Admin |
| PUT | `/admin/polls/:id` | Modifier sondage | Admin |
| DELETE | `/admin/polls/:id` | Supprimer sondage | Admin |
| POST | `/admin/polls/:id/close` | Fermer sondage | Admin |
| GET | `/admin/polls/stats` | Statistiques sondages | Admin |
| POST | `/group-admin/polls` | Créer sondage | Group Admin |
| PUT | `/group-admin/polls/:id` | Modifier sondage | Group Admin |
| DELETE | `/group-admin/polls/:id` | Supprimer sondage | Group Admin |

#### Notifications

| Méthode | Endpoint | Description | Auth |
|---------|----------|-------------|------|
| GET | `/notifications` | Liste notifications utilisateur | User |
| GET | `/notifications/unread-count` | Compteur non lues | User |
| PUT | `/notifications/:id/read` | Marquer comme lu | User |
| PUT | `/notifications/:id/unread` | Marquer comme non lu | User |
| PUT | `/notifications/mark-all-read` | Tout marquer comme lu | User |
| DELETE | `/notifications/:id` | Supprimer notification | User |
| DELETE | `/notifications/clear-all` | Supprimer toutes | User |
| POST | `/admin/notifications` | Créer notification | Admin |
| POST | `/admin/notifications/broadcast` | Diffuser aux groupes | Admin |
| GET | `/admin/notifications/stats` | Statistiques notifications | Admin |

### Exemples Requêtes

**Login :**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

**Créer Application (Group Admin) :**
```bash
curl -X POST http://localhost:8080/api/v1/group-admin/applications \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name":"Mon App",
    "url":"https://app.example.com",
    "icon":"mdi:rocket",
    "color":"#4F46E5",
    "app_group_id":1
  }'
```

**Créer Article :**
```bash
curl -X POST http://localhost:8080/api/v1/editor/news \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Mon Article",
    "content":"{\"type\":\"doc\",\"content\":[...]}",
    "category_id":1,
    "type":"article",
    "tag_ids":[1,2]
  }'
```

**Créer Sondage :**
```bash
curl -X POST http://localhost:8080/api/v1/admin/polls \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Quel est votre framework préféré?",
    "description":"Aidez-nous à améliorer notre stack",
    "poll_type":"single",
    "is_anonymous":false,
    "show_results":"after",
    "options":[
      {"text":"Vue.js","order":1},
      {"text":"React","order":2},
      {"text":"Angular","order":3}
    ]
  }'
```

**Voter sur Sondage :**
```bash
curl -X POST http://localhost:8080/api/v1/polls/1/vote \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "poll_option_id":1
  }'
```

**Obtenir Notifications :**
```bash
curl -X GET http://localhost:8080/api/v1/notifications \
  -H "Authorization: Bearer <token>"
```

**Marquer Notification comme Lue :**
```bash
curl -X PUT http://localhost:8080/api/v1/notifications/1/read \
  -H "Authorization: Bearer <token>"
```

Documentation complète : Voir [CLAUDE.md](CLAUDE.md) pour tous les endpoints.

---

## Dépannage

### SSO ne fonctionne pas

**Symptômes :** Page de login affichée au lieu de l'auto-login

**Solutions :**
1. Vérifier `SSO_ENABLED=true` dans variables d'environnement
2. Vérifier headers Authentik dans DevTools (Network > Headers)
3. Consulter logs backend : `docker-compose logs backend | grep SSO`
4. Vérifier NPM forward headers (Advanced config)
5. Tester avec curl :
```bash
curl -H "X-authentik-email: test@example.com" \
     -H "X-authentik-username: test" \
     http://localhost:8080/api/v1/auth/sso/auto-login
```

### Erreurs 404 API

**Symptômes :** Appels API retournent 404

**Solutions :**
1. Vérifier `proxy_set_header Host $host;` dans NPM
2. Vérifier domaine Coolify correct
3. Vérifier `nginx.conf` frontend proxy `/api/v1/`
4. Tester directement backend : `curl http://backend:8080/health`

### Connexion BDD échoue

**Symptômes :** Backend crash au démarrage

**Solutions :**
1. Vérifier `DB_HOST=postgres` (nom service Docker)
2. Vérifier `DB_PASSWORD` identique dans postgres et backend
3. Attendre health check : `docker-compose logs postgres`
4. Tester connexion : `docker-compose exec postgres psql -U airboard`

### Conflits de ports

**Symptômes :** "port already allocated"

**Solutions :**
1. Vérifier ports 80, 8080, 5432 libres : `sudo netstat -tulpn`
2. Arrêter services conflictuels : `sudo systemctl stop nginx`
3. Utiliser `docker-compose.prod.yaml` avec `expose` au lieu de `ports`

### Application lente

**Symptômes :** Temps de réponse élevés

**Solutions :**
1. Vérifier ressources Docker : `docker stats`
2. Augmenter RAM : Modifier limites Docker
3. Vérifier index BDD : `docker-compose exec postgres psql -U airboard -c "\d+ news"`
4. Réduire `BCRYPT_COST` si login lent (min 10, max 13)
5. Vérifier logs : `docker-compose logs --tail=100`

### Échec migrations BDD

**Symptômes :** Erreurs au démarrage backend

**Solutions :**
1. Supprimer volume : `docker-compose down -v`
2. Recréer BDD : `docker-compose up -d postgres`
3. Vérifier logs migrations : `docker-compose logs backend | grep "migrating"`
4. Connexion manuelle :
```bash
docker-compose exec postgres psql -U airboard -d airboard
\dt   # Lister tables
\q    # Quitter
```

### Images Docker volumineuses

**Symptômes :** Build lent, espace disque

**Solutions :**
1. Nettoyer images : `docker system prune -a`
2. Optimiser build : Multi-stage déjà implémenté
3. Vérifier layers : `docker images`

### Logs incomplets

**Solutions :**
```bash
# Logs complets
docker-compose logs

# Logs temps réel
docker-compose logs -f

# Logs backend uniquement
docker-compose logs -f backend

# Logs dernières 100 lignes
docker-compose logs --tail=100
```

### Reset complet

Pour recommencer à zéro :
```bash
# Arrêter et supprimer tout
docker-compose down -v

# Supprimer images
docker rmi airboard-backend airboard-frontend

# Rebuild
docker-compose build --no-cache
docker-compose up -d
```

---

## Documentation

### Guides Utilisateur

- [**HOW-TO-USE.md**](HOW-TO-USE.md) - Guide complet d'utilisation
  - Premier login et configuration
  - Utilisation quotidienne
  - Administration
  - Gestion Group Admin
  - Bonnes pratiques



- [**DOCKER_INSTALL.md**](DOCKER_INSTALL.md) - Guide installation Docker
  - Installation détaillée Docker
  - Configuration réseau
  - Troubleshooting Docker
  - Optimisations

### Fichiers de Référence

- `.env.example` - Template variables d'environnement
- `docker-compose.yaml` - Orchestration locale
- `docker-compose.prod.yaml` - Orchestration production (Coolify)

---

## Contribution

Les contributions sont les bienvenues ! Suivez ces étapes :

### 1. Fork & Clone

```bash
# Fork sur GitHub, puis :
git clone https://github.com/VOTRE_USERNAME/airboard.git
cd airboard
```

### 2. Créer une Branche

```bash
git checkout -b feature/ma-fonctionnalite
```

### 3. Développer

- Suivre le style de code existant
- Ajouter tests si applicable
- Mettre à jour la documentation

### 4. Commit

```bash
git add .
git commit -m "feat: Ajout de ma fonctionnalité"
```

**Convention commits :**
- `feat:` - Nouvelle fonctionnalité
- `fix:` - Correction bug
- `docs:` - Documentation
- `style:` - Formatage
- `refactor:` - Refactorisation
- `test:` - Tests
- `chore:` - Maintenance

### 5. Push & PR

```bash
git push origin feature/ma-fonctionnalite
```

Ouvrir une Pull Request sur GitHub avec :
- Description claire
- Screenshots si UI
- Tests effectués

### Directives Développement

- Code propre et commenté
- Pas de secrets hardcodés
- Variables d'environnement pour config
- Validation côté serveur
- Gestion d'erreurs appropriée
- Logs informatifs

---

## Roadmap

| Fonctionnalité | Statut |
|----------------|--------|
| Auth JWT | Complet |
| SSO Authentik | Complet |
| Group Admin Role | Complet |
| Groupes Privés/Publics | Complet |
| OAuth2 (Google/MS) | Complet |
| Hub Actualités | Complet |
| Sondages & Enquêtes | Complet |
| Notifications In-App | Complet |
| Analytics | Complet |
| Multi-langue | Complet |
| Docker | Complet |
| Sécurité OWASP 2025 | Complet |
| Tests automatisés | En cours |
| Intégration LDAP | Prévu |
| Support Kubernetes | Prévu |
| Application mobile | Prévu |
| Notifications push | Prévu |
| API publique | Prévu |
| WebSocket temps réel | Prévu |

---

## Support

- **Documentation** : Ce README + [HOW-TO-USE.md](HOW-TO-USE.md) + [CLAUDE.md](CLAUDE.md)
- **Issues** : [GitHub Issues](../../issues)
- **Discussions** : [GitHub Discussions](../../discussions)

---

## License

Ce projet est sous license **MIT**. Voir [LICENSE](LICENSE) pour détails.

---

## Remerciements

- [Vue.js](https://vuejs.org/) - Framework JavaScript progressif
- [Gin](https://gin-gonic.com/) - Framework web Go
- [Docker](https://docker.com/) - Plateforme de conteneurisation
- [TailwindCSS](https://tailwindcss.com/) - Framework CSS
- [Authentik](https://goauthentik.io/) - Provider d'identité open-source
- [Coolify](https://coolify.io/) - Alternative self-hosted à Heroku
- [GORM](https://gorm.io/) - ORM pour Go
- [Pinia](https://pinia.vuejs.org/) - Gestion d'état Vue
- [Tiptap](https://tiptap.dev/) - Éditeur WYSIWYG
- Tous les contributeurs open-source

---

## Statut du Projet

**Version actuelle :** 1.0.0

**Dernière mise à jour :** Décembre 2025

**Maintenance :** Active

**Production ready :** Oui

---

<div align="center">

**Si ce projet vous est utile, donnez-lui une étoile !**

[⬆ Retour en haut](#airboard---portail-dapplications-moderne)

</div>
