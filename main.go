package main

import (
	"rareGin/routes"
	//"github.com/gin-gonic/gin"

	"rareGin/controllers/cron"
	"rareGin/utils"
	"strconv"
)

func main()  {
	//gin.SetMode(gin.ReleaseMode)  //生产模式

	// 启动定时任务
	cronFlag, _:= strconv.ParseBool(utils.ReadConfigInfo("cronFlag", "cronFlagInfo"))
	 if cronFlag == true {
		 cronController := new (cron.Crontab)
		 cronController.CronBegin()
	 }
	routes.Init()

}
