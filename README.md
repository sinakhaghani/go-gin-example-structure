# Go Gin Example Structure

A minimal and clean boilerplate project using **Go** and **Gin** to help developers quickly start a scalable web application with best practices and organized structure.

This project includes:
- 🧱 Modular folder layout (controllers, models, routes, configs)
- 🔐 JWT authentication middleware
- 🌐 Multi-language support using `go-i18n`
- 🛢️ GORM ORM and MySQL integration
- ⚙️ `.env` configuration via `godotenv`

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-gin-example-structure.git
cd go-gin-example-structure
```


### 1. Clean Go module cache (Optional)

```bash
go clean -modcache
```

### 3. Download and install dependencies

```bash
go mod tidy
```
This will install all required packages and create the go.sum file.

### 5. Create a .env file

```bash
touch .env
```

Then add the following content:
```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=go_example_db
DB_CHARSET=utf8mb4

SECRET_KEY=your-secret-key

```

### 6. Migrate

```bash
migrate -path database/migrations -database "mysql://user:pass@tcp(localhost:3306)/dbname" up
```

### 7. Run the project

```bash
go run .
```

or with air

```bash
air
```
---


### CLI Commands:

create controller:

```bash
go run cli/forge.go make:controller user_controller
```
---

create model:

```bash
go run cli/forge.go make:model user_model
```
---

create migration:

```bash
go run cli/forge.go make:migration create_user_table
```
---

create middleware:

```bash
go run cli/forge.go make:middleware user_middleware
```
---

create validator:

```bash
go run cli/forge.go make:validator user_validator
```
---

### 🌍 Internationalization (i18n)
This project supports multi-language output using go-i18n.

Translation files are located in the translations/ directory.

Persian (fa.json) and English (en.json) are included by default.

You can add more languages by adding more JSON files and loading them in config/i18n.go.




