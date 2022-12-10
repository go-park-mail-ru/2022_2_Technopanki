package utils

import (
	"HeadHunter/internal/entity/models"
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"strings"
)

var BASIC_IMAGES = []string{"basic_applicant_avatar.webp", "basic_employer_avatar.webp"}

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
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	pdfg.AddPage(page)

	createErr := pdfg.Create()
	if createErr != nil {
		return nil, createErr
	}

	return pdfg.Buffer().Bytes(), nil
}

func parseImagePath(path string) string {
	if path == BASIC_IMAGES[0] || path == BASIC_IMAGES[1] {
		return path
	}

	return strings.Split(path, "?")[0]
}

func GenerateResumeInPDF(resume *models.ResumeInPDF) ([]byte, error) {
	resume.Image = parseImagePath(resume.Image)

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
