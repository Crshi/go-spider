package routers

import (
	"github.com/Crshi/go-spider/pkg/setting"
	v1 "github.com/Crshi/go-spider/routers/api/v1"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//爬取小说
		apiv1.GET("/crawl", v1.CrawlBooks)
		//下载小说
		apiv1.GET("/download", v1.DownloadBooks)
	}

	return r
}
