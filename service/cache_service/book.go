package cache_service

import "strconv"

func GetBookKey(id int) string {
	return "Book_" + strconv.Itoa(id)
}
