package controller

import (
	"errors"
	"invoice-service/internal/entity"
	"invoice-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router.DELETE("/invoices", h.DeleteInvoice)
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
}

func (h *InvoiceController) UpdateInvoice(c *gin.Context) {
}

func (h *InvoiceController) DeleteInvoice(c *gin.Context) {
}
