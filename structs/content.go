package structs

type Db struct {
	Email string
	Name string
	RiskScore string
	Module string
	Content string
	ContentType string
	Duration string
	Status string
	Past string
	Completed string
	TimeSpent string
	Score string
}

type DbEmail struct {
	Email string
}

type Csv struct {
	Email int
	Firstname int
	Lastname int
	RiskScore int
	Module int
	Content int
	ContentType int
	Duration int
	Status int
	Past int
	Completed int
	TimeSpent int
	Score int
}
