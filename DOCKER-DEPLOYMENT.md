# 🐳 Guide de déploiement Docker - Airboard

## 📋 Prérequis

- Docker 20.10+
- Docker Compose 2.0+
- 2 GB RAM minimum
- 5 GB espace disque

## 🚀 Installation rapide

### Option 1 : Script automatisé (Recommandé)

```bash
# Cloner ou extraire le projet
cd airboard-airboard

# Démarrer avec Docker
./airboard.sh start-docker
```

### Option 2 : Docker Compose manuel

```bash
# Cloner ou extraire le projet
cd airboard-airboard

# Créer le fichier .env (optionnel, des valeurs par défaut existent)
cp .env.example .env

# Démarrer tous les services
docker-compose up -d

# Vérifier que tout fonctionne
docker-compose ps
docker-compose logs -f
```

## 🌐 Accès à l'application

Une fois démarré, accédez à :

- **Interface Web** : http://localhost:3000
- **API Backend** : http://localhost:8080
- **Health Check** : http://localhost:8080/health

### 👤 Comptes par défaut

- **Administrateur** :
  - Email : `admin@airboard.com`
  - Mot de passe : `admin123`

- **Utilisateur standard** :
  - Email : `user@airboard.com`
  - Mot de passe : `user123`

## ⚙️ Configuration

### Variables d'environnement

Modifiez le fichier `.env` ou `docker-compose.yaml` :

```env
# Base de données
DB_HOST=postgres
DB_PORT=5432
DB_USER=airboard
DB_PASSWORD=airboard123
DB_NAME=airboard

# JWT (IMPORTANT: Changer en production)
JWT_SECRET=airboard-super-secret-key-2024

# Mode
GIN_MODE=release
```

### Ports personnalisés

Pour changer les ports, modifiez `docker-compose.yaml` :

```yaml
services:
  frontend:
    ports:
      - "8080:80"  # Changer 8080 par le port souhaité

  backend:
    ports:
      - "9090:8080"  # Changer 9090 par le port souhaité
```

## 🔧 Gestion des services

### Commandes du script airboard.sh

```bash
# Démarrer
./airboard.sh start-docker

# Arrêter
./airboard.sh stop

# Redémarrer
./airboard.sh restart

# Voir les logs
./airboard.sh logs

# Voir l'état
./airboard.sh status

# Nettoyer (supprimer données)
./airboard.sh clean
```

### Commandes Docker Compose

```bash
# Démarrer en arrière-plan
docker-compose up -d

# Voir les logs en temps réel
docker-compose logs -f

# Voir les logs d'un service spécifique
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f postgres

# Arrêter les services
docker-compose stop

# Arrêter et supprimer les conteneurs
docker-compose down

# Arrêter et supprimer conteneurs + volumes (ATTENTION: perte de données)
docker-compose down -v

# Redémarrer un service
docker-compose restart backend

# Voir l'état des services
docker-compose ps

# Reconstruire les images
docker-compose build --no-cache
docker-compose up -d
```

## 🔄 Mise à jour de l'application

```bash
# 1. Arrêter les services
docker-compose down

# 2. Sauvegarder les données (optionnel)
docker-compose exec postgres pg_dump -U airboard airboard > backup.sql

# 3. Mettre à jour le code source
git pull  # ou extraire la nouvelle version

# 4. Reconstruire les images
docker-compose build --no-cache

# 5. Redémarrer
docker-compose up -d

# 6. Vérifier
docker-compose logs -f
```

## 💾 Sauvegarde et restauration

### Sauvegarde de la base de données

```bash
# Sauvegarde complète
docker-compose exec postgres pg_dump -U airboard airboard > backup-$(date +%Y%m%d-%H%M%S).sql

# Sauvegarde avec Docker
docker-compose exec postgres pg_dump -U airboard -F c airboard > backup.dump
```

### Restauration

```bash
# Depuis un fichier SQL
docker-compose exec -T postgres psql -U airboard airboard < backup.sql

# Depuis un dump
docker-compose exec -T postgres pg_restore -U airboard -d airboard < backup.dump
```

