package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"

	"github.com/edr3x/tern-sqlc-poc/cmd/user"
	"github.com/edr3x/tern-sqlc-poc/internal/db/connection"
	"github.com/edr3x/tern-sqlc-poc/internal/helpers"
)

func main() {
	r := chi.NewRouter()

	conn, err := connection.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r.Get("/users", GetAllUsers)
	r.Get("/users/{id}", GetUserById)
	r.Post("/users", CreateUser)
	r.Delete("/users/{id}", DeleteUserById)

	log.Println("Started server on port 8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body user.UserCreateInput
	if err := helpers.DecodeJSONBody(w, r, &body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid input"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userService := user.UserService()
	user, err := userService.CreateUser(ctx, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating user"))
	}

	usr, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(usr)
}

func GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	userService := user.UserService()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	users, err := userService.GetAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	usr, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(usr)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	userService := user.UserService()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := userService.GetUserById(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	usr, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(usr)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	userService := user.UserService()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := userService.DeleteUser(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted user"))
}
