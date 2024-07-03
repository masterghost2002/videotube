package repository

import (
	"github.com/masterghost2002/videotube/internals/models"
	"github.com/masterghost2002/videotube/internals/registry"
)

func FindUser(email string) *models.User {
	storage := registry.GetDB()

	var user models.User

	result := storage.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil
	}
	return &user
}
