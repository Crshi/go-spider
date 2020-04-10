package models

type Book struct {
	Model

	Name string `json:"Name"`
	Type int    `json:"Type"`
}

func GetBooks() (books []Book) {
	db.Find(&books)

	return
}

func GetBookCount() (count int) {
	db.Model(&Book{}).Count(&count)

	return
}

func CreateBook(name string, bookType int) bool {
	db.Create(&Book{
		Name: name,
		Type: bookType,
	})

	return true
}
