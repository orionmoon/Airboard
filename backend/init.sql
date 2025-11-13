-- Script d'initialisation de la base de données Airboard
-- Ce script s'exécute automatiquement au premier démarrage

-- Création de la base de données (optionnel, déjà créée par docker-compose)
-- CREATE DATABASE airboard;

-- Extensions PostgreSQL utiles
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Index pour améliorer les performances
-- Ces index seront créés automatiquement par GORM, mais on peut les prévoir

-- Index sur les utilisateurs
-- CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
-- CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
-- CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
-- CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);

-- Index sur les groupes d'applications
-- CREATE INDEX IF NOT EXISTS idx_app_groups_name ON app_groups(name);
-- CREATE INDEX IF NOT EXISTS idx_app_groups_is_active ON app_groups(is_active);
-- CREATE INDEX IF NOT EXISTS idx_app_groups_order ON app_groups("order");

-- Index sur les applications
-- CREATE INDEX IF NOT EXISTS idx_applications_name ON applications(name);
-- CREATE INDEX IF NOT EXISTS idx_applications_app_group_id ON applications(app_group_id);
-- CREATE INDEX IF NOT EXISTS idx_applications_is_active ON applications(is_active);
-- CREATE INDEX IF NOT EXISTS idx_applications_order ON applications("order");

-- Table pour les paramètres de l'application
CREATE TABLE IF NOT EXISTS app_settings (
    id SERIAL PRIMARY KEY,
    app_name VARCHAR(255) NOT NULL DEFAULT 'Airboard',
    app_icon VARCHAR(255) NOT NULL DEFAULT 'mdi:view-dashboard',
    dashboard_title VARCHAR(255) NOT NULL DEFAULT 'Dashboard',
    welcome_message TEXT NOT NULL DEFAULT 'Welcome to your application portal',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insérer les paramètres par défaut
INSERT INTO app_settings (app_name, app_icon, dashboard_title, welcome_message) 
VALUES ('Airboard', 'mdi:view-dashboard', 'Dashboard', 'Welcome to your application portal')
ON CONFLICT DO NOTHING;

-- Créer les groupes par défaut
INSERT INTO groups (name, description, color, is_active, created_at, updated_at) 
VALUES 
    ('admin', 'Groupe administrateur avec tous les privilèges', '#EF4444', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('common', 'Groupe par défaut pour tous les nouveaux utilisateurs', '#3B82F6', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (name) DO NOTHING;

-- Message de confirmation
SELECT 'Base de données Airboard initialisée avec succès!' AS message;