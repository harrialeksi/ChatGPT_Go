Sure, here's an example of how to use the markbates/goth package in Go to authenticate with GitHub using OAuth:

First, install the markbates/goth package by running go get github.com/markbates/goth.

Create a new GitHub OAuth application by following the instructions in the GitHub Developer Guide. Make a note of the Client ID and Client Secret values.

In your Go code, import the github.com/markbates/goth and github.com/markbates/goth/providers/github packages. Create a new goth.Provider object 
with the github.New() function, passing in your Client ID and Client Secret values. Set the CallbackPath field to the path that your OAuth callback 
handler will use. For example:

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/markbates/goth"
    "github.com/markbates/goth/providers/github"
)

func main() {
    // Create GitHub provider with client ID and secret
    githubProvider := github.New(
        "client_id",
        "client_secret",
        "http://localhost:3000/auth/github/callback",
    )

    // Add GitHub provider to Goth
    goth.UseProviders(githubProvider)

    // Create a new HTTP handler for the GitHub OAuth callback
    http.HandleFunc("/auth/github/callback", func(res http.ResponseWriter, req *http.Request) {
        // Complete authentication and get user data
        user, err := githubProvider.FetchUser(goth.ContextForRequest(req))
        if err != nil {
            fmt.Fprintf(res, "Error authenticating with GitHub: %s", err.Error())
            return
        }

        // Display user data
        fmt.Fprintf(res, "Welcome, %s!", user.Name)
    })

    // Start the HTTP server
    http.ListenAndServe(":3000", nil)
}
```

Run the Go code and visit http://localhost:3000/auth/github to start the authentication process. You should be redirected to the GitHub login page. 
After logging in, you will be redirected back to your callback handler (/auth/github/callback) and the user's data will be displayed.
Note that this is just a basic example and you should customize it to suit your needs, such as storing user data in a database or using a session 
cookie to remember the user's login status.
