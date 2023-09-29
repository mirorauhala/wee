package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func setupRoutes(db *sql.DB) {

	http.HandleFunc("/api/url/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only POST requests are allowed.")
			return
		}

		// parse the body of the request
		err1 := r.ParseForm()
		if err1 != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unable to parse form.")
			return
		}

		// Access form values
		requestUrl := r.Form.Get("url")

		id, err2 := sid.Generate()

		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unable to generate a short id.")
			fmt.Println(err2)
			return
		}

		var url = ShortenedUrl{
			ID:        id,
			URL:       requestUrl,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		var resultUrl = ShortenedUrl{}

		err3 := db.QueryRow("INSERT INTO urls (id, url, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, url, created_at, updated_at", url.ID, url.URL, url.CreatedAt, url.UpdatedAt).Scan(&resultUrl.ID, &resultUrl.URL, &resultUrl.CreatedAt, &resultUrl.UpdatedAt)

		if err3 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unable to insert into database.")
			fmt.Println("Error inserting into database: ", err3)
			return
		}

		jsonData, err := json.Marshal(resultUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unable to marshal json.")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	})

	http.HandleFunc("/api/url", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only GET requests are allowed.")
			return
		}

		// return all urls as json
		// Query all rows from the "users" table
		rows, err := db.Query("SELECT * FROM urls")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Loop through the result set and retrieve the data
		var urls []ShortenedUrl
		for rows.Next() {
			var url ShortenedUrl
			err := rows.Scan(&url.ID, &url.URL, &url.CreatedAt, &url.UpdatedAt)
			if err != nil {
				log.Fatal(err)
			}
			urls = append(urls, url)
		}

		// Handle any errors encountered during iteration
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		// Do something with the retrieved rows
		for _, user := range urls {
			fmt.Println(user)
		}

		// turn urls into json
		w.Header().Set("Content-Type", "application/json")
		err2 := json.NewEncoder(w).Encode(urls)

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}

	})

	http.HandleFunc("/api/follow-url/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Only GET requests are allowed.")
			return
		}

		urlParts := strings.Split(r.URL.Path, "/")
		id := urlParts[len(urlParts)-1]

		http.Redirect(w, r, "https://www.google.com/"+id, http.StatusTemporaryRedirect)
		w.WriteHeader(http.StatusOK)
	})

}

func main() {
	SetupShortId()
	dbConnection, err := setupDatabase()

	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	setupRoutes(dbConnection)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
