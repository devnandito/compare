package structs

type Db struct {
	Email string
	Name string
	Module string
	Content string
	Status string
	Past string
}

type Content struct {
	Email string
	Fullname string
	Module string
	Module1 string
	Module2 string
	Module3 string
	Module4 string
	Module5 string
	Status string
	Past string
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

type DbEmail struct {
	Email string
}

