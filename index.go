package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Blank import for PostgreSQL driver
)

var db *sql.DB

// Initialize the database connection
func initDB() {
	var err error
	connectionstring := "postgres://postgres:password@localhost:5440/mydb?sslmode=disable" // Correct connection string
	db, err = sql.Open("postgres", connectionstring)
	if err != nil {
		log.Fatal("Error in connection: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	fmt.Println("Database connection is successful")
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func createUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	sqlStatement := `INSERT INTO users(name, email, age) VALUES($1, $2, $3) RETURNING id`
	err = db.QueryRow(sqlStatement, user.Name, user.Email, user.Age).Scan(&user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, user)
}
func createusers(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "error while creating the user"})

		return
	}
	sqlstatement := `INSERT INTO users (name,email,age)VALUES($1,$2,$3) RETURNING id`
	err = db.QueryRow(sqlstatement, user.Name, user.Email, user.Age).Scan(&user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "user cant't be created"})
		return
	}
	c.JSON(201, user)

}

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	// Create a Gin router
	r := gin.Default()

	// Define the GET route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is up and running!"})
	})
	r.POST("/users", createusers)

	// Start the server on port 8080
	r.Run(":8080")
}
