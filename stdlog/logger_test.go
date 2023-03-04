
package stdlog_test

import (
	"testing"

	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/stdlog"
)

var StdLogger logger.Logger = stdlog.Logger

func TestStdLogger(t *testing.T){
	StdLogger.SetLevel(logger.TraceLevel)
	StdLogger.Print("defaults.StdLogger.Print")
	StdLogger.Printf("defaults.StdLogger.Printf(%v)", 0xffff)
	StdLogger.Trace("trace test")
	StdLogger.Debug("debug test")
	StdLogger.Info("info test")
	StdLogger.Warn("warn test")
	StdLogger.Error("error test")
}
