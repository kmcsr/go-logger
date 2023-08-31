
package logger_test

import (
	"testing"

	. "github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/logrus"
)

func TestPrefixLogger(t *testing.T){
	PrefixLogger := NewPrefixLogger(logrus.Logger, "Test:")
	PrefixLogger.Print("defaults.PrefixLogger.Print")
	PrefixLogger.Printf("defaults.PrefixLogger.Printf(%v)", 0xffff)
	PrefixLogger.Trace("trace test")
	PrefixLogger.Debug("debug test")
	PrefixLogger.Info("info test")
	PrefixLogger.Warn("warn test")
	PrefixLogger.Error("error test")
}
