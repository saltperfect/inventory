package services

import (
	"inventory/models"
	"inventory/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) ListProducts(supplierID string, minPrice string, maxPrice string) ([]models.Product, error) {
	return s.Repo.ListProducts(supplierID, minPrice, maxPrice)
}

func (s *ProductService) AdjustStock(id uint, amount int) (models.Product, error) {
	product, err := s.Repo.GetProductByID(id)
	if err != nil {
		return product, err
	}
	product.StockQty += amount
	err = s.Repo.UpdateProduct(&product)
	return product, err
}
