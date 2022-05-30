package main

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	Id       int64  `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Age      int64
	Language string `gorm:"size:255"`
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "db"
)

var db *gorm.DB

func main() {

	u := url.URL{
		Scheme: "auth",
		User:   url.UserPassword("postgres", "postgres"),
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "postgres",
	}

	conn := u.String()
	fmt.Println(conn)
	var err error
	db, err = gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create table
	db.Debug().AutoMigrate(&Employee{})

	// Insert value
	db.Create(&Employee{Id: 1, Name: "John", Age: 35, Language: "Golang-GORM"})
	db.Create(&Employee{Id: 2, Name: "Smith", Age: 24, Language: "Golang-GORM"})

	// Display input data
	var employees []Employee
	db.Find(&employees)
	for _, employee := range employees {
		fmt.Printf("Employee ID:%d\nName:%s\nAge:%d\nLanguage:%s\n", employee.Id, employee.Name, employee.Age, employee.Language)
		fmt.Printf("--------------------------------------------------------------\n")
	}

}
