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

func AddUser(user_id int, user_name, password string) error{
	log.Println("User:", user_id)
	query := "INSERT INTO users (user_id, username, password) VALUES (?, ?, ?)"

	_, err := config.DB.Exec(query, user_id, user_name, password)
	if err != nil {
  		return err
	}
	return nil
}

func UserLogin(username string) (int, string, error) {
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

func UpdateUsername(userID int, newUsername models.UpdateUser) error {
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

func DeleteUser(userID int) error {
	query := "DELETE FROM users WHERE user_id = ?";
	result, err := config.DB.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated, user not found")
	}
	return nil
}

func GetAllUsers() (*sql.Rows, error)  {
	query:= "SELECT * FROM users"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil ,err
  	}
	return rows, nil
}