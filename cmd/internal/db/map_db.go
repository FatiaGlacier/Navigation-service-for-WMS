package db

import (
	"github.com/FatiaGlacier/navigation-service/cmd/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

var Neo4JDrivarer neo4j.DriverWithContext

func InitNeo4j() {
	uri := config.GetEnv("NEO4J_URI", "bolt://localhost:7687")
	username := config.GetEnv("NEO4J_USERNAME", "bolt")
	password := config.GetEnv("NEO4J_PASSWORD", "bolt")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}

	Neo4JDrivarer = driver
}
