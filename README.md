# log
Dead stupid leveled logger for Golang

Meant to shadow the actual log package to not lose habit of typing "log." :).

Usage

```golang
package main

import (
	"github.com/dpatrie/log"
)

func main() {
	log.Debug("Some debug")        //[DEBUG] - Some debug
	log.DebugF("Some %s", "debug") //[DEBUG] - Some debug
	log.Info("Some info")          //[INFO] - Some info
	log.InfoF("Some %s", "info")   //[INFO] - Some info

	log.Error("Some error")        //[ERROR] - Some error
	log.ErrorF("Some %s", "error") //[ERROR] - Some error

	log.DebugLogger = false
	log.Debug("Should not be seen") //Nothing...
	log.InfoLogger = false
	log.Info("Should not be seen") //Nothing...

	log.Fatal("Should Exit...") //[FATAL] - Should Exit...
}
```
