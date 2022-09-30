
package logger

import (
	"io"
)

type (
	BasicLogger interface{
		SetOutput(w io.Writer)

		Print(v ...any)
		Printf(format string, v ...any)
	}
	DebugLogger interface{
		Debug(v ...any)
	}
	DebugfLogger interface{
		Debugf(format string, v ...any)
	}
	InfoLogger interface{
		Info(v ...any)
	}
	InfofLogger interface{
		Infof(format string, v ...any)
	}
	WarnLogger interface{
		Warn(v ...any)
	}
	WarnfLogger interface{
		Warnf(format string, v ...any)
	}
	ErrorLogger interface{
		Error(v ...any)
	}
	ErrorfLogger interface{
		Errorf(format string, v ...any)
	}
	TraceLogger interface{
		Trace(v ...any)
	}
	TracefLogger interface{
		Tracef(format string, v ...any)
	}
	FatalLogger interface{
		Fatal(v ...any)
	}
	FatalfLogger interface{
		Fatalf(format string, v ...any)
	}
	PanicLogger interface{
		Panic(v ...any)
	}
	PanicfLogger interface{
		Panicf(format string, v ...any)
	}
	LevelsLogger interface{
		SetLevel(Level)
		Level()(Level)
	}

	Logger interface{
		BasicLogger
		DebugLogger
		DebugfLogger
		InfoLogger
		InfofLogger
		WarnLogger
		WarnfLogger
		ErrorLogger
		ErrorfLogger
		TraceLogger
		TracefLogger
		FatalLogger
		FatalfLogger
		PanicLogger
		PanicfLogger

		LevelsLogger
	}
)

func Unwrap(l BasicLogger)(v any){
	if u, ok := l.(interface{ Unwrap()(BasicLogger) }); ok {
		return Unwrap(u.Unwrap())
	}
	if u, ok := l.(interface{ Unwrap()(any) }); ok {
		return u.Unwrap()
	}
	return l
}

