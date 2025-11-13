# 🚀 Airboard - Modern Application Portal

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.0+-4FC08D.svg)](https://vuejs.org)

*A modern, secure application portal with user management and Docker support*

[🚀 Quick Start](#-quick-start) • [📖 Documentation](#-documentation) • [🤝 Contributing](#-contributing) • [📄 License](#-license)

</div>

## ✨ Features

- 🔐 **Secure Authentication** - JWT-based auth with refresh tokens
- 👥 **User Management** - Complete CRUD operations for users and groups
- 🎯 **Application Portal** - Organize apps by groups with custom icons
- 🌐 **Multi-language** - Support for French, English, Spanish, Arabic
- 🐳 **Docker Ready** - One-command deployment
- 🎨 **Modern UI** - Clean, responsive design with dark mode
- ⚡ **Fast** - Vue.js frontend with Go backend
- 🔒 **Secure** - CSP headers, CORS protection, password hashing

## 🚀 Quick Start

### Option 1: Docker (Recommended)

```bash
git clone https://github.com/orionmoon/airboard.git
cd airboard
./airboard.sh start-docker
```

Visit **http://localhost:3000**

### Option 2: Development Mode

```bash
./airboard.sh start-dev
```

Visit **http://localhost:5173**

### Option 3: Docker (Simplified)

```bash
cd docker-simple/
./docker-run.sh
```

Visit **http://localhost** (port 80)

### 👤 Default Accounts
- **Admin**: `admin@airboard.com` / `admin123`
- **User**: `user@airboard.com` / `user123`

## 📋 Requirements

- **Docker** & Docker Compose (for Docker mode)
- **Go 1.21+** (for development)
- **Node.js 18+** (for development)

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| **Backend** | Go + Gin + GORM |
| **Frontend** | Vue.js 3 + Vite + TailwindCSS |
| **Database** | PostgreSQL 15 |
| **Auth** | JWT with refresh tokens |
| **Containerization** | Docker + Docker Compose |
| **Icons** | Iconify |
| **UI** | TailwindCSS + Custom components |

## 🏗️ Architecture

```
airboard/
├── backend/              # Go API server
│   ├── handlers/         # HTTP handlers
│   ├── models/          # Database models
│   ├── middleware/      # Auth & CORS middleware
│   └── config/          # Configuration
├── frontend/            # Vue.js application
│   ├── src/
│   │   ├── components/  # Vue components
│   │   ├── views/       # Page components
│   │   ├── stores/      # Pinia stores
│   │   └── services/    # API services
│   └── nginx.conf       # Production nginx config
├── docker-compose.yml   # Docker services
└── airboard.sh         # Management script
```

## 📖 Documentation

Everything you need is in this README! For Docker deployment, see the [docker-simple/](docker-simple/) directory.

## 🎯 Available Commands

```bash
# Start in development mode
./airboard.sh start-dev

# Start with Docker
./airboard.sh start-docker

# Start locally (100% local)
./airboard.sh start-local

# Check status
./airboard.sh status

# View logs
./airboard.sh logs

# Stop services
./airboard.sh stop

# Clean and restart
./airboard.sh restart
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5433` |
| `JWT_SECRET` | JWT signing secret | Auto-generated |
| `GIN_MODE` | Gin mode | `debug` |
| `VITE_API_URL` | Frontend API URL | `http://localhost:8080/api/v1` |

### Database

PostgreSQL with automatic migrations and test data initialization.

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** your feature branch (`git checkout -b feature/AmazingFeature`)
3. **Commit** your changes (`git commit -m 'Add some AmazingFeature'`)
4. **Push** to the branch (`git push origin feature/AmazingFeature`)
5. **Open** a Pull Request

## 🐛 Issue Reporting

If you find a bug or have a feature request:

1. Check [existing issues](../../issues)
2. Create a [new issue](../../issues/new) with detailed description
3. Include environment details and reproduction steps

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Vue.js](https://vuejs.org/) team for the amazing framework
- [Gin](https://gin-gonic.com/) for the fast Go web framework
- [Docker](https://docker.com/) for containerization
- [TailwindCSS](https://tailwindcss.com/) for the utility-first CSS
- All open source contributors

## 📊 Project Status

- ✅ **Complete**: Authentication, User Management, Application Portal
- ✅ **Complete**: Docker Support, Multi-language, Modern UI
- ✅ **Complete**: Security (CSP, CORS, JWT)
- 🔄 **In Progress**: API Documentation, Automated Tests
- 📋 **Planned**: SSO Integration, Advanced Permissions

---

<div align="center">

**Star ⭐ this repository if you find it useful!**

Made with ❤️ for the developer community

</div>
