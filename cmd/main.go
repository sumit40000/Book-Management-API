package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sumit40000/book-management-api/internal/config"
	"github.com/sumit40000/book-management-api/internal/controllers"
	"github.com/sumit40000/book-management-api/internal/routes"
)

func startServer(){
	r:= routes.CreateRoutes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

func main() {
	fmt.Println("welcome to the book management api")
	err:= godotenv.Load()
	if err!= nil{
		log.Fatal("failed starting , error : error to load the env file", err)
	}

	// connecting the database
	config.ConnectDb()
	controllers.InitCollection()

	// starting the server
	fmt.Println("Starting the server")
	fmt.Println("server started successfully, listening on port 8000")
	startServer()
}