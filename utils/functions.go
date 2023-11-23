package utils

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	"github.com/devnandito/compare/structs"
	"github.com/xuri/excelize/v2"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Split(str, delimiter string) []string {
	spl := strings.Split(str, delimiter)
	return spl
}

func LoadCsv(filename string) (data map[int]*structs.Db) {
	content, err := os.Open(filename)
	CheckError(err)
	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	content.Close()
	var dat *structs.Db
	myDict := make(map[int]*structs.Db)
	for key, line := range fileLines {
		if key !=0 {
			spl := strings.Split(line, ",")
			dat = &structs.Db{
				Email: spl[0],
				Name: spl[1] + " " + spl[2],
				Module: spl[3],
				Content: spl[4],
				Status: spl[6],
				Past: spl[7],
			}
			myDict[key] = dat
		}
	}
	return myDict
}

func LoadEmailCsv(filename string) (data map[int]*structs.DbEmail) {
	content, err := os.Open(filename)
	CheckError(err)
	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	content.Close()
	var dat *structs.DbEmail
	myDict := make(map[int]*structs.DbEmail)
	for key, line := range fileLines {
		if key != 0 {
			dat = &structs.DbEmail{
				Email: line,
			}
			myDict[key] = dat
		}
	}
	return myDict
}

func LoadEmailXlsx(filename string) (data map[int]*structs.DbEmail){
	xlsx, err := excelize.OpenFile(filename)
	CheckError(err)
	// cell, err := f.GetCellValue("Hoja2", "B2")
	// Check(err)
	// fmt.Println(cell)
	rows, err := xlsx.GetRows("Hoja2")
	var dat *structs.DbEmail
	myDict := make(map[int]*structs.DbEmail)
	for i, row := range rows {
		if i != 0 {
			for key, value := range row {
				if key == 1 {
					dat = &structs.DbEmail {
						Email : value,
					}
					myDict[i] = dat
				}
			}
		}
	}
	return myDict
}

func CompareDict(dict1 map[int]*structs.DbEmail, dict2 map[int]*structs.Db) (map[int]*structs.Db) {
	var tmp *structs.Db
	myDictC := make(map[int]*structs.Db)
	var key = 0
	for _, value1 := range dict2 {
		for _, value2 := range dict1 {
			if value1.Email == value2.Email {
				tmp = &structs.Db{
					Email: value1.Email,
					Name: value1.Name,
					Module: value1.Module,
					Content: value1.Content,
					Status: value1.Status,
					Past: value1.Past,
				}
				myDictC[key] = tmp
				key +=1
			}
		}
	}
	return myDictC
}

func WriteCsv(records map[int]*structs.Db) {
	file, err := os.Create("phishing-enter.csv")
	CheckError(err)
	w := csv.NewWriter(file)
	var data [][]string
	row := []string{"Email", "Name", "Module", "Content", "Status", "Past"}
	data = append(data, row)
	for _, record := range records {
		row := []string{record.Email, record.Name, record.Module, record.Content, record.Status, record.Past}
		data = append(data, row)
	}
	w.WriteAll(data)
}