package database

import (
	"github.com/sirupsen/logrus"
	"os"
)


var Logger = func() *logrus.Logger {
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat : "2006-01-02 15:04:05.000",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.TraceLevel)

	//to log caller info
	logger.SetReportCaller(true)

	return logger
}()

var Log = func() *logrus.Entry {
	return logrus.NewEntry(Logger)
}()

func SetLogger(fun func(logger *logrus.Logger, log *logrus.Entry)) {
	fun(Logger, Log)
}

