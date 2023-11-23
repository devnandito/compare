package main

import (
	"github.com/devnandito/compare/structs"
	"github.com/devnandito/compare/utils"
)

func main() {
	data1 := make(map[int]*structs.Db)
	data1 = utils.LoadCsv("./csv/moduleall.csv")
	data2 := make(map[int]*structs.DbEmail)
	data2 = utils.LoadEmailXlsx("phishing-acciones.xlsx")
	equal := utils.CompareDict(data2, data1)
	utils.WriteCsv(equal)
	utils.SaveData("phishing-enter.csv")	
}