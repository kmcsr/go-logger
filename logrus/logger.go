
package logrus_logger

import (
	"github.com/sirupsen/logrus"
	"github.com/kmcsr/go-logger"
)

var Logger = initLogrusLogger()

type LogrusWrapper struct{
	*logrus.Logger
}

var _ logger.Logger = (*LoggerWrapper)(nil)

func initLogrusLogger()(*LogrusWrapper){
	return &LogrusWrapper{logrus.StandardLogger()}
}

func New()(*LogrusWrapper){
	return &LogrusWrapper{logrus.New()}
}

func (l *LogrusWrapper)SetLevel(lvl logger.Level){
	l.Logger.SetLevel(level2logrus(lvl))
}

func (l *LogrusWrapper)Level()(logger.Level){
	return logrus2level(l.Logger.GetLevel())
}

func Unwrap(l logger.BasicLogger)(*LogrusWrapper){
	return logger.Unwrap(l).(*LogrusWrapper)
}

func level2logrus(lvl logger.Level)(logrus.Level){
	switch lvl {
	case logger.PanicLevel: return logrus.PanicLevel
	case logger.FatalLevel: return logrus.FatalLevel
	case logger.ErrorLevel: return logrus.ErrorLevel
	case logger.WarnLevel:  return logrus.WarnLevel
	case logger.InfoLevel:  return logrus.InfoLevel
	case logger.DebugLevel: return logrus.DebugLevel
	case logger.TraceLevel: return logrus.TraceLevel
	}
	return logrus.TraceLevel
}

func logrus2level(lvl logrus.Level)(logger.Level){
	switch lvl {
	case logrus.PanicLevel: return logger.PanicLevel
	case logrus.FatalLevel: return logger.FatalLevel
	case logrus.ErrorLevel: return logger.ErrorLevel
	case logrus.WarnLevel:  return logger.WarnLevel
	case logrus.InfoLevel:  return logger.InfoLevel
	case logrus.DebugLevel: return logger.DebugLevel
	case logrus.TraceLevel: return logger.TraceLevel
	}
	panic("Unknown level")
}
