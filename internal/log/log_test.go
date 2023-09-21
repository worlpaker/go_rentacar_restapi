package Log

import (
	"errors"
	"io"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	log.SetOutput(io.Discard)
	err := errors.New("new error")
	if !Err(err) {
		t.Errorf("Log failed to detect error")
	}
	if Err(nil) {
		t.Errorf("Log returned true for no error")
	}
}
