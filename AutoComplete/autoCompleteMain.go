package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	Email    string
	Password string
}

func main() {

	fmt.Println("Server is running on http://localhost:9000")

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/welcome", indexHandler)
	http.HandleFunc("/", registerHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cityName := r.FormValue("cityName")
		if cityName != "" {
			fmt.Println(cityName)
			cities, err := fetchCities(cityName)
			if err != nil {
				http.Error(w, "Error fetching cities from the database", http.StatusInternalServerError)
				return
			}

			citiesJSON, _ := json.Marshal(cities)
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(citiesJSON)
			return
		}
	}

	renderTemplate(w, "welcome", nil)
}

func fetchCities(prefix string) ([]string, error) {
	rows, err := db.Query("SELECT name FROM cities WHERE country_code = 'US' AND name ILIKE $1 LIMIT 10", prefix+"%")
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

func fetchUser(email string) *User {
	var user User
	err := db.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil
	}
	return &user
}

func saveUser(user User) error {
	_, err := db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func comparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Simple email validation
		if !isValidEmail(email) {
			renderTemplate(w, "register", "Invalid email address")
			return
		}

		// Check if user with provided email already exists
		if userExists(email) {
			renderTemplate(w, "register", "User with this email already exists")
			return
		}

		hashedPassword, err := hashPassword(password)
		if err != nil {
			renderTemplate(w, "register", "Error hashing the password")
			return
		}

		// Save the user with hashed password to the database
		saveUser(User{Email: email, Password: hashedPassword})

		// Redirect to login page after successful registration
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "register", nil)
}

func isValidEmail(email string) bool {
	// Add your email validation logic here
	// This is a simple example; you might want to use a more robust validation library
	return len(email) > 0 && strings.Contains(email, "@") && strings.Contains(email, ".")
}

func userExists(email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Println("Error checking if user exists:", err)
		return false
	}
	return count > 0
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := fetchUser(email)
		if user != nil && comparePasswords(user.Password, password) {
			// Successful login, redirect or perform necessary actions
			// For now, we'll just log a message
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
			return
		} else {
			// Invalid login, show an error message
			renderTemplate(w, "login", "Invalid login credentials")
			return
		}
	}

	renderTemplate(w, "login", nil)
}

func renderTemplate(w http.ResponseWriter, templateName string, errorMessage interface{}) {
	const loginTemplate = `
		<!-- Your login template content with Bootstrap -->
		<!DOCTYPE html>
<html>
<head>
    <title>Login</title>
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
</head>
<body>
    <div class="container mt-5">
        {{ if .ErrorMessage }}
            <div class="alert alert-danger" role="alert">
                {{ .ErrorMessage }}
            </div>
        {{ end }}
        <h1>Login</h1>
        <form method="post" action="/login">
            <label for="email">Email:</label>
            <input type="text" id="email" name="email" required>
            <br>
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required>
            <br>
            <button type="submit" class="btn btn-primary">Login</button>
        </form>
    </div>
    <!-- Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</body>
</html>
`

	const registerTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Registration</title>
			<!-- Bootstrap CSS -->
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
		</head>
		<body>
			<div class="container mt-5">
				{{ if .ErrorMessage }}
					<div class="alert alert-danger" role="alert">
						{{ .ErrorMessage }}
					</div>
				{{ end }}
				<h1>Registration</h1>
				<form method="post" action="/register">
					<label for="email">Email:</label>
					<input type="text" id="email" name="email" required>
					<br>
					<label for="password">Password:</label>
					<input type="password" id="password" name="password" required>
					<br>
					<button type="submit" class="btn btn-primary">Register</button>
				</form>
				<p class="mt-3">Already have an account? <a href="/login">Login</a></p>
			</div>
			<!-- Bootstrap JS -->
			<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js"></script>
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
		</body>
		</html>
	`

	const indexTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>City Autocomplete</title>
			<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
			<script src="https://code.jquery.com/ui/1.12.1/jquery-ui.js"></script>
			<link rel="stylesheet" href="https://code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
			<!-- Bootstrap CSS -->
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
		</head>
		<body>
			<h1>City Autocomplete</h1>
			<form method="post" action="/welcome">
				<label for="cityName">Enter City Name:</label>
				<input type="text" id="cityName" name="cityName" autocomplete="off">
			</form>
			<div id="cityList"></div>

			<script>
			$(function () {
				$("#cityName").autocomplete({
					source: function (request, response) {
						$.ajax({
							url: "/welcome",
							type: "POST",
							data: { cityName: request.term },
							success: function (data) {
								response(data);
							}
						});
					},
					minLength: 2,
					select: function (event, ui) {
						console.log(ui.item.value);
					}
				});
			});			
			</script>
			<!-- Bootstrap JS -->
			<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js"></script>
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
		</body>
		</html>
	`

	templateText := ""
	switch templateName {
	case "login":
		templateText = loginTemplate
	case "register":
		templateText = registerTemplate
	case "welcome":
		templateText = indexTemplate
	default:
		http.Error(w, "Invalid template name", http.StatusBadRequest)
		return
	}

	tmpl, err := template.New(templateName).Parse(templateText)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	templateData := struct {
		Title        string
		Content      interface{}
		ErrorMessage interface{}
	}{
		Title:        templateName,
		Content:      nil,
		ErrorMessage: errorMessage,
	}

	if err := tmpl.Execute(w, templateData); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
