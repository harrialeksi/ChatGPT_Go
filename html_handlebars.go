/*
Sure, here's an example of how to use the handlebars package to render Handlebars templates in Go:

In this example, we define a Person struct with three fields: Name, Age, and Hobbies. We then define a handler function that creates a Person object 
and a Handlebars template as a string. We use the raymond.Render function from the handlebars package to render the template with the Person object.

The Handlebars template uses {{name}}, {{age}}, and {{#each hobbies}} syntax to reference the Name, Age, and Hobbies fields of the Person object. 
The {{#each}} directive is used to iterate over the Hobbies slice.

When the handler function is called, it will render the Handlebars template with the replaced values and serve it as HTML.

Note that you can also load Handlebars templates from external files using the raymond.ParseFile function. This can be useful for larger templates 
that you don't want to include as string literals in your Go code.
*/

package main

import (
	"fmt"
	"github.com/aymerick/raymond"
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

		// Define the Handlebars template
		template := `
			<!DOCTYPE html>
			<html>
			  <head>
			    <title>{{name}}</title>
			  </head>
			  <body>
			    <h1>{{name}}</h1>
			    <p>Age: {{age}}</p>
			    <h2>Hobbies:</h2>
			    <ul>
			      {{#each hobbies}}
			      <li>{{this}}</li>
			      {{/each}}
			    </ul>
			  </body>
			</html>
		`

		// Render the template with the person object
		output, err := raymond.Render(template, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Serve the rendered HTML
		fmt.Fprint(w, output)
	}

	// Register the handler with the server
	http.HandleFunc("/", handler)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
