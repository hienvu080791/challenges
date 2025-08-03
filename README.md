# Demo Challenges

A Go project demonstrating algorithms and backend modules using the Gin web framework.

## Features
Task 1:
- Algorithm utilities (see `source/algorithms/`)
Task 2:
- Modular backend structure with Gin (see `source/module/`)
- RESTful API with authentication and file upload
- JWT-based authentication
- Modular code structure

## Project Structure
- `main.go` — Entry point, starts the Gin server
- `source/router/` — API route definitions
- `source/algorithms/` — Algorithm implementations
- `source/module/` — Application modules:
  - `authentication/` — Auth logic and JWT
  - `file/` — File upload and handling
  - `middleware/` — Middleware services
  - `storage/` — Data storage logic, use local RAM storage for simplicity
  - `utils/` — Utility functions

## Requirements
- Go 1.18 or newer

## Sample CURLs:
1: register:
curl --location '0.0.0.0:8888/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
"username": "abcd",
"password":"123456"
}'

2: login:
curl --location '0.0.0.0:8888/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
"username": "abcd",
"password":"123456"
}'

3: upload file:
curl --location '0.0.0.0:8888/api/v1/upload' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY2QiLCJleHAiOjE3NTQyMDk5Njd9.SlfHqzyZ4eicFVVJ1Im9hJoRapu92GlhRjo38FlbwTM' \
--form 'data=@"/Users/hienvu080791/Downloads/icon.png"'