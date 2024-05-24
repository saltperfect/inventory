# Inventory Management System

This is an Inventory Management System built with Go, utilizing the Chi router and GORM for ORM. The application provides API endpoints to manage products and suppliers, including creating, updating, deleting, and viewing products and suppliers.

## Prerequisites

Before running this application, ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Running the Application

Follow these steps to run the application using Docker Compose:

### Step 1: Clone the Repository

```sh
git clone https://github.com/saltperfect/inventory.git
cd inventory
```

### Step 2: Run the command

```sh
docker-compose up --build
```

### Step 3: Use the postman collection given to run the apis


# Inventory Management System

This document explains the design decisions made for the Inventory Management System, including the architecture, API design, and technology choices.

## Table of Contents
1. [Architecture](#architecture)
2. [API Design](#api-design)
3. [Technology Stack](#technology-stack)
4. [Database Design](#database-design)
5. [Routing](#routing)
6. [Service Layer](#service-layer)
7. [Repository Layer](#repository-layer)
8. [Docker and Deployment](#docker-and-deployment)
9. [Running the Application](#running-the-application)
10. [Conclusion](#conclusion)

## Architecture

The application follows a layered architecture consisting of:

- **Endpoint Layer**: Handles HTTP requests and responses.
- **Service Layer**: Contains business logic.
- **Repository Layer**: Interacts with the database.

This separation of concerns makes the application easier to maintain, test, and extend.

## API Design

The API design follows REST principles, ensuring each endpoint corresponds to a specific resource or action. Key endpoints include:

- **Product Management**:
  - `POST /product`: Create a new product.
  - `PUT /product/{id}`: Update an existing product.
  - `DELETE /product/{id}`: Delete a product.
  - `GET /product/{id}`: View details of a single product.
  - `GET /products`: List all products with optional filtering by supplier or price range.
  - `PATCH /product/{id}/stock`: Adjust stock levels for a product.

- **Supplier Management**:
  - `POST /supplier`: Create a new supplier.
  - `PUT /supplier/{id}`: Update an existing supplier.
  - `DELETE /supplier/{id}`: Delete a supplier.
  - `GET /supplier/{id}`: View details of a single supplier.

## Technology Stack

- **Programming Language**: Go (Golang)
- **Router**: Chi for handling HTTP routing.
- **ORM**: GORM for database interactions.
- **Database**: SQLite (for simplicity and ease of setup)
- **Docker**: For containerizing the application and managing dependencies.

## Database Design

The database consists of two main tables: `Product` and `Supplier`. Each product is associated with a supplier via a foreign key.

### Product Table

- `ID` (primary key)
- `Name`
- `SupplierID` (foreign key)
- `Price`
- `StockQuantity`
- `Images` (optional, list of image URLs)

### Supplier Table

- `ID` (primary key)
- `Name`
- `ContactInformation`

## Routing

The application uses the Chi router to define routes and extract URL parameters. Chi is chosen for its lightweight and idiomatic approach to routing in Go.

Example route definition:

```go
r.Get("/supplier/{id}", supplierHandler.ViewSupplier)
```

## Service Layer
The service layer contains the business logic for managing products and suppliers. This layer interacts with the repository layer to perform CRUD operations.

## Example service function:

```go
func (s *ProductService) CreateProduct(product *models.Product) error {
    return s.Repo.CreateProduct(product)
}
```

## Repository Layer
The repository layer abstracts the database operations, making it easier to switch database implementations if needed. It uses GORM for ORM.

## Example repository function:

```go
func (r *ProductRepository) CreateProduct(product *models.Product) error {
    return r.DB.Create(product).Error
}
```

