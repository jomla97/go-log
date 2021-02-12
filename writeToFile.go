package log

import (
	"errors"
	"fmt"
)

//WriteToFile writes the specified string to the set file
func WriteToFile(s string) error {
	if file == nil {
		return errors.New("output file not set")
	}

	//Write the string to the file
	_, err := file.WriteString(s)
	if err != nil {
		return fmt.Errorf("failed to write to output file: %v", err.Error())
	}

	return nil
}
