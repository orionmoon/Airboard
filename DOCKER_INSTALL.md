# Installation Docker - Airboard

## Prérequis

- Docker installé et démarré
- Docker Compose installé
- Port 80 disponible sur votre machine

## Étapes d'installation

### 1. Vérifier que Docker est démarré

```bash
# Démarrer Docker si nécessaire
sudo systemctl start docker

# Vérifier le statut
sudo systemctl status docker

# Ajouter votre utilisateur au groupe docker (optionnel, pour éviter sudo)
sudo usermod -aG docker $USER
# Puis déconnectez-vous et reconnectez-vous
```

### 2. Cloner le repository

```bash
git clone https://github.com/orionmoon/airboard.git
cd airboard
```

### 3. Configurer les variables d'environnement (optionnel)

```bash
# Copier le fichier exemple
cp .env.example .env

# Éditer le fichier .env si vous voulez personnaliser la configuration
nano .env
```

**Variables importantes :**
- `DB_PASSWORD` : Mot de passe de la base de données (changez-le en production)
- `JWT_SECRET` : Clé secrète pour les tokens JWT (changez-la en production)
- `SSO_ENABLED` : true/false pour activer/désactiver le SSO
- `PUBLIC_URL` : URL publique de votre application

### 4. Construire et démarrer les services

```bash
# Construire les images Docker
docker-compose build

# Démarrer tous les services
docker-compose up -d

# Vérifier que tous les services sont démarrés
docker-compose ps
```

**Sortie attendue :**
```
NAME                  IMAGE                  STATUS
airboard-backend      airboard-backend       Up
airboard-frontend     airboard-frontend      Up
airboard-postgres     postgres:15-alpine     Up (healthy)
```

### 5. Vérifier les logs

```bash
# Voir les logs de tous les services
docker-compose logs -f

# Voir les logs d'un service spécifique
docker-compose logs -f frontend
docker-compose logs -f backend
docker-compose logs -f postgres
```

### 6. Accéder à l'application

Ouvrez votre navigateur et allez sur :
- **Frontend** : http://localhost
- **Backend API** : http://localhost/api/v1 (proxyfié via nginx)

### 7. Connexion

**Comptes par défaut :**
- **Admin** : admin@airboard.com / admin123
- **Utilisateur** : user@airboard.com / user123

## Commandes utiles

### Arrêter les services

```bash
docker-compose down
```

### Arrêter et supprimer les volumes (base de données)

```bash
docker-compose down -v
```

### Redémarrer un service spécifique

```bash
docker-compose restart frontend
docker-compose restart backend
```

### Reconstruire une image après modification du code

```bash
# Frontend
docker-compose build frontend
docker-compose up -d frontend

# Backend
docker-compose build backend
docker-compose up -d backend
```

### Voir les logs en temps réel

```bash
docker-compose logs -f
```

### Exécuter une commande dans un conteneur

```bash
# Backend
docker-compose exec backend sh

# Frontend
docker-compose exec frontend sh

# PostgreSQL
docker-compose exec postgres psql -U airboard -d airboard
```

## Dépannage

### Port 80 déjà utilisé

Si le port 80 est déjà utilisé, modifiez `docker-compose.yaml` :

```yaml
frontend:
  ports:
    - "8080:80"  # Utilisez le port 8080 à la place
```

Puis accédez à http://localhost:8080

### Le frontend affiche une page vide

1. Vérifiez que le backend est bien démarré :
   ```bash
   docker-compose logs backend
   ```

2. Vérifiez que nginx est configuré correctement :
   ```bash
   docker-compose exec frontend cat /etc/nginx/nginx.conf
   ```

3. Vérifiez les logs du frontend :
   ```bash
   docker-compose logs frontend
   ```

### La base de données ne démarre pas

1. Vérifiez les logs :
   ```bash
   docker-compose logs postgres
   ```

2. Supprimez le volume et recréez :
   ```bash
   docker-compose down -v
   docker-compose up -d
   ```

### Le backend ne peut pas se connecter à la base de données

Attendez que PostgreSQL soit prêt (healthcheck) :
```bash
docker-compose ps
```

Le service postgres doit afficher `Up (healthy)`.

## Configuration avancée

### Utiliser un port différent

Modifiez `docker-compose.yaml` :

```yaml
frontend:
  ports:
    - "3000:80"  # Frontend sur le port 3000
```

### Personnaliser la configuration nginx

Modifiez `frontend/nginx.conf` puis reconstruisez :

```bash
docker-compose build frontend
docker-compose up -d frontend
```

### Configuration SSO (Authentik)

1. Définissez les variables d'environnement dans `.env` :
   ```env
   SSO_ENABLED=true
   SSO_AUTO_PROVISION=true
   SSO_DEFAULT_ROLE=user
   SSO_DEFAULT_GROUP=Common
   SSO_ADMIN_GROUPS=airboard-admins
   ```

2. Configurez Authentik pour forward les headers :
   - X-authentik-email
   - X-authentik-username
   - X-authentik-name
   - X-authentik-groups
   - X-authentik-uid

3. Redémarrez le backend :
   ```bash
   docker-compose restart backend
   ```

## Mise à jour

Pour mettre à jour Airboard :

```bash
# Récupérer les dernières modifications
git pull

# Reconstruire les images
docker-compose build

# Redémarrer les services
docker-compose up -d

# Vérifier les logs
docker-compose logs -f
```

## Support

En cas de problème :
1. Consultez les logs : `docker-compose logs -f`
2. Vérifiez la configuration : `docker-compose config`
3. Ouvrez une issue sur GitHub
