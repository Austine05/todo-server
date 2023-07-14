package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Austine05/todo-server/database"
	"github.com/Austine05/todo-server/config"
	"github.com/Austine05/todo-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
)

var userCollection *mongo.Collection

func init() {
	userCollection = database.DBClient.Database(config.GetConfig().MongoDBName).Collection("user")
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Register a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = primitive.NewObjectID()

	_, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login a user and generate a JWT token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&credentials)

	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if user.Password != credentials.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(response)
}

// Update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	claims, ok := r.Context().Value("claims").(*Claims)
	if !ok || claims.Username != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{
		"firstName": updatedUser.FirstName,
		"lastName":  updatedUser.LastName,
		"email":     updatedUser.Email,
		"mobile":    updatedUser.Mobile,
	}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Update the password of a user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	claims, ok := r.Context().Value("claims").(*Claims)
	if !ok || claims.Username != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{
		"password": updatedUser.Password,
	}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
