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
├── config/                     # Configuration logic (e.g., loading .ini files)
│
├── config-file.ini            # Main config file
├── config-file.example.ini    # Example config for reference
│
├── docker-compose.yml         # Docker Compose configuration
│
├── .git/                      # Git metadata
├── .gitignore                 # Files and folders to ignore in Git
│
├── go.mod                     # Go modules (project dependencies)
├── go.sum                     # Dependency checksums
│
├── handler/                   # Handler functions for each domain
│   ├── user.go                # User-related request handlers
│   └── product.go             # Product-related request handlers
│
├── main.go                    # Application entry point
│
├── router/                    # Routing setup
│   ├── api.go                 # API groupings and subroutes
│   └── routes.go              # Base router and middleware config
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
);

select * from product p

insert into product (name, price) values ('Product 1', 100)
```
