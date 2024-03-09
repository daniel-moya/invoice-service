package repository

import (
	// "invoice-service/internal/db"
	"invoice-service/internal/entity"
)

type InvoiceRepository struct {
	// EntityManager *db.EntityManager
}

// func NewInvoiceRepository(em *db.EntityManager) (*InvoiceRepository, error) {
func NewInvoiceRepository() (*InvoiceRepository, error) {
	return &InvoiceRepository{}, nil
	// return &InvoiceRepository{EntityManager: em}, nil
}

func (ir *InvoiceRepository) CreateInvoice(i *entity.Invoice) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) GetAllInvoices() []entity.Invoice {
	return []entity.Invoice{{Total: 23}}
}

func (ir *InvoiceRepository) GetInvoiceByID(id int) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) UpdateInvoice(i *entity.Invoice) *entity.Invoice {
	return &entity.Invoice{}
}

func (ir *InvoiceRepository) DeleteInvoice(id int) {
}
