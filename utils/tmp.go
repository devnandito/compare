package utils

// func Count1(records map[int]*structs.Db) map[string]structs.Content {
// 	var module1 string
// 	var module2 string
// 	var module3 string
// 	var module4 string
// 	var module5 string
// 	// data1 := make(map[string]ContentRecord)
// 	content := make(map[string]structs.Content)
// 	for key := range records {
// 		if records[key].Module == "Modulo I Ciberseguridad Conceptos básicos" {
// 			module1 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
// 			content[module1] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Modulo II Uso seguro de activos de información" {
// 			module2 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
// 			content[module2] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo III Gestión de acceso y firma electrónica" {
// 			module3 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
// 			content[module3] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo IV Riesgos y Amenazas" {
// 			module4 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
// 			content[module4] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo V Gestión de Incidentes" {
// 			module5 = records[key].Module + "-" + records[key].Content + "-" + records[key].Email
// 			content[module5] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 	}

// 	return content
// }

// func RemoveDuplicate(records map[string]structs.Content, allemail map[int]*structs.DbEmail) () {
// 	data := make(map[string]structs.SumRecord)
// 	tmp := make(map[string]structs.SumRecord)
// 	// var tmp0 SumRecord
// 	var mod1 string
// 	var mod2 string
// 	var mod3 string
// 	var mod4 string
// 	var mod5 string
// 	// var user []string
// 	// var tmp1 string
// 	// var tmp2 string
// 	// var tmp3 string
// 	// var tmp4 string
// 	// var tmp5 string
//  	for key, line := range records {
// 		data[key] = structs.SumRecord {Email: line.Email, Name: line.Fullname, Content: line.Module + " " + line.Status}
// 		// user = append(user, line.Email)
// 	}

// 	// fmt.Println(data)

// 		for k, v := range data {
// 			spl := strings.Split(k, "-")
// 			// fmt.Println(k)
// 			if spl[0] == "Modulo I Ciberseguridad Conceptos básicos" {
// 				mod1 = mod1 + " - " + v.Content
// 				// tmp1 = spl[0] + " - " + mod1
// 			}
// 			if spl[0] == "Modulo II Uso seguro de activos de información" {
// 				mod2 = mod2 + " - " + v.Content
// 				// tmp2 = spl[0] + " - " + mod2
// 			}
// 			if spl[0] == "Módulo III Gestión de acceso y firma electrónica" {
// 				mod3 = mod3 + " - " + v.Content
// 				// tmp3 = spl[0] + " - " + mod3
// 			}
// 			if spl[0] == "Módulo IV Riesgos y Amenazas" {
// 				mod4 = mod4 + " - " + v.Content
// 				// tmp4 = spl[0] + " - " + mod4
// 			}
// 			if spl[0] == "Módulo V Gestión de Incidentes" {
// 				mod5 = mod5 + " - " + v.Content
// 				// tmp5 = spl[0] + " - " + mod5
// 			}
// 			tmp[spl[2]] = structs.SumRecord {
// 				Email: v.Email,
// 				Name: v.Name,
// 				Module1: mod1,
// 				Module2: mod2,
// 				Module3: mod3,
// 				Module4: mod4,
// 				Module5: mod5,
// 			}
// 	}
// 	// fmt.Println(tmp0)

// 	// values := reflect.ValueOf(tmp)
// 	// types := values.Type()
// 	// for i := 0; i < values.NumField(); i++ {
// 	// 	fmt.Println(types.Field(i).Name, values.Field(i))
// 	// }

// 	file, err := os.Create("phishing-enter2.csv")
// 	CheckError(err)
// 	w := csv.NewWriter(file)
// 	var data1 [][]string
// 	row := []string{"Email", "Nombre", "Modulo I Ciberseguridad Conceptos básicos", "Modulo II Uso seguro de activos de información", "Módulo III Gestión de acceso y firma electrónica", "Módulo IV Riesgos y Amenazas", "Módulo V Gestión de Incidentes"}
// 	data1 = append(data1, row)
// 	for _, record := range tmp {
// 		row := []string{record.Email, record.Name, record.Module1, record.Module2, record.Module3, record.Module4, record.Module5}
// 		data1 = append(data1, row)
// 	}
// 	w.WriteAll(data1)

// }

