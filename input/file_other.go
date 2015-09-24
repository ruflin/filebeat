// +build !windows

package input

import (
	"github.com/elastic/libbeat/logp"
	"os"
	"syscall"
)

type FileStateOS struct {
	Inode  uint64 `json:"inode,omitempty"`
	Device uint64 `json:"device,omitempty"`
}

// GetOSFileState returns the FileStateOS for non windows systems
func GetOSFileState(info *os.FileInfo) *FileStateOS {

	stat := (*(info)).Sys().(*syscall.Stat_t)

	// Convert inode and dev to uint64 to be cross platform compatible
	fileState := &FileStateOS{
		Inode:  uint64(stat.Ino),
		Device: uint64(stat.Dev),
	}

	return fileState
}

// IsSame file checks if the files are identical
func (fs *FileStateOS) IsSame(state *FileStateOS) bool {
	return fs.Inode == state.Inode && fs.Device == state.Device
}

// SafeFileRotate safely rotates an existing file under path and replaces it with the tempfile
func SafeFileRotate(path, tempfile string) error {
	if e := os.Rename(tempfile, path); e != nil {
		logp.Err("Rotate error: %s", e)
		return e
	}
	return nil
}

// Open to open files
func Open(path string, write bool) (*os.File, error) {

	flag := os.O_RDONLY
	var perm os.FileMode

	perm = 0

	if (write) {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		perm = 0666
	}
	return os.OpenFile(path, flag, perm)
}
