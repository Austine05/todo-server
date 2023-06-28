package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Todo struct, this represents a TODO item
type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Deadline    time.Time          `json:"deadline,omitempty" bson:"deadline,omitempty"`
}

var collection *mongo.Collection

func main() {

	// MongoDB connection
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing MongoDB collection
	collection = client.Database("todosDB").Collection("todos")

	// Initialize the router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}/deadline", changeDeadline).Methods("PUT")
	router.HandleFunc("/todos/{id}/delete", deleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}/status", updateStatus).Methods("PUT")
	router.HandleFunc("/todos", listTodos).Methods("GET")

	// Starting the server
	fmt.Println("server started at port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Helper function to write JSON response
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// POST /todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	result, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)
	respondWithJSON(w, http.StatusCreated, "todo created successfully")
}

// PUT /todos/{id}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": todo})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	todo.ID = id
	respondWithJSON(w, http.StatusOK, todo)
}

// PUT /todos/{id}/deadline
func changeDeadline(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var deadline struct {
		Deadline time.Time `json:"deadline"`
	}
	json.NewDecoder(r.Body).Decode(&deadline)

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"deadline": deadline.Deadline}})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Deadline changed successfully")
}

// DELETE /todos/{id}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Todo deleted successfully")
}

// PUT /todos/{id}/status
func updateStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var status struct {
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&status)

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status.Status}})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Status updated successfully")
}

// GET /todos
func listTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	respondWithJSON(w, http.StatusOK, todos)
}
