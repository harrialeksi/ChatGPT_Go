/*
show me a login handler in go language (/login endpoint) meeting owasp security requirement

Sure, here's an example of a login handler in Go that meets OWASP security requirements for a /login endpoint:
*/

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username" schema:"username,required"`
	Password string `json:"password" schema:"password,required"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var schemaDecoder = schema.NewDecoder()

func init() {
	// Register the "password" schema decoder hook to hash the password
	schemaDecoder.RegisterConverter(Credentials{}, "password", func(value string) reflect.Value {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
		return reflect.ValueOf(string(hashedPassword))
	})
}

func main() {
	// Create a new router with the login endpoint
	r := mux.NewRouter()
	r.HandleFunc("/login", handleLogin).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Parse the credentials from the request body
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Lookup the user in the database
	user, err := lookupUser(creds.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create the JWT token
	claims := Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "myapp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("mysecret"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the token as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func lookupUser(username string) (*User, error) {
	// Replace with your own database lookup code
	db, err := sql.Open("postgres", "user=postgres dbname=myapp sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)
	var user User
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &
