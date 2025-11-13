#!/bin/bash

# Airboard - Script de gestion principal
# Utilisation: ./airboard.sh [commande]

show_help() {
    cat << EOF
🚀 Airboard - Gestionnaire de portail d'applications

Usage: ./airboard.sh [COMMANDE]

COMMANDES:
    start-docker     Démarrer avec Docker Compose (recommandé)
    start-dev        Démarrer en mode développement mixte (BDD Docker + Code local)
    start-local      Démarrer en mode développement 100% local
    stop            Arrêter tous les services
    restart         Redémarrer les services
    fix-db          Corriger les problèmes de base de données
    test            Tester les fonctionnalités
    test-settings   Tester spécifiquement les settings
    logs            Afficher les logs
    status          Afficher l'état des services
    clean           Nettoyer les conteneurs et volumes
    setup           Configuration initiale
    help            Afficher cette aide

EXEMPLES:
    ./airboard.sh start-docker    # Démarrage standard avec Docker
    ./airboard.sh start-dev       # Développement mixte (recommandé)
    ./airboard.sh start-local     # Développement 100% local
    ./airboard.sh test-settings   # Test des paramètres
    ./airboard.sh fix-db          # Correction BDD

SERVICES:
    - Frontend: http://localhost:3000
    - Backend: http://localhost:8080
    - Health: http://localhost:8080/health

COMPTES PAR DÉFAUT:
    - Admin: admin@airboard.com / admin123
    - User: user@airboard.com / user123
EOF
}

start_docker() {
    echo "🚀 Démarrage d'Airboard avec Docker..."
    if [ ! -f ".env" ]; then
        echo "📋 Copie du fichier .env..."
        cp .env.example .env
    fi
    docker-compose up -d --build
    echo "⏳ Attente du démarrage des services..."
    sleep 10
    show_status
}

start_dev() {
    echo "🚀 Démarrage d'Airboard en mode développement mixte..."
    ./start-dev.sh
}

start_local() {
    echo "🚀 Démarrage d'Airboard en mode local..."
    ./start-local.sh
}

stop_services() {
    echo "🛑 Arrêt des services Airboard..."
    docker-compose down
    pkill -f "airboard" 2>/dev/null || true
    pkill -f "vite" 2>/dev/null || true
    echo "✅ Services arrêtés"
}

restart_services() {
    echo "🔄 Redémarrage des services..."
    stop_services
    sleep 3
    start_docker
}

fix_database() {
    echo "🔧 Correction de la base de données..."
    ./fix-database.sh
}

test_all() {
    echo "🧪 Tests complets d'Airboard..."
    ./test-airboard.sh
}

test_settings() {
    echo "🧪 Test des fonctionnalités Settings..."
    ./test-settings.sh
}

show_logs() {
    echo "📋 Logs des services:"
    if docker-compose ps | grep -q "airboard"; then
        echo "=== BACKEND LOGS ==="
        docker-compose logs --tail=30 backend
        echo ""
        echo "=== FRONTEND LOGS ==="
        docker-compose logs --tail=30 frontend
        echo ""
        echo "=== POSTGRES LOGS ==="
        docker-compose logs --tail=20 postgres
    else
        echo "❌ Services Docker non démarrés"
    fi
}

show_status() {
    echo "📊 État des services Airboard:"
    echo ""
    
    # Vérifier Docker
    if docker-compose ps 2>/dev/null | grep -q "airboard"; then
        echo "🐳 Services Docker:"
        docker-compose ps
        echo ""
    fi
    
    # Vérifier les ports
    echo "🌐 Connectivité:"
    
    # Test Frontend
    if curl -s http://localhost:3000 >/dev/null 2>&1; then
        echo "   ✅ Frontend: http://localhost:3000"
    else
        echo "   ❌ Frontend: http://localhost:3000"
    fi
    
    # Test Backend
    if curl -s http://localhost:8080/health >/dev/null 2>&1; then
        echo "   ✅ Backend: http://localhost:8080"
    else
        echo "   ❌ Backend: http://localhost:8080"
    fi
    
    # Test Health
    health_response=$(curl -s http://localhost:8080/health 2>/dev/null)
    if echo "$health_response" | grep -q "ok"; then
        echo "   ✅ Health Check: OK"
    else
        echo "   ❌ Health Check: FAIL"
    fi
}

clean_all() {
    echo "🧹 Nettoyage complet..."
    docker-compose down -v
    docker system prune -f
    docker volume prune -f
    echo "✅ Nettoyage terminé"
}

setup_project() {
    echo "⚙️ Configuration initiale d'Airboard..."
    
    # Vérifier les prérequis
    echo "🔍 Vérification des prérequis..."
    
    command -v docker >/dev/null 2>&1 || { echo "❌ Docker requis"; exit 1; }
    command -v docker-compose >/dev/null 2>&1 || { echo "❌ Docker Compose requis"; exit 1; }
    command -v node >/dev/null 2>&1 || { echo "❌ Node.js requis"; exit 1; }
    command -v go >/dev/null 2>&1 || { echo "❌ Go requis"; exit 1; }
    
    echo "✅ Prérequis OK"
    
    # Configuration des environnements
    if [ ! -f ".env" ]; then
        cp .env.example .env
        echo "📋 Fichier .env créé"
    fi
    
    # Installation des dépendances frontend
    echo "📦 Installation des dépendances frontend..."
    cd frontend && npm install && cd ..
    
    # Compilation du backend
    echo "🔨 Compilation du backend..."
    cd backend && go mod tidy && go build -o main . && cd ..
    
    echo "✅ Configuration terminée"
    echo "🚀 Utilisez './airboard.sh start-docker' pour démarrer"
}

# Script principal
case "${1:-help}" in
    "start-docker"|"start")
        start_docker
        ;;
    "start-dev"|"dev")
        start_dev
        ;;
    "start-local"|"local")
        start_local
        ;;
    "stop")
        stop_services
        ;;
    "restart")
        restart_services
        ;;
    "fix-db"|"fix-database")
        fix_database
        ;;
    "test")
        test_all
        ;;
    "test-settings")
        test_settings
        ;;
    "logs")
        show_logs
        ;;
    "status")
        show_status
        ;;
    "clean")
        clean_all
        ;;
    "setup")
        setup_project
        ;;
    "help"|"-h"|"--help")
        show_help
        ;;
    *)
        echo "❌ Commande inconnue: $1"
        echo "💡 Utilisez './airboard.sh help' pour voir les commandes disponibles"
        exit 1
        ;;
esac