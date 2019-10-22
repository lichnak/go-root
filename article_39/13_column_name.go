package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Customer struct {
	Id      int `gorm:"Column:ID"`
	Name    string
	Surname string
	Address string
	Country string
	Phone   string
}

func main() {
	dbDriver := flag.String("dbdriver", "sqlite3", "database driver specification")
	storageSpecification := flag.String("storage", "./test.db", "storage specification")
	flag.Parse()

	db, err := gorm.Open(*dbDriver, *storageSpecification)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Customer{})

	var customers []Customer
	db.Find(&customers)
	for _, customer := range customers {
		fmt.Printf("%2d %-10s %-10s %-12s %-12s %s\n", customer.Id, customer.Name, customer.Surname, customer.Address, customer.Country, customer.Phone)
	}
}