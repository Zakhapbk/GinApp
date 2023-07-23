package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine
var DB *sql.DB

func main() {
	ConnectStr := "user = postgres password = Xpolkwfv13 dbname = Gin sslmode = disable" // change to the relevant value
	db, err := sql.Open("postgres", ConnectStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	DB = db
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", MainPage)
	router.GET("/emp", List)
	router.GET("/emp/:name", Employee)
	router.POST("/emp/:name", DeleteEmployee)
	router.GET("/new", AddEmployeeShow)
	router.POST("/new", AddEmployeePost)
	router.Run()

}
