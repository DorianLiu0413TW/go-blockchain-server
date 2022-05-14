package log

import (
	"time"
)

func SetTimezone(loc *time.Location) {
	timezone = loc
}

func Info(args ...interface{}) {

}
	

