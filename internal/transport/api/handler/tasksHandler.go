package handler

import (
	"Actium_Todo/internal/models"
	"Actium_Todo/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	t.Creator_Id = int(user.ID)
	taskID, err := repository.CreateTask(t.Title, t.Description, t.Status, t.Creator_Id, t.Deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Id = taskID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}
	var t models.Task
	t.Creator_Id = int(user.ID)
	tasks, err := repository.GetTasksByCreator(t.Creator_Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	t.Creator_Id = int(user.ID)

	err := repository.UpdateStatus(t.Id, t.Status, t.Creator_Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "status updated",
		"task_id":   fmt.Sprintf("%d", t.Id),
		"newStatus": t.Status,
	})
}

func DeletTask(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	t.Creator_Id = int(user.ID)
	err := repository.DeleteTask(t.Id, t.Creator_Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "task deleted successfully",
	})
}
