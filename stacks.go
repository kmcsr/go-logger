
package gologger

import (
	bytes   "bytes"
	debug   "runtime/debug"
)

func getStack(_n ...int)(stack []byte){
	var n int = 0
	if len(_n) > 0 {
		n = _n[0]
	}
	buf := bytes.NewBuffer(debug.Stack())
	stacks := make([][]byte, 0, buf.Len() / 32)
	for buf.Len() > 0 {
		line, _ := buf.ReadBytes('\n')
		if len(line) > 1 {
			stacks = append(stacks, line)
		}
	}
	if n < -2 || n * 2 + 5 >= len(stacks) {
		return bytes.Join(stacks[0:1], []byte{})
	}
	return bytes.Join(append(stacks[0:1], stacks[n * 2 + 5:]...), []byte{})
}
