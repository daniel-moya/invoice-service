package main

import "github.com/gin-gonic/gin"
import "net/http"
import "invoice-service/internal/logger"
import "invoice-service/internal/config"
import "invoice-service/internal/controller"
import "invoice-service/internal/db"
import "invoice-service/internal/repository"
import "invoice-service/internal/service"
import "github.com/joho/godotenv"
import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := godotenv.Load(".env.dev")

	if err != nil {
		check(err)
	}

	dbConfig := &config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
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

	router.Run(":" + os.Getenv("IS_PORT"))
}
