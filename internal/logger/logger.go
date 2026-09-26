package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var logMutex sync.Mutex

func Log(level string, context string, message string) {
	logMutex.Lock()
	defer logMutex.Unlock()

	logLine := fmt.Sprintf(
		"[%s] %-5s (%s): %s\n",
		time.Now().Format(time.RFC822),
		level,
		context,
		message,
	)

	fmt.Print(logLine)

	f, err := os.OpenFile("fablebound.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(logLine)
}

func LogError(context string, err error) {
	if err == nil {
		return
	}
	Log("ERROR", context, err.Error())
}

func LogInfo(context string, message string) {
	Log("INFO", context, message)
}
