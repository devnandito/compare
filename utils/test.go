package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/devnandito/compare/models"
	"github.com/devnandito/compare/structs"
	"github.com/xuri/excelize/v2"
)

var usr models.User
func SaveData(filename string){
	file, err := os.Open(filename)
	CheckError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Saltar la primera línea (encabezados)
	CheckError(err)

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		email := string(row[0])
		response, err := usr.CheckUser(email)
		if response.Email == "" {
			// fmt.Println("EL user no existe ", email)
			if row[2] == "Modulo I Ciberseguridad Conceptos básicos" {
				data := models.User {
					Email: email,
					Fullname: row[1],
					Module1: row[3] + " - " + row[4] + " - " + row[5] + " - " + row[6], 
				}
				response, err := usr.CreateUser(&data)
				CheckError(err)
				log.Println("Record inserted:", response)
			}

			if row[2] == "Modulo II Uso seguro de activos de información" {
				data := models.User {
					Email: email,
					Fullname: row[1],
					Module2: row[3] + " - " + row[4] + " - " + row[5] + " - " + row[6],
				}
				response, err := usr.CreateUser(&data)
				CheckError(err)
				log.Println("Record inserted:", response)
			}

			if row[2] == "Módulo III Gestión de acceso y firma electrónica" {
				data := models.User {
					Email: email,
					Fullname: row[1],
					Module3: row[3] + " - " + row[4] + " - " + row[5] + " - " + row[6],
				}
				response, err := usr.CreateUser(&data)
				CheckError(err)
				log.Println("Record inserted:", response)
			}

			if row[2] == "Módulo IV Riesgos y Amenazas" {
				data := models.User {
					Email: email,
					Fullname: row[1],
					Module4:  row[3] + " - " + row[4] + " - " + row[5] + " - " + row[6],
				}
				response, err := usr.CreateUser(&data)
				CheckError(err)
				log.Println("Record inserted:", response)
			}

			if row[2] == "Módulo V Gestión de Incidentes" {
				data := models.User {
					Email: email,
					Fullname: row[1],
					Module5: row[3] + " - " + row[4] + " - " + row[5] + " - " + row[6],
				}
				response, err := usr.CreateUser(&data)
				CheckError(err)
				log.Println("Record inserted:", response)
			}

		} else {
			// fmt.Println("El user existe")
			if row[2] == "Modulo I Ciberseguridad Conceptos básicos" {
				data := models.User {
					Module1: row[3] + " - " + row[4] + " - " + row[5],
				}
				response, err := usr.UpdateUser(email, &data, "module1")
				CheckError(err)
				log.Println("Record updated:", response)
			}

			if row[2] == "Modulo II Uso seguro de activos de información" {
				data := models.User {
					Module2: row[3] + " - " + row[4] + " - " + row[5],
				}
				response, err := usr.UpdateUser(email, &data, "module2")
				CheckError(err)
				log.Println("Record updated:", response)
			}

			if row[2] == "Módulo III Gestión de acceso y firma electrónica" {
				data := models.User {
					Module3: row[3] + " - " + row[4] + " - " + row[5],
				}
				response, err := usr.UpdateUser(email, &data, "module3")
				CheckError(err)
				log.Println("Record updated:", response)
			}

			if row[2] == "Módulo IV Riesgos y Amenazas" {
				data := models.User {
					Module4: row[3] + " - " + row[4] + " - " + row[5],
				}
				response, err := usr.UpdateUser(email, &data, "module4")
				CheckError(err)
				log.Println("Record updated:", response)
			}

			if row[2] == "Módulo V Gestión de Incidentes" {
				data := models.User {
					Module5: row[3] + " - " + row[4] + " - " + row[5],
				}
				response, err := usr.UpdateUser(email, &data, "module5")
				CheckError(err)
				log.Println("Record updated:", response)
			}
		}
	}
}

func SaveToXlsx(filename, outputname string){
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	f := excelize.NewFile()
	sheetName := "Sheet1"
	row := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		cell, err := excelize.CoordinatesToCellName(1, row)
		if err!= nil {
			fmt.Println(err)
			break
		}

		if row == 1 {
			if err := f.SetSheetRow(sheetName, cell, &record); err!= nil {
				fmt.Println(err)
				break
			}
			row++
			continue
			}

			data, err := ConvertSlice(record)
			if err!= nil {
				fmt.Println(err)
				break
			}

			if err := f.SetSheetRow(sheetName, cell, &data); err!= nil {
				fmt.Println(err)
				break
			}
			row++
		}

		if err := f.SaveAs(outputname); err!= nil {
			fmt.Println(err)
		}
	}

	func ConvertSlice(record []string) (data []interface{}, err error) {
		for _, arg := range record {
			// var n float64
			// if n, err = strconv.ParseFloat(arg, 64); err == nil {
			// 	numbers = append(numbers, n)
			// }
			data = append(data, arg)
		}
		return
	}

