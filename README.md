# Go E-commerce Server

A simple RESTful API server built with Go for managing products.

## Features
- Get all products
- Create new products
- CORS enabled
- In-memory storage

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| POST | `/create-Product` | Create a new product |
| GET | `/hello` | Health check |
| GET | `/about` | About info |

## Running the Server

```bash
go run main.go
```

Server runs on `http://localhost:8080`

## Example Usage

**Get Products:**
```bash
curl http://localhost:8080/products
```

**Create Product:**
```bash
curl -X POST http://localhost:8080/create-Product \
  -H "Content-Type: application/json" \
  -d '{"title":"Apple","description":"Red","price":150,"ImageUrl":"url"}'
```

## Tech Stack
- Go (Golang)
- net/http package
- JSON encoding/decoding
