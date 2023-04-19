/*
Fluent interface, also known as method chaining, is a design pattern in object-oriented programming where a series of methods are chained together 
on the same object to create a concise and expressive way of writing code. The methods in a fluent interface typically return the same object on 
which they were called, allowing for consecutive method calls on the same object without the need for temporary variables.

In Go, a fluent interface can be implemented using method receivers and return values. Method receivers are similar to the this or self pointer 
in other languages, and they allow a method to be attached to a particular type or struct. Return values can be used to chain subsequent method 
calls on the same object.

In this example, the Person struct has two methods SetName and SetAge that set the name and age of a person, respectively, and return the person 
object to allow for method chaining. The PrintInfo method simply prints the name and age of the person. In the main function, we create a new person, 
set the name and age using method chaining, and then print the person's info.
*/

package main

import "fmt"

// Person represents a person with a name and age.
type Person struct {
	name string
	age  int
}

// NewPerson creates a new Person instance with the given name and age.
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// SetName sets the name of the person and returns the person object for method chaining.
func (p *Person) SetName(name string) *Person {
	p.name = name
	return p
}

// SetAge sets the age of the person and returns the person object for method chaining.
func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}

// PrintInfo prints the name and age of the person.
func (p *Person) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func main() {
	// Create a new person and set the name and age using method chaining
	person := NewPerson("", 0).SetName("Alice").SetAge(30)

	// Print the person's info
	person.PrintInfo()
}
