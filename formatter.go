
package gologger

import (
	time "time"
	strings "strings"
)

type Formatter interface{
	TimeFormat()(string)
	SetTimeFormat(string)
	Format(lgr Logger, l Level, src string)(string)
}

type defaultFormatter struct{
	tf string
}

func (f *defaultFormatter)timeNow()(string){
	return time.Now().Format(f.tf)
}

func (f *defaultFormatter)TimeFormat()(string){
	return f.tf
}

func (f *defaultFormatter)SetTimeFormat(tf string){
	f.tf = tf
}

func (f *defaultFormatter)Format(lgr Logger, l Level, src string)(str string){
	sb := &strings.Builder{}
	sb.Grow(25 + len(f.tf) + len(src))
	sb.WriteByte('[')
	sb.WriteString(lgr.Name())
	sb.WriteString("][")
	sb.WriteString(l.String())
	sb.WriteString("][")
	sb.WriteString(f.timeNow())
	sb.WriteString("]: ")
	sb.WriteString(src)
	if len(src) > 0 && src[len(src) - 1] != '\n' {
		sb.WriteByte('\n')
	}
	if l == LEVEL_STACK {
		sb.Write(getStack(1))
	}
	return sb.String()
}


