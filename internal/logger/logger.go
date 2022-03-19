package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/defry256/pokemon-helper/internal/errors"
	"go.uber.org/zap"
)

func Error(errMsg string, err error) {
	err = errors.New(errMsg, err)
	config := zap.NewProductionConfig()
	logger, _ := config.Build()
	trace := generateTrace(err)
	logger.Info(
		errMsg,
		zap.String("trace", trace),
	)
	logMessage(fmt.Sprintf("%s -> %s", errMsg, trace))
}

func Print(message string) {
	logMessage(message)
	fmt.Println(message)
}

func logMessage(message string) {
	filePath, _ := filepath.Abs(fmt.Sprintf("../logs/pokemon-helper_%s.log", time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(time.Now().Format("15:04") + ": " + message)
}

func generateTrace(err error) string {
	trace := ""
	locations := errors.TracedToStackTrace(err)
	for _, loc := range locations {
		trace += fmt.Sprintf("%s:%d\n", loc.File(), loc.Line())
	}

	return trace
}
