#!/bin/bash

# 🚀 Script de Préparation Coolify - Airboard
# Aide à préparer le déploiement sur Coolify

echo "🚀 PRÉPARATION DÉPLOIEMENT COOLIFY - AIRBOARD"
echo "=============================================="

# Couleurs
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

print_step() { echo -e "${BLUE}📋 $1${NC}"; }
print_success() { echo -e "${GREEN}✅ $1${NC}"; }
print_warning() { echo -e "${YELLOW}⚠️ $1${NC}"; }
print_error() { echo -e "${RED}❌ $1${NC}"; }

# 1. Vérifications préalables
print_step "Vérification des fichiers nécessaires..."

required_files=(
    "docker-simple/Dockerfile"
    "docker-compose.coolify.yml"
    ".env.coolify.example"
    "backend/init.sql"
)

all_good=true
for file in "${required_files[@]}"; do
    if [ -f "$file" ]; then
        print_success "$file ✓"
    else
        print_error "$file manquant"
        all_good=false
    fi
done

if [ "$all_good" = false ]; then
    print_error "Certains fichiers sont manquants. Vérifiez votre installation."
    exit 1
fi

# 2. Génération des secrets
print_step "Génération des secrets de sécurité..."

if command -v openssl >/dev/null 2>&1; then
    db_password=$(openssl rand -base64 32 | tr -d '/')
    jwt_secret=$(openssl rand -base64 32 | tr -d '/')
    
    print_success "Secrets générés automatiquement"
else
    print_warning "OpenSSL non disponible. Générez manuellement vos secrets."
    db_password="GENERATE_STRONG_PASSWORD_HERE"
    jwt_secret="GENERATE_STRONG_JWT_SECRET_HERE"
fi

# 3. Création du fichier de configuration
print_step "Création de la configuration Coolify..."

cat > .env.coolify.generated << EOF
# 🚀 Configuration générée pour Coolify - $(date)

# ===== SECRETS GÉNÉRÉS =====
DB_PASSWORD=$db_password
JWT_SECRET=$jwt_secret

# ===== CONFIGURATION =====
DOMAIN=votre-domaine.com
TZ=Europe/Paris
GIN_MODE=release
NODE_ENV=production

# ===== INSTRUCTIONS =====
# 1. Copiez DB_PASSWORD et JWT_SECRET dans Coolify
# 2. Changez DOMAIN par votre vrai domaine
# 3. Supprimez ce fichier après utilisation (contient des secrets)
EOF

print_success "Configuration créée : .env.coolify.generated"

# 4. Vérification Docker
print_step "Test de construction Docker..."

if docker build -f docker-simple/Dockerfile -t airboard-test . >/dev/null 2>&1; then
    print_success "Image Docker se construit correctement"
    docker rmi airboard-test >/dev/null 2>&1
else
    print_error "Problème de construction Docker"
    print_warning "Vérifiez que Docker fonctionne et que les fichiers sont présents"
fi

# 5. Informations pour Coolify
echo ""
print_step "📋 INFORMATIONS POUR COOLIFY"
echo "============================"

echo ""
echo "🔗 REPOSITORY:"
echo "  URL: $(git remote get-url origin 2>/dev/null || echo 'Configurez votre repository Git')"
echo "  Branche: main"
echo ""

echo "🐳 CONFIGURATION DOCKER:"
echo "  Type: Docker Compose"
echo "  Fichier: docker-compose.coolify.yml"
echo "  Port: 80"
echo "  Health Check: /health"
echo ""

echo "🔒 VARIABLES D'ENVIRONNEMENT À AJOUTER DANS COOLIFY:"
echo "  DB_PASSWORD: $db_password"
echo "  JWT_SECRET: $jwt_secret"
echo "  DOMAIN: votre-domaine.com"
echo ""

echo "⚙️ RESSOURCES RECOMMANDÉES:"
echo "  CPU: 1 vCPU minimum"
echo "  RAM: 1GB minimum"  
echo "  Storage: 10GB minimum"
echo ""

# 6. Étapes de déploiement
echo ""
print_step "📝 ÉTAPES DE DÉPLOIEMENT COOLIFY"
echo "================================"

echo ""
echo "1️⃣ PRÉPARER LE REPOSITORY:"
echo "   git add ."
echo "   git commit -m 'Ready for Coolify deployment'"
echo "   git push origin main"
echo ""

echo "2️⃣ DANS COOLIFY:"
echo "   • Nouveau Service → Git Repository"
echo "   • URL: $(git remote get-url origin 2>/dev/null || echo 'votre-repo-url')"
echo "   • Branch: main"
echo "   • Build: Automatic (détection docker-compose.coolify.yml)"
echo ""

echo "3️⃣ CONFIGURATION:"
echo "   • Environment Variables → Ajouter DB_PASSWORD et JWT_SECRET"
echo "   • Domain → Configurer votre domaine"
echo "   • SSL → Activer Let's Encrypt"
echo ""

echo "4️⃣ DÉPLOYER:"
echo "   • Cliquer sur Deploy"
echo "   • Attendre 2-3 minutes"
echo "   • Tester : https://votre-domaine.com"
echo ""

# 7. Conseils de sécurité
print_warning "🔒 SÉCURITÉ IMPORTANTE"
echo "======================"
echo ""
echo "• Supprimez .env.coolify.generated après utilisation"
echo "• Changez les mots de passe par défaut (admin123)"
echo "• Activez HTTPS obligatoire"
echo "• Configurez les backups de base de données"
echo "• Surveillez les logs régulièrement"

echo ""
print_step "✅ VÉRIFICATION POST-DÉPLOIEMENT"
echo "================================"
echo ""
echo "🌐 Application: https://votre-domaine.com"
echo "🔍 Health Check: https://votre-domaine.com/health"
echo "👤 Login Admin: admin@airboard.com / admin123"
echo "📊 Logs: Interface Coolify → Logs"

echo ""
print_success "🎊 PRÉPARATION TERMINÉE !"
print_step "Votre projet est prêt pour Coolify. Suivez les étapes ci-dessus."

echo ""
print_warning "💡 CONSEIL: Gardez cette sortie comme référence pour le déploiement"