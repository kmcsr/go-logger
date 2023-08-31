
package logger

import (
	"fmt"
)

type PrefixLogger struct {
	Logger
	prefix string
}

var _ Logger = (*PrefixLogger)(nil)

func NewPrefixLogger(origin Logger, prefix string, args ...any)(l *PrefixLogger){
	return &PrefixLogger{
		prefix: fmt.Sprintf(prefix, args...),
		Logger: origin,
	}
}

func (l *PrefixLogger)Prefix()(string){
	return l.prefix
}

func (l *PrefixLogger)SetPrefix(prefix string, args ...any){
	l.prefix = fmt.Sprintf(prefix, args...)
}

func (l *PrefixLogger)Unwrap()(BasicLogger){
	return l.Logger
}

func (l *PrefixLogger)Print(v ...any){
	l.Logger.Print(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Printf(format string, v ...any){
	l.Logger.Printf(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Trace(v ...any){
	l.Logger.Trace(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Tracef(format string, v ...any){
	l.Logger.Tracef(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Debug(v ...any){
	l.Logger.Debug(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Debugf(format string, v ...any){
	l.Logger.Debugf(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Info(v ...any){
	l.Logger.Info(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Infof(format string, v ...any){
	l.Logger.Infof(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Warn(v ...any){
	l.Logger.Warn(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Warnf(format string, v ...any){
	l.Logger.Warnf(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Error(v ...any){
	l.Logger.Error(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Errorf(format string, v ...any){
	l.Logger.Errorf(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Fatal(v ...any){
	l.Logger.Fatal(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Fatalf(format string, v ...any){
	l.Logger.Fatalf(l.prefix + " " + format, v...)
}

func (l *PrefixLogger)Panic(v ...any){
	l.Logger.Panic(addOrInsertPrefix(l.prefix, v)...)
}

func (l *PrefixLogger)Panicf(format string, v ...any){
	l.Logger.Panicf(l.prefix + " " + format, v...)
}
