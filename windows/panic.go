package windows

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const (
	kdll = "kernel32.dll"
	ssh  = "SetStdHandle"
)

var globalf *os.File

// Catch catch panic
// path is output address
func Catch(path string) error {
	// windows uncomment
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	globalf = file
	if err != nil {
		return err
	}

	t := fmt.Sprintf("panic time: %s", time.Now().Format("2006-01-02 15:04:05"))
	file.Write([]byte(t))
	file.Write([]byte("\r\n"))

	lp := syscall.NewLazyDLL(kdll).NewProc(ssh)

	seh := syscall.STD_ERROR_HANDLE
	v, _, err := lp.Call(uintptr(seh), uintptr(file.Fd()))
	if v == 0 {
		return err
	}
	return nil
}
