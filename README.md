# Gratitude List

A simple full-stack web application for creating and managing gratitude lists.

## Tech Stack

- Go
- SQLite
- Vite
- JavaScript
- Podman
- Podman Compose
- Nginx

## Development

Backend:

```bash
cd backend
go run .
```

Frontend:

```bash
cd frontend
npm install
npm run dev
```

Backend: `localhost:8888`

Frontend: `localhost:5173`

## Containerized

Build and run:

```bash
podman-compose up --build
```

Application: `localhost:8080`

Stop:

```bash
podman-compose down
```
