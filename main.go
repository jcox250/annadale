package main

import (
	"log"
	"os"
)

var DB_TYPE = os.Getenv("DB_TYPE")
var DB_URL = os.Getenv("DATABASE_URL")
var PORT = os.Getenv("PORT")

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
