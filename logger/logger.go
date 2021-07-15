package logger

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"
)

// Log log
var Log = logrus.New()

// LogError log error
var LogError = logrus.New()

// Debug debug logger
var Debug = Log.Debug

// Debugf debug formatting logger
var Debugf = Log.Debugf

// Info info logger
var Info = Log.Info

// Infof info formatting logger
var Infof = Log.Infof

// Warn warning logger
var Warn = Log.Warn

// Warnf warning formatting logger
var Warnf = Log.Warnf

// Error error logger
var Error = LogError.Error

// Errorf error formatting logger
var Errorf = LogError.Errorf

// Fatal fatal logger
var Fatal = LogError.Fatal

// Fatalf fatal formatting logger
var Fatalf = LogError.Fatalf

// Init init
func Init() {
	Log = logrus.New()
	LogError = logrus.New()

	err := initLogger(Log, "../logs/info.log")
	if err != nil {
		LogError.Error(err)
		LogError.Info("failed to log to file, using default stderr")
	}

	err = initLogger(LogError, "../logs/error.log")
	if err != nil {
		LogError.Error(err)
		LogError.Info("failed to log to file, using default stderr")
	}
}

func initLogger(logger *logrus.Logger, logPath string) error {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	// create dir
	logPath, _ = filepath.Abs(logPath)
	logDir := path.Dir(logPath)
	err := os.MkdirAll(logDir, 0755)
	if os.IsExist(err) {
		stat, err := os.Stat(logDir)
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}

	// create file
	stat, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err2 := os.Create(logPath)
			if err2 != nil {
				return err2
			}
		} else {
			return err
		}
	}
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		logger.Out = logFile
	}

	return nil
}
