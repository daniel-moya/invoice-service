package controller

import "github.com/gin-gonic/gin"
import "invoice-service/internal/service"
import "invoice-service/internal/entity"
import "net/http"

// Handlers struct to hold handler functions
type InvoiceController struct {
	InvoiceService *service.InvoiceService
}

func NewInvoiceController(is *service.InvoiceService) (*InvoiceController, error) {
	return &InvoiceController{InvoiceService: is}, nil
}

// Define routes for CRUD operations
func (h *InvoiceController) SetupRoutes(router *gin.Engine) {
	router.POST("/invoices", h.CreateInvoice)
	router.GET("/invoices", h.GetAllInvoices)
	router.GET("/invoices/:id", h.GetInvoiceByID)
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
	invoice := h.InvoiceService.CreateInvoice(bindInvoice(c))
	c.IndentedJSON(http.StatusOK, invoice)
}

func (h *InvoiceController) GetAllInvoices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.InvoiceService.GetAll())
}

func (h *InvoiceController) GetInvoiceByID(c *gin.Context) {
}

func (h *InvoiceController) UpdateInvoice(c *gin.Context) {
}

func (h *InvoiceController) DeleteInvoice(c *gin.Context) {
}
