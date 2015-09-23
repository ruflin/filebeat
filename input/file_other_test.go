// +build !windows

package input

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOSFileState(t *testing.T) {
	file, err := ioutil.TempFile("", "")
	assert.Nil(t, err)

	fileinfo, err := file.Stat()
	assert.Nil(t, err)

	state := GetOSFileState(&fileinfo)

	assert.True(t, state.Inode > 0)
	assert.True(t, state.Device > 0)
}
