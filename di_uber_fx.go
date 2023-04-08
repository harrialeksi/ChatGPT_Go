/*
show dependency injection example using uber/fx

In this example, we define an interface called Database that has a single method called GetData. We then define two types that implement the Database 
interface: ProductionDatabase and TestDatabase.

We define a Server struct that has a single field called Database that is of type Database. We also define a constructor function called NewServer 
that takes a Database as an argument and returns a new Server instance.

We create a new fx.App instance and use the fx.Provide function to provide our NewServer constructor function and a function that returns either a 
ProductionDatabase or a TestDatabase based on the environment.

We then call app.Run and pass in a function that takes a *Server as an argument. The fx library will automatically inject the appropriate Database 
implementation into the NewServer constructor function, and then inject the resulting Server instance into our app.Run function.

In the app.Run function, we call the GetData method on the Database instance that was injected into our Server instance, and print the result.

Note that we don't need to explicitly choose which implementation of the Database interface to use. Instead, we use the fx.Provide function to provide 
a function that returns the appropriate implementation based on the environment. This makes it easy to swap out implementations without changing any 
other code.
*/

package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Database interface {
	GetData() string
}

type ProductionDatabase struct{}

func (db *ProductionDatabase) GetData() string {
	return "Production data"
}

type TestDatabase struct{}

func (db *TestDatabase) GetData() string {
	return "Test data"
}

type Server struct {
	Database Database
}

func NewServer(database Database) *Server {
	return &Server{Database: database}
}

func main() {
	app := fx.New(
		fx.Provide(NewServer),
		fx.Provide(func() Database {
			// Return either a ProductionDatabase or a TestDatabase based on the environment
			return &TestDatabase{}
		}),
	)

	app.Run(func(server *Server) {
		fmt.Println(server.Database.GetData())
	})
}
