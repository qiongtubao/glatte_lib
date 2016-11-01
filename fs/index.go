package fs

import (
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

/**
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
*/
func WriteFile(filePath string, data string) error {
	return ioutil.WriteFile(filePath,  []byte(data), os.ModePerm)
}