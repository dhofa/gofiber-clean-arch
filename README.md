# GoFiber Clean Architecture

This project is a Go application that follows the Clean Architecture principles using the [GoFiber](https://gofiber.io/) web framework. It is structured to separate concerns and define clear boundaries between different components of the application.

## Getting Started

These instructions will help you set up and run the project locally.

### Prerequisites

- Go 1.24 or later
- PostgreSQL
- [GoFiber](https://gofiber.io/) library

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/gofiber-clean-arch.git
   cd gofiber-clean-arch
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   Ensure you have a PostgreSQL instance running and update the `.env` file with your database credentials:

   ```bash
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=clean_arch_db
   DB_HOST=localhost
   DB_PORT=5432
   ```

### Running the Application

To run the application, execute the following command from the root directory:

```bash
go run cmd/main.go

```

The application will start the server and listen on port `3000`. You can access the API at `http://localhost:3000/api/v1/`.

### API Endpoints

- **GET** `/api/v1/users`: Retrieve a list of users.
- **GET** `/api/v1/users/:id`: Retrieve a specific user by ID.
- **POST** `/api/v1/users`: Create a new user.
- **PUT** `/api/v1/users/:id`: Update an existing user.
- **DELETE** `/api/v1/users/:id`: Delete a user.

### CRUD Generator command
```bash
cd tools/generator
```

then running:

```bash
go run main.go modul_name
```
or
```bash
   go run main.go modulName
```
or
```bash
   go run main.go ModulName
```

### Project Structure

- `cmd/`: Contains the entry point for the application.
- `internal/`: Contains the core business logic and domain entities.
- `infrastructure/`: Contains framework and driver code such as database setup.
- `config/`: Contains configuration loading logic.
- `templates/`: Contains code generation templates.

### Articles
https://medium.com/@ahmadmundhofa/list/gofiber-clean-architecture-18c72e241fbb

### Contributing

Feel free to open issues or submit pull requests for any improvement suggestions.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
