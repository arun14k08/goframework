package logging

import (
	"fmt"
	"log"
	"log/slog"
)

var (
	SERVICE_INFO = LogTemplate{slog.LevelInfo, "%s [%s]"}
)

type LogTemplate struct {
	Level slog.Level
	Message string
}

func (lt LogTemplate) Log(arges... interface{}) {
	log.Printf("[%s] %s", lt.Level.String(), fmt.Sprintf(lt.Message, arges...))
}

func Log(Level slog.Level, Message string) {
	log.Printf("[%s] %s", Level.String(), Message)
}

func LogMessage(Message string) {
	log.Printf("%s",  Message)
}

