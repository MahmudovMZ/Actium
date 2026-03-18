package handler

import (
	"Actium_Todo/internal/models"
	"Actium_Todo/internal/repository"
	"Actium_Todo/internal/service"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SighIn_handler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	service.SignUp(u.UserName, u.Password)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
	w.Write([]byte("You signed in succesfully! Login to have access to the service"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	userID, ok := service.Login(u.UserName, u.Password)

	if !ok {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	claim := models.Claims{
		UserId: int(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(models.Secret)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successful login",
		"token":   signedToken,
	})
}
func MeHandler(w http.ResponseWriter, r *http.Request) {

	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("token")
		if key == "" {
			http.Error(w, "Token is absent", http.StatusUnauthorized)
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(key, claims, func(t *jwt.Token) (interface{}, error) {
			return models.Secret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		user, err := repository.GetByID(claims.UserId)
		if err != nil {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("user")
	user, ok := value.(models.User)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err := repository.DeleteMyAccount(user.ID)
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("account deleted successfully"))
}
