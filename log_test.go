package log_test

import (
	"errors"
	"testing"

	"github.com/ayuxsec-org/log"
)

func TestLogger(t *testing.T) {
	log.WithTimeStamp = false
	log.WithTrace = true
	log.Info("this is an info message")
	log.Warn("warning")
	log.Warnf("Warning the kid just farted: %d times", 2)
	log.Errorf("error: %v", errors.New("test error"))
	log.Fatalf("error: %v", errors.New("fatal error"))
}
