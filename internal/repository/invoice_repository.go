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
	now, err := time.Now().MarshalText()

	if err != nil {
		return nil, errors.New("Failed to add new Task")

	}

	_, error := ir.EntityManager.DB.Exec(
		"INSERT INTO invoices (name, position, archived, created_at, updated_at, total, subtotal, description, vatpercentage) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		invoice.Name,
		invoice.Position,
		invoice.Archived,
		now,
		now,
		invoice.Total,
		invoice.SubTotal,
		invoice.Description,
		invoice.VatPercentage,
	)

	if error != nil {
		return nil, errors.New("Failed to add new Task")

	}
	return invoice, nil
}

func (ir *InvoiceRepository) GetAllInvoices() ([]entity.Invoice, error) {
	invoices := []entity.Invoice{}

	rows, err := ir.EntityManager.DB.Query("SELECT id, name, position, archived, created_at, updated_at, total, subtotal, description, vatpercentage FROM invoices;")

	defer rows.Close()

	if err != nil {
		ir.Logger.Error(err)
		return invoices, errors.New("Error getting all invoices")
	}

	for rows.Next() {
		var cInvoice entity.Invoice
		rows.Scan(
			&cInvoice.Id,
			&cInvoice.Name,
			&cInvoice.Position,
			&cInvoice.Archived,
			&cInvoice.CreatedAt,
			&cInvoice.UpdatedAt,
			&cInvoice.Total,
			&cInvoice.SubTotal,
			&cInvoice.Description,
			&cInvoice.VatPercentage,
		)
		invoices = append(invoices, cInvoice)
	}

	return invoices, nil
}

func (ir *InvoiceRepository) GetInvoiceByID(id int64) (*entity.Invoice, error) {
	var invoice entity.Invoice

	err := ir.EntityManager.DB.QueryRow("SELECT id, name, position, archived, created_at, updated_at, subtotal, total, vatpercentage, description FROM invoices WHERE id = ($1);", id).Scan(
		&invoice.Id,
		&invoice.Name,
		&invoice.Position,
		&invoice.Archived,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
		&invoice.SubTotal,
		&invoice.Total,
		&invoice.VatPercentage,
		&invoice.Description,
	)

	if err != nil {
		return nil, errors.New("Not found")

	}

	return &invoice, nil
}

func (ir *InvoiceRepository) UpdateInvoice(i *entity.Invoice) (*entity.Invoice, error) {
	_, err := ir.EntityManager.DB.Exec(
		"UPDATE invoices SET name = ($1), position = ($2), archived = ($3), total = ($4), subtotal = ($5), description = ($6), vatpercentage = ($7) , updated_at = ($8) WHERE id = ($9);",
		i.Name,
		i.Position,
		i.Archived,
		i.Total,
		i.SubTotal,
		i.Description,
		i.VatPercentage,
		time.Now(),
		i.Id,
	)

	if err != nil {
		return nil, errors.New("Failed to update invoice")
	}

	return ir.GetInvoiceByID(int64(i.Id))
}

func (ir *InvoiceRepository) DeleteInvoice(id int64) error {
	_, err := ir.EntityManager.DB.Exec("DELETE FROM invoices WHERE id = ($1);", id)

	if err != nil {
		return errors.New("Failed to delete invoice")
	}

	return nil
}
