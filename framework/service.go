package framework

import "github.com/arun14k08/goframework/config"

type FrameWorkService struct{}

func (fws FrameWorkService) Name() string {
	return  "FrameWorkServie"
}

func (fws FrameWorkService) Start() error {
	// init
	config.InitializeAppProps()
	return nil
}