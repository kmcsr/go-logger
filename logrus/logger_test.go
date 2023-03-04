
package logrus_test

import (
	"testing"

	_ "github.com/sirupsen/logrus"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/logrus"
)

var LogrusLogger logger.Logger = logrus.Logger

func TestLogrusLogger(t *testing.T){
	LogrusLogger.SetLevel(logger.TraceLevel)
	LogrusLogger.Print("defaults.LogrusLogger.Print")
	LogrusLogger.Printf("defaults.LogrusLogger.Printf(%v)", 0xffff)
	LogrusLogger.Trace("trace test")
	LogrusLogger.Debug("debug test")
	LogrusLogger.Info("info test")
	LogrusLogger.Warn("warn test")
	LogrusLogger.Error("error test")
}
