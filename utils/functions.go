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

func SplitExtract(str, delimiter string) []string {
	spl := strings.Split(str, delimiter)
	return spl
}

func LoadCsv(filename string, head structs.Csv) (data map[int]*structs.Db) {
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
			spl := SplitExtract(line, ",")
			email := spl[head.Email]
			fistname := spl[head.Firstname]
			lastname := spl[head.Lastname]
			riskscore := spl[head.RiskScore]
			module := spl[head.Module]
			content := spl[head.Content]
			ctype := spl[head.ContentType]
			duration := spl[head.Duration]
			status := spl[head.Status]
			past := spl[head.Past]
			completed := spl[head.Completed]
			timespent := spl[head.TimeSpent]
			score := spl[head.Score]

			dat = &structs.Db{
				Email: email,
				Name: fistname + " " + lastname,
				RiskScore: riskscore,
				Module: module,
				Content: content,
				ContentType: ctype,
				Duration: duration,
				Status: status,
				Past: past,
				Completed: completed,
				TimeSpent: timespent,
				Score: score,
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
	dictUnion := make(map[int]*structs.Db)
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
					Duration: value1.Duration,
					TimeSpent: value1.TimeSpent,
					Past: value1.Past,
				}
				dictUnion[key] = tmp
				key +=1
			}
		}
	}
	return dictUnion
}

func WriteCsv(records map[int]*structs.Db, filename string) {
	file, err := os.Create(filename)
	CheckError(err)
	w := csv.NewWriter(file)
	var data [][]string
	row := []string{"Email", "Name", "Module", "Content", "Status", "Duration", "Past"}
	data = append(data, row)
	for _, record := range records {
		row := []string{record.Email, record.Name, record.Module, record.Content, record.Status, record.Duration, record.Past}
		data = append(data, row)
	}
	w.WriteAll(data)
}