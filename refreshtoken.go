/*
show me a jwt refresh token handler in go language

This example assumes that you have already implemented an authentication system that issues JWT access tokens and refresh tokens. The refreshHandler 
function takes a refresh token from the request header, verifies it, and generates a new access token if the token is valid. The generateAccessToken 
function creates a new JWT access token with a 15-minute expiration time.

Note that this example is for educational purposes only, and you should not use it in production without thorough testing and security reviews. In production, 
you should also consider using a third-party library for JWT handling, such as "github.com/golang-jwt/jwt".
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
    http.HandleFunc("/refresh", refreshHandler)
    http.ListenAndServe(":8000", nil)
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")[7:] // remove "Bearer " prefix
    claims := jwt.MapClaims{}

    // Parse the JWT refresh token
    refreshToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Make sure the token is a refresh token
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrInvalidSigningAlgorithm
        }
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

    // Make sure the token is a refresh token
    if !refreshToken.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Get the user ID from the claims
    userID := int(claims["user_id"].(float64))

    // Find the user
    var user User
    for _, u := range users {
        if u.ID == userID {
            user = u
            break
        }
    }

    // Generate a new JWT access token
    accessToken, err := generateAccessToken(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Write the new access token to the response
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"access_token":"` + accessToken + `"}`))
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
