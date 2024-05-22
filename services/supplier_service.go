package services

import (
	"inventory/models"
	"inventory/repositories"
)

type SupplierService struct {
	Repo *repositories.SupplierRepository
}

func (s *SupplierService) CreateSupplier(supplier *models.Supplier) error {
	return s.Repo.CreateSupplier(supplier)
}

func (s *SupplierService) UpdateSupplier(supplier *models.Supplier) error {
	return s.Repo.UpdateSupplier(supplier)
}

func (s *SupplierService) DeleteSupplier(id uint) error {
	return s.Repo.DeleteSupplier(id)
}

func (s *SupplierService) GetSupplierByID(id uint) (models.Supplier, error) {
	return s.Repo.GetSupplierByID(id)
}
