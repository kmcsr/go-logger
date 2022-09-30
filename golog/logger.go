
package golog_logger

import (
	"fmt"
	"io"

	"github.com/kataras/golog"
	"github.com/kmcsr/go-logger"
)

var Logger = initGologLogger()

func initGologLogger()(logger.Logger){
	return &GologWrapper{golog.Default}
}

func New()(logger.Logger){
	return &GologWrapper{golog.New()}
}

type GologWrapper struct{
	*golog.Logger
}

func (l *GologWrapper)Unwrap()(*golog.Logger){
	return l.Logger
}

func (l *GologWrapper)SetLevel(lvl logger.Level){
	l.Logger.Level = level2golog(lvl)
}

func (l *GologWrapper)Level()(logger.Level){
	return golog2level(l.Logger.Level)
}

func (l *GologWrapper)SetOutput(w io.Writer){
	l.Logger.SetOutput(w)
}

func (l *GologWrapper)Trace(v ...any){
	l.Print(v...)
}

func (l *GologWrapper)Tracef(format string, v ...any){
	l.Printf(format, v...)
}

func (l *GologWrapper)Panic(v ...any){
	l.Print(v...)
	panic(fmt.Sprint(v...))
}

func (l *GologWrapper)Panicf(format string, v ...any){
	l.Printf(format, v...)
	panic(fmt.Sprintf(format, v...))
}

func Unwrap(l logger.BasicLogger)(*GologWrapper){
	return logger.Unwrap(l).(*GologWrapper)
}

func level2golog(lvl logger.Level)(golog.Level){
	switch lvl {
	case logger.PanicLevel: return golog.DisableLevel
	case logger.FatalLevel: return golog.FatalLevel
	case logger.ErrorLevel: return golog.ErrorLevel
	case logger.WarnLevel:  return golog.WarnLevel
	case logger.InfoLevel:  return golog.InfoLevel
	case logger.DebugLevel: fallthrough
	case logger.TraceLevel: return golog.DebugLevel
	}
	return golog.DebugLevel
}

func golog2level(lvl golog.Level)(logger.Level){
	switch lvl {
	case golog.DisableLevel: return logger.PanicLevel
	case golog.FatalLevel: return logger.FatalLevel
	case golog.ErrorLevel: return logger.ErrorLevel
	case golog.WarnLevel:  return logger.WarnLevel
	case golog.InfoLevel:  return logger.InfoLevel
	case golog.DebugLevel: return logger.TraceLevel
	}
	panic("Unknown level")
}
