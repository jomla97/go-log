package log

import (
	"fmt"
	"time"
)

//WriteErrorf writes an error to the log with a timestamp appended with the specified formatting options
func WriteErrorf(formatString string, args ...interface{}) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "\u001b[31m[ERROR]\u001b[0m", fmt.Sprintf(formatString, args))
}
