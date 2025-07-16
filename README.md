# 🪵 Goosebumps

Goosebumps is a modern, high-performance [Go (Golang)](https://golang.org/) backend project designed for scalability, modularity, and developer experience.

## 📌 Features

- ✅ Clean project structure
- 📦 Go Modules support
- 📄 RESTful APIs (or gRPC, if applicable)
- 🛡️ Error handling and logging with `log`
- 🧪 Unit & integration testing (with `testing`)
- 🔧 Environment-based config management
- 🚀 Easy to deploy

## 🛠 Tech Stack

- **Language:** Go 1.22+
- **Dependency Management:** Go Modules
- **Logger:** `log` / `logrus` / `zap` (tùy bạn dùng gì)
- **Database:** (PostgreSQL, MongoDB, SQLite, etc. — thêm nếu có)
- **API Style:** REST / gRPC


> 🧠 Follows standard [Go project layout](https://github.com/golang-standards/project-layout).

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/goosebumps.git
cd goosebumps
```

### 2. Initialize dependencies

```bash
go mod tidy
```

### 3. Run the project

```bash
go run ./cmd/main.go
```

Or if your entry file is just `main.go`:

```bash
go run main.go
```

### 4. Run tests

```bash
go test ./...
```

## ⚙️ Configuration

Use `.env` or config files in `pkg/config/` (nếu bạn có).
Sử dụng `os.Getenv` hoặc một package như `github.com/joho/godotenv` để load biến môi trường.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

Made with ❤️ by [Trung Le](https://github.com/OptimusPrime2001)

