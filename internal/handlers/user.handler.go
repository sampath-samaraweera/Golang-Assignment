package handlers

import (
	"Codimite_Assignment/internal/auth"
	"Codimite_Assignment/internal/models"
	"Codimite_Assignment/internal/queries"
	"Codimite_Assignment/pkg/middleware"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//register new user
func RegisterUser(w http.ResponseWriter, r *http.Request){

	log.Println("Register called")
	var user models.User

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}	
	
	// Validate required fields
	if user.UserName == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		log.Println("Validation error: UserName or Password missing")
		return
	}

	// Hash the password using bcrypt
	passWordInBytes, errPwd := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errPwd != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		log.Println("Error hashing password", errPwd)
		return
	}

	// Generate a unique User ID
	user.UserId = uuid.New().ClockSequence()

	// Call the query function to Insert user into database
	err = queries.AddUser(user.UserId, user.UserName, string(passWordInBytes))
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println("Error inserting user",err)
		return
	}
	log.Println("User Successfully Registered")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully Registered"))
}

func LoginUser(w http.ResponseWriter, r *http.Request)  {
	var user models.UserLogin
		
	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}	

	// Validate required fields
	if user.UserName == "" || user.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		log.Println("Validation error: Missing username or password")
		return
	}

	// Retrieve user credentials from the database
	user_id, password, errQuery := queries.UserLogin(user.UserName)
	if errQuery != nil {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		log.Println("Invalid username",errQuery)
		return
	}

	// Compare hashed password with given password
	err = bcrypt.CompareHashAndPassword([]byte(password),[]byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		log.Println("Invalid password",err)
		return
	}

	//Generate JWT token
	token, errToken := auth.GenerateToken(user_id, user.UserName)
	if errToken != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		log.Println("Error generating token",errToken)
		return
	}	
	
	// Set response headers and return the token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

//update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	// Extract user ID from the request context
	user_id := r.Context().Value(middleware.UserContextKey).(int)
	log.Println("user Id:",user_id)

	var newUserName models.UpdateUser
	
	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&newUserName)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	// Call the query function to update the user
	err = queries.UpdateUsername(user_id, newUserName)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Println("error",err)
		return
	}
	log.Println("User Successfully updated")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully Updated"))
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Extract user ID from the request context
	user_id, ok := r.Context().Value(middleware.UserContextKey).(int)
	log.Println("user Id:",user_id)

	if !ok {
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
		log.Println("Error extracting user ID from context")
		return
	}

	// Call the query function to delete the user
	err := queries.DeleteUser(user_id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		log.Println("error",err)
		return
	}
	log.Println("User Successfully deleted")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully Deleted"))
}

//get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// get all users from the database
	rows, err := queries.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	
	// Iterate over query results and populate users slice
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.UserName, &user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Set response headers and encode the users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}