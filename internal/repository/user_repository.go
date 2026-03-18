package repository

import (
	db "Actium_Todo/internal/database"
	"Actium_Todo/internal/models"
	"database/sql"
	"errors"
)

func GetByUsersName(username string) ([]models.User, error) {
	rows, err := db.GetDB().Query(
		`SELECT id, user_name, password, joined_at FROM users WHERE user_name = $1`,
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID,
			&u.UserName,
			&u.Password,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetByID(id int) (models.User, error) {
	var user models.User

	query := `SELECT id, user_name, password, joined_at FROM users WHERE id = $1`

	err := db.GetDB().QueryRow(query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func SignUp_user(username, password string) error {
	_, err := db.GetDB().Exec("INSERT INTO users (user_name, password) VALUES ($1, $2)", username, password)

	return err
}
func DeleteMyAccount(userID int64) error {
	_, err := db.GetDB().Exec("DELETE FROM users WHERE id = $1", userID)

	return err
}

// GetUserIDByUsername retrieves the user ID based on the provided username
func GetUserIDByUsername(username string) (int, error) {
	var userID int
	err := db.GetDB().QueryRow("SELECT id FROM users WHERE user_name = $1", username).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
func DeleteAllTasksFromUser(userN string) error {
	userID, err := GetUserIDByUsername(userN)
	if err != nil {
		return err
	}
	_, err = db.GetDB().Exec("DELETE FROM tasks WHERE user_id = $1", userID)
	return err
}
