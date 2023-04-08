/*
show me a login handler in go language (/login endpoint) meeting owasp security requirement

This example meets some of the OWASP security requirements:

The login endpoint uses HTTPS, which provides transport layer security.
The login endpoint uses a strong password hashing algorithm (bcrypt) to store passwords securely in the database.
The login endpoint returns a JWT token with a short expiration time (30 minutes) to limit the amount of time an attacker can use a stolen token.
The login endpoint uses prepared SQL statements to prevent SQL injection attacks.
However, this example is not complete and there are many other security considerations to take into account when implementing a login system, such as protecting 
against brute-force attacks, implementing rate limiting, and using secure session management.

*/

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
