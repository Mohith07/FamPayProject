package controller

import (
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron"
)

var c *cron.Cron

func InitCronJob() {
	log.Info("Create new cron")
	c = cron.New()
	err := c.AddFunc("@every 1m", Init)
	if err != nil {
		log.Info("error adding init func to cron " + err.Error())
		return 
	}
	c.Start()
}
