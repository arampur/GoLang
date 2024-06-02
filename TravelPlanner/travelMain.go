package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type City struct {
	Name string
}

type AutocompleteData struct {
	Cities []City
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "3638"
	dbname   = "TravelPlanner"
)

var db *sql.DB
var tpl *template.Template

const createTableQuery = `
   CREATE TABLE IF NOT EXISTS users (
       id SERIAL PRIMARY KEY,
       email VARCHAR(100) UNIQUE NOT NULL,
       password VARCHAR(60) NOT NULL
   );
`

func init() {
	var err error

	// Connect to the PostgreSQL server
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Create the 'users' table if it doesn't exist
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize HTML templates
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/", register)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	port := getPort()
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		username := getUsernameFromEmail(email)
		fmt.Printf("Extracted username from email: %s\n", username)

		// Check if the email already exists
		var existingEmail string
		err := db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&existingEmail)
		if err == nil {
			// Email already exists, display alert and redirect to login page
			alertMessage := fmt.Sprintf("Email '%s' already exists. Please login or choose a different email.", email)
			tpl.ExecuteTemplate(w, "login.html", map[string]interface{}{"AlertMessage": alertMessage})
			return
		}

		// Hash the password before storing it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Store user data in the database
		_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, string(hashedPassword))
		if err != nil {
			log.Println("Error inserting user data:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Redirect to login page after successful registration
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "register.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		username := getUsernameFromEmail(email)
		fmt.Printf("Extracted username from email: %s\n", username)

		// Retrieve hashed password from the database
		var storedHashedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&storedHashedPassword)
		if err != nil {
			log.Println("Error querying user data:", err)
			// Display an alert for invalid email or password
			alertMessage := "Invalid email or password. Please check your email and password."
			tpl.ExecuteTemplate(w, "login.html", map[string]interface{}{"AlertMessage": alertMessage, "Username": username})
			return
		}

		// Compare the provided password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
		if err != nil {
			log.Println("Error comparing passwords:", err)
			log.Printf("Stored Hashed Password: %s\n", storedHashedPassword)
			log.Printf("Provided Password: %s\n", password)
			// Display an alert for invalid email or password
			alertMessage := "Invalid email or password. Please check your email and password."
			tpl.ExecuteTemplate(w, "login.html", map[string]interface{}{"AlertMessage": alertMessage, "Username": username})
			return
		}

		// Redirect to home page after successful login
		http.Redirect(w, r, "/welcome?username="+username, http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	//email := r.FormValue("email")
	username := r.URL.Query().Get("username")
	fmt.Printf("Extracted username from email inside welcome handler: %s\n", username)

	if term := r.URL.Query().Get("term"); term != "" {
		// Handle autocomplete request
		suggestions := getAutocompleteSuggestions(term)
		autocompleteData := AutocompleteData{Cities: suggestions}
		tpl.ExecuteTemplate(w, "autocomplete.html", autocompleteData)
		return
	}

	tpl.ExecuteTemplate(w, "welcome.html", map[string]interface{}{"Username": username})
}

func getUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	fmt.Println(parts)
	if len(parts) > 0 {
		username := parts[0]
		fmt.Printf("Input Email: %s, Output Username: %s\n", email, username) // Add this line for debugging
		return username
	}
	return email
}

func getAutocompleteSuggestions(term string) []City {
	// Modify this query based on your database schema
	query := "SELECT name FROM cities WHERE LOWER(name) LIKE LOWER($1) LIMIT 10"
	rows, err := db.Query(query, "%"+term+"%")
	if err != nil {
		log.Println("Error querying autocomplete suggestions:", err)
		return nil
	}
	defer rows.Close()

	var suggestions []City
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.Name); err != nil {
			log.Println("Error scanning autocomplete suggestion:", err)
			continue
		}
		suggestions = append(suggestions, city)
	}

	return suggestions
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}
	return port
}
