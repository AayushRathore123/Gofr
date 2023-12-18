package main

import (
	"database/sql"
	"fmt"
	"log"

	"gofr.dev/pkg/gofr"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	dbName := "gofr"

	// Create a DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	app := gofr.New()
	//get api to check working or not
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	//post student
	app.POST("/student/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO user (name) VALUES (?)", name)

		return nil, err
	})

	//get all students
	app.GET("/students", func(ctx *gofr.Context) (interface{}, error) {
		var students []Student

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM user")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var student Student
			if err := rows.Scan(&student.ID, &student.Name); err != nil {
				return nil, err
			}

			students = append(students, student)
		}

		// return the customer
		return students, nil
	})

	app.Start()
}
