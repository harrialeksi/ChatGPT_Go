Sure, here's an example of how to use the built-in html/template package in Go:

go
Copy code
package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	// Define the handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Create a person object
		person := Person{
			Name:    "Alice",
			Age:     25,
			Hobbies: []string{"reading", "hiking", "swimming"},
		}

		// Parse the template file
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the person object
		err = tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Register the handler with the server
	http.HandleFunc("/", handler)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
In this example, we define a Person struct with three fields: Name, Age, and Hobbies. We then define a handler function that creates a Person object and parses an HTML template file using template.ParseFiles. The template file contains placeholders for the Name, Age, and Hobbies fields, which we fill in using tmpl.Execute and the Person object.

When the handler function is called, it will serve the HTML file with the replaced values. Note that the template.Execute method automatically escapes any user-provided content to prevent cross-site scripting (XSS) attacks.

Additionally, you can use the {{range}} template directive to iterate over slices, such as the Hobbies field in the Person struct. Here's an updated version of the HTML template file that includes a {{range}} directive:

html
Copy code
<!DOCTYPE html>
<html>
  <head>
    <title>Person</title>
  </head>
  <body>
    <h1>{{.Name}}</h1>
    <p>Age: {{.Age}}</p>
    <h2>Hobbies:</h2>
    <ul>
      {{range .Hobbies}}
      <li>{{.}}</li>
      {{end}}
    </ul>
  </body>
</html>
This will generate an unordered list of the Hobbies field for the Person object.
