package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvoices(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := time.Second
	client := textmagic.NewClient(username, token)

	time.Sleep(interval)
	// Get Invoice List

	invoices, _ := client.GetInvoiceList(map[string]string{})

	assert.NotEmpty(t, invoices.Page)
	assert.NotEmpty(t, invoices.Limit)
	assert.NotEmpty(t, invoices.PageCount)
	assert.NotEqual(t, 0, len(invoices.Resources))

	inv := invoices.Resources[0]
	assert.NotEmpty(t, inv.Id)
	assert.NotEmpty(t, inv.Bundle)
	assert.NotEmpty(t, inv.Currency)
	assert.NotEmpty(t, inv.PaymentMethod)
}
