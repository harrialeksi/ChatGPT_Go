/*
show me a jwt access token handler in go language

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
    http.HandleFunc("/home", homeHandler)
    http.ListenAndServe(":8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the request body to get the username and password
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Find the user with the matching username and password
    var user User
    for _, u := range users {
        if u.Username == username && u.Password == password {
            user = u
            break
        }
    }

    if user.ID == 0 {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Create the JWT access token
    accessToken, err := generateAccessToken(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Write the access token to the response
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"access_token":"` + accessToken + `"}`))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")[7:] // remove "Bearer " prefix
    claims := jwt.MapClaims{}

    // Parse the JWT access token
    accessToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Make sure the token is an access token
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

    // Make sure the token is an access token
    if !accessToken.Valid {
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

    // Write a welcome message to the response
    w.Write([]byte("Welcome, " + user.Username + "!"))
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
