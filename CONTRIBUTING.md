# Contributing to Airboard

Thank you for your interest in contributing to **Airboard**! We welcome contributions from the community and appreciate your time and effort in helping us improve this project.

This document provides guidelines for contributing to Airboard. By following these guidelines, you help maintain code quality and make the review process smoother for everyone.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Guidelines](#code-guidelines)
- [Commit Message Convention](#commit-message-convention)
- [Pull Request Process](#pull-request-process)
- [Testing Guidelines](#testing-guidelines)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)
- [Questions and Support](#questions-and-support)

---

## Code of Conduct

This project adheres to a code of conduct that all contributors are expected to follow. Please be respectful, inclusive, and considerate in all interactions.

**Our Standards:**
- Use welcoming and inclusive language
- Be respectful of differing viewpoints and experiences
- Gracefully accept constructive criticism
- Focus on what is best for the community
- Show empathy towards other community members

---

## How Can I Contribute?

There are many ways to contribute to Airboard:

### 🐛 Reporting Bugs
Found a bug? Please [create an issue](../../issues/new) with detailed information about the problem.

### 💡 Suggesting Enhancements
Have an idea for a new feature or improvement? We'd love to hear it! Open an issue with the "enhancement" label.

### 📝 Improving Documentation
Documentation improvements are always welcome, whether it's fixing typos, clarifying instructions, or adding examples.

### 💻 Code Contributions
- Fix bugs
- Implement new features
- Improve performance
- Refactor code
- Add tests

### 🌍 Translations
Help us make Airboard accessible to more users by contributing translations for new languages or improving existing ones.

---

## Getting Started

### Prerequisites

Before you start contributing, make sure you have:

- **Go 1.21+** installed
- **Node.js 18+** and npm installed
- **PostgreSQL 15+** installed (or Docker)
- **Git** installed
- **A GitHub account**

### Fork and Clone

1. **Fork the repository** on GitHub by clicking the "Fork" button
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/airboard.git
   cd airboard
   ```
3. **Add the upstream repository**:
   ```bash
   git remote add upstream https://github.com/ORIGINAL_OWNER/airboard.git
   ```

### Setup Development Environment

#### Backend Setup

```bash
cd backend

# Install dependencies
go mod download

# Create .env file
cp ../.env.example .env

# Edit .env with your local database credentials
nano .env

# Run the backend server
go run main.go
```

Backend will start on `http://localhost:8080`

#### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

Frontend will start on `http://localhost:3000`

#### Database Setup

```bash
# Create PostgreSQL database
sudo -u postgres psql

CREATE DATABASE airboard;
CREATE USER airboard WITH PASSWORD 'airboard123';
GRANT ALL PRIVILEGES ON DATABASE airboard TO airboard;
\q
```

The backend will automatically run migrations on startup.

---

## Development Workflow

### 1. Create a Feature Branch

Always create a new branch for your work:

```bash
# Update your main branch
git checkout main
git pull upstream main

# Create a new feature branch
git checkout -b feature/your-feature-name
```

**Branch naming conventions:**
- `feature/` - New features (e.g., `feature/add-dark-mode`)
- `fix/` - Bug fixes (e.g., `fix/login-redirect-issue`)
- `docs/` - Documentation changes (e.g., `docs/update-readme`)
- `refactor/` - Code refactoring (e.g., `refactor/simplify-auth`)
- `test/` - Adding tests (e.g., `test/add-user-tests`)
- `chore/` - Maintenance tasks (e.g., `chore/update-dependencies`)

### 2. Make Your Changes

- Write clean, readable code
- Follow the [Code Guidelines](#code-guidelines) below
- Add tests for new features
- Update documentation as needed

### 3. Test Your Changes

```bash
# Backend
cd backend
go build    # Ensure it compiles
go run main.go  # Test locally

# Frontend
cd frontend
npm run lint    # Check for linting errors
npm run build   # Ensure it builds
npm run dev     # Test locally
```

### 4. Commit Your Changes

```bash
git add .
git commit -m "feat: add dark mode toggle"
```

Follow the [Commit Message Convention](#commit-message-convention) below.

### 5. Push and Create a Pull Request

```bash
# Push to your fork
git push origin feature/your-feature-name
```

Then, go to GitHub and create a Pull Request from your fork to the main repository.

---

## Code Guidelines

### General Principles

- **Keep it simple** - Avoid over-engineering
- **DRY (Don't Repeat Yourself)** - Reuse code where possible
- **SOLID principles** - Write maintainable, extensible code
- **Security first** - Always validate user input and follow OWASP best practices
- **Performance matters** - Write efficient code, especially for database queries

### Backend (Go)

#### File Organization

```
backend/
├── handlers/       # HTTP request handlers (one file per resource)
├── middleware/     # Authentication, CORS, logging
├── models/         # Database models (GORM)
├── services/       # Business logic
├── config/         # Configuration loading
└── main.go         # Entry point
```

#### Coding Standards

**1. Naming Conventions**
- Use **camelCase** for variables and functions: `getUserByID`, `newUserData`
- Use **PascalCase** for exported functions and types: `CreateUser`, `User`
- Use **snake_case** for database column names: `created_at`, `user_id`

**2. Error Handling**
```go
// Good
user, err := GetUserByID(userID)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
    return
}

// Bad - Silent error
user, _ := GetUserByID(userID)
```

**3. Input Validation**
```go
// Always validate user input
if req.Email == "" || req.Password == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
    return
}
```

**4. Database Queries**
```go
// Use GORM safely to prevent SQL injection
db.Where("email = ?", email).First(&user)  // Good
// Never: db.Where("email = " + email)      // SQL injection risk!
```

**5. Comments**
```go
// Add comments for complex logic
// GetUserGroups retrieves all groups associated with a user,
// including both direct memberships and inherited group access.
func GetUserGroups(userID uint) ([]Group, error) {
    // Implementation
}
```

### Frontend (Vue.js)

#### File Organization

```
frontend/src/
├── views/          # Page components (one file per route)
├── components/     # Reusable components
├── stores/         # Pinia state management
├── services/       # API calls (centralized in api.js)
├── router/         # Vue Router configuration
└── locales/        # i18n translations
```

#### Coding Standards

**1. Component Naming**
- Use **PascalCase** for components: `UserCard.vue`, `NewsEditor.vue`
- Use **kebab-case** in templates: `<user-card>`, `<news-editor>`

**2. Props and Events**
```vue
<!-- Good: Descriptive prop names -->
<UserCard
  :user="currentUser"
  :show-actions="true"
  @update:user="handleUserUpdate"
/>

<!-- Bad: Unclear prop names -->
<UserCard :data="u" :flag="true" @update="x" />
```

**3. API Calls**
```javascript
// Always use the centralized API service
import api from '@/services/api'

// Good
const users = await api.getUsers()

// Bad - Direct axios call
const users = await axios.get('/api/v1/users')
```

**4. State Management**
```javascript
// Use Pinia stores for shared state
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const isAdmin = authStore.isAdmin
```

**5. Code Formatting**
- Run `npm run lint` before committing
- Use 2 spaces for indentation
- Add semicolons at the end of statements
- Use single quotes for strings

### Security Guidelines

**CRITICAL:** Follow these security practices at all times:

1. **Never commit secrets** - No passwords, API keys, or tokens in code
2. **Validate all inputs** - Server-side validation is mandatory
3. **Use parameterized queries** - GORM handles this, never concatenate SQL
4. **Sanitize user content** - Especially for rich text (Tiptap)
5. **Implement proper authentication** - Check JWT tokens on all protected routes
6. **Verify permissions** - Use middleware to check roles (admin, group_admin, editor)
7. **Use HTTPS in production** - Never send tokens over HTTP
8. **Hash passwords properly** - bcrypt with cost >= 12

### Documentation Standards

- Update the README if you add features
- Document complex functions with comments
- Update CLAUDE.md if you change architecture
- Add JSDoc comments for JavaScript functions
- Update API documentation if you add/change endpoints

---

## Commit Message Convention

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, no logic change)
- **refactor**: Code refactoring (no feature or bug fix)
- **test**: Adding or updating tests
- **chore**: Maintenance tasks (dependencies, build, etc.)
- **perf**: Performance improvements
- **ci**: CI/CD pipeline changes

### Examples

```bash
# Feature
git commit -m "feat: add dark mode toggle to settings"

# Bug fix
git commit -m "fix: resolve login redirect loop on SSO"

# Documentation
git commit -m "docs: update installation guide for Docker"

# Refactoring
git commit -m "refactor: simplify user authentication logic"

# With body and breaking change
git commit -m "feat: redesign user management UI

- Replace table with card layout
- Add advanced filters
- Improve mobile responsiveness

BREAKING CHANGE: API endpoint /users now returns paginated results"
```

### Scope (Optional)

Add scope to provide additional context:

```bash
feat(auth): add OAuth2 support for Google
fix(news): correct slug generation for special characters
docs(readme): add troubleshooting section
```

---

## Pull Request Process

### Before Submitting

1. **Update your branch** with the latest upstream changes:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Test your changes** thoroughly:
   - Run the backend: `cd backend && go run main.go`
   - Run the frontend: `cd frontend && npm run dev`
   - Test all affected features
   - Check for console errors

3. **Lint your code**:
   ```bash
   # Frontend
   cd frontend
   npm run lint

   # Backend
   cd backend
   gofmt -w .
   go mod tidy
   ```

4. **Update documentation** if needed:
   - README.md for user-facing changes
   - CLAUDE.md for architecture changes
   - Code comments for complex logic

### Creating the Pull Request

1. **Push your branch** to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create a Pull Request** on GitHub with:
   - **Clear title** following commit convention (e.g., "feat: add dark mode")
   - **Detailed description** explaining:
     - What problem does this solve?
     - How does this change address it?
     - Any breaking changes?
     - Screenshots (for UI changes)
     - Testing performed

3. **Fill out the PR template** (if available)

4. **Link related issues** (e.g., "Fixes #123", "Closes #456")

### PR Template Example

```markdown
## Description
Brief description of what this PR does.

## Motivation
Why is this change needed? What problem does it solve?

## Changes Made
- Added dark mode toggle in settings
- Updated theme colors for dark mode
- Added user preference storage in localStorage

## Screenshots (if applicable)
[Add before/after screenshots for UI changes]

## Testing Performed
- [ ] Tested locally on development server
- [ ] Tested dark mode toggle functionality
- [ ] Tested theme persistence across page reloads
- [ ] No console errors
- [ ] Responsive design verified

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex logic
- [ ] Documentation updated (if needed)
- [ ] No breaking changes (or documented if present)
- [ ] Tests pass locally

## Related Issues
Fixes #123
```

### Review Process

1. **Automated checks** will run (linting, build)
2. **Maintainers will review** your code
3. **Address feedback** by pushing additional commits to your branch
4. Once approved, a maintainer will **merge your PR**

### After Your PR is Merged

1. **Delete your branch** (optional):
   ```bash
   git branch -d feature/your-feature-name
   git push origin --delete feature/your-feature-name
   ```

2. **Update your main branch**:
   ```bash
   git checkout main
   git pull upstream main
   ```

---

## Testing Guidelines

### Backend Testing

While we don't currently have automated tests, please perform these manual tests:

- **Test all affected endpoints** with curl or Postman
- **Verify database migrations** work correctly
- **Check error handling** for invalid inputs
- **Test role-based access control** (admin, group_admin, editor, user)
- **Verify SSO integration** if authentication changes are made

**Example manual tests:**
```bash
# Test login endpoint
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'

# Test protected endpoint
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer <your_token>"
```

### Frontend Testing

- **Check for console errors** in browser DevTools
- **Test all user interactions** (clicks, forms, navigation)
- **Verify responsive design** on different screen sizes
- **Test with different roles** (admin, group_admin, editor, user)
- **Check i18n** - Test in all supported languages (fr, en, es, ar)
- **Run linter**: `npm run lint`
- **Build production bundle**: `npm run build`

### Integration Testing

- **Test full user flows** (e.g., login → view news → add favorite → logout)
- **Test SSO authentication** if applicable
- **Verify Docker deployment** works correctly
- **Check database migrations** don't break existing data

---

## Reporting Bugs

Found a bug? Please help us fix it by providing detailed information.

### Before Reporting

1. **Check existing issues** - Your bug might already be reported
2. **Verify it's reproducible** - Can you consistently reproduce it?
3. **Test on latest version** - Update to the latest release

### Bug Report Template

```markdown
## Bug Description
A clear and concise description of the bug.

## Steps to Reproduce
1. Go to '...'
2. Click on '...'
3. Scroll down to '...'
4. See error

## Expected Behavior
What you expected to happen.

## Actual Behavior
What actually happened.

## Screenshots
If applicable, add screenshots to help explain the problem.

## Environment
- OS: [e.g., Ubuntu 22.04]
- Browser: [e.g., Chrome 120]
- Airboard Version: [e.g., 1.0.0]
- Deployment: [Docker / Coolify / Development]

## Additional Context
Any other relevant information (logs, error messages, etc.)

## Logs
```
Paste relevant logs here
```
```

---

## Suggesting Enhancements

We welcome feature requests and enhancement suggestions!

### Enhancement Template

```markdown
## Feature Description
A clear and concise description of the feature you'd like to see.

## Problem Statement
What problem does this feature solve? Why is it needed?

## Proposed Solution
How do you envision this feature working?

## Alternatives Considered
What alternative solutions have you considered?

## Additional Context
Screenshots, mockups, or examples from other applications.

## User Benefit
Who will benefit from this feature? How?
```

---

## Questions and Support

### Getting Help

- **Documentation**: Check [README.md](README.md), [HOW-TO-USE.md](HOW-TO-USE.md), and [CLAUDE.md](CLAUDE.md)
- **Discussions**: Use [GitHub Discussions](../../discussions) for questions
- **Issues**: For bug reports and feature requests

### Communication Channels

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and community support
- **Pull Requests**: Code contributions and reviews

---

## License

By contributing to Airboard, you agree that your contributions will be licensed under the same [MIT License](LICENSE) that covers this project.

---

## Thank You!

Thank you for taking the time to contribute to Airboard. Every contribution, no matter how small, helps make this project better for everyone.

**Happy coding!** 🚀

---

<div align="center">

[⬆ Back to top](#contributing-to-airboard)

</div>
