// +build !windows

package logp

import (
	"os"
)

// Open to open files
func Open(path string, write bool) (*os.File, error) {

	flag := os.O_RDONLY
	var perm os.FileMode

	perm = 0

	if write {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		perm = 0666
	}
	return os.OpenFile(path, flag, perm)
}
