package main

import (
	//"context"
	"fmt"
	"log"

	//"github.com/jackc/pgx/v4"

	//"project/src/models"
	"project/src/pkg/repository"
)

type book struct {
	id       int
	name     string
	authorID int
	genreID  int
}

const connStr = "postgres://Admin:admin@localhost:8090/courses"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := db.GetBooks()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, it := range data {
		fmt.Printf("%d: %s, author id: %d, genre id: %d\n", it.ID, it.Name, it.AuthorID, it.GenreID)
	}

	item, err := db.GetBookByID(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%d: %s, author id: %d, genre id: %d\n", item.ID, item.Name, item.AuthorID, item.GenreID)

	//book := models.Book{
	//	Name:     "Удушье",
	//	AuthorID: 1,
	//	GenreID:  1,
	//}
	//err = db.NewBook(book)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
}
