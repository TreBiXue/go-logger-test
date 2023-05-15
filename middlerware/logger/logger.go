package logger

import (
	"cloud.google.com/go/logging"
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var rootDir string

type Logger interface {
	Info(msg string)
	Error(msg string)
}

// LocalLogger 是本地环境下的日志记录器
type LocalLogger struct {
	*zap.SugaredLogger
}

// Info 实现了 Logger 接口的 Info 方法
func (l *LocalLogger) Info(msg string) {
	l.SugaredLogger.Infof("[Local] %s", msg)
}

// Error 实现了 Logger 接口的 Error 方法
func (l *LocalLogger) Error(msg string) {
	l.SugaredLogger.Errorf("[Local] %s", msg)
}

// GCPLogger 是GCP环境下的日志记录器
type GCPLogger struct {
	*logging.Logger
}

// Info 实现了 Logger 接口的 Info 方法
func (l *GCPLogger) Info(msg string) {
	l.Logger.Log(logging.Entry{Severity: logging.Info, Payload: msg})
}

// Error 实现了 Logger 接口的 Error 方法
func (l *GCPLogger) Error(msg string) {
	l.Logger.Log(logging.Entry{Severity: logging.Error, Payload: msg})
}

func InitLogger() (Logger, error) {
	var mode string
	flag.StringVar(&mode, "mode", "dev", "mode flag")
	flag.Parse()
	if mode == "dev" {
		fmt.Println("aaaaaaaaaaaaaa")
		return initDevelopmentLogger()
	} else {
		fmt.Println("bbbbbbbbbbbbbbbbb")
		return initProductionLogger()
	}
}

func initDevelopmentLogger() (Logger, error) {
	// logger path
	logFilePath := filepath.Join(".", "logger", "log.txt")
	err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.OutputPaths = []string{logFilePath}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &LocalLogger{logger.Sugar()}, nil
}

func initProductionLogger() (Logger, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "horizontal-ally-385009")
	if err != nil {
		fmt.Println("ccccccccccccccccccc")
		return nil, err
	}

	loggingClient := client.Logger("test-logger")

	// Logs a basic entry.
	loggingClient.Log(logging.Entry{Payload: "hello world"})
	fmt.Println("hello world ddddddddddddddddddd")

	return &GCPLogger{loggingClient}, nil
}
