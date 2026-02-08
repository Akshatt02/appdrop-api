# AppDrop API - Postman Testing Guide

## Base URL
```
http://localhost:8080
```

---

# 1. HEALTH CHECK

### Test 1.1: Health Check
```
GET http://localhost:8080/health
```

**Expected Response:** `200 OK`
```
API + DB working
```

---

# 2. PAGE ENDPOINTS

## 2.1 GET /pages - List All Pages

### Test 2.1.1: Get Pages (Empty Database)
```
GET http://localhost:8080/pages
```

**Expected Response:** `200 OK`
```json
[]
```

### Test 2.1.2: Get Pages (With Data)
```
GET http://localhost:8080/pages
```

**Expected Response:** `200 OK`
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Home",
    "route": "/home",
    "is_home": true,
    "created_at": "2025-02-07T10:30:00Z",
    "updated_at": "2025-02-07T10:30:00Z"
  },
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "Collection",
    "route": "/collection",
    "is_home": false,
    "created_at": "2025-02-07T10:31:00Z",
    "updated_at": "2025-02-07T10:31:00Z"
  }
]
```

---

## 2.2 POST /pages - Create Page

### Test 2.2.1: Create Page (Valid)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Home",
  "route": "/home",
  "is_home": true
}
```

**Expected Response:** `201 Created`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Home",
  "route": "/home",
  "is_home": true,
  "created_at": "2025-02-07T10:30:00Z",
  "updated_at": "2025-02-07T10:30:00Z"
}
```

### Test 2.2.2: Create Page (Missing Name)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```json
{
  "route": "/home",
  "is_home": true
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "name and route are required"
  }
}
```

### Test 2.2.3: Create Page (Missing Route)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Home",
  "is_home": true
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "name and route are required"
  }
}
```

### Test 2.2.4: Create Page (Duplicate Route)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Another Home",
  "route": "/home",
  "is_home": false
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "page route already exists"
  }
}
```

### Test 2.2.5: Create Page (Invalid JSON)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```
{invalid json}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "INVALID_JSON",
    "message": "Invalid request body"
  }
}
```

### Test 2.2.6: Create Second Home Page (Replaces First)
```
POST http://localhost:8080/pages
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "New Home",
  "route": "/newhome",
  "is_home": true
}
```

**Expected Response:** `201 Created` (Old home page is_home becomes false)
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "New Home",
  "route": "/newhome",
  "is_home": true,
  "created_at": "2025-02-07T10:35:00Z",
  "updated_at": "2025-02-07T10:35:00Z"
}
```

---

## 2.3 GET /pages/:id - Get Single Page with Widgets

### Test 2.3.1: Get Page with Widgets
```
GET http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000
```

**Expected Response:** `200 OK`
```json
{
  "page": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Home",
    "route": "/home",
    "is_home": true,
    "created_at": "2025-02-07T10:30:00Z",
    "updated_at": "2025-02-07T10:30:00Z"
  },
  "widgets": [
    {
      "id": "660e8400-e29b-41d4-a716-446655440000",
      "page_id": "550e8400-e29b-41d4-a716-446655440000",
      "type": "banner",
      "position": 0,
      "config": {
        "image_url": "https://example.com/banner.jpg",
        "title": "Welcome"
      },
      "created_at": "2025-02-07T10:30:30Z",
      "updated_at": "2025-02-07T10:30:30Z"
    }
  ]
}
```

### Test 2.3.2: Get Page Without Widgets
```
GET http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440002
```

**Expected Response:** `200 OK`
```json
{
  "page": {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "name": "Empty Page",
    "route": "/empty",
    "is_home": false,
    "created_at": "2025-02-07T10:40:00Z",
    "updated_at": "2025-02-07T10:40:00Z"
  },
  "widgets": []
}
```

### Test 2.3.3: Get Non-existent Page
```
GET http://localhost:8080/pages/00000000-0000-0000-0000-000000000000
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Page not found"
  }
}
```

