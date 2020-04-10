module github.com/Crshi/go-spider

go 1.14

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.54.0
	github.com/jinzhu/gorm v1.9.12
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace (
	github.com/Crshi/go-spider/conf => ./conf
	github.com/Crshi/go-spider/docs => ./docs
	github.com/Crshi/go-spider/models => ./models
	github.com/Crshi/go-spider/pkg/e => ./pkg/e
	github.com/Crshi/go-spider/pkg/logging => ./pkg/logging
	github.com/Crshi/go-spider/pkg/setting => ./pkg/setting
	github.com/Crshi/go-spider/routers => ./routers
	github.com/Crshi/go-spider/routers/api => ./routers/api
)
