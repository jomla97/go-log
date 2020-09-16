package log

import (
	"fmt"
	"time"
)

//Errorf writes an error to the log with a timestamp and [ERROR] tag appended with the specified formatting options
func Errorf(format string, args ...interface{}) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "\u001b[31m[ERROR]\u001b[0m", fmt.Sprintf(format, args...))
}
