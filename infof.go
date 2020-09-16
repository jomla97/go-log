package log

import (
	"fmt"
	"time"
)

//Infof writes to the log with a timestamp and [INFO] tag appended
func Infof(format string, args ...interface{}) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "[INFO]", fmt.Sprintf(format, args...))
}
