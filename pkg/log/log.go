package log

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Logger instance
var Logger log.Logger

func init() {
	l := log.NewLogfmtLogger(os.Stdout)
	if os.Getenv("DEBUG") == "true" {
		l = level.NewFilter(l, level.AllowAll())
	} else {
		l = level.NewFilter(l, level.AllowInfo())
	}
	Logger = log.With(l, "timestamp", log.DefaultTimestamp, "caller", log.Caller(5))
}
