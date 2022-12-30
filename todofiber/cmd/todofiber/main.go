package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/template/html"

	"github.com/dkr290/go-devops/todofiber/handlers"
	"github.com/dkr290/go-devops/todofiber/repository"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var Repo = repository.NewRepo()

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	Repo.DbUser = os.Getenv("DATABASE_USER")
	Repo.DbHost = os.Getenv("DATABASE_HOST")
	Repo.DbName = "todos"
	Repo.DbPass = os.Getenv("DATABASE_PASS")
	Repo.DbPort = os.Getenv("APP_PORT")

	connInfo := "user=postgres password=Password123 host=172.19.158.144 sslmode=disable"
	initdb, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	dbName := "todos"
	_, err = initdb.Exec("create database " + dbName)
	if strings.Contains(err.Error(), "already exists") && err != nil {
		//handle the error
		fmt.Println("Database already created")

	} else {
		log.Fatalln("Cannot create database", err)
	}
	initdb.Close()

	connStr := "postgresql://postgres:Password123@172.19.158.144/todos?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//Then execute your query for creating table
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE todos( %s )", "item text"))

	if strings.Contains(err.Error(), "already exists") && err != nil {
		fmt.Println("Table already created")
	} else {
		log.Fatal("Create table failed ", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(db, c)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return handlers.PostHandler(db, c)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return handlers.UpdateHandler(db, c)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return handlers.DeleteHandler(db, c)
	})

	if Repo.DbPort == "" {
		Repo.DbPort = "3000"
	}

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", Repo.DbPort)))
}
