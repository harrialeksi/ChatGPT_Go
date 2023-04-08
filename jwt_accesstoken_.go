/*
show me a jwt access token handler in go language

This example demonstrates how to implement JWT access token authentication using a middleware function (authMiddleware) that checks the validity of the token before
*/

package main

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
)

type User struct {
    ID       int
    Username string
    Password string
}

var (
    jwtKey = []byte("my_secret_key")
    users  = []User{
        {ID: 1, Username: "user1", Password: "password1"},
        {ID: 2, Username: "user2", Password: "password2"},
    }
)

func main() {
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/protected", authMiddleware(protectedHandler))
    http.ListenAndServe(":8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Get the username and password from the request body
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Find the user with the given username and password
    var user User
    for _, u := range users {
        if u.Username == username && u.Password == password {
            user = u
            break
        }
    }

    // If the user is not found, return an error
    if user.ID == 0 {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Generate a new JWT access token
    accessToken, err := generateAccessToken(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Write the access token to the response
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"access_token":"` + accessToken + `"}`))
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get the JWT access token from the request header
        tokenString := r.Header.Get("Authorization")[7:] // remove "Bearer " prefix
        claims := jwt.MapClaims{}

        // Parse the JWT access token
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        // Make sure the token is a JWT access token
        if !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        // Call the next handler function
        next.ServeHTTP(w, r)
    }
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    // Write the protected content to the response
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"This content is protected"}`))
}

func generateAccessToken(user User) (string, error) {
    // Set the expiration time of the token to 15 minutes
    expirationTime := time.Now().Add(15 * time.Minute)

    // Create the JWT claims
    claims := jwt.MapClaims{}
    claims["user_id"] = user.ID
    claims["exp"] = expirationTime.Unix()

    // Create the JWT access token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    accessToken, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return accessToken, nil
}
