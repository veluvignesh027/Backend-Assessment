package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var totalCustomers int64
var totalOrders int64
var totalRevenue float64
var averageOrderValue float64

func CustomerAnalysisHandler(ctx *gin.Context) {
	db, err := GetDBInstance()
	if err != nil {
		log.Panic("Error getting the DB instance. Error: ", err)
	}

	startDate := ctx.Param("startdate")
	endDate := ctx.Param("enddate")

	query := ctx.Param("query")
	if query == "totalcustomers" {
		err = db.Model(&Customer{}).
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Distinct("customer_id").Count(&totalCustomers).Error
		if err != nil {
			log.Fatalf("failed to count customers: %v", err)
		}
		ctx.JSON(http.StatusOK, gin.H{"Total Number of Customers": totalCustomers})
	}
	if query == "totalorders" {
		err = db.Model(&Order{}).
			Where("order_date BETWEEN ? AND ?", startDate, endDate).
			Count(&totalOrders).Error
		if err != nil {
			log.Fatalf("failed to count orders: %v", err)
		}
		ctx.JSON(http.StatusOK, gin.H{"Total Number of Orders": totalOrders})
	}

	if query == "averageorders" {
		err = db.Table("order_details").
			Select("SUM(unit_price * quantity_sold * (1 - discount))").
			Joins("JOIN orders ON order_details.order_id = orders.order_id").
			Where("orders.order_date BETWEEN ? AND ?", startDate, endDate).
			Scan(&totalRevenue).Error
		if err != nil {
			log.Fatalf("failed to calculate total revenue: %v", err)
		}
		if totalOrders > 0 {
			averageOrderValue = totalRevenue / float64(totalOrders)
		}
		ctx.JSON(http.StatusOK, gin.H{"Average value of Orders": averageOrderValue})
	}

}
