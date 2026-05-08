package log_test

import (
	"errors"
	"testing"

	"github.com/ayuxsec/log"
)

func TestLogger(t *testing.T) {
	log.WithTimeStamp = false
	log.WithTrace = true
	log.Info("this is an info message")
	log.Errorf("error: %v", errors.New("test error"))
}
