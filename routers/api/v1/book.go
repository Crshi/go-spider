package v1

import (
	"net/http"

	"github.com/Crshi/go-spider/models"
	"github.com/Crshi/go-spider/pkg/e"
	"github.com/Crshi/go-spider/pkg/logging"
	"github.com/Crshi/go-spider/utils/crawl"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	results := make(map[string]interface{})
	// maps := make(map[string] interface{})

	code := e.INVALID_PARAMS
	code = e.SUCCESS

	results["lists"] = models.GetBooks()
	// results["total"] = models.GetBookCount()

	c.JSON(http.StatusOK, gin.H{
		"Code":    code,
		"Msg":     e.GetMsg(code),
		"Results": results,
	})
}

//爬取Book
func CrawlBooks(c *gin.Context) {
	bookSiteId := c.Query("bookSiteId")

	if bookSiteId == "" {
		bookSiteId = "book_10459"
	}

	bookUrl := "https://www.52bqg.com/" + bookSiteId + "/"

	s, err := crawl.NewSpider("52bqg")
	if err != nil {
		logging.Fatal("new Spider error: ", err.Error())
	}

	s.CrawlBook(bookUrl)
}

//下载Book
func DownloadBooks(c *gin.Context) {
	//从redis取数据
	//如果没有从数据库取数据并更新到redis
	//返回
}
