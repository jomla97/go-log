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

	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "[INFO]", fmt.Sprintf(format, args...))
}
