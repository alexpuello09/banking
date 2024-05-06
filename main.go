package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("Starting the application..")
	//log.Println("Starting application")
	app.Start()
}
