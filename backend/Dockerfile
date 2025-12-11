# Build stage
FROM golang:1.21-alpine AS builder

# Installation des dépendances système
RUN apk add --no-cache git ca-certificates tzdata

# Définition du répertoire de travail
WORKDIR /app

# Copie des fichiers de dépendances
COPY go.mod go.sum ./

# Téléchargement des dépendances
RUN go mod download

# Copie du code source
COPY . .

# Compilation de l'application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

# Installation des certificats CA
RUN apk --no-cache add ca-certificates tzdata

# Création d'un utilisateur non-root
RUN adduser -D -s /bin/sh airboard

# Définition du répertoire de travail
WORKDIR /app

# Copie de l'exécutable depuis le stage de build
COPY --from=builder /app/main .

# Copie du fichier version.json
COPY --from=builder /app/version.json .

# Changement de propriétaire
RUN chown airboard:airboard /app/main /app/version.json

# Utilisation de l'utilisateur non-root
USER airboard

# Exposition du port
EXPOSE 8080

# Variables d'environnement par défaut
ENV GIN_MODE=release
ENV PORT=8080

# Commande de démarrage
CMD ["./main"]