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
gin-boilerplate/
├── cmd/
│   └── main.go                    # Entry point of the application, initializes the server
│
├── config/
│   └── config.go                  # Loads and parses the .ini configuration file
│
├── config-file.ini               # Configuration file used by the app (e.g., port, DB)
├── config-file.example.ini       # Example configuration to help set up new environments
│
├── controller/
│   ├── product_controller.go     # Handles HTTP requests related to products
│   └── user_controller.go        # Handles HTTP requests related to users
│
├── db/
│   └── conn.go                   # Manages PostgreSQL database connection
│   └── redis.go                  # Manages Redis connection
│
├── docker-compose.yml            # Docker setup for services like PostgreSQL
│
├── go.mod                        # Declares the Go module and its dependencies
├── go.sum                        # Verifies integrity of Go module dependencies
│
├── model/
│   ├── product.go                # Structs and types related to products
│   └── user.go                   # Structs and types related to users
│
├── README.md                     # Project documentation and usage instructions
│
├── repository/
│   ├── product_repository.go     # Database operations for products
│   └── user_repository.go        # Database operations for users
│
├── router/
│   ├── api.go                    # Initializes and returns the Gin router
│   └── routes.go                 # Declares all route groups and endpoints
│
└── usecase/
    ├── product_usecase.go        # Business logic and orchestration for products
    └── user_usecase.go           # Business logic and orchestration for users

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
