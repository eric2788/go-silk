//go:build windows
// +build windows

package multiplat

import (
	"os/exec"
	"syscall"
)

func HideWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
