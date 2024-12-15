package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"hexagonal-potato/domain"
	"net/http"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.AuthUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	CreateUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.AuthUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := GetUser(user)

	json.NewEncoder(w).Encode(res)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var updatedUser domain.AuthUser
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	UpdateUser(updatedUser)
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var deletedUser domain.AuthUser
	deletedUser.ID = id
	DeleteUser(deletedUser)
	w.WriteHeader(http.StatusNoContent)
}

func listUsers(w http.ResponseWriter, _r *http.Request) {
	users := ListUsers()

	w.WriteHeader(http.StatusOK)
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		errors.New("An error happened when listing users")
	}
	w.Write(data)
}

func InitHTTPService() {
	http.HandleFunc("/listUsers", listUsers)
	http.HandleFunc("/createUser", createUserHandler)
	http.HandleFunc("/getUser", getUserHandler)
	http.HandleFunc("/updateUser", updateUserHandler)
	http.HandleFunc("/deleteUser", deleteUserHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
