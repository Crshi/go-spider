# go-spider

![在线地址](http://47.94.164.184:8089/swagger/index.html)

* 介绍

1.基于Gin框架搭建，使用Mysql作为数据库
2.使用Redis为Chapter做缓存
3.goquery解析html,使用mahonia完成编码格式转换。
4.多线程爬虫
5.Swagger作为接口文档
6.Docker部署

* 使用
以 https://www.52bqg.com/book_127354/ 为例
1./api/v1/crawl 爬取数据 bookSiteId:book_127354 threadCount:100
2./api/v1/books 查看爬取的book
3./api/v1/chapter 根据BookId，Order 获取章节
