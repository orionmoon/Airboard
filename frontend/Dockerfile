# Build stage
FROM node:18-alpine AS builder

# Installation des dépendances système
RUN apk add --no-cache git

# Définition du répertoire de travail
WORKDIR /app

# Copie des fichiers de dépendances
COPY package*.json ./

# Installation des dépendances
RUN npm ci --only=production --silent

# Copie du code source
COPY . .

# Build de l'application
RUN npm run build

# Production stage
FROM nginx:alpine

# Installation des certificats CA
RUN apk --no-cache add ca-certificates

# Copie de la configuration Nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Copie des fichiers buildés
COPY --from=builder /app/dist /usr/share/nginx/html

# Création d'un utilisateur non-root
RUN addgroup -g 1001 -S airboard && \
    adduser -S airboard -u 1001

# Ajustement des permissions
RUN chown -R airboard:airboard /usr/share/nginx/html && \
    chown -R airboard:airboard /var/cache/nginx && \
    chown -R airboard:airboard /var/log/nginx && \
    chown -R airboard:airboard /etc/nginx/conf.d

# Création du répertoire pour le PID
RUN touch /var/run/nginx.pid && \
    chown -R airboard:airboard /var/run/nginx.pid

# Utilisation de l'utilisateur non-root
USER airboard

# Exposition du port
EXPOSE 80

# Variables d'environnement
ENV NODE_ENV=production

# Commande de démarrage
CMD ["nginx", "-g", "daemon off;"]