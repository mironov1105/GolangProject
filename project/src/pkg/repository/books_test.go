package repository

import (
	"testing"

	"project/src/pkg/models"
)

const connStr = "postgres://Admin:admin@localhost:8090/courses"

func TestBooksCRUD(t *testing.T) {
	db, err := New(connStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = db.GetBooks()
	if err != nil {
		t.Fatal(err.Error())
	}

	book := models.Book{
		Name:     "Удушье",
		AuthorID: 1,
		GenreID:  1,
	}
	id, err := db.NewBook(book)
	if err != nil || id == 0 {
		t.Fatal(err.Error())
	}

	_, err = db.GetBookByID(id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
