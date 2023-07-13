package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Austine05/todo-server/config"
	"github.com/gorrila/mux"
	"github.com/joho/godotenv"
	"gitub.com/Austine05/todo-server/controllers"
	"gitub.com/Austine05/todo-server/middleware"
)

func main() {

	cfg := config.LoadConfig()

	//Middleware
	router.Use(middleware.AuthMiddleware)

	// Initialize the router
	router := mux.NewRouter()

	// Todo endpoints
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}/deadline", controllers.ChangeDeadline).Methods("PUT")
	router.HandleFunc("/todos/{id}/delete", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}/status", controllers.UpdateOnepdateStatus).Methods("PUT")
	router.HandleFunc("/todos", controllers.ListTodos).Methods("GET")

	// User endpoints
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}/password", controllers.UpdatePassword).Methods("PUT")

	// Starting the server
	fmt.Println("server started at port :8000")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
