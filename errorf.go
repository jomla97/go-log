package log

import (
	"fmt"
	"time"
)

//Errorf writes an error to the log with a timestamp and [ERROR] tag appended with the specified formatting options
func Errorf(format string, args ...interface{}) {
	msg := "[" + time.Now().Format(DateFormat) + "] %v[ERROR]%v " + fmt.Sprintf(format, args...) + "\n"

	fmt.Printf(msg, "\u001b[31m", "\u001b[0m")

	if file != nil {
		//Write to the file
		WriteToFile(fmt.Sprintf(msg, "", ""))
	}
}
