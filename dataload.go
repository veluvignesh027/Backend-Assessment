package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type SalesData struct {
	OrderID         uint
	ProductID       uint
	CustomerID      uint
	ProductName     string
	Category        string
	Region          string
	DateOfSale      time.Time
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
}

func LoadData() error {
	db, err := GetDBInstance()
	if err != nil {
		log.Panic(err)
		return err
	}

	// Sample data
	salesDatabyt, err := os.ReadFile("data.csv")
	if err != nil {
		log.Println(err)
	}
	var salesData []SalesData
	err = json.Unmarshal(salesDatabyt, &salesData)
	if err != nil {
		log.Panic(err)
	}

	for _, data := range salesData {
		customer := Customer{
			CustomerID:      data.CustomerID,
			CustomerName:    data.CustomerName,
			CustomerEmail:   data.CustomerEmail,
			CustomerAddress: data.CustomerAddress,
			Region:          data.Region,
		}
		db.Where(Customer{CustomerID: data.CustomerID}).Assign(customer).FirstOrCreate(&customer)

		product := Product{
			ProductID:   data.ProductID,
			ProductName: data.ProductName,
			Category:    data.Category,
			UnitPrice:   data.UnitPrice,
		}
		db.Where(Product{ProductID: data.ProductID}).Assign(product).FirstOrCreate(&product)

		order := Order{
			OrderID:       data.OrderID,
			CustomerID:    customer.CustomerID,
			OrderDate:     data.DateOfSale.Format("2006-01-02"),
			PaymentMethod: data.PaymentMethod,
			ShippingCost:  data.ShippingCost,
			QuantitySold:  data.QuantitySold,
			Discount:      data.Discount,
		}
		db.Create(&order)
	}

	fmt.Println("Data loaded successfully!")
	return nil
}
