package models

type Book struct {
	Model

	Name   string `json:"Name"`
	Type   int    `json:"Type"`
	Author string `json:"Author"`
}

func GetBooks() (books []Book) {
	db.Find(&books)

	return
}

func GetBookByName(name string) (book Book) {
	db.Where("Name = ?", name).Find(&book)

	return
}

func CreateBook(book Book) Book {
	db.Create(&book)
	return book
}
