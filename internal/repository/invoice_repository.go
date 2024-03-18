package repository

import (
	"errors"
	"invoice-service/internal/db"
	"invoice-service/internal/entity"
	"invoice-service/internal/logger"
	"time"
)

type InvoiceRepository struct {
	EntityManager *db.EntityManager
	Logger        *logger.Logger
}

func NewInvoiceRepository(em *db.EntityManager, logger *logger.Logger) (*InvoiceRepository, error) {
	return &InvoiceRepository{EntityManager: em, Logger: logger}, nil
}

func (ir *InvoiceRepository) CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	newInvoice := &entity.Invoice{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := ir.EntityManager.DB.Exec(
		"INSERT INTO invoices (name, position, archived, created_at, updated_at, total, subtotal, description, vatpercentage) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		newInvoice.Name,
		newInvoice.Position,
		newInvoice.Archived,
		newInvoice.CreatedAt,
		newInvoice.UpdatedAt,
		newInvoice.Total,
		newInvoice.SubTotal,
		newInvoice.Description,
		newInvoice.VatPercentage,
	)
	if err != nil {
		return nil, errors.New("Failed to add new Task")

	}
	return &entity.Invoice{}, nil
}

func (ir *InvoiceRepository) GetAllInvoices() ([]entity.Invoice, error) {
	invoices := []entity.Invoice{}

	rows, err := ir.EntityManager.DB.Query("SELECT id FROM invoices;")

	defer rows.Close()

	if err != nil {
		ir.Logger.Error(err)
		return invoices, errors.New("Error getting all invoices")
	}

	for rows.Next() {
		var cInvoice entity.Invoice
		rows.Scan(&cInvoice.Id)
		invoices = append(invoices, cInvoice)
	}

	return invoices, nil
}

func (ir *InvoiceRepository) GetInvoiceByID(id int) (*entity.Invoice, error) {
	return nil, errors.New("method not implemented")
}

func (ir *InvoiceRepository) UpdateInvoice(i *entity.Invoice) (*entity.Invoice, error) {
	return nil, errors.New("method not implemented")
}

func (ir *InvoiceRepository) DeleteInvoice(id int) error {
	return errors.New("method not implemented")
}
