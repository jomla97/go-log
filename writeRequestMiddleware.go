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

//WriteRequestMiddleware writes information about a HTTP request to the log with a timestamp appended
func WriteRequestMiddleware(c *gin.Context) {
	//Save start time
	startTime := time.Now()

	//Clone the request body
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	bodyBytes, err := ioutil.ReadAll(tee)
	if err != nil {
		WriteError(fmt.Errorf("failed to read body: %v", err.Error()))
	} else {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	go func(c *gin.Context, bodyBytes []byte) {
		select {
		case <-c.Request.Context().Done():
			//Get time elapsed
			elapsed := time.Since(startTime)
			var elapsedString string
			if elapsed.Minutes() >= 1 {
				//Format as minutes and seconds
				elapsedString = fmt.Sprintf("%vm%vs", math.Round(elapsed.Minutes()), math.Round(math.Mod(elapsed.Seconds(), 60)))
			} else if elapsed.Seconds() >= 1 {
				//Format as seconds
				elapsedString = fmt.Sprintf("%gs", float64(elapsed.Milliseconds())/float64(1000))
			} else if elapsed.Milliseconds() >= 1 {
				//Format as milliseconds
				elapsedString = fmt.Sprintf("%gms", float64(elapsed.Microseconds())/float64(1000))
			} else if elapsed.Microseconds() >= 1 {
				//Format as microseconds
				elapsedString = fmt.Sprintf("%gµs", float64(elapsed.Nanoseconds())/float64(1000))
			} else {
				//Format as nanoseconds
				elapsedString = fmt.Sprintf("%vns", elapsed.Nanoseconds())
			}

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

			//Write in console
			fmt.Println("["+time.Now().Format(time.UnixDate)+"] [REQUEST]", statusString, "|", fmt.Sprintf("%-10s", elapsedString), "|", c.Request.Method, c.Request.URL.String(), string(bodyBytes))
			return
		}
	}(c, bodyBytes)
}
