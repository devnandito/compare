package models

import (
	"github.com/devnandito/compare/lib"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	Fullname string
	Module1 string
	Module2 string
	Module3 string
	Module4 string
	Module5 string
}

func(u User) CreateUser(usr *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DnsStringGorm()
	db.AutoMigrate(&User{})
	response := db.Create(&usr)
	data := User {
		Email: usr.Email,
		Fullname: usr.Fullname,
		Module1: usr.Module1,
		Module2: usr.Module2,
		Module3: usr.Module3,
		Module4: usr.Module4,
		Module5: usr.Module5,
	}

	return data, response.Error
}

func (u User) UpdateUser(email string, data *User, col string) (usr User, err error) {
	conn := lib.NewConfig()
	db := conn.DnsStringGorm()
	db.AutoMigrate(&User{})
	if col == "module1" {
		res, err := u.GetUser(email)
		if err != nil {
			panic(err)
		}
		mod1 := res.Module1 + " - " + data.Module1
		response := db.Model(&u).Where("email = ?", email).Updates(User{Module1: mod1,})
		err = response.Error
	}
	
	if col == "module2" {
		res, err := u.GetUser(email)
		if err != nil {
			panic(err)
		}
		mod2 := res.Module2 + " - " + data.Module2
		response := db.Model(&u).Where("email = ?", email).Updates(User{Module2: mod2,})
		err = response.Error
	}

	if col == "module3" {
		res, err := u.GetUser(email)
		if err != nil {
			panic(err)
		}
		mod3 := res.Module3 + " - " + data.Module3
		response := db.Model(&u).Where("email = ?", email).Updates(User{Module3: mod3,})
		err = response.Error
	}

	if col == "module4" {
		res, err := u.GetUser(email)
		if err != nil {
			panic(err)
		}
		mod4 := res.Module4 + " - " + data.Module4
		response := db.Model(&u).Where("email = ?", email).Updates(User{Module4: mod4,})
		err = response.Error
	}

	if col == "module5" {
		res, err := u.GetUser(email)
		if err != nil {
			panic(err)
		}
		mod5 := res.Module5 + " - " + data.Module5
		response := db.Model(&u).Where("email = ?", email).Updates(User{Module5: mod5,})
		err = response.Error
	}

	return u, err
}

func (u User) ShowUsers() ([]User, error) {
	conn := lib.NewConfig()
	db := conn.DnsStringGorm()
	var objects []User
	response := db.Find(&objects)
	return objects, response.Error
}

func (u User) GetUser(email string) (User, error) {
	conn := lib.NewConfig()
	db := conn.DnsStringGorm()
	db.AutoMigrate(&User{})
	response := db.Find(&u, "email = ?", email)
	
	return u, response.Error
}

func (u User) CheckUser(email string) (User, error) {
	conn := lib.NewConfig()
	db := conn.DnsStringGorm()
	db.AutoMigrate(&User{})
	response := db.Find(&u, "email = ?", email)
	return u, response.Error
}