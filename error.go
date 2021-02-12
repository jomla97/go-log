package log

import (
	"fmt"
	"time"
)

//Error writes an error to the log with a timestamp and [ERROR] tag appended
func Error(err error) {
	msg := "[" + time.Now().Format(DateFormat) + "] %v[ERROR]%v " + err.Error() + "\n"

	fmt.Printf(msg, "\u001b[31m", "\u001b[0m")

	if file != nil {
		//Write to the file
		WriteToFile(fmt.Sprintf(msg, "", ""))
	}
}
