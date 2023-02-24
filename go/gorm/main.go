package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Device struct {
	Code string
}

func main() {
	dsn := "host=localhost user=db-whoosh password=db-whoosh dbname=db-whoosh sslmode=disable TimeZone=Asia/Yekaterinburg"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var result []Device
	db.Raw("SELECT code from denormalized_search.device where is_online").Scan(&result)
	fmt.Printf("%+v\n", result)
}
