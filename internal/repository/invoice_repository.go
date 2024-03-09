package repository

import "invoice-service/internal/db"
import "invoice-service/internal/entity"

type InvoiceRepository struct {
	EntityManager *db.EntityManager
}

func NewInvoiceRepository(em *db.EntityManager) (*InvoiceRepository, error) {
	return &InvoiceRepository{EntityManager: em}, nil
}

func (ir *InvoiceRepository) CreateInvoice(i *entity.Invoice) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) GetAllInvoices() []entity.Invoice {
	invoices := []entity.Invoice{}

	rows, _ := ir.EntityManager.DB.Query("SELECT * FROM invoices;")

	defer rows.Close()

	for rows.Next() {
		var cInvoice entity.Invoice
		rows.Scan(&cInvoice.Id, &cInvoice.Name, &cInvoice.Archived, cInvoice.Position, &cInvoice.Total, &cInvoice.SubTotal)
		invoices = append(invoices, cInvoice)
	}

	return invoices
}

func (ir *InvoiceRepository) GetInvoiceByID(id int) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) UpdateInvoice(i *entity.Invoice) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) DeleteInvoice(id int) {
}
