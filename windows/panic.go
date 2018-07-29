package windows

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const (
	kernel32dll = "kernel32.dll"
)

var globalFile *os.File

// Catch catch panic
// path is output address
func Catch(path string) error {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	globalFile = file
	if err != nil {
		return err
	}

	t := fmt.Sprintf("time: %s", time.Now().Format("2006-01-02 15:04:05"))
	file.Write([]byte(t))
	file.Write([]byte("\r\n"))

	kernel32 := syscall.NewLazyDLL(kernel32dll)
	setStdHandle := kernel32.NewProc("SetStdHandle")
	sh := syscall.STD_ERROR_HANDLE
	v, _, err := setStdHandle.Call(uintptr(sh), uintptr(file.Fd()))
	if v == 0 {
		return err
	}
	return nil
}
