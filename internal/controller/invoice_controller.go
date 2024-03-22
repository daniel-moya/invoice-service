package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"invoice-service/internal/entity"
	"invoice-service/internal/service"
	"net/http"
	"strconv"
)

// Handlers struct to hold handler functions
type InvoiceController struct {
	InvoiceService *service.InvoiceService
}

func NewInvoiceController(is *service.InvoiceService) (*InvoiceController, error) {
	return &InvoiceController{InvoiceService: is}, nil
}

// Define routes for CRUD operations
func (h *InvoiceController) SetupRoutes(router *gin.Engine) {
	router.GET("/invoices", h.GetAllInvoices)
	router.GET("/invoices/:id", h.GetInvoiceByID)
	router.POST("/invoices", h.CreateInvoice)
	router.PUT("/invoices", h.UpdateInvoice)
	router.DELETE("/invoices/:id", h.DeleteInvoice)
}

func bindInvoice(c *gin.Context) *entity.Invoice {
	var invoice entity.Invoice
	if err := c.BindJSON(&invoice); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON payload"})
		return nil
	}
	return &invoice
}

func (h *InvoiceController) CreateInvoice(c *gin.Context) {
	invoice, error := h.InvoiceService.CreateInvoice(bindInvoice(c))
	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.New("Create invoice failed"))
	}

	c.IndentedJSON(http.StatusOK, invoice)
}

func (h *InvoiceController) GetAllInvoices(c *gin.Context) {
	invoices, err := h.InvoiceService.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
	}

	c.IndentedJSON(http.StatusOK, invoices)
}

func (h *InvoiceController) GetInvoiceByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	invoice, err := h.InvoiceService.GetInvoiceByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.New("Getting invoice failed"))
	}

	c.IndentedJSON(http.StatusOK, invoice)
}

func (h *InvoiceController) UpdateInvoice(c *gin.Context) {
	invoice, error := h.InvoiceService.UpdateInvoice(bindInvoice(c))

	if error != nil {
		fmt.Println(error)
		c.IndentedJSON(http.StatusInternalServerError, errors.New("Failed to update invoice"))
	}

	c.IndentedJSON(http.StatusOK, invoice)
}

func (h *InvoiceController) DeleteInvoice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid param")
	}

	error := h.InvoiceService.DeleteInvoice(id)

	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Something went wrong trying do delete invoice")
	}

	c.IndentedJSON(http.StatusOK, "Invoice Deleted Successfully")
}
