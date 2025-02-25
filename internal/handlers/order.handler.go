package handlers

import (
	"Codimite_Assignment/internal/models"
	"Codimite_Assignment/internal/queries"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//create a new order
func CreateOrder(w http.ResponseWriter, r *http.Request){
	var order models.CreateOrder

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	// Call the query function to Insert order into database
	err = queries.AddOrder(order.UserID, order.PID, order.OrderedQuantity)
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		log.Println("Error creating order",err)
		return
	}
	log.Println("order Successfully created")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully created"))
}

//update order
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	
	// Extract order_id from path params
	vars := mux.Vars(r)
	orderIdStr, exists := vars["id"]
	
	if !exists {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Convert order_id from string to int
	order_id, err := strconv.Atoi(orderIdStr)

	log.Println("order Id:", order_id)

	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	var updatedOrder models.CreateOrder
	// Decode JSON request body
	err = json.NewDecoder(r.Body).Decode(&updatedOrder)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("err", err)
		return
	}
	// Call the query function to update the order
	err = queries.UpdateOrder(order_id, updatedOrder)
	if err != nil {
		http.Error(w, "Error updating order", http.StatusInternalServerError)
		log.Println("err", err)
		return
	}

	log.Println("order successfully updated:", order_id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - order successfully updated"))
}

// delete order
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Extract order_id from path params
	vars := mux.Vars(r)
	orderIdStr, exists := vars["id"]
	
	if !exists {
		http.Error(w, "order ID is required", http.StatusBadRequest)
		return
	}

	// Convert order_id from string to int
	order_id, err := strconv.Atoi(orderIdStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	log.Println("orderId:", order_id)

	// Call the query function to delete the order
	err = queries.DeleteOrder(order_id)
	if err != nil {
		http.Error(w, "Error deleting order ", http.StatusInternalServerError)
		log.Println("error",err)
		return
	}
	log.Println("order Successfully deleted")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - order Successfully Deleted"))
}

//get all orders
func GetAllOrder(w http.ResponseWriter, r *http.Request) {

	var orders []models.Order

	// get all orders from the database
	rows, err := queries.GetAllOrders()
	if err != nil {
		http.Error(w, "Error getting orders", http.StatusInternalServerError)
		return
	}

	// Iterate over query results and populate orders slice
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID , &order.UserID, &order.PID, &order.OrderedQuantity, &order.TotalPrice, &order.OrderDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orders = append(orders, order)
	}

	// Set response headers and encode the orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}