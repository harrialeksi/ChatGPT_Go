To demonstrate a Cross-Site Scripting (XSS) attack in Go, we can create a simple web application that takes user input and displays it on a web page 
without proper validation and sanitization.

Here's an example of a vulnerable web application that allows an attacker to inject malicious JavaScript code into the page:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the value of the "name" parameter from the query string
		name := r.URL.Query().Get("name")

		// Display the name on the page
		fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
	})

	http.ListenAndServe(":8080", nil)
}
```

In this example, the main function defines a route for the root URL / that retrieves the value of the name parameter from the query string and 
displays it on the page using fmt.Fprintf. However, this code is vulnerable to XSS attacks because it does not validate or sanitize the input from 
the user.

An attacker can inject malicious JavaScript code into the page by passing a name parameter with a script tag in the query string, like this:

```http://localhost:8080/?name=<script>alert("XSS!");</script>```

This will display an alert box with the message "XSS!" when the page loads, because the script tag was not properly validated or sanitized.

To prevent XSS attacks, you can use the html/template package to escape any user input before displaying it on the page. Here's an example of a 
secure version of the same web application that uses the html/template package to prevent XSS attacks:

```go
package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the value of the "name" parameter from the query string
		name := r.URL.Query().Get("name")

		// Define the template with the {{.}} placeholder that will be replaced with the name
		tmpl := template.Must(template.New("example").Parse(`<h1>Hello, {{.}}!</h1>`))

		// Use the template to render the name, automatically escaping any HTML characters
		err := tmpl.Execute(w, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
```

In this example, we use the html/template package to create a template that automatically escapes any HTML characters in the user input. We define 
the template with the {{.}} placeholder, which will be replaced with the name parameter from the query string.

We then use the Execute method to render the template with the name and automatically escape any HTML characters. This prevents any potentially 
malicious scripts from being executed in the browser.

By using the html/template package and properly escaping any user input, we can prevent XSS attacks in our Go web applications.
