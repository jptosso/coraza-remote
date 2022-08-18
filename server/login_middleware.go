package server

import (
	"log"
	"net/http"

	"github.com/jptosso/coraza-center/database"
)

func loginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid authentication", http.StatusUnauthorized)
			return
		}
		var userModel *database.User
		tx := database.DB.Model(&database.User{}).First(&userModel, "user_name = ? and password = ?", user, pass)
		if tx.Error != nil || userModel == nil {
			log.Printf("Invalid credentials for user %s", user)
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
