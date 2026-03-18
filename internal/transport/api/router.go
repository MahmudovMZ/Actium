package api

import (
	"Actium_Todo/internal/transport/api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() error {
	r := mux.NewRouter()
	r.HandleFunc("/signIn", handler.SighIn_handler).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST").Headers("Content-Type", "application/json")

	user := r.PathPrefix("/user").Subrouter() //Using middleware to prevent unauthorized users from the service
	user.Use(handler.ValidateMiddleware)
	user.HandleFunc("/me", handler.MeHandler).Methods("GET")
	user.HandleFunc("/deleteMy_account", handler.DeleteUser).Methods("DELETE")

	task := r.PathPrefix("/task").Subrouter()
	task.Use(handler.ValidateMiddleware)
	task.HandleFunc("/create", handler.CreateTask).Methods("GET")
	task.HandleFunc("/show", handler.ShowTasks).Methods("POST")
	task.HandleFunc("/update", handler.UpdateTaskStatus).Methods("PUT")
	task.HandleFunc("/delete", handler.DeletTask).Methods("DELETE")

	// user.HandleFunc("/logout", handler.Logout).Methods("POST").Headers("Content-Type", "application/json")
	return http.ListenAndServe(":8000", r)

}
