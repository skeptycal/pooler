package pooler

import (
	"log"
	"os"
)

const envDebugName = "GO_DEBUG"

var DEBUG bool = false

func init() {
	if os.Getenv(envDebugName) == "true" {
		DEBUG = true
	}
}

// dbLog prints unformatted log messages if the variable
// DEBUG is true. The default is false. DEBUG can be set
// by creating an environment variable named "GO_DEBUG"
// and setting the value to "true".
func dbLog(args ...interface{}) {
	if DEBUG {
		log.Print(args...)
	}
}

// dbLogf prints formatted log messages if the variable
// DEBUG is true. The default is false. DEBUG can be set
// by creating an environment variable named "GO_DEBUG"
// and setting the value to "true".
func dbLogf(format string, args ...interface{}) {
	if DEBUG {
		log.Printf(format, args...)
	}
}
