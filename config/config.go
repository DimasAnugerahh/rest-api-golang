package config

import (
	"Rest-api-golang/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	godotenv.Load(".env")
	var res = new(Config)

	if val, found := os.LookupEnv("DBPORT"); found {
		res.DB_Port = val
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DB_Host = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DB_Username = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DB_Password = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DB_Name = val
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		res.DB_Username,
		res.DB_Password,
		res.DB_Host,
		res.DB_Port,
		res.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{})
}
