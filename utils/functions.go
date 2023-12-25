package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/devnandito/compare/structs"
	"github.com/xuri/excelize/v2"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckErrorMsg(msg string, e error){
	if e != nil {
		fmt.Println(msg,":",e)
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

func GetEnrollmentsKb4(uri, per_page, xlsxName string, totalPage int){
	var responseData []structs.Enrollment
	f := excelize.NewFile()
	headers := []string{"Email", "Name", "Module", "Content", "Status",}
	
	for i, header := range headers {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	if err := f.SaveAs(xlsxName+".xlsx"); err != nil {
		fmt.Println(err)
	}

	count := 0
	for page := 1; page <= totalPage; page++{
		params := url.Values{}
		p := strconv.Itoa(page)
		params.Add("page", p)
		params.Add("per_page", per_page)
		tk := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJzaXRlIjoidHJhaW5pbmcua25vd2JlNC5jb20iLCJ1dWlkIjoiNWQzYTY4ZGEtYTNhNi00OWY2LTk5NDAtNDU4ZDFkYjNhNzBkIiwic2NvcGVzIjpbImVsdmlzIl0sImFpZCI6NTgyNzIwfQ.j3iaBsTrWdyl6qZ8FM_DG4GLlJAB0UK39SVNut5rFSQNOgRD6Teeyy9iHhbMCKhKbObDsrH26HqhWQbnI2t9IQ"
	
		req, err := http.NewRequest("GET", uri+"?"+params.Encode(), nil)
		CheckErrorMsg("Error creating request: ", err)
			
		req.Header.Set("accept", "application/json")
		req.Header.Set("Authorization", tk)
	
		client := &http.Client{}
		res, err := client.Do(req)
		CheckErrorMsg("Error sending request: ", err)
		defer res.Body.Close()
	
		if res.StatusCode != http.StatusOK {
			fmt.Println("Error: ", res.Status)
		} else {
			fmt.Printf("Page %d load Successfully \n", page)
		}
	
		err = json.NewDecoder(res.Body).Decode(&responseData)
		CheckErrorMsg("Error decoding JSON: ", err)

		data := [][]interface{}{}
		for _, value := range responseData {
			row := []interface{}{
				value.User.Email,
				value.User.FirstName+" "+value.User.LastName,
				value.CampaignName,
				value.ModuleName,
				value.Status,
			}
			data = append(data, row)
		}

		file, err := excelize.OpenFile(xlsxName+".xlsx")
		CheckErrorMsg("Error opening file: ", err)
	
		for _, row := range data {
			dataRow := count + 2
			for j, col := range row {
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
			}
			count++
		}
	
		err = file.Save()
	
		CheckErrorMsg("Error to save file: ", err)

		fmt.Println("Excel update successfuly")
	}
}