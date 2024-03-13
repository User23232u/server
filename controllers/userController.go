package controllers

import (
	"project-websocket/models"
	"project-websocket/repositories"
)

// GetAllUsers obtiene todos los usuarios
func GetAllUsers() ([]*models.User, error) {
	return repositories.GetAllUsers()
}

// GetUser obtiene un usuario por su ID
func GetUser(id string) (*models.User, error) {
	return repositories.GetUser(id)
}

// CreateUser crea un nuevo usuario
func CreateUser(newUser *models.User) (*models.User, error) {
	return repositories.CreateUser(newUser)
}

func UpdateUser(id string, updatedUser *models.User) (*models.User, error) {
	return repositories.UpdateUser(id, updatedUser)
}

// DeleteUser elimina un usuario
func DeleteUser(id string) (*models.User, error) {
    user, err := repositories.DeleteUser(id)
    if err != nil {
        return nil, err
    }
    return user, nil
}
