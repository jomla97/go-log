package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var file *os.File = nil
var path string

//SetPath sets the path to the file to write to
func SetPath(p string) error {
	if p[len(p)-1] == '/' {
		//Path does not include a filename and extension
		return errors.New("the specified path does not include a filename and a filetype extension")
	} else if !strings.Contains(filepath.Base(p), ".") {
		//Path does not include an extension
		return errors.New("the specified path does not include a filetype extension and/or a filename")
	}

	//Open the file
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err.Error())
	}

	file = f
	path = p

	return nil
}

//GetPath gets the path to the file to write to
func GetPath() string {
	return path
}
