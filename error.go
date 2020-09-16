package log

import (
	"fmt"
	"time"
)

//Error writes an error to the log with a timestamp and [ERROR] tag appended
func Error(err error) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "\u001b[31m[ERROR]\u001b[0m", err.Error())
}
