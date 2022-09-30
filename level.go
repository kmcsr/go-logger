
package logger

type Level uint32
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel

	allLevel
)

func (l Level)String()(string){
	switch l {
	case PanicLevel: return "PANIC"
	case FatalLevel: return "FATAL"
	case ErrorLevel: return "ERROR"
	case WarnLevel:  return "WARN"
	case InfoLevel:  return "INFO"
	case DebugLevel: return "DEBUG"
	case TraceLevel: return "TRACE"
	}
	panic("Unknown level")
}
