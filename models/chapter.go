package models

type Chapter struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
	Book_Id int    `json:"book_id"`
	Order   int    `json:"order"`
	Url     string `json:"url"`
}

func GetChapters(bookId int) (chapters []Chapter) {
	db.Where("book_id = ?", bookId).Order("order").Find(&chapters)
	return
}

func AddChapter(chapter Chapter) bool {
	db.Create(&chapter)

	return true
}

func GetChapterByTitle(title string, bookId int) (chapter Chapter) {
	db.Where("title = ? AND book_id = ?", title, bookId).First(&chapter)

	return
}

func GetChapterByOrder(order int, bookId int) (chapter Chapter) {
	db.Where("`order` = ? AND book_id = ?", order, bookId).First(&chapter)

	return
}
