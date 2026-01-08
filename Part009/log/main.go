package log

import (
	"fmt"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	cfg "testPackages/config"
	"time"
)

type ILogger interface {
	Log(message string)
}

type ZapLogger struct {
	z *zap.SugaredLogger
}
type FmtLogger struct {
}

func (f FmtLogger) Log(message string) {
	fmt.Println(message)
}

func (z *ZapLogger) Log(message string) {
	z.z.Infow(message)
}

func newZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return &ZapLogger{z: sugar.With("time", time.Now())}
}

func newFmtLogger() *FmtLogger {
	return &FmtLogger{}
}

func NewLogger() ILogger {
	switch cfg.Api.Log.Driver {
	case "fmt":
		return newFmtLogger()
	case "zap":
		return newZapLogger()
	}
	return nil
}