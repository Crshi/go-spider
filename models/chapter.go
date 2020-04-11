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
	db.Where("BookId = ?", bookId).Find(&chapters)
	return
}

func AddChapter(chapter Chapter) bool {
	db.Create(&chapter)

	return true
}

func CheckChapter(order int, bookId int) (count int) {
	db.Model(&Chapter{}).Where("Order = ? AND BookId = ?", order, bookId).Count(&count)

	return
}
