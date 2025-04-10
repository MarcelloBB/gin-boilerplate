# 🧪 Gin Boilerplate

This is a simple Go project using [Gin](https://github.com/gin-gonic/gin) as the web framework. The project includes a basic folder structure, route organization, handlers, and Docker support with a PostgreSQL container.

---

## 🚀 Prerequisites

Before starting, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (recommended: 1.20+)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🏗️ Project Structure
```text
gin-boilerplate/
├── cmd/
│   └── main.go                     # Application entry point
│
├── config/
│   └── config.go                   # Loads and parses the .ini config file
│
├── config-file.ini                # Main config file (e.g., port, env)
├── config-file.example.ini        # Example config file for reference
│
├── controller/
│   ├── product_controller.go      # HTTP handlers for product routes
│   └── user_controller.go         # HTTP handlers for user routes
│
├── db/
│   └── conn.go                    # Database connection and initialization
│
├── docker-compose.yml             # Docker services setup (e.g., PostgreSQL)
│
├── go.mod                         # Go module definition
├── go.sum                         # Checksums for module dependencies
│
├── model/
│   ├── product.go                 # Product domain model definition
│   └── user.go                    # User domain model definition
│
├── README.md                      # Project documentation
│
├── repository/
│   └── product_repository.go      # Repository layer for product persistence
│
├── router/
│   ├── router.go                  # Gin engine setup with middleware
│   └── routes.go                  # Route groupings and route registration
│
└── usecase/
    ├── product_usecase.go         # Business logic for products
    └── user_usecase.go            # Business logic for users

```
---

## 🛠️ Running Locally (without Docker)

### 1. Clone the project

```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

### 2. Run
```bash
go mod tidy
go run main.go
```

The application should be running at: http://localhost:8080

## ⚙️ Configuration File
The application uses a config file named config-file.ini. Base example:

```bash
[server]
port=8080
```

## 🐳 Running with Docker Compose
Currently, the docker-compose.yml only starts a PostgreSQL container. You can extend it later to include the Go application container.

### 1. Start the services
```bash
docker-compose up -d go_db
```
### 2. PostgreSQL access
Configure the database by inserting your credentials into compose:
- Host
- Port
- User
- Password
- Database

### 3. Setting up a demo database
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
