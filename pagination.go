/*
In this example, we first connect to the MySQL database using the gorm.Open function. We then define a User struct to represent the table in our
database. We set the current page to 1 and the number of records per page to 10.

We use the Count function to count the total number of records in the User table. We then calculate the offset and limit based on the current page
and the page size. We retrieve the records using the Offset and Limit functions, and store them in a slice of User structs.

Finally, we display the current page and the retrieved records. The (total+pageSize-1)/pageSize expression calculates the total number of pages based
on the total number of records and the page size.

Note that this is just a basic example of pagination in Go using the gorm package. You may need to modify the code to fit your specific requirements.
*/

package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// HERE! HERE!
	// Initialize variables
	page := 1      // Current page
	pageSize := 10 // Number of records per page
	var total int  // Total number of records

	// Count total number of records
	db.Model(&User{}).Count(&total)

	// Calculate offset and limit
	offset := (page - 1) * pageSize
	limit := pageSize

	// Retrieve records
	var users []User
	db.Offset(offset).Limit(limit).Find(&users)

	// Display records
	fmt.Printf("Page %d/%d\n", page, (total+pageSize-1)/pageSize)
	for _, user := range users {
		fmt.Println(user)
	}
}
