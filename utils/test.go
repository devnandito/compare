package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/devnandito/compare/models"
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