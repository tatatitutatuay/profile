package services

import (
	"profile-backend/internal/models"
	"profile-backend/internal/repositories"
)

func UpdateProfile(profile models.Profile) (*models.Profile, error) {
	// Call repository to update the profile
	return repositories.UpdateProfile(profile)
}
