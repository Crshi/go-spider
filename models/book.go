package models

type Book struct {
	Model

	Name   string `json:"name"`
	Type   int    `json:"type"`
	Author string `json:"author"`
}

func GetBooks() (books []Book) {
	db.Find(&books)

	return
}

func GetBookById(id int) Book {
	var book Book
	db.First(&book, id)

	return book
}

func GetBookByName(name string) Book {
	var book Book
	db.Where("Name = ?", name).First(&book)

	return book
}

func CreateBook(book Book) Book {
	db.Create(&book)
	return book
}