func GetApiKb4(uri, page, per_page, excelName string){
	params := url.Values{}
	params.Add("page", page)
	params.Add("per_page", per_page)
	tk := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJzaXRlIjoidHJhaW5pbmcua25vd2JlNC5jb20iLCJ1dWlkIjoiNWQzYTY4ZGEtYTNhNi00OWY2LTk5NDAtNDU4ZDFkYjNhNzBkIiwic2NvcGVzIjpbImVsdmlzIl0sImFpZCI6NTgyNzIwfQ.j3iaBsTrWdyl6qZ8FM_DG4GLlJAB0UK39SVNut5rFSQNOgRD6Teeyy9iHhbMCKhKbObDsrH26HqhWQbnI2t9IQ"

	req, err := http.NewRequest("GET", uri+"?"+params.Encode(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", tk)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
	} else {
		fmt.Println("User List Loaded Successfully")
	}

	var responseData []structs.Enrollment
  err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	data := [][]interface{}{}
	for _, value := range responseData {
		row := []interface{}{
			value.User.Email,
			value.User.FirstName+" "+value.User.LastName,
			value.ModuleName,
			value.CampaignName,
			value.Status,
		}
		data = append(data, row)
		// fmt.Println(value.User.FirstName, value.User.LastName, value.User.Email, value.TimeSpent, value.StartDate)
	}


	f := excelize.NewFile()
	headers := []string{"Email", "Name", "Module", "Content", "Status", "CampaignID"}
	for i, header := range headers {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)),1 ), header)
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := f.SaveAs(excelName+".xlsx"); err != nil{
		fmt.Println(err)
	}
	
}

func GetEnrollmentApiKb4(uri, xlsxName, output string){
	f, err := excelize.OpenFile(xlsxName+".xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	f1 := excelize.NewFile()
	headers := []string{"Email", "Name", "Module", "Content", "Status",}
	for i, header := range headers {
		f1.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)),1 ), header)
	}

	if err := f1.SaveAs(output+".xlsx"); err != nil{
		fmt.Println(err)
	}

	count := 0
	for i, row := range rows {
		for j, colCell := range row {
			cell := string(rune(65+j))
			if cell == "A" && i != 0 {
				fmt.Println(colCell)
				params := url.Values{}
				params.Add("user_id", colCell)
				tk := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJzaXRlIjoidHJhaW5pbmcua25vd2JlNC5jb20iLCJ1dWlkIjoiNWQzYTY4ZGEtYTNhNi00OWY2LTk5NDAtNDU4ZDFkYjNhNzBkIiwic2NvcGVzIjpbImVsdmlzIl0sImFpZCI6NTgyNzIwfQ.j3iaBsTrWdyl6qZ8FM_DG4GLlJAB0UK39SVNut5rFSQNOgRD6Teeyy9iHhbMCKhKbObDsrH26HqhWQbnI2t9IQ"
			
				req, err := http.NewRequest("GET", uri+"?"+params.Encode(), nil)
				if err != nil {
					fmt.Println("Error creating request:", err)
					return
				}
			
				req.Header.Set("accept", "application/json")
				req.Header.Set("Authorization", tk)
			
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return
				}
				defer resp.Body.Close()
			
				if resp.StatusCode != http.StatusOK {
					fmt.Println("Error:", resp.Status)
				} else {
					fmt.Println("User List Loaded Successfully")
				}
			
				var responseData []structs.Enrollment
				err = json.NewDecoder(resp.Body).Decode(&responseData)
				if err != nil {
					fmt.Println("Error decoding JSON:", err)
					return
				}
			
				data := [][]interface{}{}
				for _, value := range responseData {
					row := []interface{}{
						value.User.Email,
						value.User.FirstName+" "+value.User.LastName,
						value.ModuleName,
						value.CampaignName,
						value.Status,
					}
					data = append(data, row)
				}
			
				file, err := excelize.OpenFile(output+".xlsx")
				if err != nil {
					fmt.Println(err)
				}

				for _, row := range data {
					dataRow := count + 2
					for j, col := range row {
						file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
					}
					count++
				}

				err = file.Save()

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Excel update")
			}
		}
	}
	
}

func GetUsersKb4(uri, per_page, xlsxName string, totalPage int){
	var users []structs.User
	f := excelize.NewFile()
	headers := []string{"Id", "Name", "Email"}
	
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
		// params.Add("group_id", "4298854")
		tk := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJzaXRlIjoidHJhaW5pbmcua25vd2JlNC5jb20iLCJ1dWlkIjoiNWQzYTY4ZGEtYTNhNi00OWY2LTk5NDAtNDU4ZDFkYjNhNzBkIiwic2NvcGVzIjpbImVsdmlzIl0sImFpZCI6NTgyNzIwfQ.j3iaBsTrWdyl6qZ8FM_DG4GLlJAB0UK39SVNut5rFSQNOgRD6Teeyy9iHhbMCKhKbObDsrH26HqhWQbnI2t9IQ"
	
		req, err := http.NewRequest("GET", uri+"?"+params.Encode(), nil)
		if err != nil {
			fmt.Println("Error creating request: ", err)
			return
		}
	
		req.Header.Set("accept", "application/json")
		req.Header.Set("Authorization", tk)
	
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request: ", err)
			return
		}
		defer res.Body.Close()
	
		if res.StatusCode != http.StatusOK {
			fmt.Println("Error: ", res.Status)
		} else {
			fmt.Println("User list Successfully")
		}
	
		err = json.NewDecoder(res.Body).Decode(&users)
		if err != nil {
			fmt.Println("Error decoding JSON: ", err)
			return
		}

		data := [][]interface{}{}
		for _, value := range users {
			row := []interface{}{
				value.UserId,
				value.FirstName+" "+value.LastName,
				value.Email,
			}
			data = append(data, row)
		}

		file, err := excelize.OpenFile(xlsxName+".xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
	
		for _, row := range data {
			dataRow := count + 2
			for j, col := range row {
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
			}
			count++
		}
	
		err = file.Save()
	
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Excel update")
	}
}