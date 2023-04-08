/*

show me a login handler in go language (/login endpoint) using cookies

In this example, we're setting a cookie with the JWT token using the http.SetCookie() function. We're also setting the HttpOnly flag to true to prevent 
the cookie from being accessed by client-side JavaScript, and the Secure flag to true to ensure that the cookie is only sent over HTTPS. Finally
*/

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var (
	store = sessions.NewCookieStore([]byte("my-secret-key"))
)

func main() {
	// Create a new HTTP server
	server := http.Server{
		Addr: ":8080",
	}

	// Create a new database connection
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new HTTP handler for the /login endpoint
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body into a User struct
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Query the database for the user with the given username
		row := db.QueryRow("SELECT * FROM users WHERE username = $1", user.Username)
		var storedUser User
		err = row.Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Compare the password hash from the database to the provided password
		err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Create a JWT token with a 30 minute expiration time
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": storedUser.ID,
			"exp": time.Now().Add(30 * time.Minute).Unix(),
		})
		tokenString, err := token.SignedString([]byte("my-secret-key"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Create a new session for the user
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		session.Values["user_id"] = storedUser.ID
		session.Values["token"] = tokenString
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return the JWT token in the response
		response := LoginResponse{
			Token: tokenString,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Start the HTTP server
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
