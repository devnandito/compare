package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Db struct {
	Email string
	Name string
	Module string
	Content string
	Status string
	Past string
}

type SumRecord struct {
	Email string
	Name string
	Module1 string
	Module2 string
	Module3 string
	Module4 string
	Module5 string
}

type DbEmail struct {
	Email string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadCsv(filename string) (data map[int]*Db) {
	content, err := os.Open(filename)
	Check(err)
	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	content.Close()
	var dat *Db
	myDict := make(map[int]*Db)
	for key, line := range fileLines {
		if key !=0 {
			spl := strings.Split(line, ",")
			dat = &Db{
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

func LoadEmail(filename string) (data map[int]*DbEmail) {
	content, err := os.Open(filename)
	Check(err)
	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	content.Close()
	var dat *DbEmail
	myDict := make(map[int]*DbEmail)
	for key, line := range fileLines {
		if key != 0 {
			dat = &DbEmail{
				Email: line,
			}
			myDict[key] = dat
		}
	}
	return myDict
}

func CompareDict(dict1 map[int]*DbEmail, dict2 map[int]*Db) (map[int]*Db) {
	var tmp *Db
	myDictC := make(map[int]*Db)
	var key = 0
	for _, value1 := range dict2 {
		for _, value2 := range dict1 {
			if value1.Email == value2.Email {
				tmp = &Db{
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

func LoadEmailXlsx(filename string) (data map[int]*DbEmail){
	xlsx, err := excelize.OpenFile(filename)
	Check(err)
	// cell, err := f.GetCellValue("Hoja2", "B2")
	// Check(err)
	// fmt.Println(cell)
	rows, err := xlsx.GetRows("Hoja2")
	var dat *DbEmail
	myDict := make(map[int]*DbEmail)
	for i, row := range rows {
		if i != 0 {
			for key, value := range row {
				if key == 1 {
					dat = &DbEmail {
						Email : value,
					}
					myDict[i] = dat
					// fmt.Println(i," ", key, " ", value)
				}
			}
		}
	}
	return myDict
}

func WriteCsv(records map[int]*Db) {
	file, err := os.Create("phishing-enter.csv")
	Check(err)
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

// func Unique(records map[int]SumRecord) {
// 	result := []SumRecord{}
// 	encountered := map[int]bool{}

// 	for key := range records {
// 		encountered[records[key]] = true
// 	}

// 	for _, value := range encountered {
// 		result = append(result, value.Email)
// 	}
// }

func Count(records map[int]*Db) (map[string]string) {
	var module1 string
	var module2 string
	var module3 string
	var module4 string
	var module5 string
	// srMap := make(map[string]int)
	srMap := make(map[SumRecord]int)
	strMap := make(map[int]SumRecord)
	// sr := make(map[string]SumRecord)
	email := map[string]string{}
	for key := range records {
		email[records[key].Email] = records[key].Email
		if records[key].Module == "Modulo I Ciberseguridad Conceptos básicos" {
			module1 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
		}
		if records[key].Module == "Modulo II Uso seguro de activos de información" {
			module2 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
		}
		if records[key].Module == "Módulo III Gestión de acceso y firma electrónica" {
			module3 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
		}
		if records[key].Module == "Módulo IV Riesgos y Amenazas" {
			module4 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
		}
		if records[key].Module == "Módulo V Gestión de Incidentes" {
			module5 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
		}
		
		sr := SumRecord {
			Email: records[key].Email,
			Name: records[key].Name,
			Module1: module1,
			Module2: module2,
			Module3: module3,
			Module4: module4,
			Module5: module5,
		}
		// srMap[sr] +=1
		// srMap[records[key].Email] +=1
		srMap[sr] +=1
		strMap[key] = sr

	}
	fmt.Println(srMap)
	// for _, record := range strMap {
	// 	fmt.Println(record.Email, " ", record.Module1, " ", record.Module2, " ", record.Module3, " ", record.Module4, " ", record.Module5)
	// }
	
	file, err := os.Create("phishing-enter1.csv")
	Check(err)
	w := csv.NewWriter(file)
	var data [][]string
	row := []string{"Email", "Name", "Module1", "Module2", "Module3", "Module4", "Module5"}
	data = append(data, row)
	for _, record := range strMap {
		row := []string{record.Email, record.Name, record.Module1, record.Module2, record.Module3, record.Module4, record.Module5}
		data = append(data, row)
	}
	w.WriteAll(data)
	
	// dict := make(map[string]int)
	// for _, record := range records {
	// 	dict[record.Email] = dict[record.Email]+1
	// }
	// fmt.Println(dict)
	return email
}

func RemoveDuplicate(filename string) []string {
	content, err := os.Open(filename)
	Check(err)
	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	content.Close()
	var data []string
	result := []string{}
	encountered := map[string]bool{}

	for key, line := range fileLines {
		if key !=0 {
			spl := strings.Split(line, ",")
			data = append(data, spl[0])
		}
	}

	for v := range data {
		encountered[data[v]] = true
	}

	for key, _ := range encountered {
		result = append(result, key)
	}

	// for _, str := range data {
	// 	dat1 := strings.Split(str, ",")
	// 	if _, ok := encountered[dat1[0]]; !ok {
	// 		encountered[str] = true
	// 		result = append(result, str)
	// 	}
	// }

	return result

}


func main() {
	data1 := make(map[int]*Db)
	data1 = LoadCsv("./csv/moduleall.csv")
	data2 := make(map[int]*DbEmail)
	data2 = LoadEmailXlsx("phishing-acciones.xlsx")
	equal := CompareDict(data2, data1)
	WriteCsv(equal)
	Count(equal)
	// for _, record := range data1 {
	// 	fmt.Println(record.Module)
	// }
	// data3 := RemoveDuplicate("phishing-enter1.csv")

	// for key, line := range data3 {
	// 	if key != 0 {
	// 		// dat1 := strings.Split(line, ",")
	// 		fmt.Println(line)
	// 	}
	// }
}


// for _, value := range equal {
// 	fmt.Println(value.Email, " ", value.Name, " ", value.Module, " ", value.Status, " ", value.Past)
// }