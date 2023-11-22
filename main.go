package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/devnandito/compare/structs"
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
	Content string
	Module1 string
	Module2 string
	Module3 string
	Module4 string
	Module5 string
}
type ContentRecord struct {
	Email string
	Name string
	Module1 map[string]Content
	Module2 map[string]Content
	Module3 map[string]Content
	Module4 map[string]Content
	Module5 map[string]Content
}

type Content struct {
	Email string
	Fullname string
	Name string
	Status string
	Past string
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


func Count1(records map[int]*Db) map[string]Content {
	var module1 string
	var module2 string
	var module3 string
	var module4 string
	var module5 string
	// data1 := make(map[string]ContentRecord)
	content := make(map[string]Content)
	for key := range records {
		if records[key].Module == "Modulo I Ciberseguridad Conceptos básicos" {
			module1 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
			content[module1] = Content {
				Email: records[key].Email,
				Fullname: records[key].Name,
				Name: records[key].Content,
				Status: records[key].Status,
				Past: records[key].Past,
			}
		}
		if records[key].Module == "Modulo II Uso seguro de activos de información" {
			module2 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
			content[module2] = Content {
				Email: records[key].Email,
				Fullname: records[key].Name,
				Name: records[key].Content,
				Status: records[key].Status,
				Past: records[key].Past,
			}
		}
		if records[key].Module == "Módulo III Gestión de acceso y firma electrónica" {
			module3 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
			content[module3] = Content {
				Email: records[key].Email,
				Fullname: records[key].Name,
				Name: records[key].Content,
				Status: records[key].Status,
				Past: records[key].Past,
			}
		}
		if records[key].Module == "Módulo IV Riesgos y Amenazas" {
			module4 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
			content[module4] = Content {
				Email: records[key].Email,
				Fullname: records[key].Name,
				Name: records[key].Content,
				Status: records[key].Status,
				Past: records[key].Past,
			}
		}
		if records[key].Module == "Módulo V Gestión de Incidentes" {
			module5 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email 
			content[module5] = Content {
				Email: records[key].Email,
				Fullname: records[key].Name,
				Name: records[key].Content,
				Status: records[key].Status,
				Past: records[key].Past,
			}
		}
	}
	
	return content
}

func RemoveDuplicate(records map[string]Content, allemail map[int]*DbEmail) () {
	data := make(map[string]SumRecord)
	tmp := make(map[string]SumRecord)
	// var tmp0 SumRecord
	var mod1 string
	var mod2 string
	var mod3 string
	var mod4 string
	var mod5 string
	// var user []string
	// var tmp1 string
	// var tmp2 string
	// var tmp3 string
	// var tmp4 string
	// var tmp5 string
 	for key, line := range records {
		data[key] = SumRecord {Email: line.Email, Name: line.Fullname, Content: line.Name + " " + line.Status}
		// user = append(user, line.Email)
	}

	// fmt.Println(data)

		for k, v := range data {
			spl := strings.Split(k, "-")
			// fmt.Println(k)
			if spl[0] == "Modulo I Ciberseguridad Conceptos básicos" {
				mod1 = mod1 + " - " + v.Content
				// tmp1 = spl[0] + " - " + mod1
			}
			if spl[0] == "Modulo II Uso seguro de activos de información" {
				mod2 = mod2 + " - " + v.Content
				// tmp2 = spl[0] + " - " + mod2
			}
			if spl[0] == "Módulo III Gestión de acceso y firma electrónica" {
				mod3 = mod3 + " - " + v.Content
				// tmp3 = spl[0] + " - " + mod3
			}
			if spl[0] == "Módulo IV Riesgos y Amenazas" {
				mod4 = mod4 + " - " + v.Content
				// tmp4 = spl[0] + " - " + mod4
			}
			if spl[0] == "Módulo V Gestión de Incidentes" {
				mod5 = mod5 + " - " + v.Content
				// tmp5 = spl[0] + " - " + mod5
			}
			tmp[spl[2]] = SumRecord {
				Email: v.Email,
				Name: v.Name,
				Module1: mod1,
				Module2: mod2,
				Module3: mod3,
				Module4: mod4,
				Module5: mod5,
			} 
	}
	// fmt.Println(tmp0)

	// values := reflect.ValueOf(tmp)
	// types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i).Name, values.Field(i))
	// }

	file, err := os.Create("phishing-enter2.csv")
	Check(err)
	w := csv.NewWriter(file)
	var data1 [][]string
	row := []string{"Email", "Nombre", "Modulo I Ciberseguridad Conceptos básicos", "Modulo II Uso seguro de activos de información", "Módulo III Gestión de acceso y firma electrónica", "Módulo IV Riesgos y Amenazas", "Módulo V Gestión de Incidentes"}
	data1 = append(data1, row)
	for _, record := range tmp {
		row := []string{record.Email, record.Name, record.Module1, record.Module2, record.Module3, record.Module4, record.Module5}
		data1 = append(data1, row)
	}
	w.WriteAll(data1)

}

func main() {
	var msg structs.EmailStore
	var email structs.EmailStore
	emailS := make(map[string]structs.EmailStore)
	data1 := make(map[int]*Db)
	data1 = LoadCsv("./csv/moduleall.csv")
	data2 := make(map[int]*DbEmail)
	data2 = LoadEmailXlsx("phishing-acciones.xlsx")
	equal := CompareDict(data2, data1)
	WriteCsv(equal)
	RemoveDuplicate(Count1(equal), data2)
	data3 := Count1(equal)
	// mod := make(map[string]string)
	for k, v := range data3 {
		// mod[v.Module] = mod[v.Module] + v.Content
		val := structs.EmailStore {
			Email: v.Email,
			Name: v.Name,
			Module: k,
		}
		email := email.InsertData(&val)
		emailS[v.Email] = email
	}
	// fmt.Println(emailS)
	fmt.Println(emailS["ccuevas@bcp.gov.py"])
	fmt.Println(len(emailS))
	// e := msg.Verifity(equal[0].Email)
	// fmt.Println(e)
	for _, v := range data3 {
		e := msg.Verifity(v.Email, emailS)
		if e {
			fmt.Println(v.Email)
		}else {
			continue
		}
	}
}


// func Count(records map[int]*Db) {
// 	var module1 string
// 	var module2 string
// 	var module3 string
// 	var module4 string
// 	var module5 string
// 	srMap := make(map[SumRecord]int)
// 	strMap := make(map[int]SumRecord)
// 	for key := range records {
// 		if records[key].Module == "Modulo I Ciberseguridad Conceptos básicos" {
// 			module1 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
// 		}
// 		if records[key].Module == "Modulo II Uso seguro de activos de información" {
// 			module2 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
// 		}
// 		if records[key].Module == "Módulo III Gestión de acceso y firma electrónica" {
// 			module3 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
// 		}
// 		if records[key].Module == "Módulo IV Riesgos y Amenazas" {
// 			module4 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
// 		}
// 		if records[key].Module == "Módulo V Gestión de Incidentes" {
// 			module5 = records[key].Module + " - " + records[key].Content + " - " + records[key].Status
// 		}
		
// 		sr := SumRecord {
// 			Email: records[key].Email,
// 			Name: records[key].Name,
// 			Module1: module1,
// 			Module2: module2,
// 			Module3: module3,
// 			Module4: module4,
// 			Module5: module5,
// 		}
// 		// srMap[sr] +=1
// 		// srMap[records[key].Email] +=1
// 		srMap[sr] +=1
// 		strMap[key] = sr

// 	}

// 	file, err := os.Create("phishing-enter1.csv")
// 	Check(err)
// 	w := csv.NewWriter(file)
// 	var data [][]string
// 	row := []string{"Email", "Name", "Module1", "Module2", "Module3", "Module4", "Module5"}
// 	data = append(data, row)
// 	for _, record := range strMap {
// 		row := []string{record.Email, record.Name, record.Module1, record.Module2, record.Module3, record.Module4, record.Module5}
// 		data = append(data, row)
// 	}
// 	w.WriteAll(data)
// }


// for _, value := range equal {
// 	fmt.Println(value.Email, " ", value.Name, " ", value.Module, " ", value.Status, " ", value.Past)
// }



// func RemoveDuplicate(filename string) (/*[]string,*/ map[string]SumRecord) {
// 	content, err := os.Open(filename)
// 	Check(err)

// 	fileScanner := bufio.NewScanner(content)
// 	fileScanner.Split(bufio.ScanLines)
// 	var fileLines []string

// 	for fileScanner.Scan() {
// 		fileLines = append(fileLines, fileScanner.Text())
// 	}
// 	content.Close()
// 	// var data []string
// 	// data := make(map[string][]interface{})
// 	data := make(map[string]SumRecord)
// 	// result := []string{}
// 	// encountered := map[string]bool{}

// 	for key, line := range fileLines {
// 		if key !=0 {
// 			spl := strings.Split(line, ",")
// 			// data = append(data, spl[0])
// 			// data[spl[0]] = []interface{}{spl[1], spl[2], spl[3], spl[4], spl[5], spl[6]}
// 			data[spl[0]] = SumRecord {Email: spl[0], Name: spl[1], Module1: spl[2], Module2: spl[3], Module3: spl[4], Module4: spl[5], Module5: spl[6]}
// 		}
// 	}

// 	file, err := os.Create("phishing-enter2.csv")
// 	Check(err)
// 	w := csv.NewWriter(file)
// 	var data1 [][]string
// 	row := []string{"Email", "Name", "Module1", "Module2", "Module3", "Module4", "Module5"}
// 	data1 = append(data1, row)
// 	for _, record := range data {
// 		row := []string{record.Email, record.Name, record.Module1, record.Module2, record.Module3, record.Module4, record.Module5}
// 		data1 = append(data1, row)
// 	}
// 	w.WriteAll(data1)

// 	// for v := range data {
// 	// 	encountered[v] = true
// 	// 	// encountered[data[v]] = true
// 	// }

// 	// for key, _ := range encountered {
// 	// 	result = append(result, key)
// 	// }

// 	// fmt.Println(data)
// 	// for _, str := range data {
// 	// 	dat1 := strings.Split(str, ",")
// 	// 	if _, ok := encountered[dat1[0]]; !ok {
// 	// 		encountered[str] = true
// 	// 		result = append(result, str)
// 	// 	}
// 	// }

// 	return /*result,*/ data

// }