package repository

import (
	db "Actium_Todo/internal/database"
	"Actium_Todo/internal/models"
	"fmt"
	"log"
)

var tasks []models.Task

func CreateTask(title, description, status string, creatorId int, deadline string) error {

	_, err := db.GetDB().Exec(
		"INSERT INTO tasks(title,description,status,creator_id,deadline) VALUES($1,$2,$3,$4,$5)",
		title, description, status, creatorId, deadline,
	)

	return err
}

func GetTasksByCreator(creatorId int) ([]models.Task, error) {

	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1  ORDER BY id ", creatorId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tasks = []models.Task{}

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}

func ShowCompletedTasks(creatorId int) ([]models.Task, error) {
	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1 AND status = 'Completed' ORDER BY id ", creatorId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tasks = []models.Task{}

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}

func UpdateStatus(taskId int, newStatus string, creatorId int) {

	db.GetDB().Exec("UPDATE tasks SET status = $1 WHERE creator_id = $2 AND id = $3", newStatus, creatorId, taskId)
	fmt.Printf("Task's status has been successfully changed to '%s'", newStatus)

}

func LoadAllTasks(creatorId int) ([]models.Task, error) {

	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1  ORDER BY id ", creatorId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tasks = []models.Task{}

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}
func DeleteTask(taskId, creatorId int) {
	db.GetDB().Exec("DELETE FROM tasks WHERE id = $1 AND creator_id = $2", taskId, creatorId)
}

func SearchTask_byId(creatorId, taskID int) ([]models.Task, error) {
	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1  AND id = $2 ORDER BY id ", creatorId, taskID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var foundTasks []models.Task

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		foundTasks = append(foundTasks, t)
	}
	return foundTasks, err
}

func SearchTask_byTitle(creatorId int, taskTitle string) ([]models.Task, error) {
	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1  AND id = $2 ORDER BY id ", creatorId, taskTitle)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var foundTasks []models.Task

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		foundTasks = append(foundTasks, t)
	}
	return foundTasks, err
}
func SearchTask_byStatus(creatorId int, taskStatus string) ([]models.Task, error) {
	rows, err := db.GetDB().Query("SELECT * FROM tasks WHERE creator_id = $1  AND id = $2 ORDER BY id ", creatorId, taskStatus)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var foundTasks []models.Task

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Creator_Id, &t.CreatedAt, &t.Deadline); err != nil {
			log.Fatal(err)
		}
		foundTasks = append(foundTasks, t)
	}
	return foundTasks, err
}
