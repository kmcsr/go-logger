
package logger

import (
	"fmt"
	"os"
)

type LoggerWrapper struct{
	BasicLogger
	AddPrefixOnMissing bool
	level Level
}

var _ Logger = (*LoggerWrapper)(nil)

func WrapLogger(bl BasicLogger)(*LoggerWrapper){
	if l, ok := bl.(*LoggerWrapper); ok {
		return l
	}
	return &LoggerWrapper{
		BasicLogger: bl,
		AddPrefixOnMissing: true,
		level: InfoLevel,
	}
}

func (l *LoggerWrapper)Unwrap()(BasicLogger){
	return l.BasicLogger
}

func (l *LoggerWrapper)SetLevel(lvl Level){
	if ll, ok := l.BasicLogger.(LevelsLogger); ok {
		ll.SetLevel(lvl)
	}else{
		l.level = lvl
	}
}

func (l *LoggerWrapper)Level()(Level){
	if ll, ok := l.BasicLogger.(LevelsLogger); ok {
		return ll.Level()
	}
	return l.level
}

func (l *LoggerWrapper)Trace(v ...any){
	if il, ok := l.BasicLogger.(TraceLogger); ok {
		il.Trace(v...)
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Trace'")
		}
		if l.level < TraceLevel {
			return
		}
		v = addOrInsertPrefix("TRACE:", v)
		l.Print(v...)
	}
}

func (l *LoggerWrapper)Tracef(format string, v ...any){
	if il, ok := l.BasicLogger.(TracefLogger); ok {
		il.Tracef(format, v...)
	}else if il, ok := l.BasicLogger.(TraceLogger); ok {
		il.Trace(fmt.Sprintf(format, v...))
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Tracef'")
		}
		if l.level < TraceLevel {
			return
		}
		l.Printf("TRACE: " + format, v...)
	}
}

func (l *LoggerWrapper)Debug(v ...any){
	if il, ok := l.BasicLogger.(DebugLogger); ok {
		il.Debug(v...)
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Debug'")
		}
		if l.level < DebugLevel {
			return
		}
		v = addOrInsertPrefix("DEBUG:", v)
		l.Print(v...)
	}
}

func (l *LoggerWrapper)Debugf(format string, v ...any){
	if il, ok := l.BasicLogger.(DebugfLogger); ok {
		il.Debugf(format, v...)
	}else if il, ok := l.BasicLogger.(DebugLogger); ok {
		il.Debug(fmt.Sprintf(format, v...))
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Debugf'")
		}
		if l.level < DebugLevel {
			return
		}
		l.Printf("DEBUG: " + format, v...)
	}
}

func (l *LoggerWrapper)Info(v ...any){
	if il, ok := l.BasicLogger.(InfoLogger); ok {
		il.Info(v...)
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Info'")
		}
		if l.level < InfoLevel {
			return
		}
		v = addOrInsertPrefix("INFO:", v)
		l.Print(v...)
	}
}

func (l *LoggerWrapper)Infof(format string, v ...any){
	if il, ok := l.BasicLogger.(InfofLogger); ok {
		il.Infof(format, v...)
	}else if il, ok := l.BasicLogger.(InfoLogger); ok {
		il.Info(fmt.Sprintf(format, v...))
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Infof'")
		}
		if l.level < InfoLevel {
			return
		}
		l.Printf("INFO: " + format, v...)
	}
}

func (l *LoggerWrapper)Warn(v ...any){
	if il, ok := l.BasicLogger.(WarnLogger); ok {
		il.Warn(v...)
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Warn'")
		}
		if l.level < WarnLevel {
			return
		}
		v = addOrInsertPrefix("WARN:", v)
		l.Print(v...)
	}
}

func (l *LoggerWrapper)Warnf(format string, v ...any){
	if il, ok := l.BasicLogger.(WarnfLogger); ok {
		il.Warnf(format, v...)
	}else if il, ok := l.BasicLogger.(WarnLogger); ok {
		il.Warn(fmt.Sprintf(format, v...))
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Warnf'")
		}
		if l.level < WarnLevel {
			return
		}
		l.Printf("WARN: " + format, v...)
	}
}

func (l *LoggerWrapper)Error(v ...any){
	if il, ok := l.BasicLogger.(ErrorLogger); ok {
		il.Error(v...)
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Error'")
		}
		if l.level < ErrorLevel {
			return
		}
		v = addOrInsertPrefix("ERROR:", v)
		l.Print(v...)
	}
}

func (l *LoggerWrapper)Errorf(format string, v ...any){
	if il, ok := l.BasicLogger.(ErrorfLogger); ok {
		il.Errorf(format, v...)
	}else if il, ok := l.BasicLogger.(ErrorLogger); ok {
		il.Error(fmt.Sprintf(format, v...))
	}else{
		if !l.AddPrefixOnMissing {
			panic("Logger: Cannot found method 'Errorf'")
		}
		if l.level < ErrorLevel {
			return
		}
		l.Printf("ERROR: " + format, v...)
	}
}

func (l *LoggerWrapper)Fatal(v ...any){
	if il, ok := l.BasicLogger.(FatalLogger); ok {
		il.Fatal(v...)
	}else{
		if l.level < FatalLevel {
			return
		}
		l.Print(v...)
		os.Exit(1)
	}
}

func (l *LoggerWrapper)Fatalf(format string, v ...any){
	if il, ok := l.BasicLogger.(FatalfLogger); ok {
		il.Fatalf(format, v...)
	}else if il, ok := l.BasicLogger.(FatalLogger); ok {
		il.Fatal(fmt.Sprintf(format, v...))
	}else{
		if l.level < FatalLevel {
			return
		}
		l.Printf(format, v...)
		os.Exit(1)
	}
}

func (l *LoggerWrapper)Panic(v ...any){
	if il, ok := l.BasicLogger.(PanicLogger); ok {
		il.Panic(v...)
	}else{
		l.Print(v...)
		panic(fmt.Sprint(v...))
	}
}

func (l *LoggerWrapper)Panicf(format string, v ...any){
	if il, ok := l.BasicLogger.(PanicfLogger); ok {
		il.Panicf(format, v...)
	}else if il, ok := l.BasicLogger.(PanicLogger); ok {
		il.Panic(fmt.Sprintf(format, v...))
	}else{
		panic(fmt.Sprintf(format, v...))
	}
}

func addOrInsertPrefix(prefix string, v []any)([]any){
	if len(v) == 0 {
		return []any{prefix}
	}
	if v0, ok := v[0].(string); ok {
		v[0] = prefix + " " + v0
	}else{
		v = append([]any{prefix}, v...)
	}
	return v
}
