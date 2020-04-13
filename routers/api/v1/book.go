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
// @Summary Get all books
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
// @Summary Crawl Book
// @Produce  json
// @Param bookSiteId query string true "BookSiteId"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/crawl [post]
func CrawlBook(c *gin.Context) {
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

// 获取章节
// @Summary GetChapter
// @Produce  json
// @Param bookid query int false "BookId"
// @Param order query int false "Order"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/chapter [get]
func GetChapter(c *gin.Context) {
	//从redis取数据
	//如果没有从数据库取数据并更新到redis

	bookid, _ := strconv.Atoi(c.Query("bookid"))
	order, _ := strconv.Atoi(c.Query("order"))
	key := cache_service.GetChapterKey(bookid, order)
	res := make(map[string]string)
	res["bookid"] = strconv.Itoa(bookid)
	res["order"] = strconv.Itoa(order)

	//有缓存
	if gredis.Exists(key) {
		//获取缓存
		tmpContent, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		}
		res["content"] = string(tmpContent)
	} else {
		chapter := models.GetChapterByOrder(order, bookid)
		res["content"] = chapter.Content
		gredis.Set(key, chapter.Content, 3600)
	}

	gredis.SetExpire(key, 3600)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": res,
	})
}
