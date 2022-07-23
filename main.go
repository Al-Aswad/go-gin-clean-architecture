package main

import "go-gin-clean-architecture/app/config"

func main() {
	db, err := config.DBConnect().DB()

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

}
