package queries

import (
	"Codimite_Assignment/config"
	"Codimite_Assignment/internal/models"
	"database/sql"
	"fmt"
	"log"
)

// Check if a product exists in the database
func UserExists(user_id int) (bool, error) {
	var count int
	log.Println("count:", count)
	query := "SELECT COUNT(*) FROM users WHERE user_id = ?"
	err := config.DB.QueryRow(query, user_id).Scan(&count)
	if err != nil {
		log.Println("Error checking user existence:", err)
		return false, err
	}
	log.Println("count:", count)
	return count > 0, nil
}

//create user
func AddUser(user_id int, user_name, password string) error{
	log.Println("User:", user_id)
	// SQL query to insert a new user
	query := "INSERT INTO users (user_id, username, password) VALUES (?, ?, ?)"

	_, err := config.DB.Exec(query, user_id, user_name, password)
	if err != nil {
  		return err
	}
	return nil
}

// Log in a user by validating their username and password
func UserLogin(username string) (int, string, error) {
	
	// SQL query to get user_id and password for a given username
	query := "SELECT user_id, password FROM users WHERE username = ?"

	var user_id int
	var password string

	err := config.DB.QueryRow(query, username).Scan(&user_id, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found")
			return 0, "", nil
		}
		return 0, "", err
	}

	return user_id, password, nil
}

//update username
func UpdateUsername(userID int, newUsername models.UpdateUser) error {
	
	// Check if user exists
	exists, err := UserExists(userID)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	// sql query to update the username for a given user ID
	query := "UPDATE users SET username = ? WHERE user_id = ?"

	result, err := config.DB.Exec(query, newUsername.NewUserName, userID)
	if err != nil {
		return err
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated, user not found")
	}

	return nil
}

//delete user
func DeleteUser(userID int) error {

	// Check if user exists
	exists, err := UserExists(userID)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	// SQL query to delete a user
	query := "DELETE FROM users WHERE user_id = ?";
	result, err := config.DB.Exec(query, userID)
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated, user not found")
	}
	return nil
}

// get all users
func GetAllUsers() (*sql.Rows, error)  {

	// SQL query to get all users
	query:= "SELECT * FROM users"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil ,err
  	}
	return rows, nil
}