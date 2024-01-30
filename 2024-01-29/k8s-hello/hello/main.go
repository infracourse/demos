package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func renderGreetingsPage(c *gin.Context, name string) {
	// Generate the cat image URL with the user's name
	catImageUrl := fmt.Sprintf("https://cataas.com/cat/says/Hello%%20%s", name)

	// Render the greeting HTML page
	c.HTML(http.StatusOK, "greeting.html", gin.H{
		"name":     name,
		"catImage": catImageUrl,
		"greeting": fmt.Sprintf("Hello, %s!", name),
	})
}

func NewUser(c *gin.Context) {
	user := User{Name: c.PostForm("name")}

	// Create a new user record in the database
	db.Create(&user)

	// Render the greetings page
	renderGreetingsPage(c, user.Name)
}

func GreetPage(c *gin.Context) {
	name := c.Query("name")

	if db.Where("name = ?", name).First(&User{}).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Render the greetings page
	renderGreetingsPage(c, name)
}

func MainPage(c *gin.Context) {
	var users []User
	db.Find(&users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

func initDB() {
	var err error
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}

func main() {
	initDB()

	router := gin.Default()
	router.GET("/", MainPage)
	router.GET("/greet", GreetPage)
	router.POST("/user", NewUser)
	router.LoadHTMLGlob("templates/*")

	router.Run(":8080")
}
