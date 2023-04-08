/*
show me a logout handler in go language (/logout endpoint) while using cookies and clear out the session

In this example, we're using the http.SetCookie() function to clear the session cookie. We're setting the MaxAge field to -1 to tell the browser 
to immediately expire the cookie. We're also setting the Value field to an empty string and the Path field to / to ensure that the cookie is deleted 
for all paths on the domain. Finally, we're redirecting the user to the login page using the http.Redirect() function with the http.StatusSeeOther status code.
*/

package main

import (
	"net/http"
)

func main() {
	// Create a new HTTP server
	server := http.Server{
		Addr: ":8080",
	}

	// Create a new HTTP handler for the /logout endpoint
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		// Clear the session cookie
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})

		// Redirect the user to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	// Start the HTTP server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
