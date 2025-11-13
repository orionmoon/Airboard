# 🚀 Démarrage rapide Docker - Airboard

## Installation en 3 étapes

### 1️⃣ Prérequis
- Docker et Docker Compose installés
- Port 3000 (frontend) et 8080 (backend) disponibles

### 2️⃣ Lancer l'application

```bash
# Méthode simple
./airboard.sh start-docker

# OU avec Docker Compose
docker-compose up -d
```

### 3️⃣ Accéder à l'application

Ouvrez votre navigateur : **http://localhost:3000**

**Connexion admin** :
- Email : `admin@airboard.com`
- Mot de passe : `admin123`

---

## ⚡ Commandes essentielles

```bash
# Démarrer
./airboard.sh start-docker

# Arrêter
docker-compose down

# Voir les logs
docker-compose logs -f

# Redémarrer
docker-compose restart

# Nettoyer tout
docker-compose down -v
```

---

## 🎯 Fonctionnalités

- ✅ Portail d'applications avec groupes
- ✅ Gestion des utilisateurs et permissions
- ✅ Interface multilingue (FR, EN, ES, AR)
- ✅ Mode sombre
- ✅ Réinitialisation de la base de données (Admin → Zone de danger)

---

## 📚 Documentation complète

Voir [DOCKER-DEPLOYMENT.md](DOCKER-DEPLOYMENT.md) pour plus de détails.
