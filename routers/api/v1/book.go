package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Crshi/go-spider/pkg/logging"

	"github.com/Crshi/go-spider/models"
	"github.com/Crshi/go-spider/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	results := make(map[string]interface{})
	// maps := make(map[string] interface{})

	code := e.INVALID_PARAMS
	code = e.SUCCESS

	results["lists"] = models.GetBooks()
	results["total"] = models.GetBookCount()

	c.JSON(http.StatusOK, gin.H{
		"Code":    code,
		"Msg":     e.GetMsg(code),
		"Results": results,
	})
}

//爬取Book
func CrawlBooks(c *gin.Context) {
	bookSiteId := c.Query("bookSiteId")

	resp, err := http.Get("https://www.52bqg.com/" + bookSiteId + "/")

	if err != nil {
		logging.Info("err:Crawl Book Filed, err.message: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Info("err:Book IO Read Filed, err.message: %s", err)
		return
	}
	fmt.Println(string(body))
	//根据bookSiteId爬取数据
	//create book
	//create chapters
	//return
}

//下载Book
func DownloadBooks(c *gin.Context) {
	//从redis取数据
	//如果没有从数据库取数据并更新到redis
	//返回
}
