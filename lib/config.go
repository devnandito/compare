package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Name string
}

func GetEnv() (conf *Config) {
	if _, err := os.Stat("lib/.env"); err == nil {
		err := godotenv.Load("lib/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbName := os.Getenv("DB_NAME")

	data := &Config {
		Name: dbName,
	}
	return data
}

func NewConfig() *Config {
	conf := GetEnv()
	return &Config {
		Name: conf.Name,
	}
}

func (c *Config) DnsStringGorm() (conn *gorm.DB) {
	dns := fmt.Sprintf("%s", c.Name)
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}