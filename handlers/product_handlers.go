package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory/models"
	"inventory/services"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	product.ID = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err := h.Service.DeleteProduct(uint(id)); err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	supplierID := r.URL.Query().Get("supplier_id")
	minPrice := r.URL.Query().Get("min_price")
	maxPrice := r.URL.Query().Get("max_price")
	products, err := h.Service.ListProducts(supplierID, minPrice, maxPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) ViewProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) AdjustStock(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	var stockAdjustment struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&stockAdjustment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.Service.AdjustStock(uint(id), stockAdjustment.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}
