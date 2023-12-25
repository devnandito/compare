package main

import "github.com/devnandito/compare/utils"

func main() {
	// data1 := make(map[int]*structs.Db)
	// csv := structs.Csv {
	// 	Email: 0,
	// 	Firstname: 1,
	// 	Lastname: 2,
	// 	RiskScore: 3,
	// 	Module: 4,
	// 	Content: 5,
	// 	ContentType: 6,
	// 	Duration: 7,
	// 	Status: 8,
	// 	Past: 9,
	// 	Completed: 10,
	// 	TimeSpent: 11,
	// 	Score: 12,
	// }

	// filename := "./csv/20231222_campaign_refuerzo.csv"
	// data1 = utils.LoadCsv(filename, csv)
	// data2 := make(map[int]*structs.DbEmail)
	// data2 = utils.LoadEmailXlsx("./csv/phishing-acciones.xlsx")
	// equal := utils.CompareDict(data2, data1)
	// file := "20231222_report_campaign_all.csv"
	// utils.WriteCsv(equal, file)
	// utils.WriteCsv(data1, "20231222_report_campaign_refuerzo.csv")
	// utils.SaveData(file)
	// utils.SaveToXlsx(file, "20231222_report_campaign_all.xlsx")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "1", "500", "enrollments1")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "2", "500", "enrollments2")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "3", "500", "enrollments3")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "4", "500", "enrollments4")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "5", "500", "enrollments5")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "6", "500", "enrollments6")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "7", "500", "enrollments7")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "8", "500", "enrollments8")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "9", "500", "enrollments9")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "10", "500", "enrollments10")
	// utils.GetApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "11", "500", "enrollments11")
	// utils.GetUsersKb4("https://us.api.knowbe4.com/v1/users", "500", "users", 2)
	// utils.GetEnrollmentApiKb4("https://us.api.knowbe4.com/v1/training/enrollments", "users", "campaign_test")
	utils.GetEnrollmentsKb4("https://us.api.knowbe4.com/v1/training/enrollments", "500", "campaign2", 1)
}