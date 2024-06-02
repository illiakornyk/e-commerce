# E-Commerce System API

## Overview

This project is a simple e-commerce system developed as part of a job application test task. It is built with Go's standard library and enhanced with several third-party libraries for database migration, authentication, configuration management, and more.

## Libraries Used

- **golang-migrate**: For applying database migrations.
- **golang-jwt**: For handling JSON Web Tokens.
- **godotenv**: For loading environment variables from `.env` files.
- **PostgreSQL DB driver**: For database interactions.
- **Cryptographic libraries**: For hashing user passwords.

## Makefile Commands

The project includes a `Makefile` to simplify common tasks:

- `make build`: Compiles the application to `bin/ecom`.
- `make test`: Runs all tests.
- `make run`: Builds and runs the application.
- `make migration`: Creates a new SQL migration file.
- `make migrate-up`: Applies migrations to the database.
- `make migrate-down`: Rolls back migrations.
- `make init-admin`: Creates an admin user in the database.

## Database Design

The database schema includes the following tables:

- **Users**: Admin records.
- **Customers**: Customer details.
- **Sellers**: Seller information.
- **Orders**: Order summaries (customer ID, total amount, status).
- **Products**: Product details.
- **Order Items**: Detailed order line items.

## API Routes

The server handles HTTP requests using custom routing built with Go's standard `net/http` package.

## Setup Instructions

1. Clone the repository.
2. Install Docker and Docker Compose.
3. Use `docker-compose up` to initialize the server and database.

## Usage

Interact with the API using the following routes:

### User Endpoints

- **Registration**: Restricted to prevent new user registrations.
- **Login**: Admin login to access and manage the API.

### Customer Endpoints

- **List All Customers**: `GET /customers`
- **Retrieve Customer by ID**: `GET /customers/:id`
- **Create New Customer**: `POST /customers`

### Seller Endpoints

- **List All Sellers**: `GET /sellers`
- **Retrieve Seller by ID**: `GET /sellers/:id`
- **Create New Seller**: `POST /sellers`

### Product Endpoints (Full CRUD)

- **Create**: `POST /products`
- **Read**: `GET /products/{id}`
- **Update**: `PUT /products/{id}`
- **Delete**: `DELETE /products/{id}`

### Checkout Endpoint

- **Checkout Cart**: `POST /cart/checkout`

### Admin Login

- `POST /admin/login`

## Testing with Postman

A Postman collection is included in this repository to facilitate testing of the API. To use it:

1. Download the collection from the repository.
2. Import it into Postman.
3. Use the `/login` endpoint with the correct admin credentials to obtain a JWT token.
4. Include the JWT token in the `Authorization` header for subsequent requests.

## Environment Setup

Create a `.env` file in the root directory of the project with the following structure, which is also demonstrated in the `.env.example` file.

Fill in the .env file with your own credentials and database connection information.

## Setup Instructions

1. Clone the repository.
2. Install Docker and Docker Compose.
3. Run `docker-compose build` to build the Docker images.
4. Use `docker-compose up` to initialize the server and database.
