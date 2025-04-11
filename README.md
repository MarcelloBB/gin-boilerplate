# 🧪 Gin boilerplate

Go API project using [Gin](https://github.com/gin-gonic/gin) as the main framework.

🧱 Clean Architecture – Clear separation of layers: model, usecase, repository, controller, and router.

🐘 PostgreSQL – Relational database running via Docker.

🧠 Redis – Data caching with expiration support and performance improvement.

📦 Makefile – Simplified commands for build, run, and test.

📚 Swagger – Automatic API routes documentation.

---

## 🚀 Prerequisites

Before starting, make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🏗️ Project Structure
```text
gin-boilerplate
├── cmd
│   └── main.go                   # Entry point of the application
├── config
│   └── config.go                 # Loads and parses the .ini configuration file
├── config-file.example.ini       # Example configuration file
├── config-file.ini               # Main configuration file used by the app
├── controller
│   ├── product_controller.go     # Handles product-related HTTP requests
│   └── user_controller.go        # Handles user-related HTTP requests
├── db
│   ├── conn.go                   # PostgreSQL database connection logic
│   └── redis.go                  # Redis client setup and cache utility functions
├── docker-compose.yml            # Docker Compose configuration for PostgreSQL and Redis
├── docs
│   ├── docs.go                   # Auto-generated Swagger documentation
│   ├── swagger.json              # Swagger JSON output
│   └── swagger.yaml              # Swagger YAML output
├── go.mod                        # Go module definition
├── go.sum                        # Go module dependencies checksums
├── makefile                      # Simplified CLI commands for build, run, and docs
├── model
│   ├── product.go                # Product data structure
│   └── user.go                   # User data structure
├── README.md                     # Project documentation
├── repository
│   ├── product_repository.go     # Data access layer for product entity
│   └── user_repository.go        # Data access layer for user entity
├── router
│   ├── api.go                    # Initializes route groups
│   └── routes.go                 # Defines all available routes
└── usecase
    ├── product_usecase.go        # Business logic for product operations
    └── user_usecase.go           # Business logic for user operations

```
---

## 🛠️ Running

### 1. Clone the project

```bash
git clone https://github.com/MarcelloBB/gin-boilerplate.git
cd gin-boilerplate
```

### 2. 📦 Run with Go
```bash
go mod tidy
go run main.go
```

### 3. 🛠️ Running with Makefile
To simplify common development tasks, you can use the provided Makefile:

```bash
# Run the application with Swagger docs generation
make run

# Build the application binary
make build

# Generate Swagger documentation only
make docs

# Remove the generated binary and Swagger docs
make clean
```
ℹ️ These commands assume you have swag (from github.com/swaggo/swag/cmd/swag) installed.

The application should be running at: http://localhost:8080

## ⚙️ Configuration File
The application uses a config file named config-file.ini. Base example:

```bash
[server]
port = 8081

[db]
host     = 192.168.3.4
port     = 5432
user     = postgres
password = 1234
database = postgres
name     = postgres

[redis]
host       = localhost:6379
db         = 0
password   = 
expiration = 10

```

## 🐳 Running with Docker Compose
Currently, the docker-compose.yml starts PostgreSQL and Redis container.


### 1. Start the services
```bash
docker compose up
```
### 2. PostgreSQL access
Configure the database by inserting your credentials into compose:
- Host
- Port
- User
- Password
- Database

### 3. Redis access
Configure the database by inserting your credentials into compose:
- Host
- Db
- Password
- Expiration (in minutes)

### 4. Setting up a demo database
If you want to test the model and API, run the script:
```sql
create table product (
	id SERIAL primary key,
	name varchar(2000),
	price numeric(10, 2)
	quantity int
);

select * from product p

insert into product (name, price) values ('Product 1', 100, 10)
```
