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

// create a product
func CreateProduct(w http.ResponseWriter, r *http.Request){
	var product models.Product

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("error",err)
		return
	}

	// Call the query function to Insert product into database
	err = queries.AddProduct(product.Name, product.PType, product.Price, product.Quantity)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		log.Println("Error creating product",err)
		return
	}
	log.Println("Product Successfully created", product.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully created"))
}

//update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	
	// Extract p_id from path params
	vars := mux.Vars(r)
	pidStr, exists := vars["id"]
	
	if !exists {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Convert p_id from string to int
	p_id, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	log.Println("pid:", p_id)

	// Decode JSON request body
	var updatedProduct models.UpdateProduct
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		log.Println("err", err)
		return
	}

	// Call the query function to update the product
	err = queries.UpdateProduct(p_id, updatedProduct)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		log.Println("err", err)
		return
	}

	log.Println("Product successfully updated:", p_id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Product successfully updated"))
}

//delete order
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	// Extract p_id from path params
	vars := mux.Vars(r)
	pidStr, exists := vars["id"]
	
	if !exists {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Convert p_id from string to int
	p_id, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	log.Println("pid:", p_id)

	// Call the query function to delete the product
	err = queries.DeleteProduct(p_id)
	if err != nil {
		http.Error(w, "Error deleting product ", http.StatusInternalServerError)
		log.Println("error",err)
		return
	}
	log.Println("product Successfully deleted")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Successfully Deleted"))
}

//get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	// get all products from the database
	var products []models.Product
	rows, err := queries.GetAllProducts()
	if err != nil {
		http.Error(w, "Error getting products", http.StatusInternalServerError)
		return
	}

	// Iterate over query results and populate products slice
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.PId, &product.Name, &product.PType, &product.Price, &product.Quantity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	// Set response headers and encode the products as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}