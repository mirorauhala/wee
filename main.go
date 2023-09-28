package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func setupRoutes(dbConnection *sql.DB) {
	// server current time on /time
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The time is %s", time.Now().Format(time.RFC1123))
	})

	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {

		id, err := id()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error generating id")
			return
		}

		fmt.Fprintf(w, "The id is %s", id)
	})

	// fetch value from database
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only POST requests are allowed.")
			return
		}
		value, err := getKey(dbConnection, "foo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "The value of foo is %s", value)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

}

func main() { // setup database
	dbConnection, err := setupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	setupRoutes(dbConnection)

	// listen on port 8080
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
