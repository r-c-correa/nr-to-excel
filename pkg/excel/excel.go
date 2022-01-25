package excel

import (
	"github.com/r-c-correa/nr-to-excel/pkg/errr"
	"github.com/xuri/excelize/v2"
)

func SaveDataInExcel(filename string, rows []map[string]interface{}) {
	if len(rows) == 0 {
		return
	}

	file := excelize.NewFile()
	defer file.Close()

	sheetIndex := file.GetActiveSheetIndex()
	sheetName := file.GetSheetName(sheetIndex)

	headers := map[string]int{}
	rowHeader := rows[0]
	for key, _ := range rowHeader {
		headers[key] = len(headers) + 1
		setValue(file, key, sheetName, len(headers), 1)
	}

	for rowIndex, rowValues := range rows {
		for fieldName, fieldValue := range rowValues {
			setValue(file, fieldValue, sheetName, headers[fieldName], rowIndex+2)
		}
	}

	file.SetActiveSheet(sheetIndex)

	errr.PanicIfIsNotNull(file.SaveAs(filename))
}

func setValue(file *excelize.File, value interface{}, sheetName string, col, row int) {
	coordinate, err := excelize.CoordinatesToCellName(col, row)
	errr.PanicIfIsNotNull(err)

	err = file.SetCellValue(sheetName, coordinate, value)
	errr.PanicIfIsNotNull(err)
}
