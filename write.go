package log

import (
	"fmt"
	"time"
)

//Write writes a string to the log with a timestamp appended
func Write(message string) {
	fmt.Println("["+time.Now().Format(time.UnixDate)+"]", "[INFO]", message)
}
