package service

import (
	"log"

	"github.com/signintech/gopdf"
)

type PdfService struct {
}

func NewPdfService() PdfService {
	return PdfService{}
}

func (p PdfService) GeneratePdf() error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "static/fonts/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("static/pdf/chinese.pdf")

	return nil
}
