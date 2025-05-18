package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	dbhandler "github.com/richmont/GoDemoDBClient/DbHandler"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port_str := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	port, errPort := strconv.ParseInt(port_str, 10, 64)

	if errPort == nil {
		var m = dbhandler.NewMariadbHandler(
			user,
			passwd,
			hostname,
			port,
			database,
		)
		var err = m.Connect()
		if err != nil {
			fmt.Println(err)
		} else {
			m.ShowTables()
		}
	} else {
		log.Panic("Failed to get port from env")
	}

}
