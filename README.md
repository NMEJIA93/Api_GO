# Api_GO

## 📄 Project Description

This project is an API developed in Go that allows managing users. It uses GORM for interaction with the MySQL database and Gorilla Mux for routing HTTP requests.

## 🛠️ Language and Technologies

- **Language**: Go
- **Database**: MySQL
- **ORM**: GORM
- **Routing**: Gorilla Mux
- **Environment Variables**: godotenv
- **Containers**: Docker

## 📥 Steps to Download

1. Clone the repository:
    ```sh
    git clone https://github.com/NMEJIA93/Api_GO.git
    ```
2. Navigate to the project directory:
    ```sh
    cd Api_GO
    ```

## 🚀 Steps to Run

1. Create a `.env` file in the root directory with the following content:
    ```env
    DATABASE_USER=root
    DATABASE_PASSWORD=root
    DATABASE_HOST=localhost
    DATABASE_PORT=3320
    DATABASE_NAME=go_course
    DATABASE_DEBUG=true
    DATABASE_AUTO_MIGRATE=true
    ```
2. Build and start the Docker containers:
    ```sh
    docker-compose up --build
    ```
3. Run the application:
    ```sh
    go run main.go
    ```

## 📡 Exposed Endpoints and Examples

- **GET /user/{id}**: Retrieves a user by ID.
    ```sh
    curl -X GET http://127.0.0.1:8000/user/{id}
    ```
- **GET /user**: Retrieves all users.
    ```sh
    curl -X GET http://127.0.0.1:8000/user
    ```
- **POST /user**: Creates a new user.
    ```sh
    curl -X POST http://127.0.0.1:8000/user -d '{"first_name":"John", "last_name":"Doe", "email":"john.doe@example.com", "phone":"1234567890"}'
    ```
- **PUT /user**: Updates an existing user.
    ```sh
    curl -X PUT http://127.0.0.1:8000/user -d '{"first_name":"John", "last_name":"Doe", "email":"john.doe@example.com", "phone":"1234567890"}'
    ```
- **PATCH /user/{id}**: Partially updates a user by ID.
    ```sh
    curl -X PATCH http://127.0.0.1:8000/user/{id} -d '{"first_name":"John"}'
    ```
- **DELETE /user/{id}**: Deletes a user by ID.
    ```sh
    curl -X DELETE http://127.0.0.1:8000/user/{id}
    ```

## 🏗️ Architecture

The project follows a layered architecture that includes the following main layers:

- **Controllers (Endpoints)**: Handle HTTP requests and call the corresponding services.
- **Services**: Contain business logic.
- **Repositories**: Interact with the database.

## 🗂️ Project Structure

```plaintext
Api_GO
├── .docker/
│   └── mysql/
│       ├── Dockerfile
│       └── init.sql
├── .idea/
│   ├── .gitignore
│   ├── Api_GO.iml
│   ├── dataSources.local.xml
│   ├── dataSources.xml
│   ├── material_theme_project_new.xml
│   ├── modules.xml
│   ├── vcs.xml
│   └── workspace.xml
├── docker/
│   └── mysql/
│       └── init.sql
├── pkg/
│   └── bootstrap/
│       └── bootstrap.go
├── src/
│   └── user/
│       ├── domain.go
│       ├── dto.go
│       ├── endpoint.go
│       ├── repository.go
│       └── service.go
├── .env
├── .gitignore
├── [docker-compose.yml](http://_vscodecontentref_/1)
├── [go.mod](http://_vscodecontentref_/2)
├── [go.sum](http://_vscodecontentref_/3)
├── LICENSE
├── [main.go](http://_vscodecontentref_/4)
└── [README.md](http://_vscodecontentref_/5)