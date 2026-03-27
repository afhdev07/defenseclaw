package gateway

import (
	"fmt"
	"os"
	"sync"
)

// The WebSocket read loop must stay responsive: if stderr is a slow writer
// (e.g. daemon mode appending both streams to one log file), synchronous
// fmt.Fprintf(os.Stderr) can block and stall RPC delivery (including connect).
const readLoopStderrQueue = 2048

var readLoopStderrOnce sync.Once
var readLoopStderrCh chan string

func startReadLoopStderrDrainer() {
	readLoopStderrOnce.Do(func() {
		readLoopStderrCh = make(chan string, readLoopStderrQueue)
		go func() {
			for line := range readLoopStderrCh {
				_, _ = fmt.Fprint(os.Stderr, line)
			}
		}()
	})
}

// readLoopLogf queues one log line for stderr (adds newline if missing).
// Used only from readLoop so inbound frames keep being read even if stderr blocks.
func readLoopLogf(format string, args ...interface{}) {
	startReadLoopStderrDrainer()
	s := fmt.Sprintf(format, args...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += "\n"
	}
	readLoopStderrCh <- s
}
