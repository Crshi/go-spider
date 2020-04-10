package models

import "strconv"

type Chapter struct {
	Model

	Content string `json:"Content"`
	BookId  int    `json:"BookId"`
}

func GetChapters(bookId int) (chapters []Chapter) {
	db.Where("BookId = ?", bookId).Find(&chapters)
	return
}

func AddChapters(datas []string, bookId int) bool {
	sql := "Insert into `chapters` (`BookId`,`Content`) Values"
	for _, v := range datas {
		sql += "(" + strconv.Itoa(bookId) + ",'" + v + "'),"
	}

	sql = sql[:len(sql)-1] + ";"

	db.Exec(sql)

	return true
}