// func Count2(records map[int]*structs.Db) map[string]structs.Content {
// 	var module1 string
// 	var module2 string
// 	var module3 string
// 	var module4 string
// 	var module5 string
// 	// data1 := make(map[string]ContentRecord)
// 	content := make(map[string]structs.Content)
// 	for key := range records {
// 		if records[key].Module == "Modulo I Ciberseguridad Conceptos básicos" {
// 			module1 = records[key].Module + "*" + records[key].Content + "*" + records[key].Email
// 			content[module1] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Module1: module1,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Modulo II Uso seguro de activos de información" {
// 			module2 = records[key].Module + "*" + records[key].Content + "*" + records[key].Email
// 			content[module2] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Module2: module2,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo III Gestión de acceso y firma electrónica" {
// 			module3 = records[key].Module + "*" + records[key].Content + "*" + records[key].Email
// 			content[module3] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Module3: module3,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo IV Riesgos y Amenazas" {
// 			module4 = records[key].Module + "*" + records[key].Content + "*" + records[key].Email
// 			content[module4] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Module4: module4,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 		if records[key].Module == "Módulo V Gestión de Incidentes" {
// 			module5 = records[key].Module + "*" + records[key].Content + "*" + records[key].Email
// 			content[module5] = structs.Content {
// 				Email: records[key].Email,
// 				Fullname: records[key].Name,
// 				Module: records[key].Content,
// 				Module5: module5,
// 				Status: records[key].Status,
// 				Past: records[key].Past,
// 			}
// 		}
// 	}

// 	return content
// }

// func ReplaceString(input string) string {
// 	var module string

// 	switch input {
// 	case "Modulo I Ciberseguridad Conceptos básicos":
// 		module = "Module1"
// 	case "Modulo II Uso seguro de activos de información":
// 		module = "Module2"
// 	case "Módulo III Gestión de acceso y firma electrónica":
// 		module = "Module3"
// 	case "Módulo IV Riesgos y Amenazas":
// 		module = "Module4"
// 	case "Módulo V Gestión de Incidentes":
// 		module = "Module5"
// 	default:
// 		module = ""
// 	}

// 	return module
// }

// func Count3(records string){

// 	counts := make(map[string]int)
// 	activities := make(map[string]string)
// 	list1 := []map[string]string{}

// 	file, err := os.Open(records)
// 	if err != nil {
// 		fmt.Println("Error al abrir el archivo:", err)
// 		return
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	_, err = reader.Read() // Saltar la primera línea (encabezados)
// 	if err != nil {
// 		fmt.Println("Error al leer el encabezado:", err)
// 		return
// 	}

// 	for {
// 		row, err := reader.Read()
// 		if err != nil {
// 			break
// 		}

// 		// activity := ReplaceString(row[2])
// 		activity := row[2]
// 		email := row[0]

// 		counts[email]++

// 		if counts[email] == 1 {
// 			list1 = append(list1, map[string]string{
// 				"email":  email,
// 				"module": activities[activity],
// 			})
// 		}

// 		if _, ok := activities[activity]; !ok {
// 			activities[activity] = activity
// 		}
// 	}

// 	fmt.Println("Counts:")
// 	for email, count := range counts {
// 		fmt.Printf("Email: %s, Count: %d\n", email, count)
// 	}

// 	fmt.Println("Email and Module:")
// 	for _, item := range list1 {
// 		fmt.Println(item["email"], item["module"])
// 	}
// }

// data3 := utils.Count1(equal)
// mod := make(map[string]string)
// for k, v := range data3 {
// 	// mod[v.Module] = mod[v.Module] + v.Content
// 	val := structs.EmailStore {
// 		Email: v.Email,
// 		Name: v.Module,
// 		Module: k,
// 	}
// 	email := email.InsertData(&val)
// 	emailS[v.Email] = email
// }
// fmt.Println(emailS)
// fmt.Println(emailS["ccuevas@bcp.gov.py"])
// fmt.Println(len(emailS))
// e := msg.Verifity(equal[0].Email)
// fmt.Println(e)
// for _, v := range data3 {
// 	e := msg.Verifity(v.Email, emailS)
// 	if e {
// 		fmt.Println(v.Email)
// 	}else {
// 		continue
// 	}
// }
// data4 := utils.Count2(equal)
// for k, v := range data4 {
// 	spl := utils.Split(k, "*")
// 	if spl[2] == v.Email {
// 		// fmt.Println(k, v.Email, v.Status, spl[2], v.Fullname, v.Module1, v.Module2, v.Module3)
// 		fmt.Println(v.Module1, v.Module2, v.Module3, v.Module4, v.Module5)
// 	}
// }
// utils.Count3("phishing-enter.csv")

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