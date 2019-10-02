package util

import (
	"github.com/tealeg/xlsx"
)

func Parse() [][]string{
	var form [][]string
	excelFileName := "./doc/产品链接.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var rows []string
			for _, cell := range row.Cells {
				text := cell.String()
				rows = append(rows,text)
			}
			form = append(form,rows)
		}
	}
	return form
}