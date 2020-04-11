package crawl

import (
	"strings"

	"github.com/Crshi/go-spider/models"
	"github.com/Crshi/go-spider/pkg/logging"
	"github.com/Crshi/go-spider/utils/mahonia"
	"github.com/PuerkitoBio/goquery"
)

type BookTextSpider struct {
}

func (bs *BookTextSpider) CrawlBook(url string) bool {

	doc, err := goquery.NewDocument(url)

	if err != nil {
		logging.Fatal("get document error: ", err.Error())
		return false
	}

	bookname := mahonia.GbkToUtf8(doc.Find("#info h1").Text())
	author := mahonia.GbkToUtf8(doc.Find("#info p a").First().Text())

	book := models.GetBookByName(bookname)
	book.Author = author
	book.Name = bookname

	if book.Id == 0 {
		book = models.CreateBook(book)
	} else {
		logging.Fatal("book alredy exist!")
		return false
	}

	channel := make(chan struct{}, 100)

	doc.Find("#list dd").Each(func(i int, contentSelection *goquery.Selection) {
		title := mahonia.GbkToUtf8(contentSelection.Find("a").Text())
		href, _ := contentSelection.Find("a").Attr("href")
		chapter := models.Chapter{
			Title:   title,
			Order:   i + 1,
			Book_Id: book.Id,
			Url:     url + href,
		}

		channel <- struct{}{}
		go CrawlChaptersInfo(chapter, channel)
	})

	for i := 0; i < 100; i++ {
		channel <- struct{}{}
	}
	close(channel)

	return true
}

type ChanTag struct{}

func CrawlChaptersInfo(chapter models.Chapter, c chan struct{}) {
	defer func() { <-c }()
	doc, err := goquery.NewDocument(chapter.Url)
	if err != nil {
		logging.Error("get chapter details error: ", err.Error())
		return
	}
	content := doc.Find("#content").Text()
	content = mahonia.GbkToUtf8(content)
	content = strings.Replace(content, "聽", " ", -1)
	chapter.Content = content
	models.AddChapter(chapter)
}
