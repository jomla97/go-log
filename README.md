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
