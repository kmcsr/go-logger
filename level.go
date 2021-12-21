
package gologger

type Level int8

const (
	_ Level = iota
	LEVEL_DEBUG
	LEVEL_DISABLE
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_STACK
)

func (l Level)String()(string){
	switch l {
	case LEVEL_DEBUG:   return "DBUG"
	case LEVEL_DISABLE: return "<DISABLE>"
	case LEVEL_INFO:    return "INFO"
	case LEVEL_WARN:    return "WARN"
	case LEVEL_ERROR:   return "EROR"
	case LEVEL_STACK:   return "STAK"
	}
	return "<UNKNOWN>"
}

