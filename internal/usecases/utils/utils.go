package utils

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"bytes"
	"encoding/base64"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/webp"
	"html/template"
	"image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func encodeWebpToPng(webpReader io.Reader) ([]byte, error) {
	img, err := webp.Decode(webpReader, &decoder.Options{
		UseThreads: true,
	})
	if err != nil {
		return nil, err
	}

	buffer := bytes.Buffer{}
	err = png.Encode(&buffer, img)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func YearPostfix(year uint) string {
	switch year % 100 {
	case 1:
		return "год"
	case 12, 13, 14:
		return "лет"
	}

	switch year % 10 {
	case 2, 3, 4:
		return "года"
	default:
		return "лет"
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func getBase64MimeType(file []byte) (string, error) {
	mimeType := http.DetectContentType(file)

	switch mimeType {
	case "image/jpeg":
		return "data:image/jpeg;base64,", nil
	case "image/png":
		return "data:image/png;base64,", nil
	default:
		return "", errorHandler.ErrInvalidMimeType
	}
}

func generateBase64ForImage(imagePath string) (template.URL, error) {
	image, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	image, err = encodeWebpToPng(bytes.NewReader(image))
	if err != nil {
		return "", err
	}

	base64ImageMimeType, err := getBase64MimeType(image)
	if err != nil {
		return "", err
	}

	return template.URL(strings.Join([]string{base64ImageMimeType, toBase64(image)}, "")), nil
}

type resumeTemplate struct {
	Resume            *models.ResumeInPDF
	ExperiencePostfix string
	AgePostfix        string
	ImageBase64       template.URL
}

func generateHTMLFromResume(resume *models.ResumeInPDF, cfg *configs.Config, style string) (bytes.Buffer, error) {
	templ, errParse := template.ParseGlob(strings.Join([]string{cfg.PDFConfig.HTMLPath, style}, ""))
	if errParse != nil {
		return bytes.Buffer{}, errParse
	}

	experienceYear, errAtoi := strconv.ParseInt(resume.ExperienceInYears, 10, 64)
	if errAtoi != nil {
		return bytes.Buffer{}, errAtoi
	}

	base64Image, err := generateBase64ForImage(strings.Join([]string{cfg.Image.Path, "avatar/", resume.Image}, ""))
	if err != nil {
		return bytes.Buffer{}, err
	}

	templateStruct := &resumeTemplate{
		Resume:            resume,
		ExperiencePostfix: YearPostfix(uint(experienceYear)),
		AgePostfix:        YearPostfix(resume.Age),
		ImageBase64:       base64Image,
	}

	buffer := bytes.Buffer{}
	if errExec := templ.Execute(&buffer, templateStruct); errExec != nil {
		return bytes.Buffer{}, errExec
	}

	return buffer, nil
}

func generatePDFFromHTML(html bytes.Buffer, cfg *configs.Config) ([]byte, error) {
	pdfg, generatorErr := wkhtmltopdf.NewPDFGenerator()
	if generatorErr != nil {
		return nil, generatorErr
	}

	page := wkhtmltopdf.NewPageReader(&html)
	page.Zoom.Set(cfg.PDFConfig.ZoomSize)

	pdfg.AddPage(page)

	createErr := pdfg.Create()
	if createErr != nil {
		return nil, createErr
	}

	return pdfg.Buffer().Bytes(), nil
}

func GenerateResumeInPDF(resume *models.ResumeInPDF, cfg *configs.Config, style string) ([]byte, error) {
	resume.Image = strings.Split(resume.Image, "?")[0]
	html, htmlErr := generateHTMLFromResume(resume, cfg, style)
	if htmlErr != nil {
		return nil, htmlErr
	}

	pdf, pdfErr := generatePDFFromHTML(html, cfg)
	if pdfErr != nil {
		return nil, pdfErr
	}

	return pdf, nil
}
