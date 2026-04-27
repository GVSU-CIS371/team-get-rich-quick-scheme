package database

import (
	"errors"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Note           string
	Items          []InvoiceItem
	OrganizationID uint          `json:"-"`
	Organization   *Organization `json:"-"`
}

type InvoiceItem struct {
	gorm.Model
	Description string
	Quantity    uint
	Price       uint
	InvoiceID   uint     `json:"-"`
	Invoice     *Invoice `json:"-"`
}

func (db *Database) CreateInvoice(org *Organization, note string) *Invoice {
	inv := &Invoice{
		Note:           note,
		OrganizationID: org.ID,
	}

	db.client.Create(inv)

	return inv
}

func (db *Database) GetInvoices(org *Organization) []Invoice {
	var invoices []Invoice
	db.client.Where(Invoice{OrganizationID: org.ID}).Find(&invoices)

	return invoices
}

func (db *Database) GetInvoice(org *Organization, id uint) (*Invoice, error) {
	inv := new(Invoice)
	tx := db.client.Preload("Items").Where("id = ? AND organization_id = ?", id, org.ID).First(inv)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("invoice not found")
	}

	return inv, nil
}

func (db *Database) CreateInvoiceItem(invoice *Invoice, description string, quantity, price uint) *InvoiceItem {
	item := &InvoiceItem{
		Description: description,
		Quantity:    quantity,
		Price:       price,
		InvoiceID:   invoice.ID,
	}

	db.client.Create(item)
	return item
}
