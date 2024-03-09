package service

import "invoice-service/internal/entity"
import "invoice-service/internal/repository"

type InvoiceService struct {
	InvoiceRepository *repository.InvoiceRepository
}

func NewInvoiceService(ir *repository.InvoiceRepository) (*InvoiceService, error) {
	return &InvoiceService{InvoiceRepository: ir}, nil
}

func (ir *InvoiceService) CreateInvoice(invoice *entity.Invoice) *entity.Invoice {
	return ir.InvoiceRepository.CreateInvoice(invoice)
}

func (ir *InvoiceService) GetAll() []entity.Invoice {
	return ir.InvoiceRepository.GetAllInvoices()
}

func (ir *InvoiceService) GetInvoiceByID(invoiceId int) *entity.Invoice {
	return ir.InvoiceRepository.GetInvoiceByID(invoiceId)
}

func (ir *InvoiceService) UpdateInvoice(invoice *entity.Invoice) *entity.Invoice {
	return ir.InvoiceRepository.UpdateInvoice(invoice)
}

func (ir *InvoiceService) DeleteInvoice(invoiceId int) {
	ir.InvoiceRepository.DeleteInvoice(invoiceId)
}
