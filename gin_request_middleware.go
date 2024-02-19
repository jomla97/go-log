package log

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

// GinResponse if true logs entire response body
var GinResponse = false

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GinRequestMiddleware writes information about a HTTP request to the log with a timestamp appended
func GinRequestMiddleware(c *gin.Context) {
	colorReset := "\u001b[0m"
	red := "\u001b[31m"
	orange := "\u001b[38;5;209m"
	yellow := "\u001b[33m"

	//Save start time
	startTime := time.Now()

	//Clone the request body
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	bodyBytes, err := ioutil.ReadAll(tee)
	if err != nil {
		Errorf("failed to read body: %v", err)
	} else {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// create a body log writer, for potentially logging response body
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	go func(c *gin.Context, bodyBytes []byte) {
		select {
		case <-c.Request.Context().Done():
			//Get time elapsed
			elapsed := time.Since(startTime)
			var elapsedString string
			if elapsed.Minutes() >= 1 {
				//Format as minutes and seconds
				elapsedString = red + fmt.Sprintf("%vm%vs", math.Round(elapsed.Minutes()), math.Round(math.Mod(elapsed.Seconds(), 60)))
			} else if elapsed.Seconds() >= 1 {
				//Format as seconds
				elapsedString = red + fmt.Sprintf("%gs", float64(elapsed.Milliseconds())/float64(1000))
			} else if elapsed.Milliseconds() >= 1 {
				//Format as milliseconds
				ms := float64(elapsed.Microseconds()) / float64(1000)

				if elapsed.Milliseconds() >= 700 {
					elapsedString = fmt.Sprintf("%s%gms", red, ms)
				} else if elapsed.Milliseconds() >= 400 {
					elapsedString = fmt.Sprintf("%s%gms", orange, ms)
				} else if elapsed.Milliseconds() >= 100 {
					elapsedString = fmt.Sprintf("%s%gms", yellow, ms)
				} else {
					elapsedString = fmt.Sprintf("%gms", ms)
				}
			} else if elapsed.Microseconds() >= 1 {
				//Format as microseconds
				elapsedString = fmt.Sprintf("%gµs", float64(elapsed.Nanoseconds())/float64(1000))
			} else {
				//Format as nanoseconds
				elapsedString = fmt.Sprintf("%vns", elapsed.Nanoseconds())
			}

			elapsedString += colorReset

			statusString := fmt.Sprintf(" %v ", c.Writer.Status())
			if c.Writer.Status() >= 200 && c.Writer.Status() <= 299 {
				//Green background
				statusString = "\u001b[42;1m" + statusString + "\u001b[0m"
			} else if c.Writer.Status() >= 500 && c.Writer.Status() <= 599 {
				//Red background
				statusString = "\u001b[41;1m" + statusString + "\u001b[0m"
			} else {
				//Blue background
				statusString = "\u001b[44;1m" + statusString + "\u001b[0m"
			}

			var msg = ""
			if GinResponse {
				//Write in console with response body
				msg = fmt.Sprintf("%v %v %v %v %v %v %v %v %v %v\n", "["+time.Now().Format(DateFormat)+"] [REQUEST]", statusString, "|", fmt.Sprintf("%-10s", elapsedString), "|", c.Request.Method, c.Request.URL.String(), string(bodyBytes), "|", blw.body.String())
			} else {
				//Write in console
				msg = fmt.Sprintf("%v %v %v %v %v %v %v %v\n", "["+time.Now().Format(DateFormat)+"] [REQUEST]", statusString, "|", fmt.Sprintf("%-10s", elapsedString), "|", c.Request.Method, c.Request.URL.String(), string(bodyBytes))
			}

			if msg != "" {
				fmt.Print(msg)

				if file != nil {
					//Write to the OutputFile
					err = WriteToFile(msg)
					if err != nil {
						Errorf("failed to write to output file: %v", err.Error())
					}
				}
			}

			return
		}
	}(c, bodyBytes)
}
