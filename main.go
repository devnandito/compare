package main

import (
	"github.com/devnandito/compare/structs"
	"github.com/devnandito/compare/utils"
)

func main() {
	data1 := make(map[int]*structs.Db)
	csv := structs.Csv {
		Email: 0,
		Firstname: 1,
		Lastname: 2,
		RiskScore: 3,
		Module: 4,
		Content: 5,
		ContentType: 6,
		Duration: 7,
		Status: 8,
		Past: 9,
		Completed: 10,
		TimeSpent: 11,
		Score: 12,
	}

	filename := "./csv/modules.csv"
	data1 = utils.LoadCsv(filename, csv)
	data2 := make(map[int]*structs.DbEmail)
	data2 = utils.LoadEmailXlsx("./csv/phishing-acciones.xlsx")
	equal := utils.CompareDict(data2, data1)
	file := "20231123_report_camping.csv"
	utils.WriteCsv(equal, file)
	utils.WriteCsv(data1, "20231123_report_camping_all.csv")
	// utils.SaveData(file)	
}