---

## 2.4 PUT /pages/:id - Update Page

### Test 2.4.1: Update Page (Change Name)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Home Page Updated",
  "route": "/home",
  "is_home": true
}
```

**Expected Response:** `200 OK`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Home Page Updated",
  "route": "/home",
  "is_home": true,
  "created_at": "2025-02-07T10:30:00Z",
  "updated_at": "2025-02-07T10:45:00Z"
}
```

### Test 2.4.2: Update Page (Change Route)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440001
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Collection",
  "route": "/products",
  "is_home": false
}
```

**Expected Response:** `200 OK`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "Collection",
  "route": "/products",
  "is_home": false,
  "created_at": "2025-02-07T10:31:00Z",
  "updated_at": "2025-02-07T10:46:00Z"
}
```

### Test 2.4.3: Update Page (Duplicate Route for Different Page)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440001
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Collection",
  "route": "/home",
  "is_home": false
}
```

**Expected Response:** `409 Conflict`
```json
{
  "error": {
    "code": "CONFLICT",
    "message": "Page route already exists"
  }
}
```

### Test 2.4.4: Update Page (Keep Same Route)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440001
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Collection Updated",
  "route": "/products",
  "is_home": false
}
```

**Expected Response:** `200 OK` (Allowed because it's the same page)
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "Collection Updated",
  "route": "/products",
  "is_home": false,
  "created_at": "2025-02-07T10:31:00Z",
  "updated_at": "2025-02-07T10:47:00Z"
}
```

### Test 2.4.5: Update Page (Set is_home to true)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440001
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "New Home",
  "route": "/products",
  "is_home": true
}
```

**Expected Response:** `200 OK` (Old home page is_home becomes false)
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "New Home",
  "route": "/products",
  "is_home": true,
  "created_at": "2025-02-07T10:31:00Z",
  "updated_at": "2025-02-07T10:48:00Z"
}
```

### Test 2.4.6: Update Non-existent Page
```
PUT http://localhost:8080/pages/00000000-0000-0000-0000-000000000000
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Non-existent",
  "route": "/nonexistent",
  "is_home": false
}
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Page not found"
  }
}
```

### Test 2.4.7: Update Page (Missing Required Fields)
```
PUT http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "",
  "route": "/home",
  "is_home": true
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "name and route are required"
  }
}
```

---

## 2.5 DELETE /pages/:id - Delete Page

### Test 2.5.1: Delete Non-Home Page
```
DELETE http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440001
```

**Expected Response:** `200 OK`
```json
{
  "message": "Page deleted"
}
```

### Test 2.5.2: Delete Home Page (Should Fail)
```
DELETE http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000
```

**Expected Response:** `409 Conflict`
```json
{
  "error": {
    "code": "CONFLICT",
    "message": "Cannot delete home page"
  }
}
```

### Test 2.5.3: Delete Non-existent Page
```
DELETE http://localhost:8080/pages/00000000-0000-0000-0000-000000000000
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Page not found"
  }
}
```

### Test 2.5.4: Delete Page (Cascades Widgets)
**Setup:** Create a page with widgets, then delete the page
```
DELETE http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440003
```

**Expected Response:** `200 OK` (All widgets are deleted automatically)
```json
{
  "message": "Page deleted"
}
```

---

# 3️. WIDGET ENDPOINTS

## 3.1 POST /pages/:id/widgets - Create Widget

### Test 3.1.1: Create Widget (Banner with Config)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "banner",
  "position": 0,
  "config": {
    "image_url": "https://example.com/banner.jpg",
    "title": "Summer Sale",
    "description": "50% off everything"
  }
}
```

**Expected Response:** `201 Created`
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "page_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "banner",
  "position": 0,
  "config": {
    "image_url": "https://example.com/banner.jpg",
    "title": "Summer Sale",
    "description": "50% off everything"
  },
  "created_at": "2025-02-07T10:50:00Z",
  "updated_at": "2025-02-07T10:50:00Z"
}
```

### Test 3.1.2: Create Widget (Product Grid)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "product_grid",
  "position": 1,
  "config": {
    "columns": 3,
    "items_per_page": 12
  }
}
```

