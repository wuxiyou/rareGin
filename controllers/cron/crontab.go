package cron

import (
	cron2 "github.com/robfig/cron"
	"fmt"
	"rareGin/utils"
)

type Crontab struct {

}

type Test2Job struct {
}

func (this Test2Job)Run() {
	fmt.Println("testJob2...")
}

func (c *Crontab)CronBegin()  {
	cron01 := cron2.New()
	// 从配置文件读取
	// 此程序是阻塞进行的
	cron01.AddFunc(utils.ReadConfigInfo("cron","specCrontab"), func() {
		fmt.Println("cron666666")
	})

	cron01.AddJob(utils.ReadConfigInfo("cron", "specTest2Job"), Test2Job{})

	cron01.Start()

	/*t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}*/

	defer cron01.Stop()

	select {}
}
