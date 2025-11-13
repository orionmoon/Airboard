#!/bin/bash

# Script de nettoyage complet du projet Airboard
# Supprime tous les fichiers non essentiels pour ne garder que le code source

set -e

echo "🧹 Nettoyage complet du projet Airboard..."
echo ""

# Fonction pour supprimer en toute sécurité
safe_remove() {
    if [ -e "$1" ]; then
        echo "  🗑️  Suppression: $1"
        rm -rf "$1"
    fi
}

# Fonction pour afficher la taille d'un répertoire
show_size() {
    if [ -e "$1" ]; then
        du -sh "$1" 2>/dev/null || echo "0"
    else
        echo "N/A"
    fi
}

echo "📊 Tailles avant nettoyage:"
echo "  Frontend node_modules: $(show_size frontend/node_modules)"
echo "  Frontend dist: $(show_size frontend/dist)"
echo "  Backend binaires: $(show_size backend/backend) $(show_size backend/backend-test)"
echo ""

echo "🗑️  Suppression des node_modules..."
safe_remove "frontend/node_modules"

echo "🗑️  Suppression des fichiers de build frontend..."
safe_remove "frontend/dist"
safe_remove "frontend/dist-ssr"
safe_remove "frontend/.vite"

echo "🗑️  Suppression des binaires Go..."
safe_remove "backend/backend"
safe_remove "backend/backend-test"
safe_remove "backend/main"

echo "🗑️  Suppression des fichiers temporaires..."
find . -type f \( -name "*.log" -o -name "*.swp" -o -name "*.swo" -o -name "*~" \) -delete 2>/dev/null || true
find . -type f -name ".DS_Store" -delete 2>/dev/null || true
find . -type f -name "Thumbs.db" -delete 2>/dev/null || true

echo "🗑️  Suppression des fichiers d'environnement locaux..."
safe_remove ".env"
safe_remove ".env.local"
safe_remove ".env.*.local"
safe_remove "backend/.env"
safe_remove "frontend/.env"
safe_remove "frontend/.env.local"

echo "🗑️  Suppression des caches..."
safe_remove ".cache"
safe_remove "frontend/.eslintcache"
safe_remove "backend/.cache"

echo "🗑️  Suppression des fichiers de coverage..."
safe_remove "coverage"
safe_remove "frontend/coverage"
safe_remove "backend/coverage"

echo "🗑️  Suppression des fichiers IDE (VSCode, Cursor, etc.)..."
safe_remove ".vscode"
safe_remove ".cursor"
safe_remove ".idea"
safe_remove "backend/.cursor-free-vip"

echo "🗑️  Suppression des fichiers de lock npm (on garde package.json)..."
# On garde package-lock.json pour la reproductibilité mais on peut le supprimer si souhaité
# safe_remove "frontend/package-lock.json"

echo "🗑️  Suppression des anciens fichiers de release..."
safe_remove "releases"
safe_remove "create-release.sh"
safe_remove "../create-release.sh"

echo ""
echo "✅ Nettoyage terminé!"
echo ""

echo "📊 Tailles après nettoyage:"
TOTAL_SIZE=$(du -sh . 2>/dev/null | cut -f1)
echo "  Taille totale du projet: $TOTAL_SIZE"
echo ""

echo "📁 Fichiers conservés:"
echo "  ✓ Code source backend (Go)"
echo "  ✓ Code source frontend (Vue.js)"
echo "  ✓ Fichiers Docker (Dockerfile, docker-compose.yaml)"
echo "  ✓ Configuration (nginx.conf, vite.config.js, etc.)"
echo "  ✓ Documentation (.md)"
echo "  ✓ Scripts de gestion (airboard.sh)"
echo "  ✓ .env.example (template)"
echo ""

echo "💡 Pour reconstruire:"
echo "  Frontend: cd frontend && npm install && npm run build"
echo "  Backend:  cd backend && go build -o backend main.go"
echo ""

echo "🐳 Pour déployer avec Docker:"
echo "  ./airboard.sh start-docker"
echo "  OU: docker-compose up -d"
echo ""
