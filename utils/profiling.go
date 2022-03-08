package utils

import (
	"main/logging"
	"time"
)

func RunWithProfiler(tag string, p func() error) error {
	startTime := time.Now()
	err := p()
	if err != nil {
		return err
	}
	endTime := time.Now()
	logging.DebugFormat("Run function %s for %d ms", tag, endTime.Sub(startTime).Milliseconds())
	return nil
}
