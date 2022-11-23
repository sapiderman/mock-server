package logger

import (
	"context"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sapiderman/mock-server/internal/config"
	"github.com/sapiderman/mock-server/internal/contextkeys"
	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

// ConfigureLogging set logging lever from config
func ConfigureLogging() {
	lLevel := config.Get("server.log.level")
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting log level to: ", lLevel)
	switch strings.ToUpper(lLevel) {
	default:
		log.Info("Unknown level [", lLevel, "]. Log level set to ERROR")
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	}

	currentTime := time.Now()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/payment-" + currentTime.Format("2006-01-02") + ".log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
		Level:      log.GetLevel(),
		Formatter: &log.JSONFormatter{
			TimestampFormat: time.RFC3339,
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     false,
	})
	log.AddHook(rotateFileHook)
}

// GetRequestID will get reqID from a http request and return it as a string
func GetRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextkeys.XRequestID)
	if id, ok := reqID.(string); ok {
		return id
	}
	return "-"
}

// GetLogger returns our logger with some custom fields
func GetLogger(ctx context.Context, logf *log.Entry, key string, value interface{}) *log.Entry {
	return logf.WithFields(log.Fields{key: value, "request-id": GetRequestID(ctx)})
}
