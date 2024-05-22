package models

type Product struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `json:"name"`
	SupplierID uint    `json:"supplier_id"`
	Price      float64 `json:"price"`
	StockQty   int     `json:"stock_qty"`
	Image      string  `json:"image"`
}

type Supplier struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Products []Product
}
