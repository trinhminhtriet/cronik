package jobs

import (
	"os/exec"
)

// CmdJob spawns a process.
// Example usage:
// cmd := exec.Command("echo", "Hello world")
// cronik.Schedule(
//
//	"echo",
//	&jobs.CmdJob{Cmd: cmd},
//	&t.When{Every: t.Every(1).Seconds()})
type CmdJob struct {
	Cmd *exec.Cmd
}

// Runs the command.
func (j *CmdJob) Run() error {
	return j.Cmd.Run()
}
