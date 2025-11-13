#!/bin/bash

# 🐳 Script Simple - Démarrage Airboard avec Docker
# Usage: ./docker-run.sh

echo "🚀 DÉMARRAGE AIRBOARD AVEC DOCKER"
echo "================================="

# Couleurs
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_step() { echo -e "${BLUE}📋 $1${NC}"; }
print_success() { echo -e "${GREEN}✅ $1${NC}"; }
print_warning() { echo -e "${YELLOW}⚠️ $1${NC}"; }

print_step "Préparation de l'environnement..."
if [ ! -f ".env" ]; then
    cp .env.example .env
    print_success "Fichier .env créé depuis .env.example"
fi

print_step "Construction de l'image Airboard..."
docker build -t airboard ..

if [ $? -eq 0 ]; then
    print_success "Image construite avec succès"
else
    echo "❌ Erreur lors de la construction"
    exit 1
fi

echo ""
print_step "Démarrage des services..."
docker-compose up -d

if [ $? -eq 0 ]; then
    print_success "Services démarrés"
    
    echo ""
    print_step "⏳ Attente du démarrage complet..."
    sleep 10
    
    echo ""
    print_success "🎉 Airboard est prêt !"
    echo ""
    echo "🌐 Accès: http://localhost"
    echo "👤 Admin: admin@airboard.com / admin123"
    echo "👤 User: user@airboard.com / user123"
    echo ""
    print_warning "Commandes utiles:"
    echo "• Voir les logs: docker-compose logs -f"
    echo "• Arrêter: docker-compose down"
    echo "• Redémarrer: docker-compose restart"
    
else
    echo "❌ Erreur lors du démarrage"
    exit 1
fi