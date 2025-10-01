package error_logs

import (
	"log/slog"
	"github.com/arun14k08/goframework/logging"
)



var (
	BAD_REQUEST = logging.LogTemplate{Level: slog.LevelInfo, Message: "Error in parsing the payload/param %v"}
)