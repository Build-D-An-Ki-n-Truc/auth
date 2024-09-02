package auth

import (
	"github.com/Build-D-An-Ki-n-Truc/auth/internal/db/mongodb"
)

// Claim structure
/*
	claim := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
*/

// Register a user account
func RegisterAccount(username string, password string, name string, email string, role string, phone string, isLocked bool) (bool, error) {
	// Check if user already exists
	_, err := mongodb.ReadUser(username)
	if err == nil {
		return false, nil
	}

	// Hash the password
	// hashedPassword, err := hash.HashPassword(password)
	// if err != nil {
	// 	return false, err
	// }

	// Create a new user
	user := mongodb.UserStruct{
		Username: username,
		Password: password,
		Name:     name,
		Email:    email,
		Role:     role,
		Phone:    phone,
		IsLocked: isLocked,
	}

	// Insert the user into the database
	err = mongodb.CreateUser(user)
	if err != nil {
		return false, err
	}

	return true, nil
}
