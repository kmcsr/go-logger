
package std_logger_test

import (
	"testing"

	"github.com/kmcsr/go-logger"
	stdl "github.com/kmcsr/go-logger/std"
)

var StdLogger logger.Logger = stdl.Logger

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
