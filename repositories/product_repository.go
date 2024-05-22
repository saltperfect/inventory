package repositories

import (
	"gorm.io/gorm"
	"inventory/models"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) ListProducts(supplierID string, minPrice string, maxPrice string) ([]models.Product, error) {
	var products []models.Product
	query := r.DB
	if supplierID != "" {
		query = query.Where("supplier_id = ?", supplierID)
	}
	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}
	err := query.Find(&products).Error
	return products, err
}
