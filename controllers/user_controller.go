package controllers

import (
	"encoding/json"
	"go-user-crud/models"
	"net/http"
	"sync"
)

var (
	users    = make(map[int]models.User)
	nextID   = 1
	usersMux sync.Mutex
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	usersMux.Lock()
	user.Id = nextID
	users[nextID] = user
	nextID++
	usersMux.Unlock()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}
