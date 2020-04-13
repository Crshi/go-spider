package cache_service

import "strconv"

// Chapter Zset
func GetChapterKey(bookId int, chapterOrder int) string {
	return "Book_" + strconv.Itoa(bookId) + "_Chapter_" + strconv.Itoa(chapterOrder)
}
