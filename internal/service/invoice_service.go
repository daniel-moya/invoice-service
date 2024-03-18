package service

import "invoice-service/internal/entity"
import "invoice-service/internal/repository"

type InvoiceService struct {
	InvoiceRepository *repository.InvoiceRepository
}

func NewInvoiceService(ir *repository.InvoiceRepository) (*InvoiceService, error) {
	return &InvoiceService{InvoiceRepository: ir}, nil
}

func (ir *InvoiceService) CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	return ir.InvoiceRepository.CreateInvoice(invoice)
}

func (ir *InvoiceService) GetAll() ([]entity.Invoice, error) {
	return ir.InvoiceRepository.GetAllInvoices()
}

func (ir *InvoiceService) GetInvoiceByID(invoiceId int) (*entity.Invoice, error) {
	return ir.InvoiceRepository.GetInvoiceByID(invoiceId)
}

func (ir *InvoiceService) UpdateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	return ir.InvoiceRepository.UpdateInvoice(invoice)
}

func (ir *InvoiceService) DeleteInvoice(invoiceId int) error {
	return ir.InvoiceRepository.DeleteInvoice(invoiceId)
}
