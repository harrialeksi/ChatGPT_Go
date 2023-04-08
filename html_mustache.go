/*
In this example, we first define the template data as a map of string keys to interface values. We then compile the Mustache template string using 
the mustache.ParseString function, which returns a *mustache.Template value.

We then render the template using the Render method of the mustache.Template value, passing in the data map and a bytes.Buffer value to store the 
rendered output. The rendered output is then printed to the console.

Alternatively, you can also use Go's built-in HTML template package to render the same Mustache template. We first create a new HTML template using 
the template.New function, passing in a name for the template and the template string. We then execute the template using the Execute method, passing 
in the same data map and a bytes.Buffer value to store the rendered output. The rendered output is then printed to the console.

Note that both Mustache and Go's HTML template package support a wide range of template features, including conditional statements, loops, and template 
inheritance. You can consult their respective documentation for more information.
*/

package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/cbroglie/mustache"
)

func main() {
	// Define the template data
	data := map[string]interface{}{
		"Title": "Hello, Mustache!",
		"Body":  "This is a Mustache template rendered in Go.",
	}

	// Compile the Mustache template
	templateString := "<html><head><title>{{Title}}</title></head><body><p>{{Body}}</p></body></html>"
	mustacheTemplate, err := mustache.ParseString(templateString)
	if err != nil {
		panic(err)
	}

	// Render the template
	var rendered bytes.Buffer
	err = mustacheTemplate.Render(&rendered, data)
	if err != nil {
		panic(err)
	}

	// Print the rendered template
	fmt.Println(rendered.String())

	// Alternatively, you can use Go's built-in HTML template package
	htmlTemplate, err := template.New("example").Parse(templateString)
	if err != nil {
		panic(err)
	}

	// Render the template using the same data
	var htmlRendered bytes.Buffer
	err = htmlTemplate.Execute(&htmlRendered, data)
	if err != nil {
		panic(err)
	}

	// Print the rendered template
	fmt.Println(htmlRendered.String())
}
