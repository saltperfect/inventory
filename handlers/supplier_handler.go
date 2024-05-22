package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory/models"
	"inventory/services"

	"github.com/go-chi/chi/v5"
)

type SupplierHandler struct {
	Service *services.SupplierService
}

func (h *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateSupplier(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

func (h *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	supplier.ID = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateSupplier(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(supplier)
}

func (h *SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err := h.Service.DeleteSupplier(uint(id)); err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SupplierHandler) ViewSupplier(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	supplier, err := h.Service.GetSupplierByID(uint(id))
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(supplier)
}
