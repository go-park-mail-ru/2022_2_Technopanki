package utils

import (
	"HeadHunter/internal/entity/models"
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"strings"
)

func generateHTMLFromResume(resume *models.ResumeInPDF) (bytes.Buffer, error) {
	templ, errParse := template.ParseGlob("./static/html/resume.html")
	if errParse != nil {
		return bytes.Buffer{}, errParse
	}

	buffer := bytes.Buffer{}
	errExec := templ.Execute(&buffer, *resume)
	if errExec != nil {
		return bytes.Buffer{}, errExec
	}

	return buffer, nil
}

func generatePDFFromHTML(html bytes.Buffer) ([]byte, error) {
	pdfg, generatorErr := wkhtmltopdf.NewPDFGenerator()
	if generatorErr != nil {
		return nil, generatorErr
	}

	page := wkhtmltopdf.NewPageReader(&html)
	page.FooterFontSize.Set(10)
	page.Zoom.Set(1)

	pdfg.AddPage(page)

	createErr := pdfg.Create()
	if createErr != nil {
		return nil, createErr
	}

	return pdfg.Buffer().Bytes(), nil
}

func GenerateResumeInPDF(resume *models.ResumeInPDF) ([]byte, error) {
	resume.Image = strings.Split(resume.Image, "?")[0]
	html, htmlErr := generateHTMLFromResume(resume)
	if htmlErr != nil {
		return nil, htmlErr
	}

	pdf, pdfErr := generatePDFFromHTML(html)
	if pdfErr != nil {
		return nil, pdfErr
	}

	return pdf, nil
}
