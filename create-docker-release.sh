#!/bin/bash

# Script de création d'une release Docker Airboard
set -e

VERSION="${1:-1.0.0}"
RELEASE_NAME="airboard-docker-v${VERSION}"

echo "🐳 Création de la release Docker Airboard v${VERSION}..."
echo ""

# Vérifier que Docker est installé
if ! command -v docker &> /dev/null; then
    echo "❌ Docker n'est pas installé"
    exit 1
fi

# Créer le répertoire de release
mkdir -p releases

# Créer l'archive avec tous les fichiers nécessaires pour Docker
echo "📦 Création de l'archive..."
tar -czf "releases/${RELEASE_NAME}.tar.gz" \
  --exclude='node_modules' \
  --exclude='.git' \
  --exclude='backend/backend' \
  --exclude='backend/backend-test' \
  --exclude='frontend/dist' \
  --exclude='frontend/node_modules' \
  --exclude='.env' \
  --exclude='*.log' \
  --exclude='releases' \
  --transform 's,^,airboard/,' \
  backend/ \
  frontend/ \
  docker-compose.yaml \
  docker-compose.simple.yml \
  .env.example \
  airboard.sh \
  README.md \
  DOCKER-DEPLOYMENT.md \
  QUICK-START-DOCKER.md \
  LICENSE \
  2>/dev/null || true

echo ""
echo "✅ Release créée: releases/${RELEASE_NAME}.tar.gz"
ls -lh "releases/${RELEASE_NAME}.tar.gz"

# Créer un fichier d'instructions
cat > "releases/INSTALLATION-${VERSION}.txt" << 'INSTRUCTIONS'
╔══════════════════════════════════════════════════════════════╗
║           AIRBOARD - Installation Docker                    ║
╚══════════════════════════════════════════════════════════════╝

📋 PRÉREQUIS
  • Docker 20.10+
  • Docker Compose 2.0+
  • 2 GB RAM minimum
  • Ports 3000 et 8080 disponibles

🚀 INSTALLATION

1. Extraire l'archive:
   tar -xzf airboard-docker-v*.tar.gz
   cd airboard

2. Démarrer l'application:
   ./airboard.sh start-docker

   OU manuellement:
   docker-compose up -d

3. Accéder à l'application:
   http://localhost:3000

👤 CONNEXION

   Admin:
   • Email: admin@airboard.com
   • Mot de passe: admin123

   User:
   • Email: user@airboard.com
   • Mot de passe: user123

⚙️ COMMANDES UTILES

   Voir les logs:      docker-compose logs -f
   Arrêter:            docker-compose down
   Redémarrer:         docker-compose restart
   Nettoyer:           docker-compose down -v

📚 DOCUMENTATION

   Voir DOCKER-DEPLOYMENT.md pour plus de détails

🔒 SÉCURITÉ

   ⚠️  IMPORTANT pour la production:
   1. Changer JWT_SECRET dans docker-compose.yaml
   2. Changer les mots de passe PostgreSQL
   3. Changer le mot de passe admin après connexion

═══════════════════════════════════════════════════════════════

Support: https://github.com/votre-username/airboard
INSTRUCTIONS

echo ""
echo "📄 Instructions créées: releases/INSTALLATION-${VERSION}.txt"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✨ Release prête à être distribuée!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "📦 Fichiers dans releases/:"
ls -lh releases/
echo ""
echo "🎯 Pour tester la release:"
echo "   cd /tmp"
echo "   tar -xzf $(pwd)/releases/${RELEASE_NAME}.tar.gz"
echo "   cd airboard"
echo "   ./airboard.sh start-docker"
