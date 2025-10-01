package framework

import (
	"github.com/arun14k08/goframework/config"
	_ "github.com/lib/pq"
)

type FrameWorkService struct{}

func (fws FrameWorkService) Name() string {
	return  "FrameWorkService"
}

func (fws FrameWorkService) Start() error {
	// init
	config.InitializeAppProps()
	return nil
}