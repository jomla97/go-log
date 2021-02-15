# Getting started

## Installation
```
go mod init github.com/my/repo
go get github.com/jomla97/go-log
```

## Importing
``` go
import "github.com/jomla97/go-log"
```

## Print info
```go
//Regular
log.Info("write", anything, here)

//Formatter
log.Infof("some %v interesting %v", "very", format)
```

## Print error
```go
//Regular
log.Error(err)

//Formatter
log.Errorf("something went wrong: %v", err)
```

## Gin request middleware
```go
router := gin.New()
router.Use(log.GinRequestMiddleware)
```

## Set path to file to write to
```go
err := log.SetPath("/path/to/my.log")

path := log.GetPath()
```

## Change the date format used in the log file
Default value: `2006-01-02 15:04:05`.
```go
log.DateFormat = "some format"
```
