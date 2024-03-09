package main

import "net/http"
import "github.com/gin-gonic/gin"

import "invoice-service/internal/config"
import "invoice-service/internal/db"
import "invoice-service/internal/repository"
import "invoice-service/internal/service"
import "invoice-service/internal/controller"

func main() {
	config := &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		Database: "invoices_db",
		Username: "postgres",
		Password: "dev",
	}

	entityManager, _ := db.NewEntityManager(config)
	invoiceRepository, _ := repository.NewInvoiceRepository(entityManager)
	invoiceService, _ := service.NewInvoiceService(invoiceRepository)
	invoiceController, _ := controller.NewInvoiceController(invoiceService)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, "Invoices Service") })
	invoiceController.SetupRoutes(router)

	router.Run(":8080")
}
