package service

import (
	"Actium_Todo/internal/repository"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SignUp hashes the password and saves the new user in the DB
func SignUp(userN, userP string) {
	// Generate bcrypt hash from plain password
	hash, err := bcrypt.GenerateFromPassword([]byte(userP), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Save the username and hashed password to the database
	err = repository.SignUp_user(userN, string(hash))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAccount with username %s has been created\n\n", userN)
}

// Login fetches user by username and verifies password hash
func Login(userN, userP string) (int, bool) {
	// Fetch user(s) from database
	users, err := repository.GetByUsersName(userN)
	if err != nil {
		fmt.Printf("Error retrieving users: %v\n", err)
		return 0, false
	}

	if len(users) == 0 {
		return 0, false
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(userP))
	if err != nil {
		return 0, false
	}

	// Password correct, return user ID
	return int(users[0].ID), true
}

func DeleteMyAccount(UserN, UserP string) error {
	// Fetch user by username
	users, err := repository.GetByUsersName(UserN)
	if err != nil {
		fmt.Printf("Error retrieving users: %v\n", err)
		return err
	}

	if len(users) == 0 {
		fmt.Println("User not found.")
		return err
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(UserP))
	if err != nil {
		fmt.Println("Incorrect password.")
		return err
	}

	// Password correct, delete the user
	err = repository.DeleteMyAccount(users[0].ID)
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}
