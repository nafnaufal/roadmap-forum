package db

import (
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var driver neo4j.Driver

func InitDB() {
	uri := os.Getenv("NEO4J_URI")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("NEO4J_PASSWORD")
	var err error

	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}
}

func GetDriver() neo4j.Driver {
	return driver
}

func CloseDB() {
	if driver != nil {
		driver.Close()
	}
}
