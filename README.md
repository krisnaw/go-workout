# Workout App built in Go

This application is a simple **Workout App** built in **Go** to showcase my skills in utilizing the Go programming language. It demonstrates routing, database management, and health check functionality with a modular structure.

### Features
- RESTful API implementation.
- Modular structure for scalability and maintainability.
- Database integration and utility functions.
- Health check endpoint for monitoring uptime.
- Easy-to-understand codebase for learning and experimentation.

---

## Technologies Used

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Additional Tools/Packages:** Go-Chi

---


## Folder Structure

    .
    ├── internal                # Compiled files (alternatively `dist`)
    │       ├── api             # Route handler
    │       ├── app             # application's core components, including the database, logger, API handlers, and health check functionality.
    │       ├── routes          # Api End-points
    │       ├── store           # Database connections
    │       └── utils           # Utilities
    ├── migrations              # Migrations directory
    └── README.md
---

## Prerequisites

Make sure you have the following installed on your machine before running the application:

1. **Go** (version 1.23 or later)
2. (Add any other necessary installations, such as Docker/PostgreSQL if applicable.)
3. (Include details of third-party packages required by the app and how they can be installed.)

---

## Installation

1. Clone this repository to your local system:

   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Run the following command to download all dependencies:

   ```bash
   go mod tidy
   ```

3. Configure your database (if applicable) by editing the configuration file.

---

## Usage

1. Run the application:

   ```bash
   go run main.go
   ```

2. Access endpoints using tools like **curl**, **Postman**, or a web browser.

   For example:
    - `GET /health` - Returns app health status.

---

## API Endpoints

| Method | Endpoint        | Description            |
|--------|-----------------|------------------------|
| GET    | /health         | Health check endpoint |
| POST   | /workouts       | Create a workout      |
| GET    | /workouts       | List all workouts     |

(Add more as needed.)

---

## Database Migrations

If your application has migrations, include instructions for handling them:

1. To apply database migrations, run:
   ```bash
   (Add command or tool, e.g., migrate up, sqlc, etc.)
   ```

---

