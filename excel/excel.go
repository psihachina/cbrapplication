package excel

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/cbrapplication/model"
)

// CreateExcelWithValutes - function to create an excel file with rates and chart
func CreateExcelWithValutes(valutes []model.ValuteCursOnDate, times []string, path string) error {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	// Создать новый лист
	index := f.NewSheet("Valutes")
	// Установленное значение ячейки
	f.SetCellValue("Valutes", "A1", "Имя валюты")
	f.SetCellValue("Valutes", "B1", "Номинал")
	f.SetCellValue("Valutes", "C1", "Курс")
	f.SetCellValue("Valutes", "D1", "Код")
	f.SetCellValue("Valutes", "E1", "Символьный код")
	f.SetCellValue("Valutes", "F1", "Дата")

	for i, v := range valutes {
		vcurs, err := strconv.ParseFloat(v.Vcurs, 32)
		if err != nil {
			return err
		}
		f.SetCellValue("Valutes", "A"+fmt.Sprintf("%v", (i+2)), v.Vname)
		f.SetCellValue("Valutes", "B"+fmt.Sprintf("%v", (i+2)), v.Vnom)
		f.SetCellValue("Valutes", "C"+fmt.Sprintf("%v", (i+2)), vcurs)
		f.SetCellValue("Valutes", "D"+fmt.Sprintf("%v", (i+2)), v.Vcode)
		f.SetCellValue("Valutes", "E"+fmt.Sprintf("%v", (i+2)), v.VchCode)
		f.SetCellValue("Valutes", "F"+fmt.Sprintf("%v", (i+2)), times[i])
	}

	if err := f.AddChart("Valutes", "G1", `{
        "type": "line",
        "series": [
        {
            "categories": "Valutes!$F$2:$F$6",
            "values": "Valutes!$C$2:$C$6"
        }],
        "format":
        {
            "x_scale": 1.0,
            "y_scale": 1.0,
            "x_offset": 15,
            "y_offset": 10,
            "print_obj": true,
            "lock_aspect_ratio": false,
            "locked": false
        },
        "legend":
        {
            "position": "left",
            "show_legend_key": false
        },
        "plotarea":
        {
            "show_bubble_size": true,
            "show_cat_name": false,
            "show_leader_lines": false,
            "show_percent": true,
            "show_series_name": true,
            "show_val": true
        }}`); err != nil {
		fmt.Println(err)
	}

	// Установить активный лист рабочей книги
	f.SetActiveSheet(index)

	// Сохранить файл xlsx по данному пути
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}

	return nil
}
