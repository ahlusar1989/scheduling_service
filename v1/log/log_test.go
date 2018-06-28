package log_test

import (
	"testing"

	"github.com/ahlusar1989/scheduling_service/v1/log"
)

func TestDefaultLogger(t *testing.T) {
	log.INFO.Print("should not panic")
	log.WARNING.Print("should not panic")
	log.ERROR.Print("should not panic")
	log.FATAL.Print("should not panic")
}
