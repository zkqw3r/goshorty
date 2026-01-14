# 🚀 GoShorty — Modern URL Shortener

[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

<div align="center">
  <h3>⚡ Fast. Stylish. Simple.</h3>
  <img src="demo.gif" width="100%" alt="video example of work">
  <p><i>Minimalistic URL shortener written in Go</i></p>
</div>


## ✨ Features

- **🚀 High Performance** — built with Go and optimized `pgx` driver
- **🔒 Secure Validation** — strict URL checking (regex) to prevent junk data
- **🐳 Docker Ready** — effortless database deployment with containerization
- **⚡ Smart Caching** — prevents duplicates by checking existing URLs in DB



## 🛠️ Tech Stack

### Backend
- **Go 1.22+** — core application logic
- **pgx/v5** — high-performance, concurrent PostgreSQL driver
- **net/http** — standard library for robust HTTP server
- **godotenv** — secure environment configuration management

### Frontend
- **HTML5 & CSS3** — semantic layout with "Liquid Glass" design
- **Vanilla JS** — asynchronous form handling (`fetch` API) without bloat
- **CSS Animations** — smooth floating effects and validation feedback

### Infrastructure
- **PostgreSQL** — reliable storage for links and analytics
- **Docker** — isolated container environment for the database

## 🚀 Install and Run

### Prerequisites
- **Go** 1.22 or higher
- **Docker** (recommended for DB)
- **PostgreSQL** (if running locally without Docker)

### 1. Clone the repository
```bash
git clone https://github.com/zkqw3r/goshorty.git
cd goshorty
```

### 2. Setup Database (via Docker)
Start the PostgreSQL container:
```bash
docker run --name my-db -e POSTGRES_PASSWORD=mysecret -p 5432:5432 -d postgres
```

Create the required table:
```bash
docker exec -it my-db psql -U postgres -c "CREATE TABLE urls (id VARCHAR(10) PRIMARY KEY, original_url TEXT NOT NULL UNIQUE, clicks INT DEFAULT 0, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
```

### 3. Configure Environment
Create a `.env` file in the root directory:
```env
DATABASE_URL=postgres://postgres:mysecret@localhost:5432/postgres
BASE_URL=http://localhost:8080
PORT=:8080
```

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Run the Application
**Important:** Always run from the project root directory!

```bash
go run ./app
```

🎉 **The application is ready:** [http://localhost:8080](http://localhost:8080)

## 🔌 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/shorten` | Create a new short link (accepts JSON/Form) |
| `GET` | `/{id}` | Redirect to the original URL |
| `GET` | `/stats/{id}` | Get click statistics (JSON) |

---

## 📝 To-do List

- [x] Basic URL shortening logic
- [ ] Click analytics counter
- [x] Duplicate link prevention
- [x] UI/UX with Glassmorphism
- [ ] QR Code generation for links

---

<div align="center">
  <sub>Made with ❤️ by <b>zkqw3r</b></sub>
</div>
# goshorty
# goshorty
