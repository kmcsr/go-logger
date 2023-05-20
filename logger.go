
package logger

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func fileExist(filename string)(bool){
	if _, err := os.Stat(filename); err != nil {
		return errors.Is(err, os.ErrExist)
	}
	return true
}

func copyLogToGzip(filename string, r io.Reader)(err error){
	t := filename + ".2.gz"
	{
		i := 2
		n := t
		for fileExist(n) {
			i++
			n = fmt.Sprintf(filename + ".%d.gz", i)
		}
		for i > 2 {
			i--
			m := fmt.Sprintf(filename + ".%d.gz", i)
			if err = os.Rename(m, n); err != nil {
				return
			}
			n = m
		}
	}
	fd, err := os.Create(t)
	if err != nil {
		return
	}
	defer fd.Close()
	w := gzip.NewWriter(fd)
	defer w.Close()
	_, err = io.Copy(w, r)
	return
}

func checkAndMoveLogs(filename string)(err error){
	if !fileExist(filename) {
		return
	}
	t := filename + ".1"
	fd, err := os.Open(t)
	if err == nil {
		if err = copyLogToGzip(filename, fd); err != nil {
			return
		}
	}else if !errors.Is(err, os.ErrNotExist) {
		return
	}
	return os.Rename(filename, t)
}

func OutputToFile(l BasicLogger, filename string, outs ...io.Writer)(err error){
	var out io.Writer
	if err = checkAndMoveLogs(filename); err != nil {
		return
	}
	dir := filepath.Dir(filename)
	if !fileExist(dir) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return
		}
	}
	if out, err = os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_EXCL, 0666); err != nil {
		return
	}
	if len(outs) > 0 {
		out = io.MultiWriter(append(outs, out)...)
	}
	l.SetOutput(out)
	return
}

