package queries

import (
	"Codimite_Assignment/config"
	"Codimite_Assignment/internal/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// Check if product exists
func ProductExists(p_id int) (bool, error) {
	var count int
	log.Println("count:", count)
	query := "SELECT COUNT(*) FROM products WHERE p_id = ?"
	err := config.DB.QueryRow(query, p_id).Scan(&count)
	if err != nil {
		log.Println("Error checking product existence:", err)
		return false, err
	}
	log.Println("count:", count)
	return count > 0, nil
}

func AddProduct(name string, p_type string, price int, quantity int) error{
	log.Println("Product:", name)
	query := "INSERT INTO products (name, p_type, price, quantity) VALUES (?, ?, ?, ?)"

	_, err := config.DB.Exec(query, name, p_type, price, quantity)
	if err != nil {
  		return err
	}
	return nil
}

func UpdateProduct(p_id int, updatedProduct models.UpdateProduct) error {
	log.Println("Executing Update for Product ID:", p_id)
	
	// Check if product exists
	exists, err := ProductExists(p_id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("product not found")
	}

	// Create dynamic update fields
	var updateFields []string
	var args []interface{}

	log.Println("err:", err)
	if updatedProduct.Name != "" {
		updateFields = append(updateFields, "name = ?")
		args = append(args, updatedProduct.Name)
	}
	if updatedProduct.PType != "" {
		updateFields = append(updateFields, "p_type = ?")
		args = append(args, updatedProduct.PType)
	}
	if updatedProduct.Price > 0 {
		updateFields = append(updateFields, "price = ?")
		args = append(args, updatedProduct.Price)
	}
	if updatedProduct.Quantity > 0 {
		updateFields = append(updateFields, "quantity = ?")
		args = append(args, updatedProduct.Quantity)
	}

	// If no fields to update, return early
	if len(updateFields) == 0 {
		log.Println("No fields provided for update.")
		return fmt.Errorf("no fields to update")
	}

	// Construct SQL query
	query := fmt.Sprintf("UPDATE products SET %s WHERE p_id = ?", strings.Join(updateFields, ", "))
	args = append(args, p_id)

	// Execute query
	_, err = config.DB.Exec(query, args...)
	if err != nil {
		log.Println("Error executing update query:", err)
		return err
	}

	log.Println("Product updated successfully:", p_id)
	return nil
}

func DeleteProduct(p_id int) error {
	// Check if product exists
	exists, err := ProductExists(p_id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("product not found")
	}
	
	query := "DELETE FROM products WHERE p_id = ?";
	result, err := config.DB.Exec(query, p_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, product not found")
	}
	return nil
}

func GetAllProducts() (*sql.Rows, error)  {
	query:= "SELECT * FROM products"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil ,err
  	}
	return rows, nil
}