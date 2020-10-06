package silk

import (
	"os"
	"runtime"
)

// FileExist 检查文件是否存在
func fileExist(path string) bool {
	if runtime.GOOS == "windows" {
		path = path + ".exe"
	}
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
