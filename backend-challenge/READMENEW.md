# 🍔 Food Ordering Backend API (Clean Architecture in Golang)

This is a **Golang-based Food Ordering System** implemented using the Gin web framework, following clean architecture principles. It supports order placement, product listing, promo code validation (with large `.gz` files), and comes with Swagger documentation, PostgreSQL integration, and environment-based configuration.

---

## 📦 Features

- Clean Architecture: `handlers`, `services`, `repositories`, `routes`, `config`, `middlewares`, `models`
- PostgreSQL with GORM and versioned migrations
- Swagger UI auto-generated from Go code annotations
- API Key authentication middleware
- CORS & Graceful Shutdown
- Promo Code Validation from large `.gz` files
- Structured logging middleware

---

## 📁 Project Structure

```bash
backend-challenge/
│
├── cmd/                 # CLI entry points (e.g., migration runner)
├── config/              # Configuration loader using .env
├── controllers/         # HTTP handlers
├── middlewares/         # Middleware (API key, logging, CORS)
├── models/              # GORM models
├── repositories/        # DB interactions
├── routes/              # Route bindings
├── services/            # Business logic
├── migrations/          # SQL migrations (up/down)
├── utility/             # 🔒 Large promo code files (NOT INCLUDED IN GIT)
├── docs/                # Swagger-generated docs
├── main.go              # App entry point
├── go.mod/go.sum
├── .env                 # Environment variables
```
🚀 Getting Started
✅ Prerequisites
Go 1.20+

PostgreSQL

golang-migrate CLI (for DB migrations)

swag CLI (for Swagger docs)
1 ⚙️ Setup Instructions
Clone the repository
```bash
git clone https://github.com/amarsinghrathour/oolio_backend_challenge.git
cd oolio_backend_challenge
```

2 Create your .env file
```bash
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=food_ordering
API_KEY=your-api-key
```

3. Run Migrations
```bash
go run cmd/main.go migrate
```
4. Generate Swagger Docs
````
 swag init -g ./cmd/main.go -o ./docs  
````
5. Run the App
```bash
go run cmd/main.go
```

6. Access the API
- Swagger UI: http://localhost:8080/swagger/index.html

- Example endpoint: GET /product

📂 Promo Code Files Notice

⚠️ Important:
The promo code files couponbase1.gz, couponbase2.gz, and couponbase3.gz are NOT included in the Git repository due to their large size (~800MB each).

💾 Where to Put Them?

Place the files in a utility/ directory at the root of the project:
```bash

backend-challenge/
└── utility/
├── couponbase1.gz
├── couponbase2.gz
└── couponbase3.gz

```
✅ Promo Code Validation
A promo code is valid if:

It's a string of 8 to 10 characters

Appears in at least two out of the three .gz files

Validation is performed when placing an order.

🔑 Authentication

All protected routes require an API key via header:

```bash

-H "api_key: your-api-key"
```
🧪 Sample API Request

```bash

curl -X POST http://localhost:8080/order \
-H "Content-Type: application/json" \
-H "api_key: your-api-key" \
-d '{
  "couponCode": "HAPPYHRS",
  "items": [
    {
      "productId": "10",
      "quantity": 2
    }
  ]
}'

```

📜 License

MIT License

🙋‍♂️ Author

Made with ❤️ by Amar


