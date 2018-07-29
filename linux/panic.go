package linux

import (
	"os"
)

var globalFile *os.File

// Catch catch panic
func Catch(path string) error {
	// Linux Uncomment
	// file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// globalFile = file
	// if err != nil {
	// 	return err
	// }

	// t := fmt.Sprintf(" time: %s", time.Now().Format("2006-01-02 15:04:05"))
	// file.Write([]byte(t))
	// file.Write([]byte("\n"))

	// if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
	// 	return err
	// }
	return nil
}
