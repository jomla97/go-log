package log

import (
	"fmt"
	"time"
)

//Info writes to the log with a timestamp and [INFO] tag appended
func Info(args ...interface{}) {
	var format string
	for i := 0; i < len(args); i++ {
		if format != "" {
			format += " %v"
		} else {
			format += "%v"
		}
	}

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
