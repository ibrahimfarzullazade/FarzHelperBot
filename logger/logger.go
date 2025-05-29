package logger

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Log acmadi: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(multiWriter, "BOT_LOG: ", log.LstdFlags|log.Lshortfile)
}
func Log(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
