package repositories

import (
	"gorm.io/gorm"
	"inventory/models"
)

type SupplierRepository struct {
	DB *gorm.DB
}

func (r *SupplierRepository) CreateSupplier(supplier *models.Supplier) error {
	return r.DB.Create(supplier).Error
}

func (r *SupplierRepository) UpdateSupplier(supplier *models.Supplier) error {
	return r.DB.Save(supplier).Error
}

func (r *SupplierRepository) DeleteSupplier(id uint) error {
	return r.DB.Delete(&models.Supplier{}, id).Error
}

func (r *SupplierRepository) GetSupplierByID(id uint) (models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.Where("id = ?", id).First(&supplier).Error
	return supplier, err
}
