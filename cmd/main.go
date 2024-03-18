package main

import "github.com/gin-gonic/gin"
import "net/http"
import "invoice-service/internal/logger"
import "invoice-service/internal/config"
import "invoice-service/internal/controller"
import "invoice-service/internal/db"
import "invoice-service/internal/repository"
import "invoice-service/internal/service"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dbConfig := &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		Database: "postgres",
		Username: "postgres",
		Password: "dev",
	}

	logger := logger.NewLogger()
	entityManager, emError := db.NewEntityManager(dbConfig)
	invoiceRepo, repoError := repository.NewInvoiceRepository(entityManager, logger)
	invoiceService, serviceError := service.NewInvoiceService(invoiceRepo)
	invoiceController, controllerError := controller.NewInvoiceController(invoiceService)

	check(emError)
	check(repoError)
	check(serviceError)
	check(controllerError)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, "Invoices Service") })
	invoiceController.SetupRoutes(router)

	router.Run(":8080")
}
