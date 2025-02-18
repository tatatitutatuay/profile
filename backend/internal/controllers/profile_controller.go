package controllers

import (
	"encoding/json"
	"net/http"
	"profile-backend/internal/models"
	"profile-backend/internal/services"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedProfile, err := services.UpdateProfile(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProfile)
}
