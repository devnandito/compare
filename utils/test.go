package utils

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/devnandito/compare/models"
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