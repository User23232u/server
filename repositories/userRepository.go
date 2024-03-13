package repositories

import (
	"context"
	"project-websocket/database"
	"project-websocket/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllUsers obtiene todos los usuarios
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	collection := database.Client.Database("test").Collection("users")
	err := collection.Find(context.Background(), bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser obtiene un usuario por su ID
func GetUser(id string) (*models.User, error) {
	var user models.User

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := database.Client.Database("test").Collection("users")
	err = collection.Find(context.Background(), bson.M{"_id": objectID}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser crea un nuevo usuario
func CreateUser(newUser *models.User) (*models.User, error) {
	collection := database.Client.Database("test").Collection("users")
	_, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// UpdateUser actualiza un usuario
func UpdateUser(id string, updatedUser *models.User) (*models.User, error) {
	collection := database.Client.Database("test").Collection("users")

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Construir el objeto de actualización dinámicamente
	updateFields := bson.M{}
	if updatedUser.Username != "" {
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Email != "" {
		updateFields["email"] = updatedUser.Email
	}
	if updatedUser.Password != "" {
		updateFields["password"] = updatedUser.Password
	}

	// Actualizar manualmente el campo updateAt con la hora actual
	updateFields["updateAt"] = time.Now()

	// Utilizar UpdateId para actualizar el usuario por su ID
	err = collection.UpdateId(context.Background(), objectID, bson.M{"$set": updateFields})
	if err != nil {
		return nil, err
	}

	// Recuperar el usuario actualizado de la base de datos
	var updatedUserFromDB models.User
	err = collection.Find(context.Background(), bson.M{"_id": objectID}).One(&updatedUserFromDB)
	if err != nil {
		return nil, err
	}

	return &updatedUserFromDB, nil
}

// DeleteUser elimina un usuario por su ID
func DeleteUser(id string) (*models.User, error) {
	var user models.User

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := database.Client.Database("test").Collection("users")
	err = collection.Find(context.Background(), bson.M{"_id": objectID}).One(&user)
	if err != nil {
		return nil, err
	}

	// Utilizamos el método RemoveId para eliminar un documento por su ID
	err = collection.RemoveId(context.Background(), objectID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}