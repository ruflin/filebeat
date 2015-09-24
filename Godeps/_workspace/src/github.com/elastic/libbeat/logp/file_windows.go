package logp

import (
	"fmt"
	"os"
	"syscall"
)

// Open to open files
// As Windows blocks deleting a file when its open, some special params are pass here.
func Open(path string, write bool) (*os.File, error) {

	// Set all write flags
	// This indirectly calls syscall_windows::Open method https://github.com/golang/go/blob/7ebcf5eac7047b1eef2443eda1786672b5c70f51/src/syscall/syscall_windows.go#L251
	// As FILE_SHARE_DELETE cannot be passed to Open, os.CreateFile must be implemented directly

	// This is mostly the code from syscall_windows::Open. Only difference is passing the Delete flag
	// TODO: Open pull request to Golang so also Delete flag can be set
	if len(path) == 0 {
		return nil, fmt.Errorf("File '%s' not found. Error: %v", syscall.ERROR_FILE_NOT_FOUND)
	}

	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, fmt.Errorf("Error converting to UTF16: %v", err)
	}

	var access uint32
	access = syscall.GENERIC_READ | syscall.GENERIC_WRITE

	sharemode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE | syscall.FILE_SHARE_DELETE)

	var sa *syscall.SecurityAttributes

	var createmode uint32

	if write {
		createmode = syscall.CREATE_NEW
	} else {
		createmode = syscall.OPEN_EXISTING
	}

	handle, err := syscall.CreateFile(pathp, access, sharemode, sa, createmode, syscall.FILE_ATTRIBUTE_NORMAL, 0)

	if err != nil {
		return nil, fmt.Errorf("Error creating file '%s': %v", pathp, err)
	}

	return os.NewFile(uintptr(handle), path), nil
}
