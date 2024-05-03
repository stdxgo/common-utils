package logging

import (
	"context"
	"testing"
)

func TestInfo(t *testing.T) {

	Info(context.Background(), "abcd")
}
