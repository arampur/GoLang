package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "3638"
	dbname   = "TravelPlanner"
)

var db *sql.DB

func init() {
	var err error
	connStr := "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Server running on localhost:9000..")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cityName := r.FormValue("cityName")
		if cityName != "" {
			cities, err := fetchCities(cityName)
			if err != nil {
				http.Error(w, "Error fetching cities from the database", http.StatusInternalServerError)
				return
			}

			citiesJSON, _ := json.Marshal(cities)
			w.Header().Set("Content-Type", "application/json")
			w.Write(citiesJSON)
			return
		}
	}

	renderTemplate(w, r)
}

func fetchCities(prefix string) ([]string, error) {
	rows, err := db.Query("SELECT name FROM cities WHERE country_code = 'US' AND name ILIKE $1 LIMIT 5", prefix+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []string
	for rows.Next() {
		var cityName string
		if err := rows.Scan(&cityName); err != nil {
			return nil, err
		}
		cities = append(cities, cityName)
	}

	return cities, nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("index").Parse(`

	// `)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }

	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, "Error executing template", http.StatusInternalServerError)
	// 	return
	// }
	templatesDir, err := filepath.Abs("templates")
	if err != nil {
		http.Error(w, "Error getting templates directory path", http.StatusInternalServerError)
		return
	}

	indexFilePath := filepath.Join(templatesDir, "index.html")

	http.ServeFile(w, r, indexFilePath)
}

// func renderTemplate(w http.ResponseWriter) {
// 	tmpl, err := template.New("index").Parse(`
// 		<!DOCTYPE html>
// 		<html>
// 		<head>
// 			<title>City Autocomplete</title>
// 			<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
// 			<script src="https://code.jquery.com/ui/1.12.1/jquery-ui.js"></script>
// 			<link rel="stylesheet" href="https://code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
// 			<!-- Include Bootstrap CSS link here -->
// 			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
// 		</head>
// 		<body>
// 			<nav class="navbar navbar-dark bg-dark">
// 				<span class="navbar-brand mb-0 h1">TRAVEL PLANNER</span>
// 			</nav>

// 			<h1>City Autocomplete</h1>
// 			<form method="post" action="/">
// 				<label for="cityName">Enter City Name:</label>
// 				<input type="text" id="cityName" name="cityName" autocomplete="off">
// 			</form>
// 			<div id="cityList"></div>

// 			<script>
// 				$(function () {
// 					$("#cityName").autocomplete({
// 						source: function (request, response) {
// 							$.ajax({
// 								url: "/",
// 								type: "POST",
// 								data: { cityName: request.term },
// 								success: function (data) {
// 									response(data);
// 								}
// 							});
// 						},
// 						minLength: 2,
// 						select: function (event, ui) {
// 							console.log(ui.item.value);
// 						}
// 					});
// 				});
// 			</script>

// 			<!-- Include Bootstrap JS and Popper.js links here -->
// 			<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
// 			<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
// 			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
// 		</body>
// 		</html>
// 	`)
// 	if err != nil {
// 		http.Error(w, "Error rendering template", http.StatusInternalServerError)
// 		return
// 	}

// 	if err := tmpl.Execute(w, nil); err != nil {
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 		return
// 	}
// }
