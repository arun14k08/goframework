package logging

type LogService struct{}

func (ls LogService) Name() string {
	return "LogService"
}

func (ls LogService) Start() error {
	// init
	return nil
}