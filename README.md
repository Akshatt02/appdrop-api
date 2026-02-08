# AppDrop API

A professional REST API backend for AppDrop's no-code mobile app builder platform. This API manages the configuration and layout of pages and widgets for mobile applications, allowing merchants to build dynamic app layouts through drag-and-drop interfaces.

## Table of Contents

- [Overview](#overview)
- [Technology Stack](#technology-stack)
- [Setup Instructions](#setup-instructions)
- [Database Setup](#database-setup)
- [Running the Server](#running-the-server)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Project Structure](#project-structure)

---

## Overview

AppDrop API provides a complete REST interface for managing:

- **Pages**: Application screens with unique routes and home page designation
- **Widgets**: UI components placed on pages with flexible JSON configuration

The API enforces strict validation rules, maintains data integrity through transactions, and provides comprehensive error handling with consistent response formats.

### Key Features

- UUID-based resource identification
- JSONB configuration storage using PostgreSQL
- Atomic transactions for multi-step operations
- Cascade delete for data consistency
- Complete validation at handler, service, and repository layers
- Professional error responses with error codes
- Request/response logging middleware

---

## Technology Stack

- **Language**: Go 1.25
- **Database**: PostgreSQL (recommended: Neon PostgreSQL)
- **Driver**: pgx/v5 (high-performance PostgreSQL driver)
- **HTTP**: Go standard library net/http
- **Environment Management**: godotenv

---

## Setup Instructions

### Prerequisites

Ensure you have the following installed:

- Go 1.20 or higher
- PostgreSQL 13+ (or Neon PostgreSQL account)
- Git

### Step 1: Clone the Repository

```bash
git clone https://github.com/Akshatt02/appdrop-api/
cd appdrop-api
```

### Step 2: Verify Go Installation

```bash
go version
```

Expected output: `go version go1.25.X ...`

### Step 3: Create Environment File

Create a `.env` file in the project root:

```bash
# For Neon PostgreSQL
DATABASE_URL="postgresql://user:password@ep-xxxxx.region.postgres.vercel.sh/dbname?sslmode=require"
PORT=8080
```

### Step 4: Install Dependencies

```bash
go mod download
go mod tidy
```

Verify dependencies are installed:

```bash
go list -m all
```

---

## Database Setup

### Neon PostgreSQL (Recommended for Cloud)

1. Visit [Neon PostgreSQL](https://neon.com)
2. Sign up for free account
3. Create a new project
4. Copy the connection string from the dashboard
5. Update `.env`:

```env
DATABASE_URL="postgresql://user:password@ep-xxxxx.us-east-1.postgres.vercel.sh/dbname?sslmode=require"
PORT=8080
```

### Step 5: Run Database Migrations

Once database connection is configured, create tables:

```bash
# The migration file is at: migrations/schema.sql

# Using Neon Dashboard
# 1. Open Neon dashboard
# 2. Go to SQL Editor
# 3. Copy and execute migrations/schema.sql contents
```

### Running the Server

```bash
# Ensure .env file is in project root
go run main.go
```

Expected output:
```
✅ Connected to Postgres
Server running on :8080
```

### Verify Server is Running

```bash
# Test health endpoint
curl http://localhost:8080/health

# Response
API + DB working
```

---

## API Documentation

### Base URL

```
http://localhost:8080
```

### Endpoints Overview

#### Pages Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/pages` | List all pages |
| POST | `/pages` | Create a new page |
| GET | `/pages/:id` | Get page with widgets |
| PUT | `/pages/:id` | Update page |
| DELETE | `/pages/:id` | Delete page |

#### Widgets Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/pages/:id/widgets` | Create widget on page |
| PUT | `/widgets/:id` | Update widget |
| DELETE | `/widgets/:id` | Delete widget |
| POST | `/pages/:id/widgets/reorder` | Reorder page widgets |

### Example Requests

#### Create Page

```bash
curl -X POST http://localhost:8080/pages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Home",
    "route": "/home",
    "is_home": true
  }'
```

#### Create Widget

```bash
curl -X POST http://localhost:8080/pages/{pageId}/widgets \
  -H "Content-Type: application/json" \
  -d '{
    "type": "banner",
    "position": 0,
    "config": {
      "image_url": "https://example.com/banner.jpg",
      "title": "Welcome"
    }
  }'
```

### Error Response Format

All errors follow this format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message"
  }
}
```

### Validation Rules

- Page name is required and non-empty
- Page route is required, unique, and non-empty
- Only ONE page can have `is_home = true`
- Cannot delete the home page
- Widget type must be one of: `banner`, `product_grid`, `text`, `image`, `spacer`
- Widget config is optional but must be valid JSON
- Widgets reorder must include all widgets from that page

---

### Testing Guide

Complete testing documentation is available in `POSTMAN_TESTING_GUIDE.md`

This guide includes:
- 40+ test cases
- All success scenarios
- All error scenarios
- Integration test workflows
- Expected responses for each endpoint

### How to Proceed

1. Start server: `go run main.go`
2. Open `POSTMAN_TESTING_GUIDE.md`
3. Follow test cases in order
4. Copy/paste URLs and payloads into Postman or cURL
5. Verify responses match expected outputs

---

## Project Structure

```
appdrop-api/
├── main.go                          # Server entry point and routing
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── .env                             # Environment variables (local)
├── README.md                        # This file
├── POSTMAN_TESTING_GUIDE.md         # Complete API testing documentation
│
├── internal/
│   ├── db/
│   │   └── db.go                   # Database connection and initialization
│   │
│   ├── models/
│   │   ├── page.go                 # Page data structure
│   │   └── widget.go               # Widget data structure
│   │
│   ├── handlers/
│   │   ├── page_handler.go         # HTTP handlers for page endpoints
│   │   └── widget_handler.go       # HTTP handlers for widget endpoints
│   │
│   ├── services/
│   │   ├── page_service.go         # Page business logic and validation
│   │   └── widget_service.go       # Widget business logic and validation
│   │
│   ├── repository/
│   │   ├── page_repository.go      # Database operations for pages
│   │   └── widget_repository.go    # Database operations for widgets
│   │
│   ├── middleware/
│   │   └── logger.go               # HTTP request/response logging
│   │
│   └── utils/
│       ├── response.go             # Response formatting utilities
│       └── constants.go            # Application constants (widget types)
│
└── migrations/
    └── schema.sql                   # PostgreSQL schema definition
```

### Layer Descriptions

**Models**: Data structures that represent domain entities (Page, Widget)

**Handlers**: HTTP request/response layer - parses input, calls services, returns responses

**Services**: Business logic layer - validates data, enforces business rules, orchestrates operations

**Repository**: Data access layer - performs database operations using parameterized queries

**Middleware**: Cross-cutting concerns like logging

**Utils**: Shared utilities and constants

---

## Code Documentation

### Key Code Comments

Each file contains detailed comments explaining:

- Function purpose and parameters
- Complex business logic
- Database operations
- Error handling
- Validation rules

Example:
```go
// GetPageByID retrieves a page from the database by its UUID.
// Returns nil if the page is not found.
// Used to validate page existence before widget operations.
func GetPageByID(id string) (*models.Page, error) {
    ...
}
```

---

## Environment Variables Reference

```env
PORT=8080

# Example for Neon PostgreSQL
DATABASE_URL=postgresql://user:password@ep-xxxxx.region.postgres.vercel.sh/dbname?sslmode=require
```

---