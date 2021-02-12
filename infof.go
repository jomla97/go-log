package log

import (
	"fmt"
	"time"
)

//Infof writes to the log with a timestamp and [INFO] tag appended
func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf("%v %v %v\n", "["+time.Now().Format(DateFormat)+"]", "[INFO]", fmt.Sprintf(format, args...))

	fmt.Print(msg)

	if file != nil {
		//Write to the file
		err := WriteToFile(msg)
		if err != nil {
			Errorf("failed to write to output file: %v", err.Error())
		}
	}
}
