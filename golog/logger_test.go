
package golog_test

import (
	"testing"

	_ "github.com/kataras/golog"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/golog"
)

var GologLogger logger.Logger = golog.Logger

func TestGologLogger(t *testing.T){
	GologLogger.SetLevel(logger.TraceLevel)
	GologLogger.Print("defaults.GologLogger.Print")
	GologLogger.Printf("defaults.GologLogger.Printf(%v)", 0xffff)
	GologLogger.Trace("trace test")
	GologLogger.Debug("debug test")
	GologLogger.Info("info test")
	GologLogger.Warn("warn test")
	GologLogger.Error("error test")
}
