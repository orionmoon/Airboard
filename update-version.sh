#!/bin/bash

# Script de mise à jour de version pour Airboard
# Usage: ./update-version.sh <version> [message]
# Exemple: ./update-version.sh 1.1.0 "Ajout du système de versioning"

set -e

if [ -z "$1" ]; then
    echo "Usage: ./update-version.sh <version> [message]"
    echo "Exemple: ./update-version.sh 1.1.0 'Ajout du système de versioning'"
    exit 1
fi

VERSION=$1
MESSAGE=${2:-"Release version $VERSION"}
BUILD_DATE=$(date +%Y-%m-%d)
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

echo "🚀 Mise à jour de la version vers $VERSION"
echo "📅 Date: $BUILD_DATE"
echo "🔖 Commit: $GIT_COMMIT"
echo ""

# Récupérer les derniers commits pour le changelog
echo "📝 Génération du changelog..."
RECENT_COMMITS=$(git log --oneline -10 --pretty=format:'        "%s",' | sed '$ s/,$//')

# Créer le fichier version.json
cat > version.json <<EOF
{
  "version": "$VERSION",
  "buildDate": "$BUILD_DATE",
  "gitCommit": "$GIT_COMMIT",
  "changelog": [
    {
      "version": "$VERSION",
      "date": "$BUILD_DATE",
      "changes": [
$RECENT_COMMITS
      ]
    }
  ]
}
EOF

# Copier dans le backend
cp version.json backend/version.json

# Mettre à jour le package.json frontend
if [ -f frontend/package.json ]; then
    echo "📦 Mise à jour du package.json frontend..."
    sed -i "s/\"version\": \".*\"/\"version\": \"$VERSION\"/" frontend/package.json
fi

echo ""
echo "✅ Version mise à jour avec succès!"
echo ""
echo "Prochaines étapes:"
echo "1. Vérifiez les changements:"
echo "   git diff version.json backend/version.json frontend/package.json"
echo ""
echo "2. Commitez les changements:"
echo "   git add version.json backend/version.json frontend/package.json"
echo "   git commit -m 'chore: bump version to $VERSION'"
echo ""
echo "3. Créez un tag Git:"
echo "   git tag -a v$VERSION -m '$MESSAGE'"
echo ""
echo "4. Poussez les changements et le tag:"
echo "   git push origin main"
echo "   git push origin v$VERSION"
echo ""
echo "5. Créez une release sur GitHub:"
echo "   - Allez sur https://github.com/YOUR_USERNAME/YOUR_REPO/releases/new"
echo "   - Sélectionnez le tag v$VERSION"
echo "   - Ajoutez les notes de release"
echo "   - Publiez la release"
echo ""
