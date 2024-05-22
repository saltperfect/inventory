/*

|-- main.go
|-- models/
|   |-- models.go
|-- repositories/
|   |-- product_repository.go
|   |-- supplier_repository.go
|-- services/
|   |-- product_service.go
|   |-- supplier_service.go
|-- handlers/
|   |-- product_handler.go
|   |-- supplier_handler.go

*/

package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"inventory/handlers"
	"inventory/models"
	"inventory/repositories"
	"inventory/services"
)

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Product{}, &models.Supplier{})

	// Set up repositories
	productRepo := &repositories.ProductRepository{DB: db}
	supplierRepo := &repositories.SupplierRepository{DB: db}

	// Set up services
	productService := &services.ProductService{Repo: productRepo}
	supplierService := &services.SupplierService{Repo: supplierRepo}

	// Set up handlers
	productHandler := &handlers.ProductHandler{Service: productService}
	supplierHandler := &handlers.SupplierHandler{Service: supplierService}

	// Create Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// Routes for products
	r.Post("/product", productHandler.CreateProduct)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/product/{id}", productHandler.ViewProduct)
	r.Patch("/product/{id}/stock", productHandler.AdjustStock)

	// Routes for suppliers (if needed)
	r.Post("/supplier", supplierHandler.CreateSupplier)
	r.Put("/supplier/{id}", supplierHandler.UpdateSupplier)
	r.Delete("/supplier/{id}", supplierHandler.DeleteSupplier)
	r.Get("/supplier/{id}", supplierHandler.ViewSupplier)

	// Start the server
	http.ListenAndServe(":8080", r)
}
