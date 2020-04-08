package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

var RedisSetting = &Redis{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//加载配置文件
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//从配置文件加载运行模式
	LoadBase()
	//从配置文件加载端口号，配置读时间过期，写过期时间
	LoadServer()
	//从配置文件JwtSecret，PageSize
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
