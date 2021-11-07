package service

import (
	"github.com/xuri/excelize/v2"
)

type ExcelService struct {
}

func NewExcelService() ExcelService {
	return ExcelService{}
}

func (e ExcelService) WriteExcel() error {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	f.SaveAs("static/excel/write.xlsx")

	return nil
}
