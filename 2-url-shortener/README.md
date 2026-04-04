# URL Shortener

A simple URL shortening service built with Go, Gin, and PostgreSQL.

## Features

- Shorten long URLs
- Redirect from short links to original URLs
- Duplicate detection for same URLs

## Setup

### 1. Spin Up PostgreSQL Instance

### 2. Create Table
```bash
docker exec -it url_shortener_postgres psql -U postgres -d postgres -c "
CREATE TABLE IF NOT EXISTS url (
  hash VARCHAR(7) UNIQUE PRIMARY KEY,
  long_url TEXT NOT NULL
);
"
```

### 3. Set Environment
Create `.env`:
```env
DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

### 4. Run
```bash
go run ./cmd/main.go
```

## API

**Create Short URL**
```bash
POST /api/v1/urls
{ "url": "https://example.com/long/url" }
```

**Get Original URL**
```bash
GET /api/v1/urls/:hash
```

**Redirect**
```bash
GET /:hash
```

