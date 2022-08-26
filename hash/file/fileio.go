package file

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
)

type Excel struct{
	Name string
	f *excelize.File
}

func NewExcel(name string,option excelize.Options) (*Excel){
	name=name+".xlsx"
	f:=new(excelize.File)
	if _,err:=os.Stat(name);err!=nil{
		f=excelize.NewFile()
		f.SaveAs(name)
		f.Close()
	}
	f,_=excelize.OpenFile(name,option)
	return &Excel{ f:f,Name:name };
}
// Record value
func(file *Excel) Record( sheet string, cell string,value float64) error{
	if file.f==nil{
		return errors.New("file is wrong")
	}
	// 设置单元格的值
	file.f.SetCellFloat(sheet,cell,value,6,64)
	// 根据指定路径保存文件
	if err := file.f.Save(); err != nil {
		return err;
	}
	return nil
}

func(file *Excel) Close() error{
	err:=file.f.Close()
	return err;
}

func(file *Excel) Newsheet(name string) error{
	if _,err:=file.f.NewSheet(name);err!=nil{
		return err
	}
	return nil;
}

func (file *Excel) Read (sheet string, cell string) error{
	return nil;
}