package queries

import (
	"Codimite_Assignment/config"
	"Codimite_Assignment/internal/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Check if order exists
func OrderExists(order_id int) (bool, error) {
	var count int
	log.Println("count:", count)
	query := "SELECT COUNT(*) FROM orders WHERE order_id = ?"
	err := config.DB.QueryRow(query, order_id).Scan(&count)
	if err != nil {
		log.Println("Error checking order existence:", err)
		return false, err
	}
	log.Println("count:", count)
	return count > 0, nil
}

// Add a new order to the database
func AddOrder(user_id int, p_id int, ordered_quantity int) error {

	// Check if order exists
	exists, err := ProductExists(p_id)

	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("product not found")
	}
	var price int

	// get product price
	query := "SELECT price FROM products WHERE p_id = ?"
	err = config.DB.QueryRow(query, p_id).Scan(&price)
	if err != nil {
		return fmt.Errorf("error fetching product price: %v", err)
	}

	// Check if user exists
	exists, err = UserExists(user_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	// Calculate total price
	total_price := ordered_quantity * price
	log.Print("total price", total_price)

	now := time.Now()
	order_date := &now

	// Insert order into database
	query = "INSERT INTO orders (user_id, p_id, ordered_quantity, total_price, order_date) VALUES (?, ?, ?, ?, ?)"

	_, err = config.DB.Exec(query, user_id, p_id, ordered_quantity, total_price, order_date)
	if err != nil {
		return err
	}
	log.Println("order successfully added")
	return nil
}

//update order
func UpdateOrder(order_id int, updatedOrder models.CreateOrder) error {
    log.Println(" order ID:")

	// Check if order exists
    exists, err := OrderExists(order_id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("order not found")
    }

	// Check if product exists
    exists, err = ProductExists(updatedOrder.PID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("product not found")
    }

    var price int
	// get product price
    query := "SELECT price FROM products WHERE p_id = ?"
    err = config.DB.QueryRow(query, updatedOrder.PID).Scan(&price)
    if err != nil {
        return fmt.Errorf("error fetching product price: %v", err)
    }

	// Check if user exists
    exists, err = UserExists(updatedOrder.UserID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("user not found")
    }

	// Calculate total price
    total_price := updatedOrder.OrderedQuantity * price
    log.Println("Total price:", total_price)

    now := time.Now()
    order_date := now

    // Update order in database
    query = "UPDATE orders SET ordered_quantity = ?, total_price = ?, order_date = ? WHERE user_id = ? AND p_id = ?"
    _, err = config.DB.Exec(query, updatedOrder.OrderedQuantity, total_price, order_date, updatedOrder.UserID, updatedOrder.PID)
    if err != nil {
        log.Println("Error updating order:", err)
        return err
    }

    log.Println("Order updated successfully:", "user_id:",updatedOrder.UserID, "p_id:", updatedOrder.PID)
    return nil
}

//delete order
func DeleteOrder(order_id int) error {    
	log.Println(" order ID:")

	// Check if order exists
	exists, err := OrderExists(order_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("order not found")
	}

	// Delete the order
	query := "DELETE FROM orders WHERE order_id = ?"
	result, err := config.DB.Exec(query, order_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, order not found")
	}
	return nil
}

// get all orders
func GetAllOrders() (*sql.Rows, error) {

	// Delete all filed of order table
	query := "SELECT * FROM orders"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	log.Print(query, rows)
	return rows, nil
}
