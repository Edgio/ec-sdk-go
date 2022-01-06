package ectest

import (
	"fmt"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

func GetTestLogger(prefix string) logging.Logger {
	timeNowUTC := time.Now().UTC()
	timeStamp := timeNowUTC.Format("20060102150405")
	filePath := fmt.Sprintf("%s_%s.log", prefix, timeStamp)
	return logging.NewFileLogger(filePath)
}
