/*
show dependency injection example using google/wire

In this example, we define an interface called Database that has a single method called GetData. We then define two types that implement the Database 
interface: ProductionDatabase and TestDatabase.

We define a Server struct that has a single field called Database that is of type Database. We also define a constructor function called NewServer 
that takes a Database as an argument and returns a new Server instance.

We then define two sets of dependencies using wire.NewSet. The productionSet includes the NewServer constructor function and binds the Database interface to the ProductionDatabase type. The testSet is similar, but it binds the Database interface to the TestDatabase type instead.

We then define two initialization functions, initializeProductionServer and initializeTestServer, that use the wire.Build function to create instances 
of Server with the appropriate dependencies.

In the main function, we call initializeProductionServer and initializeTestServer to create instances of Server with the appropriate Database 
implementations. We then call the GetData method on each Database instance and print the result.

Note that we use wire.Bind to specify which implementation of the Database interface should be used. This allows us to easily swap out implementations 
without changing any other code.
*/

package main

import (
	"fmt"

	"github.com/google/wire"
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

var productionSet = wire.NewSet(NewServer, wire.Bind(new(Database), new(*ProductionDatabase)))
var testSet = wire.NewSet(NewServer, wire.Bind(new(Database), new(*TestDatabase)))

func main() {
	productionServer := initializeProductionServer()
	testServer := initializeTestServer()

	fmt.Println(productionServer.Database.GetData())
	fmt.Println(testServer.Database.GetData())
}

func initializeProductionServer() *Server {
	wire.Build(productionSet)
	return &Server{}
}

func initializeTestServer() *Server {
	wire.Build(testSet)
	return &Server{}
}
