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

func (bs *BookTextSpider) CrawlBook(url string, thread int) int {

	doc, err := goquery.NewDocument(url)

	if err != nil {
		logging.Fatal("get document error: ", err.Error())
		return 0
	}

	bookname := mahonia.GbkToUtf8(doc.Find("#info h1").Text())
	author := mahonia.GbkToUtf8(doc.Find("#info p a").First().Text())

	book := models.GetBookByName(bookname)
	book.Author = author
	book.Name = bookname

	if book.Id == 0 {
		book = models.CreateBook(book)
	}

	channel := make(chan struct{}, thread)

	order := 0

	doc.Find("#list dd").Each(func(i int, contentSelection *goquery.Selection) {
		title := mahonia.GbkToUtf8(contentSelection.Find("a").Text())
		if title != "" {
			order++
			href, _ := contentSelection.Find("a").Attr("href")
			chapter := models.Chapter{
				Title:   title,
				Order:   order,
				Book_Id: book.Id,
				Url:     url + href,
			}

			channel <- struct{}{}
			go CrawlChaptersInfo(chapter, channel)
		}
	})

	for i := 0; i < thread; i++ {
		channel <- struct{}{}
	}
	close(channel)

	return book.Id
}

func CrawlChaptersInfo(chapter models.Chapter, c chan struct{}) {
	defer func() { <-c }()

	checkChapter := models.GetChapterByTitle(chapter.Title, chapter.Book_Id)

	if checkChapter.Id == 0 {
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
}
