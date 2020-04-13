package v1

import (
	"net/http"
	"strconv"

	"github.com/Crshi/go-spider/models"
	"github.com/Crshi/go-spider/pkg/e"
	"github.com/Crshi/go-spider/pkg/gredis"
	"github.com/Crshi/go-spider/pkg/logging"
	"github.com/Crshi/go-spider/service/cache_service"
	"github.com/Crshi/go-spider/utils/crawl"
	"github.com/gin-gonic/gin"
)

//获取所有books
// @Summary Get multiple books
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/books [get]
func GetBooks(c *gin.Context) {
	results := make(map[string]interface{})

	code := e.SUCCESS

	results["lists"] = models.GetBooks()

	c.JSON(http.StatusOK, gin.H{
		"Code":    code,
		"Msg":     e.GetMsg(code),
		"Results": results,
	})
}

//爬取Book
// @Summary CrawlBooks
// @Produce  json
// @Param bookSiteId query string true "BookSiteId"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/crawl [post]
func CrawlBooks(c *gin.Context) {
	res := make(map[string]int)
	bookSiteId := c.Query("bookSiteId")

	if bookSiteId == "" {
		return
	}

	bookUrl := "https://www.52bqg.com/" + bookSiteId + "/"

	s, err := crawl.NewSpider("52bqg")
	if err != nil {
		logging.Fatal("new Spider error: ", err.Error())
	}

	code := e.SUCCESS

	res["bookid"] = s.CrawlBook(bookUrl)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": res,
	})
}

//下载Book
// @Summary DownloadBooks
// @Produce  json
// @Param bookid query int false "BookId"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/download [post]
func DownloadBooks(c *gin.Context) {
	//从redis取数据
	//如果没有从数据库取数据并更新到redis

	bookid, _ := strconv.Atoi(c.Query("bookid"))
	book := models.GetBookById(int(bookid))
	content := ""
	key := cache_service.GetBookKey(book.Id)

	if gredis.Exists(key) {
		tmpContent, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		}
		content = string(tmpContent)
	} else {
		chapters := models.GetChapters(bookid)

		for _, v := range chapters {
			content += v.Content
		}
	}

	gredis.Set(key, content, 3600)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": content,
	})
}
