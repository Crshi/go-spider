package crawl

import "errors"

type Spider interface {
	CrawlBook(url string) int
}

func NewSpider(from string) (Spider, error) {
	switch from {
	case "52bqg":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}
