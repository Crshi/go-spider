package models

type Chapter struct {
	Model

	Title   string `json:Title`
	Content string `json:"Content"`
	BookId  int    `json:"BookId"`
	Order   int    `json:"Order"`
	Url     string `json:Url`
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
