package file

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestNewExcel(t *testing.T) {
	var opt excelize.Options

	file:=NewExcel("record.xlsx",opt)
	file.f.NewSheet("sheet1")
	file.f.Save()
	file.f.Close()
}

func TestFileRecode(t *testing.T){
	var opt excelize.Options
	file:=NewExcel("Book1",opt)
	file.Record("Sheet1","A1",0.0018)
	file.Close()
}
