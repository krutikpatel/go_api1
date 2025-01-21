package logger

import (
	"api1/config"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func EnableFileLogging(log *logrus.Logger, logConfig config.LogConfig) {
	if logConfig.EnableFile {
		file, err := os.OpenFile(logConfig.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		log.SetOutput(file)
	}

	switch logConfig.Level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	switch logConfig.Format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{})
	}
}

/*
// Configure sets up logrus with custom formatting and settings
func Configure(env string) *logrus.Logger {
	log := logrus.New()

	// Set log level based on environment
	switch strings.ToLower(env) {
	case "production":
		log.SetLevel(logrus.InfoLevel)
	case "development":
		log.SetLevel(logrus.DebugLevel)
	case "testing":
		log.SetLevel(logrus.DebugLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// Custom formatter with caller info
	log.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Split(f.File, "/")
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename[len(filename)-1], f.Line)
		},
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
		TimestampFormat: time.RFC3339,
	})

	// Enable caller info
	log.SetReportCaller(true)

	// Set output to stdout
	log.SetOutput(os.Stdout)

	return log
}

// Fields type for structured logging
type Fields logrus.Fields

// Logger wrapper around logrus.Logger
type Logger struct {
	*logrus.Logger
}

// NewLogger creates a new logger instance
func NewLogger(env string) *Logger {
	return &Logger{Configure(env)}
}

// WithFields creates an entry with fields
func (l *Logger) WithFields(fields Fields) *logrus.Entry {
	return l.Logger.WithFields(logrus.Fields(fields))
}
*/
