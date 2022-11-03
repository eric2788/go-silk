//go:build !windows
// +build !windows

package multiplat

import "os/exec"

func HideWindow(cmd *exec.Cmd) {
}
