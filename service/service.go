package service

import (
	"log"
	"github.com/arun14k08/goframework/logging"
)


type Service interface {
	Name() string
	Start() error
}

func Start(service Service) {
	err := service.Start()
	if err != nil {
		logging.SERVICE_INFO.Log(service.Name(), "STOPPED")
		log.Fatal("err -> ", err)
	}
	logging.SERVICE_INFO.Log(service.Name(), "STARTED")
}