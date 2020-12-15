# goLogger
Very simple logger written in golang


## How to use
```golang
import(
  log "github.com/kylhuk/goLogger"
  "fmt"
)


// Just call AddLog with the desired LogLevel and the log message
log.AddLog(log.LogLevelWarning, "This is a Warning!")

// If you want the full struct
fmt.Println(fmt.Sprintf("%+v ", log.LogFull))

// If you want it as simple newLine delimitted string (Warning: This always gives you the whole log)
fmt.Println(log.LogFullAsString)
```
