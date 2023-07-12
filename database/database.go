package database

import (
	"context"
	"fmt"
	"log"
	"poll-app/ent"

	_ "github.com/lib/pq"
)

var EntClient *ent.Client

func init() {
	//Open a connection to the database
	Client, err := ent.Open("postgres", "host=localhost port=5432 user=kartikay.kaushik dbname=kartikay.kaushik sslmode=disable")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to connect to database")

	}

	fmt.Println("Connected to database successfully")
	// AutoMigration with ENT
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EntClient = Client
}
