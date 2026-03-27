package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func sendTermSignal(proc *os.Process) error {
	return proc.Signal(syscall.SIGTERM)
}

func sendKillSignal(proc *os.Process) error {
	return proc.Signal(syscall.SIGKILL)
}

func processExists(pid int) bool {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.Signal(0))
	return err == nil
}

// killStaleProcesses finds and kills any defenseclaw-gateway processes that
// are not tracked by the PID file. This prevents orphaned daemons from
// accumulating across restarts.
func (d *Daemon) killStaleProcesses() {
	self, _ := os.Executable()
	binName := filepath.Base(self)
	if binName == "" || binName == "." {
		binName = "defenseclaw-gateway"
	}

	out, err := exec.Command("pgrep", "-f", binName).Output()
	if err != nil {
		return
	}

	trackedPID := 0
	if info, err := d.readPIDInfo(); err == nil {
		trackedPID = info.PID
	}
	myPID := os.Getpid()

	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		pid, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil || pid <= 0 || pid == myPID || pid == trackedPID {
			continue
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			continue
		}
		fmt.Fprintf(os.Stderr, "[daemon] killing stale gateway process (PID %d)\n", pid)
		_ = proc.Signal(syscall.SIGTERM)
	}
}
