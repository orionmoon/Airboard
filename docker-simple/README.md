# 🐳 Airboard - Docker Simple

## 📁 Répertoire Docker Simplifié

Ce répertoire contient une version simplifiée du déploiement Docker d'Airboard :
- Une seule image contenant Frontend + Backend + Nginx
- Configuration Docker simple et claire
- Démarrage en une commande

## 🚀 Démarrage Rapide

**⚠️ Important :** Exécutez ces commandes depuis ce répertoire `docker-simple/`

### **Option 1: Script Automatique (Recommandé)**
```bash
cd docker-simple/
./docker-run.sh
```

Le script créera automatiquement un fichier `.env` avec les variables nécessaires.

### **Option 2: Commandes Manuelles**
```bash
cd docker-simple/

# 1. Construire l'image (depuis le contexte parent)
docker build -t airboard -f Dockerfile ..

# 2. Démarrer les services
docker-compose up -d

# 3. Accéder à l'application
# http://localhost
```

### **Option 3: Docker Run Direct**
```bash
cd docker-simple/

# 1. Démarrer PostgreSQL
docker run -d --name airboard-db \
  -e POSTGRES_DB=airboard \
  -e POSTGRES_USER=airboard \
  -e POSTGRES_PASSWORD=airboard123 \
  postgres:15-alpine

# 2. Construire et démarrer Airboard
docker build -t airboard -f Dockerfile ..
docker run -d --name airboard-app \
  --link airboard-db:postgres \
  -p 80:80 \
  -e DB_HOST=postgres \
  airboard
```

## 📋 Accès à l'Application

- **URL** : http://localhost
- **Admin** : admin@airboard.com / admin123
- **User** : user@airboard.com / user123

## 🛠️ Commandes Utiles

```bash
# Voir les logs
docker-compose logs -f

# Arrêter les services
docker-compose down

# Redémarrer
docker-compose restart

# Nettoyer complètement
docker-compose down -v
docker rmi airboard
```

## 📊 Services

| Service | Port | Description |
|---------|------|-------------|
| **Airboard** | 80 | Application complète (Frontend + Backend + Nginx) |
| **PostgreSQL** | 5432 | Base de données (interne) |

## 🔧 Configuration

Les variables d'environnement principales :

```bash
DB_HOST=postgres
DB_PORT=5432
DB_NAME=airboard
DB_USER=airboard
DB_PASSWORD=airboard123
```

## 🐛 Dépannage

### Problème de port
```bash
# Si le port 80 est occupé, modifiez dans docker-compose.simple.yml :
ports:
  - "8080:80"  # Accès via http://localhost:8080
```

### Voir les logs d'erreur
```bash
docker logs airboard-app
docker logs airboard-db
```

### Reconstruction complète
```bash
docker-compose -f docker-compose.simple.yml down -v
docker rmi airboard
./docker-run.sh
```

## ✅ Vérification

Si tout fonctionne, vous devriez voir :
- ✅ http://localhost accessible
- ✅ Login fonctionnel
- ✅ Dashboard avec sidebar
- ✅ Gestion des utilisateurs accessible

---

**🎯 Un seul Dockerfile, un seul docker-compose, démarrage en une commande !**