package routers

import (
	_ "github.com/Crshi/go-spider/docs"
	"github.com/Crshi/go-spider/pkg/setting"
	v1 "github.com/Crshi/go-spider/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		apiv1.GET("/books", v1.GetBooks)
		//爬取小说
		apiv1.POST("/crawl", v1.CrawlBooks)
		//下载小说
		apiv1.POST("/download", v1.DownloadBooks)
	}

	return r
}
