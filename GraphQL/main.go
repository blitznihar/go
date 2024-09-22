//TODO:simple graphsql api server fetch from mongodb

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// main is a simple web server that listens on port 8080 and responds
// to all HTTP requests with a "Hello, World!" message.
func main() {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Fetch data from MongoDB
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = client.Database("your-database-name").Collection("your-collection-name")
		// Perform your query or operations on the collection here
		// ...

		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
