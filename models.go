package main

type Customer struct {
	CustomerID      uint   `gorm:"primaryKey"`
	CustomerName    string `gorm:"size:255"`
	CustomerEmail   string `gorm:"size:255;unique"`
	CustomerAddress string `gorm:"size:255"`
	Region          string `gorm:"size:100"`
}

type Product struct {
	ProductID   uint    `gorm:"primaryKey"`
	ProductName string  `gorm:"size:255"`
	Category    string  `gorm:"size:100"`
	UnitPrice   float64 `gorm:"type:decimal(10,2)"`
}

type Order struct {
	OrderID       uint     `gorm:"primaryKey"`
	CustomerID    uint     `gorm:"primaryKey"`
	ProductID     uint     `gorm:"primaryKey"`
	OrderDate     string   `gorm:"type:date"`
	PaymentMethod string   `gorm:"size:100"`
	ShippingCost  float64  `gorm:"type:decimal(10,2)"`
	QuantitySold  int      `gorm:"not null"`
	Discount      float64  `gorm:"type:decimal(10,2)"`
	Customer      Customer `gorm:"foreignKey:customer_id;references:CustomerID"`
	Product       Product  `gorm:"foreignKey:product_id;references:ProductID"`
}
