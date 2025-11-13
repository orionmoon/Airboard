# 🚀 Configuration Coolify Correcte - Airboard

## ❌ Problème Identifié

Coolify essaie d'exécuter `docker-compose` dans le sous-répertoire `docker-simple/` au lieu d'utiliser l'image construite.

## ✅ Solution: Configuration Simple (Non-Compose)

### **Option 1: Service Simple (Recommandée)**

Dans Coolify, configurez comme **Service Simple** au lieu de **Docker Compose** :

#### **1. Type de Service**
- **Type** : `Simple Docker Service`
- **Source** : `Git Repository`
- **Repository** : `https://github.com/votre-username/airboard`
- **Branch** : `main`

#### **2. Build Configuration**
```yaml
Build Pack: Docker
Dockerfile Location: docker-simple/Dockerfile
Build Context: . (racine du repo)
Port: 80
```

#### **3. Variables d'Environnement**
```bash
DB_HOST=postgres-service-name
DB_PORT=5432
DB_NAME=airboard
DB_USER=airboard
DB_PASSWORD=your-secure-password
JWT_SECRET=your-jwt-secret
GIN_MODE=release
```

#### **4. Base de Données Séparée**
Créez un service PostgreSQL séparé :
- **Type** : `Database`
- **Engine** : `PostgreSQL 15`
- **Database Name** : `airboard`
- **Username** : `airboard`
- **Password** : `[généré par Coolify]`

---

### **Option 2: Docker Compose à la Racine**

Si vous préférez docker-compose, utilisez le fichier à la racine :

#### **1. Configuration Coolify**
- **Type** : `Docker Compose`
- **Docker Compose File** : `docker-compose.coolify.yml` (à la racine)
- **Working Directory** : `/` (racine)

#### **2. Variables d'Environnement**
```bash
DB_PASSWORD=your-secure-password
JWT_SECRET=your-jwt-secret
DOMAIN=your-domain.com
```

---

## 🔧 Solution Immédiate

### **Étape 1: Déplacer docker-compose.coolify.yml**

Le fichier doit être à la racine pour que Coolify le trouve :

```bash
# Sur votre machine
cd Documents/WindowsFiles/SI/Airboard
mv docker-simple/docker-compose.coolify.yml ./
git add .
git commit -m "Move docker-compose.coolify.yml to root for Coolify"
git push origin main
```

### **Étape 2: Reconfigurer Coolify**

Dans Coolify :
1. **Service Settings** → **Source**
2. **Docker Compose File** : `docker-compose.coolify.yml`
3. **Working Directory** : `/` (ou laisser vide)
4. **Redéployer**

---

## 🎯 Configuration Recommandée (Service Simple)

Pour éviter les complications docker-compose :

### **1. Nouveau Service dans Coolify**
- **Type** : `Service`
- **Source Type** : `Git Repository`
- **Build Type** : `Docker`

### **2. Build Settings**
```yaml
Dockerfile: docker-simple/Dockerfile
Build Context: . 
Port: 80
Health Check: /health
```

### **3. Services Requis**
```yaml
# Service 1: PostgreSQL
Type: Database
Engine: PostgreSQL 15
Name: airboard-db

# Service 2: Airboard App  
Type: Service
Image: Built from Git
Port: 80
Environment:
  DB_HOST: airboard-db
  DB_PASSWORD: [from database service]
```

---

## 🚀 Déploiement Correct

Une fois configuré correctement :

1. **Build** se fait sans docker-compose
2. **Image** est construite avec le Dockerfile
3. **Service** démarre directement avec l'image
4. **Base de données** est un service séparé connecté

## 🎊 Résultat

- ✅ Pas d'erreur `docker-compose: command not found`
- ✅ Build et déploiement réussis
- ✅ Application accessible
- ✅ Base de données connectée