module github.com/Crshi/go-spider

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.54.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.8 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/gomodule/redigo v2.0.1-0.20180401191855-9352ab68be13+incompatible
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200409092240-59c9f1ba88fa // indirect
	golang.org/x/tools v0.0.0-20200410194907-79a7a3126eef // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace (
	github.com/Crshi/go-spider/conf => ./conf
	github.com/Crshi/go-spider/docs => ./docs
	github.com/Crshi/go-spider/models => ./models
	github.com/Crshi/go-spider/pkg/e => ./pkg/e
	github.com/Crshi/go-spider/pkg/logging => ./pkg/logging
	github.com/Crshi/go-spider/pkg/path => ./pkg/path
	github.com/Crshi/go-spider/pkg/setting => ./pkg/setting
	github.com/Crshi/go-spider/routers => ./routers
	github.com/Crshi/go-spider/routers/api => ./routers/api
	github.com/Crshi/go-spider/service/cache_service => ./service/cache_service
	github.com/Crshi/go-spider/utils/crawl => ./utils/crawl
	github.com/Crshi/go-spider/utils/mahonia => ./utils/mahonia
	github.com/Crshi/go-spider/utils/response => ./utils/response
)
