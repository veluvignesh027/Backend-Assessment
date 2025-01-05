package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("HTTP Server Listening on: ", server.Addr)
	return server
}

func RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/resync", resyncHandler)
	r.GET("/customer/analysis/:startdate/:enddate/:query", CustomerAnalysisHandler)
	return r
}

func resyncHandler(ctx *gin.Context) {
	log.Println("Request to Load the data into the Database")
	err := LoadData()
	if err != nil {
		log.Println("Load Data failed! Error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}
