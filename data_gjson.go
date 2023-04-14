/*
tidwall/gjson is a Go library that provides fast and easy JSON parsing and manipulation. Here's an example of using tidwall/gjson to extract 
and manipulate data from a complex JSON document:


In this example, we have a complex JSON document containing various nested fields and an array. We use tidwall/gjson to extract values from the 
JSON document using the gjson.Get function and JSON path queries. We then update the age field and add a new contact to the contacts array using 
the gjson.Set function. Finally, we delete the phone contact from the contacts array using the gjson.Delete function. This demonstrates how 
tidwall/gjson can be used for parsing, extracting, and manipulating complex JSON data in Go.
*/


```go
package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	// Complex JSON document
	jsonData := `
	{
		"name": "Alice",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "New York",
			"country": "USA"
		},
		"contacts": [
			{
				"type": "email",
				"value": "alice@example.com"
			},
			{
				"type": "phone",
				"value": "123-456-7890"
			}
		]
	}
	`

	// Extract values from the JSON document
	name := gjson.Get(jsonData, "name").String()
	age := gjson.Get(jsonData, "age").Int()
	street := gjson.Get(jsonData, "address.street").String()
	city := gjson.Get(jsonData, "address.city").String()
	country := gjson.Get(jsonData, "address.country").String()
	email := gjson.Get(jsonData, "contacts.#(type==\"email\").value").String()
	phone := gjson.Get(jsonData, "contacts.#(type==\"phone\").value").String()

	// Print extracted values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Street:", street)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Email:", email)
	fmt.Println("Phone:", phone)

	// Update the age in the JSON document
	jsonData = gjson.Set(jsonData, "age", 31)

	// Add a new contact to the JSON document
	jsonData = gjson.Set(jsonData, "contacts.-1", map[string]interface{}{
		"type":  "website",
		"value": "alice.com",
	})

	// Delete the phone contact from the JSON document
	jsonData = gjson.Delete(jsonData, "contacts.#(type==\"phone\")")

	// Print the updated JSON document
	fmt.Println("\nUpdated JSON document:")
	fmt.Println(jsonData)
}
```
