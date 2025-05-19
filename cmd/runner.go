package cmd

import (
	"fmt"
	"graduation-system/cmd/server"
	"graduation-system/database/constructor"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	log.Println("Starting server...")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read environment variables
	serverPort := os.Getenv("SERVER_PORT")

	// Read database environment variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// Initialize database
	err = constructor.InitDB(dsn)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// Initialize log file
	initLogFile("../logs/server.log")

	server.RunDBHttpServer(serverPort)
}

func initLogFile(logFilePath string) {
	LogFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("(Error) : error opening the log file : %v", err)
	}
	log.SetOutput(LogFile)
	log.Printf("Log output is set to %s", logFilePath)
}
