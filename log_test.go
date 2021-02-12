package log

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	err := SetPath("testlog.log")
	if err == nil {
		Info("Some info!")
		Infof("Some %v info!", "other")
		Errorf("Some %v error!", "red")
	} else {
		Error(errors.New("Failed to log to file, using default stderr"))
	}
}
