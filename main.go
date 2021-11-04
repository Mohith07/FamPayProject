package main

import (
	"FamPayProject/clients"
	"FamPayProject/config"
	"FamPayProject/controller"
	"FamPayProject/model"
	"FamPayProject/server"
)

func main() {
	//read the config
	config.Init()
	//connect to db
	clients.Init()
	//create tables in db
	model.Init()
	//start background cron
	controller.InitCronJob()
	//start rest server
	server.Init()
}
