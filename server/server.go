package server

import (
	"log"
	"sync"

	"FamPayProject/config"
)

var onceRest sync.Once

// Init function to initialize the service
func Init() {
	onceRest.Do(func() {
		conf := config.GetConfig()
		log.Println("starting a new server")
		r := NewRouter()
		if err := r.Start(conf.RestServerPort); err != nil {
			log.Fatal("Unable to bring service up: " + err.Error())
		}
	})
}
