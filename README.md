go 依赖管理- govendor
Installation
---- go get -u github.com/kardianos/govendor

Sub-commands
-- init   创建 vendor 文件夹和vendor.json
-- list   列出已经存在的依赖包
-- add    从$GOPATH 中添加依赖包，会加到 vendor.json
-- update 从$GOPATH 升级依赖包
-- remove 从$GOPATH 中删除依赖
-- status 从本地丢失的, 过期的和修改的package
-- fetch 从远端库增加新的，或者更新 vendor 文件中的依赖包
-- sync  Pall package into vendor folder from remote repository with revisions
-- get
-- license
-- shell


* 判断某个参数是否存在 , 使用 FromValue()
  r.FromValue("id") == ""
* 判断某参数是否有值，使用 FromValue()
  len (r.FromValue("id")) > 0

* 得知某段代码的执行位置
  _, file, line, _ := runtime.Caller(0)



Cron 表达格式
* * * * * *  分别代表的是 秒 分 时 天 月 星期 【0-59  0-59  0-23 1-31 1-12 0-6】

