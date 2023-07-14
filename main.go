package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Austine05/todo-server/config"
	"github.com/Austine05/todo-server/controllers"
	"github.com/Austine05/todo-server/middleware"
	"github.com/gorilla/mux"
)

func main() {

	// Initialize the router
	router := mux.NewRouter()

	// Todo endpoints
	router.HandleFunc("/todos", middleware.Authenticate(controllers.CreateTodo)).Methods("POST")
	router.HandleFunc("/todos/{id}", middleware.Authenticate(controllers.UpdateTodo)).Methods("PUT")
	router.HandleFunc("/todos/{id}/deadline", middleware.Authenticate(controllers.UpdateTodoDeadline)).Methods("PUT")
	router.HandleFunc("/todos/{id}/delete", middleware.Authenticate(controllers.DeleteTodo)).Methods("DELETE")
	router.HandleFunc("/todos/{id}/status", middleware.Authenticate(controllers.UpdateTodoStatus)).Methods("PUT")
	router.HandleFunc("/todos", middleware.Authenticate(controllers.ListTodos)).Methods("GET")

	// User endpoints
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/users/{id}", middleware.Authenticate(controllers.UpdateUser)).Methods("PUT")
	router.HandleFunc("/users/{id}/password", middleware.Authenticate(controllers.UpdatePassword)).Methods("PUT")

	// Starting the server
	addr := config.Cfg.ServerHost + ":" + config.Cfg.ServerPort
	log.Println("server started on", addr)
	fmt.Println("server started on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
