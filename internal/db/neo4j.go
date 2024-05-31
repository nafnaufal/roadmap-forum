package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func ConnectToNeo4j() (neo4j.Session, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	uri := os.Getenv("NEO4J_URI")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("NEO4J_PASSWORD")

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	session := driver.NewSession(neo4j.SessionConfig{})
	return session, nil
}

func RunCypherQuery(session neo4j.Session, query string, params map[string]interface{}) (neo4j.Result, error) {
	result, err := session.Run(query, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CloseSession(session neo4j.Session) {
	session.Close()
}
