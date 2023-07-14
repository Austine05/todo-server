package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Austine05/todo-server/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection

func init() {
	todoCollection = dbClient.Database("TodosDB").Collection("todo")
}

// Create a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = primitive.NewObjectID()
	_, err := todoCollection.InsertOne(context.TODO(), newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newTodo)
}

// Update an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	updatedTodo.ID = todoID

	filter := bson.M{"_id": todoID}
	update := bson.M{"$set": bson.M{
		"name":        updatedTodo.Name,
		"description": updatedTodo.Description,
	}}
	_, err = todoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// Update the status of a todo (mark as done/undone)
func UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	updatedTodo.ID = todoID

	filter := bson.M{"_id": todoID}
	update := bson.M{"$set": bson.M{
		"status": updatedTodo.Status,
	}}
	_, err = todoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// Update the deadline of a todo
func UpdateTodoDeadline(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	updatedTodo.ID = todoID

	filter := bson.M{"_id": todoID}
	update := bson.M{"$set": bson.M{
		"deadline": updatedTodo.Deadline,
	}}
	_, err = todoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// List all todos
func ListTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	cursor, err := todoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

// Delete a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": todoID}
	_, err = todoCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
