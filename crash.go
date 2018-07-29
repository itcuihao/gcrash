package gcrash

import (
	"errors"
	"runtime"

	"github.com/itcuihao/gcrash/linux"
	"github.com/itcuihao/gcrash/windows"
)

var (
	// ErrNotGoos Unsupported Goos
	ErrNotGoos = errors.New("unsupported goos")
)

// Catch init
func Catch(path string) error {
	return Goos(path)
}

// Goos distinguish system
func Goos(path string) (err error) {
	switch runtime.GOOS {
	default:
		err = ErrNotGoos
	case "darwin":
		fallthrough
	case "linux":
		err = linux.Catch(path)
	case "windows":
		err = windows.Catch(path)
	}
	return
}
