
package gologger

import (
	io "io"
	os "os"
	fmt "fmt"
)

type Logger interface{
	Outputs()([]io.Writer)
	SetOutputs(...io.Writer)
	AddOutputs(...io.Writer)
	Errouts()([]io.Writer)
	SetErrouts(...io.Writer)
	AddErrouts(...io.Writer)

	Formatter()(Formatter)
	SetFormatter(Formatter)

	Name()(string)

	Level()(Level)
	SetLevel(Level)
	ErrLevel()(Level)
	SetErrLevel(Level)

	Log(args ...interface{})(error)
	Logf(format string, args ...interface{})(error)
	Debug(args ...interface{})(error)
	Debugf(format string, args ...interface{})(error)
	Info(args ...interface{})(error)
	Infof(format string, args ...interface{})(error)
	Warn(args ...interface{})(error)
	Warnf(format string, args ...interface{})(error)
	Error(args ...interface{})(error)
	Errorf(format string, args ...interface{})(error)
	Stack(args ...interface{})(error)
	Stackf(format string, args ...interface{})(error)
}

type PreCall func(Logger)(string)

type logger struct{
	prev Logger
	outs, errs []io.Writer
	fmtr Formatter
	name string
	lvl, errlvl Level
}

var _ Logger = (*logger)(nil)

func NewLogger(name string)(Logger){
	return &logger{
		prev: nil,
		outs: []io.Writer{os.Stdout},
		errs: []io.Writer{os.Stderr},
		fmtr: &defaultFormatter{tf: "2006-01-02|15:04:05.000"},
		name: name,
		lvl: LEVEL_INFO,
		errlvl: LEVEL_WARN,
	}
}

func NewLoggerPrev(prev Logger, name string)(Logger){
	return &logger{
		prev: prev,
		name: name,
		lvl: prev.Level(),
		errlvl: prev.ErrLevel(),
	}
}

func (l *logger)Outputs()([]io.Writer){
	return l.outs
}

func (l *logger)SetOutputs(ws ...io.Writer){
	l.outs = ws
}

func (l *logger)AddOutputs(ws ...io.Writer){
	l.outs = append(l.outs, ws...)
}

func (l *logger)Errouts()([]io.Writer){
	return l.errs
}

func (l *logger)SetErrouts(ws ...io.Writer){
	l.errs = ws
}

func (l *logger)AddErrouts(ws ...io.Writer){
	l.errs = append(l.errs, ws...)
}

func (l *logger)Formatter()(Formatter){
	return l.fmtr
}

func (l *logger)SetFormatter(fmtr Formatter){
	l.fmtr = fmtr
}

func (l *logger)Name()(string){
	return l.name
}

func (l *logger)Level()(Level){
	return l.lvl
}

func (l *logger)SetLevel(lvl Level){
	l.lvl = lvl
}

func (l *logger)ErrLevel()(Level){
	return l.errlvl
}

func (l *logger)SetErrLevel(errlvl Level){
	l.errlvl = errlvl
}

func (l *logger)format(lvl Level, src string)(string){
	if l.prev != nil {
		return l.prev.Formatter().Format(l, lvl, src)
	}
	return l.fmtr.Format(l, lvl, src)
}

func (l *logger)writeLine(lvl Level, line string)(err error){
	outs := l.outs
	if l.prev != nil {
		if l.ErrLevel() > lvl {
			outs = l.prev.Outputs()
		}else{
			outs = l.prev.Errouts()
		}
	}else if l.ErrLevel() <= lvl {
		outs = l.errs
	}
	if len(line) > 0 && line[len(line) - 1] != '\n' {
		line += "\n"
	}
	for _, o := range outs {
		_, err = io.WriteString(o, line)
		if err != nil { return }
	}
	return nil
}

func (l *logger)Log(args ...interface{})(error){
	if l.Level() > LEVEL_DISABLE {
		return nil
	}
	return l.writeLine(LEVEL_DISABLE, fmt.Sprintln(args...))
}

func (l *logger)Logf(format string, args ...interface{})(error){
	if l.Level() > LEVEL_DISABLE {
		return nil
	}
	return l.writeLine(LEVEL_DISABLE, fmt.Sprintf(format, args...))
}

func (l *logger)Debug(args ...interface{})(error){
	if l.Level() > LEVEL_DEBUG {
		return nil
	}
	return l.writeLine(LEVEL_DEBUG, l.format(LEVEL_DEBUG, fmt.Sprintln(args...)))
}

func (l *logger)Debugf(format string, args ...interface{})(error){
	if l.Level() > LEVEL_DEBUG {
		return nil
	}
	return l.writeLine(LEVEL_DEBUG, l.format(LEVEL_DEBUG, fmt.Sprintf(format, args...)))
}

func (l *logger)Info(args ...interface{})(error){
	if l.Level() > LEVEL_INFO {
		return nil
	}
	return l.writeLine(LEVEL_INFO, l.format(LEVEL_INFO, fmt.Sprintln(args...)))
}

func (l *logger)Infof(format string, args ...interface{})(error){
	if l.Level() > LEVEL_INFO {
		return nil
	}
	return l.writeLine(LEVEL_INFO, l.format(LEVEL_INFO, fmt.Sprintf(format, args...)))
}

func (l *logger)Warn(args ...interface{})(error){
	if l.Level() > LEVEL_WARN {
		return nil
	}
	return l.writeLine(LEVEL_WARN, l.format(LEVEL_WARN, fmt.Sprintln(args...)))
}

func (l *logger)Warnf(format string, args ...interface{})(error){
	if l.Level() > LEVEL_WARN {
		return nil
	}
	return l.writeLine(LEVEL_WARN, l.format(LEVEL_WARN, fmt.Sprintf(format, args...)))
}

func (l *logger)Error(args ...interface{})(error){
	if l.Level() > LEVEL_ERROR {
		return nil
	}
	return l.writeLine(LEVEL_ERROR, l.format(LEVEL_ERROR, fmt.Sprintln(args...)))
}

func (l *logger)Errorf(format string, args ...interface{})(error){
	if l.Level() > LEVEL_ERROR {
		return nil
	}
	return l.writeLine(LEVEL_ERROR, l.format(LEVEL_ERROR, fmt.Sprintf(format, args...)))
}

func (l *logger)Stack(args ...interface{})(error){
	if l.Level() > LEVEL_STACK {
		return nil
	}
	return l.writeLine(LEVEL_STACK, l.format(LEVEL_STACK, fmt.Sprintln(args...)))
}

func (l *logger)Stackf(format string, args ...interface{})(error){
	if l.Level() > LEVEL_STACK {
		return nil
	}
	return l.writeLine(LEVEL_STACK, l.format(LEVEL_STACK, fmt.Sprintf(format, args...)))
}


