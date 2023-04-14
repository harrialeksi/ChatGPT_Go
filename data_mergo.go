//mdario/mergo is a Go library that provides functionalities for merging and deep-merging maps and structs. Here's an example of using mdario/mergo to merge two structs with nested fields:

```
/*
In this example, we have two Person structs, person1 and person2, with nested fields. We want to merge person2 into person1 using mdario/mergo. 
After the merge, person1 will be updated with the fields from person2 while preserving its original values for the fields that are not present in 
person2. In this case, person1 will have its Age field updated to 25 and the City field in the Address field updated to "Los Angeles". 
The mergo.WithOverride option is used to specify that the fields from person2 should override the corresponding fields in person1.

Note that mdario/mergo provides various other options for merging structs, such as deep merging of maps and slices, specifying custom merge strategies, 
and more. You can refer to the library's documentation for more advanced usage.
*/

package main

import (
	"fmt"

	"github.com/mdario/mergo"
)

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

func main() {
	// Create two Person structs
	person1 := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	person2 := Person{
		Age:  25,
		Address: Address{
			City: "Los Angeles",
		},
	}

	// Print the original person1
	fmt.Println("Original Person1:")
	fmt.Println(person1)

	// Print the original person2
	fmt.Println("\nOriginal Person2:")
	fmt.Println(person2)

	// Merge person2 into person1
	if err := mergo.Merge(&person1, person2, mergo.WithOverride); err != nil {
		fmt.Println("Error merging structs:", err)
		return
	}

	// Print the merged person1
	fmt.Println("\nMerged Person1:")
	fmt.Println(person1)
}
```
