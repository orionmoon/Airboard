# 📘 How to Use Airboard - Getting Started Guide

This guide will help you get started with Airboard and make the most of its features.

---

## 📑 Table of Contents

1. [First Login](#1-first-login)
2. [Understanding the Dashboard](#2-understanding-the-dashboard)
3. [User Management (Admin)](#3-user-management-admin)
4. [Group Management (Admin)](#4-group-management-admin)
5. [Application Management (Admin)](#5-application-management-admin)
6. [Regular User Workflow](#6-regular-user-workflow)
7. [Common Tasks](#7-common-tasks)
8. [Best Practices](#8-best-practices)
9. [Troubleshooting](#9-troubleshooting)

---

## 1. First Login

### Default Credentials

After deploying Airboard, use these default accounts:

**Administrator Account:**
- **Email**: `admin@airboard.com`
- **Password**: `admin123`
- **Role**: Admin (full access)

**Regular User Account (for testing):**
- **Email**: `user@airboard.com`
- **Password**: `user123`
- **Role**: User (limited access)

### ⚠️ Important First Steps

1. **Change the admin password immediately**
   - Go to your profile (top-right corner)
   - Click "Change Password"
   - Use a strong password (16+ characters)

2. **Create your first real admin account**
   - Go to Users section
   - Create a new user with admin role
   - Use your real email address

3. **Delete or disable default accounts**
   - For security, disable the default `admin@airboard.com` account
   - Or change its credentials to something secure

---

## 2. Understanding the Dashboard

### Dashboard Layout

The Airboard dashboard is organized into several sections:

```
┌─────────────────────────────────────────────────────────┐
│  🚀 Airboard                    [Search] [Profile] [⚙️]  │
├─────────────────────────────────────────────────────────┤
│  Sidebar:                                               │
│  - Dashboard (home)                                     │
│  - Users (admin only)                                   │
│  - Groups (admin only)                                  │
│  - Applications (admin only)                            │
│  - Settings                                             │
│  - Logout                                               │
├─────────────────────────────────────────────────────────┤
│  Main Content:                                          │
│  ┌────────────┐ ┌────────────┐ ┌────────────┐         │
│  │ App Icon   │ │ App Icon   │ │ App Icon   │         │
│  │ App Name   │ │ App Name   │ │ App Name   │         │
│  └────────────┘ └────────────┘ └────────────┘         │
│                                                         │
│  Grouped by: [Group Name]                              │
└─────────────────────────────────────────────────────────┘
```

### Key Elements

- **Sidebar**: Navigation menu (left side)
- **Application Cards**: Clickable cards for each application
- **Groups**: Applications organized by group
- **Search Bar**: Quick search for applications
- **Profile Menu**: Access settings, profile, and logout

---

## 3. User Management (Admin)

### Creating a New User

1. **Navigate to Users**
   - Click "Users" in the sidebar
   - Click "Add User" button (top-right)

2. **Fill User Information**
   ```
   Username: john.doe
   Email: john.doe@company.com
   Password: [generate strong password]
   First Name: John
   Last Name: Doe
   Role: user (or admin for administrators)
   ```

3. **Assign to Groups**
   - Select one or more groups (e.g., "IT", "Sales", "Common")
   - Groups determine which applications the user can see

4. **Click "Create"**

### Editing a User

1. Go to Users section
2. Click the **Edit** icon (✏️) next to the user
3. Modify fields as needed
4. Click "Update"

### Deleting a User

1. Go to Users section
2. Click the **Delete** icon (🗑️) next to the user
3. Confirm deletion

### User Roles

| Role | Permissions |
|------|-------------|
| **user** | Can view assigned applications only |
| **admin** | Full access: manage users, groups, applications |

---

## 4. Group Management (Admin)

Groups are used to organize applications and control access.

### Creating a New Group

1. **Navigate to Groups**
   - Click "Groups" in the sidebar
   - Click "Add Group" button

2. **Fill Group Information**
   ```
   Name: IT Department
   Description: Technical tools and resources for IT staff
   ```

3. **Click "Create"**

### Example Group Structure

```
Common (all staff)
├── Company Intranet
├── Email Webmail
└── HR Portal

IT Department
├── GitLab
├── Monitoring Dashboard
├── Server Admin Panel
└── Documentation Wiki

Sales Department
├── CRM
├── Sales Dashboard
└── Product Catalog

HR Department
├── HR Management System
├── Payroll
└── Recruitment Platform
```

### Best Practices for Groups

- **Create a "Common" group** for applications accessible to everyone
- **Organize by department** (IT, HR, Sales, Finance, etc.)
- **Use descriptive names** (e.g., "IT - Development Team" instead of "Team1")
- **Keep it simple** - Don't create too many groups

---

## 5. Application Management (Admin)

### Adding a New Application

1. **Navigate to Applications**
   - Click "Applications" in the sidebar
   - Click "Add Application" button

2. **Fill Application Information**
   ```
   Name: GitLab
   Description: Code repository and CI/CD platform
   URL: https://gitlab.company.com
   Icon: mdi:gitlab (or URL to icon)
   Order: 1
   ```

3. **Assign to Groups**
   - Select which groups should see this application
   - Example: IT Department, DevOps Team

4. **Click "Create"**

### Finding Icons

Airboard uses [Iconify](https://icon-sets.iconify.design/) for icons.

**Popular Icon Sets:**
- `mdi:` - Material Design Icons (e.g., `mdi:gitlab`, `mdi:microsoft-office`)
- `fa:` - Font Awesome (e.g., `fa:slack`, `fa:jira`)
- `simple-icons:` - Simple Icons (e.g., `simple-icons:github`, `simple-icons:notion`)

**How to find an icon:**
1. Visit https://icon-sets.iconify.design/
2. Search for your application (e.g., "gitlab")
3. Copy the icon code (e.g., `mdi:gitlab`)
4. Paste it in the "Icon" field

### Example Applications

| Application | Icon Code | URL Example |
|-------------|-----------|-------------|
| GitLab | `mdi:gitlab` | https://gitlab.company.com |
| Microsoft 365 | `mdi:microsoft-office` | https://office.com |
| Slack | `fa:slack` | https://company.slack.com |
| Jira | `fa:jira` | https://company.atlassian.net |
| Notion | `simple-icons:notion` | https://notion.so |

### Ordering Applications

The "Order" field determines the display order on the dashboard:
- Lower numbers appear first (Order 1 → Order 2 → Order 3)
- Applications within the same group are sorted by order
- Use increments of 10 (10, 20, 30) to allow easy insertion later

---

## 6. Regular User Workflow

### Daily Usage

1. **Login**
   - Visit your Airboard URL
   - Enter your credentials (or use SSO if configured)

2. **View Your Dashboard**
   - See all applications assigned to your groups
   - Applications are organized by group

3. **Access an Application**
   - Click on an application card
   - Opens in a new tab

4. **Search for an Application**
   - Use the search bar at the top
   - Type the application name
   - Click the result

### Managing Your Profile

1. **Access Profile**
   - Click your profile icon (top-right)
   - Click "Profile"

2. **Change Password**
   - Click "Change Password"
   - Enter current password
   - Enter new password (twice)
   - Click "Update"

3. **Update Personal Information**
   - Modify first name, last name, email
   - Click "Update Profile"

---

## 7. Common Tasks

### Task 1: Onboard a New Employee

**Scenario**: A new employee "Jane Smith" joins the Sales department.

**Steps:**
1. Go to **Users** section
2. Click **Add User**
3. Fill information:
   - Username: `jane.smith`
   - Email: `jane.smith@company.com`
   - Password: [generate secure password]
   - First Name: `Jane`
   - Last Name: `Smith`
   - Role: `user`
4. Assign to groups: `Common`, `Sales Department`
5. Click **Create**
6. Send credentials to Jane (securely!)

**Result**: Jane can now login and see applications for Common + Sales groups.

---

### Task 2: Add a New Department Application

**Scenario**: IT department needs access to a new monitoring tool.

**Steps:**
1. Go to **Applications** section
2. Click **Add Application**
3. Fill information:
   - Name: `Grafana Monitoring`
   - Description: `Server and application monitoring dashboard`
   - URL: `https://grafana.company.com`
   - Icon: `simple-icons:grafana`
   - Order: `20`
4. Assign to groups: `IT Department`
5. Click **Create**

**Result**: All users in IT Department group now see the Grafana application on their dashboard.

---

### Task 3: Create a Project-Specific Dashboard

**Scenario**: A new project "Project Phoenix" needs its own set of tools.

**Steps:**
1. **Create the group:**
   - Go to **Groups** section
   - Click **Add Group**
   - Name: `Project Phoenix`
   - Description: `Tools for Project Phoenix team`
   - Click **Create**

2. **Add applications to the group:**
   - Go to **Applications** section
   - For each tool (Jira, Confluence, GitLab, etc.):
     - Edit the application
     - Add "Project Phoenix" to its groups
     - Click **Update**

3. **Assign users to the group:**
   - Go to **Users** section
   - Edit each project member
   - Add "Project Phoenix" to their groups
   - Click **Update**

**Result**: Project team members see their project tools grouped together.

---

### Task 4: Organize Applications by Department

**Recommended Structure:**

```
1. Common (for everyone)
   - Company Intranet
   - Email
   - HR Portal
   - IT Support

2. Executive
   - BI Dashboard
   - Financial Reports
   - Strategic Planning

3. IT Department
   - GitLab
   - Jenkins
   - Monitoring
   - Server Admin

4. Sales
   - CRM (Salesforce)
   - Sales Dashboard
   - Product Catalog

5. Marketing
   - Marketing Automation
   - Analytics
   - Social Media Management

6. HR
   - HRIS
   - Payroll
   - Recruitment Platform

7. Finance
   - Accounting Software
   - Expense Management
   - Invoicing
```

---

## 8. Best Practices

### For Administrators

1. **Security**
   - ✅ Change default passwords immediately
   - ✅ Use strong passwords (16+ characters)
   - ✅ Enable SSO if possible (Authentik, Microsoft 365)
   - ✅ Regularly review user access
   - ✅ Disable inactive user accounts

2. **Organization**
   - ✅ Use clear, descriptive group names
   - ✅ Create a "Common" group for universal applications
   - ✅ Keep group structure simple (5-10 groups max)
   - ✅ Use consistent naming conventions

3. **User Management**
   - ✅ Assign minimum necessary groups to each user
   - ✅ Document who has admin access
   - ✅ Offboard users promptly when they leave

4. **Application Management**
   - ✅ Use descriptive application names
   - ✅ Add helpful descriptions
   - ✅ Use consistent icon sources (Iconify)
   - ✅ Order applications logically (most-used first)

5. **Maintenance**
   - ✅ Regular backups of the database
   - ✅ Update applications URLs if they change
   - ✅ Remove unused applications
   - ✅ Monitor system logs for errors

### For Users

1. **Security**
   - ✅ Use a strong password
   - ✅ Don't share your credentials
   - ✅ Log out when using shared computers
   - ✅ Report suspicious activity to admin

2. **Usage**
   - ✅ Use the search bar for quick access
   - ✅ Bookmark frequently used applications
   - ✅ Report broken links to admin
   - ✅ Request access to applications you need

---

## 9. Troubleshooting

### Issue: Can't Login

**Symptoms**: Login fails with error message

**Solutions:**
1. Check username/email and password (case-sensitive)
2. Verify account is active (contact admin)
3. Check SSO configuration if using SSO
4. Clear browser cache and cookies
5. Try a different browser

---

### Issue: Applications Not Showing

**Symptoms**: Dashboard is empty or missing applications

**Solutions:**
1. Verify you're assigned to the correct groups
2. Check that applications are assigned to your groups
3. Refresh the page (Ctrl+F5)
4. Log out and log back in
5. Contact administrator to verify your group memberships

---

### Issue: Application Link Broken

**Symptoms**: Clicking an application shows error or wrong page

**Solutions:**
1. Report to administrator
2. Check if the target application is online
3. Verify the URL is correct (admin should update)

---

### Issue: Can't Access Admin Sections

**Symptoms**: "Users", "Groups", "Applications" menus not visible

**Solutions:**
1. Verify your role is "admin" (contact administrator)
2. Log out and log back in
3. Check with administrator if access was revoked

---

### Issue: SSO Not Working

**Symptoms**: Login page shows instead of automatic SSO login

**Solutions:**
1. Verify SSO is enabled (`SSO_ENABLED=true`)
2. Check Nginx Proxy Manager configuration
3. Verify Authentik headers are being forwarded
4. Check backend logs: `docker-compose logs backend`
5. See [README.md - Troubleshooting](README.md#-troubleshooting) for detailed SSO debugging

---

## 🎓 Learning Path

### Week 1: Setup & Configuration
- [ ] Deploy Airboard
- [ ] Change default passwords
- [ ] Create first admin account
- [ ] Create basic group structure

### Week 2: Content Population
- [ ] Add all departments as groups
- [ ] Add 10-20 main applications
- [ ] Configure application icons
- [ ] Test with a few pilot users

### Week 3: User Onboarding
- [ ] Create user accounts for all staff
- [ ] Assign users to appropriate groups
- [ ] Send welcome emails with login instructions
- [ ] Provide training session or guide

### Week 4: Optimization
- [ ] Gather user feedback
- [ ] Adjust group structure if needed
- [ ] Add missing applications
- [ ] Optimize application ordering

---

## 📞 Getting Help

- **Documentation**: [README.md](README.md)
- **API Reference**: [README.md - API Endpoints](README.md#-api-endpoints)
- **Issues**: [GitHub Issues](../../issues)
- **Community**: [GitHub Discussions](../../discussions)

---

## 🎯 Quick Reference

### Admin Checklist

- [ ] Change default admin password
- [ ] Create groups for each department
- [ ] Add applications with proper icons
- [ ] Create user accounts
- [ ] Assign users to groups
- [ ] Test with a regular user account
- [ ] Enable SSO (if available)
- [ ] Setup regular backups

### User Quick Start

1. Login with your credentials
2. Browse applications on dashboard
3. Click an application to open it
4. Use search bar for quick access
5. Update your profile as needed

---

<div align="center">

**Happy organizing with Airboard! 🚀**

[⬆ Back to top](#-how-to-use-airboard---getting-started-guide)

</div>
