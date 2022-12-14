package services_test

import (
	"legal-research-automation-go/pkg/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPDF(t *testing.T) {
	filename := "Test PDF File.pdf"
	pdfcontent := "Hello, World!"
	new_pdf := services.ExtractPDF{}
	new_pdf.New(filename)
	new_pdf.Open()
	assert.Equal(t, filename, new_pdf.FileName)
	assert.Equal(t, pdfcontent, new_pdf.Text)
}
