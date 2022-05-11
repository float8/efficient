package efficient

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	db "github.com/whf-sky/efficient/database"
	"math/rand"
	"os"
	"strconv"
	"time"
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

	//database Logger
	db.Logger = logger

	return logger
}()

var Log = func() *logrus.Entry {
	log := logrus.NewEntry(Logger)

	//database Log
	db.Log = log
	return log
}()

func SetLogger(fun func(logger *logrus.Logger, log *logrus.Entry)) {
	fun(Logger, Log)
	db.SetLogger(func(logger *logrus.Logger, log *logrus.Entry) {
		logger = Logger
	})
}

func ginLogger(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	id := strconv.FormatUint(rand.Uint64(), 10)

	Log = Log.WithFields(logrus.Fields{
		"id":        id,
	})

	Log.WithFields(logrus.Fields{
		"client_ip": ctx.ClientIP(),
		"method":    ctx.Request.Method,
		"uri":       ctx.Request.RequestURI,
		"status":    ctx.Writer.Status(),
		"agent":     ctx.Request.UserAgent(),
	}).Info("ok")
}