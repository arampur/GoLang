// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// Replace these constants with your own database credentials
const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "3638"
	DBName     = "TravelPlanner"
)

var db *sql.DB
var err error

// User model represents the structure of the 'users' table in the database
type User struct {
	gorm.Model
	Email          string `gorm:"unique_index"`
	Password       string
	HashedPassword string
}

type City struct {
	gorm.Model
	Name string `gorm:"soft_delete: false"`
}

// Initialize database connection
func initDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect to database")
	}
	return db, nil
}

// Setup web server
func main() {

	var err error
	db, err := initDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")
	router.Static("/images", "./static/images")

	// Routes
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/", rootHandler)
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)
	router.GET("/welcome", welcomeHandler)

	router.Run(":9000")
}

// Root handler redirects to the login page
func rootHandler(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/register")
}

// Register a new user
func registerHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Simple email validation
	if email == "" || password == "" {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Email and password are required"})
		return
	}

	// Check if the user with the provided email already exists
	var existingUser User
	if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&existingUser.ID, &existingUser.Email, &existingUser.Password, &existingUser.HashedPassword); err == nil {
		// User with the provided email already exists
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "User with this email already exists"})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Failed to hash password"})
		return
	}

	// Create a new user in the database
	_, err = db.Exec("INSERT INTO users (email, password, hashed_password) VALUES ($1, $2, $3)", email, password, string(hashedPassword))
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Failed to create user"})
		return
	}

	// Redirect to the login screen after successful registration
	c.Redirect(http.StatusSeeOther, "/login")
}

// Handle user login
func loginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user User
	if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.HashedPassword); err != nil {
		fmt.Println("Error retrieving user:", err)
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Invalid email or password"})
		return
	}

	// Compare the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Invalid email or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusSeeOther, "/welcome")
}

// Handle welcome screen
func welcomeHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	// Check if the user is logged in
	if userID == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Retrieve user details from the database
	var user User
	if err := db.QueryRow("SELECT * FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Email, &user.Password, &user.HashedPassword); err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	query := c.Query("query")
	limit := 5 // Set your desired limit here
	cityNames, err := fetchCityNames(db, query, limit)
	if err != nil {
		log.Println("Error fetching city names:", err)
		c.HTML(http.StatusOK, "welcome.html", gin.H{"user": user, "Error": "Failed to fetch city names"})
		return
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{"user": user, "cityNames": cityNames})
}

// Handle user logout
func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "/login")
}

// Function to fetch city names from the cities table based on user input
func fetchCityNames(db *sql.DB, query string, limit int) ([]string, error) {
	rows, err := db.Query("SELECT name FROM cities WHERE name ILIKE $1 LIMIT $2", "%"+query+"%", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cityNames []string
	for rows.Next() {
		var cityName string
		if err := rows.Scan(&cityName); err != nil {
			return nil, err
		}
		cityNames = append(cityNames, cityName)
	}

	return cityNames, nil
}
