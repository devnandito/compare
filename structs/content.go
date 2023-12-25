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

type Enrollment struct {
	EnrollmentID int `json:"enrollment_id"`
	ContentType string `json:"content_type"`
	ModuleName string `json:"module_name"`
	User struct {
		Id int `json:"id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email string `json:"email"`
	}
	CampaignName string `json:"campaign_name"`
	EnrollmentDate string `json:"enrollment_date"`
	StartDate string `json:"start_date"`
	CompletionDate string `json:"completion_date"`
	Status string `json:"status"`
	TimeSpent int `json:"time_spent"`
	PolicyAcknowledged bool `json:"policy_acknowledged"`
	Score float64 `json:"score"`
}

type User struct {
	UserId int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}