**Expected Response:** `201 Created`
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "page_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "product_grid",
  "position": 1,
  "config": {
    "columns": 3,
    "items_per_page": 12
  },
  "created_at": "2025-02-07T10:51:00Z",
  "updated_at": "2025-02-07T10:51:00Z"
}
```

### Test 3.1.3: Create Widget (Text)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "text",
  "position": 2,
  "config": {
    "content": "Welcome to our store!",
    "font_size": "18",
    "color": "#333333"
  }
}
```

**Expected Response:** `201 Created`

### Test 3.1.4: Create Widget (Image)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "image",
  "position": 3,
  "config": {
    "url": "https://example.com/image.jpg",
    "alt_text": "Product showcase",
    "width": "100%"
  }
}
```

**Expected Response:** `201 Created`

### Test 3.1.5: Create Widget (Spacer)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "spacer",
  "position": 4,
  "config": {
    "height": "20"
  }
}
```

**Expected Response:** `201 Created`

### Test 3.1.6: Create Widget (No Config - Allowed)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "spacer",
  "position": 5
}
```

**Expected Response:** `201 Created`

### Test 3.1.7: Create Widget (Invalid Type)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "invalid_widget",
  "position": 0,
  "config": {}
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "invalid widget type"
  }
}
```

### Test 3.1.8: Create Widget (Non-existent Page)
```
POST http://localhost:8080/pages/00000000-0000-0000-0000-000000000000/widgets
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "banner",
  "position": 0,
  "config": {}
}
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Page not found"
  }
}
```

### Test 3.1.9: Create Widget (Invalid JSON)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets
Content-Type: application/json
```

**Request Body:**
```
{invalid json}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "INVALID_JSON",
    "message": "Invalid request body"
  }
}
```

---

## 3.2 PUT /widgets/:id - Update Widget

### Test 3.2.1: Update Widget (Change Type)
```
PUT http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "image",
  "position": 0,
  "config": {
    "url": "https://example.com/new-banner.jpg"
  }
}
```

**Expected Response:** `200 OK`
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440000",
  "page_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "image",
  "position": 0,
  "config": {
    "url": "https://example.com/new-banner.jpg"
  },
  "created_at": "2025-02-07T10:50:00Z",
  "updated_at": "2025-02-07T11:00:00Z"
}
```

### Test 3.2.2: Update Widget (Change Config)
```
PUT http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440001
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "product_grid",
  "position": 1,
  "config": {
    "columns": 4,
    "items_per_page": 16
  }
}
```

**Expected Response:** `200 OK`

### Test 3.2.3: Update Widget (Change Position)
```
PUT http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "banner",
  "position": 5,
  "config": {
    "image_url": "https://example.com/banner.jpg",
    "title": "Summer Sale"
  }
}
```

**Expected Response:** `200 OK`

### Test 3.2.4: Update Widget (Invalid Type)
```
PUT http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "unknown",
  "position": 0,
  "config": {}
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "invalid widget type"
  }
}
```

### Test 3.2.5: Update Non-existent Widget
```
PUT http://localhost:8080/widgets/00000000-0000-0000-0000-000000000000
Content-Type: application/json
```

**Request Body:**
```json
{
  "type": "banner",
  "position": 0,
  "config": {}
}
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Widget not found"
  }
}
```

### Test 3.2.6: Update Widget (Invalid JSON)
```
PUT http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440000
Content-Type: application/json
```

**Request Body:**
```
{invalid}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "INVALID_JSON",
    "message": "Invalid request body"
  }
}
```

---

## 3.3 DELETE /widgets/:id - Delete Widget

### Test 3.3.1: Delete Widget
```
DELETE http://localhost:8080/widgets/660e8400-e29b-41d4-a716-446655440000
```

