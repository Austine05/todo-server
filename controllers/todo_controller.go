package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Austine05/todo-server/models"
)

// POST /todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
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

	var todo models.Todo
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
func ChangeDeadline(w http.ResponseWriter, r *http.Request) {
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
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
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
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
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
func ListTodos(w http.ResponseWriter, r *http.Request) {
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