## 🔒 Réinitialisation de la base de données

L'application dispose d'un bouton de réinitialisation accessible depuis le Dashboard Admin :

1. Se connecter en tant qu'administrateur
2. Aller dans **Administration** (`/admin`)
3. Descendre jusqu'à la **Zone de danger**
4. Cliquer sur **"Réinitialiser la base de données"**
5. Confirmer l'action
6. **Redémarrer le conteneur backend** :
   ```bash
   docker-compose restart backend
   ```

Ou via la ligne de commande :

```bash
# Supprimer tous les volumes (ATTENTION: perte de données)
docker-compose down -v

# Redémarrer
docker-compose up -d
```

## 🐛 Dépannage

### Les conteneurs ne démarrent pas

```bash
# Vérifier les logs
docker-compose logs

# Vérifier l'état
docker-compose ps

# Reconstruire les images
docker-compose build --no-cache
docker-compose up -d
```

### Erreur de connexion à la base de données

```bash
# Vérifier que PostgreSQL est démarré
docker-compose ps postgres

# Voir les logs PostgreSQL
docker-compose logs postgres

# Redémarrer PostgreSQL
docker-compose restart postgres

# Vérifier la santé
docker-compose exec postgres pg_isready -U airboard
```

### Port déjà utilisé

```bash
# Trouver quel processus utilise le port 3000
sudo lsof -i :3000

# Ou changer le port dans docker-compose.yaml
```

### Nettoyer complètement Docker

```bash
# Arrêter et supprimer tout
docker-compose down -v

# Supprimer les images
docker-compose down --rmi all

# Nettoyer Docker
docker system prune -a --volumes
```

## 📊 Monitoring

### Vérifier la santé des services

```bash
# Health check backend
curl http://localhost:8080/health

# Health check frontend
curl http://localhost:3000/health

# Statistiques des conteneurs
docker stats
```

### Accéder aux conteneurs

```bash
# Accéder au backend
docker-compose exec backend sh

# Accéder à la base de données
docker-compose exec postgres psql -U airboard airboard

# Voir les fichiers du frontend
docker-compose exec frontend sh
```

## 🌍 Déploiement en production

### Recommandations

1. **Changer les secrets** :
   - JWT_SECRET
   - Mots de passe PostgreSQL
   - Mots de passe admin par défaut

2. **Utiliser un reverse proxy** (Nginx, Traefik, Caddy)

3. **Activer HTTPS** avec Let's Encrypt

4. **Configurer les sauvegardes automatiques**

5. **Limiter les ressources** :
   ```yaml
   services:
     backend:
       deploy:
         resources:
           limits:
             cpus: '1'
             memory: 512M
   ```

6. **Utiliser des volumes nommés pour la persistance**

### Exemple avec Traefik

Voir `docker-compose.coolify.yaml` pour un exemple avec Traefik et SSL.

## 📝 Architecture Docker

```
┌─────────────────────────────────────────────────────────┐
│                    Docker Network                       │
│                  (airboard-network)                     │
│                                                         │
│  ┌──────────────┐  ┌──────────────┐  ┌─────────────┐ │
│  │  Frontend    │  │   Backend    │  │  PostgreSQL  │ │
│  │  (Vue.js)    │  │    (Go)      │  │   (v15)     │ │
│  │              │  │              │  │             │ │
│  │  Port: 80    │  │  Port: 8080  │  │  Port: 5432 │ │
│  │  Nginx       │  │  Gin         │  │             │ │
│  └──────┬───────┘  └──────┬───────┘  └──────┬──────┘ │
│         │                 │                  │        │
└─────────┼─────────────────┼──────────────────┼────────┘
          │                 │                  │
    Port 3000          Port 8080         Port 5433
          │                 │                  │
          └─────────────────┴──────────────────┘
                     localhost
```

## 📞 Support

Pour tout problème :

1. Vérifier les logs : `docker-compose logs -f`
2. Consulter ce guide de dépannage
3. Ouvrir une issue sur GitHub

---

**Airboard** - Portail applicatif moderne
Version 1.0.0 - 2025
