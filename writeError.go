package log

import (
	"fmt"
	"time"
)

//WriteError writes an error to the log with a timestamp appended
func WriteError(err error) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "\u001b[31m[ERROR]\u001b[0m", err.Error())
}