**Expected Response:** `200 OK`
```json
{
  "message": "Widget deleted"
}
```

### Test 3.3.2: Delete Non-existent Widget
```
DELETE http://localhost:8080/widgets/00000000-0000-0000-0000-000000000000
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Widget not found"
  }
}
```

---

## 3.4 POST /pages/:id/widgets/reorder - Reorder Widgets

### Test 3.4.1: Reorder Widgets
**Setup:** Create 3 widgets with positions 0, 1, 2
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": [
    "660e8400-e29b-41d4-a716-446655440002",
    "660e8400-e29b-41d4-a716-446655440000",
    "660e8400-e29b-41d4-a716-446655440001"
  ]
}
```

**Expected Response:** `200 OK` (Positions will be [2, 0, 1])
```json
{
  "message": "Widgets reordered"
}
```

**Verify with GET /pages/:id:**
```
GET http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000
```

Response should show widgets ordered by new positions:
```json
{
  "page": {...},
  "widgets": [
    {
      "id": "660e8400-e29b-41d4-a716-446655440002",
      "position": 0,
      ...
    },
    {
      "id": "660e8400-e29b-41d4-a716-446655440000",
      "position": 1,
      ...
    },
    {
      "id": "660e8400-e29b-41d4-a716-446655440001",
      "position": 2,
      ...
    }
  ]
}
```

### Test 3.4.2: Reorder Widgets (Non-existent Page)
```
POST http://localhost:8080/pages/00000000-0000-0000-0000-000000000000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": ["660e8400-e29b-41d4-a716-446655440000"]
}
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Page not found"
  }
}
```

### Test 3.4.3: Reorder Widgets (Non-existent Widget)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": [
    "660e8400-e29b-41d4-a716-446655440000",
    "00000000-0000-0000-0000-000000000000"
  ]
}
```

**Expected Response:** `404 Not Found`
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "One or more widgets not found"
  }
}
```

### Test 3.4.4: Reorder Widgets (Widget from Different Page)
**Setup:** Widget exists but belongs to different page
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": [
    "660e8400-e29b-41d4-a716-446655440000",
    "770e8400-e29b-41d4-a716-446655440000"
  ]
}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "widget does not belong to this page"
  }
}
```

### Test 3.4.5: Reorder Widgets (Invalid JSON)
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```
{invalid}
```

**Expected Response:** `400 Bad Request`
```json
{
  "error": {
    "code": "INVALID_JSON",
    "message": "Invalid request body"
  }
}
```

### Test 3.4.6: Reorder Single Widget
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": ["660e8400-e29b-41d4-a716-446655440000"]
}
```

**Expected Response:** `200 OK`
```json
{
  "message": "Widgets reordered"
}
```

### Test 3.4.7: Reorder Empty Widget List
```
POST http://localhost:8080/pages/550e8400-e29b-41d4-a716-446655440000/widgets/reorder
Content-Type: application/json
```

**Request Body:**
```json
{
  "widget_ids": []
}
```

**Expected Response:** `200 OK` (No widgets to reorder)
```json
{
  "message": "Widgets reordered"
}
```

---

# EXPECTED STATUS CODES SUMMARY

| Operation | Success | Validation Error | Not Found | Conflict |
|-----------|---------|------------------|-----------|----------|
| GET /pages | 200 | - | - | - |
| POST /pages | 201 | 400 | - | - |
| GET /pages/:id | 200 | - | 404 | - |
| PUT /pages/:id | 200 | 400 | 404 | 409 |
| DELETE /pages/:id | 200 | 400 | 404 | 409 |
| POST /pages/:id/widgets | 201 | 400 | 404 | - |
| PUT /widgets/:id | 200 | 400 | 404 | - |
| DELETE /widgets/:id | 200 | - | 404 | - |
| POST /pages/:id/widgets/reorder | 200 | 400 | 404 | - |
