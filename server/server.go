package server

import (
	"log"
	"os"
)

// InitializeServer --
func InitializeServer() {
	e := NewRouter()
	port := os.Getenv("APP_PORT")
	log.Print(port)
	e.Logger.Fatal(e.Start(":" + port))
}
