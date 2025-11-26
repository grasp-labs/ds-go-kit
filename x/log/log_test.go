package log_test

import (
	"context"
	"strings"
	"testing"

	"github.com/grasp-labs/ds-go-kit/internal/fakes"
	"github.com/grasp-labs/ds-go-kit/x/log"

	"github.com/google/uuid"
)

func TestInfoLogger(t *testing.T) {
	var c fakes.Ctx
	out := fakes.CaptureLogs(func() {
		ctx := c.New()
		log.Info(ctx, "hello %s", "world")
	})

	if !strings.Contains(out, "[INFO]") {
		t.Errorf("expected INFO prefix, got: %s", out)
	}

	if !strings.Contains(out, "hello world") {
		t.Errorf("missing message")
	}

	if !strings.Contains(out, c.RequestID.String()) {
		t.Errorf("missing requestID, got: %v", out)
	}
}

func TestWarningLogger(t *testing.T) {
	out := fakes.CaptureLogs(func() {
		c := fakes.Ctx{}
		ctx := c.New()
		log.Warning(ctx, "hello %s", "world")
	})

	if !strings.Contains(out, "[WARN]") {
		t.Errorf("expected WARN prefix, got: %s", out)
	}

	if !strings.Contains(out, "hello world") {
		t.Errorf("missing message")
	}
}

func TestErrorLogger(t *testing.T) {
	out := fakes.CaptureLogs(func() {
		c := fakes.Ctx{}
		ctx := c.New()
		log.Error(ctx, "hello %s", "world")
	})

	if !strings.Contains(out, "[ERROR]") {
		t.Errorf("expected ERROR prefix, got: %s", out)
	}

	if !strings.Contains(out, "hello world") {
		t.Errorf("missing message")
	}
}

func TestStackErrorLogger(t *testing.T) {
	out := fakes.CaptureLogs(func() {
		c := fakes.Ctx{}
		ctx := c.New()
		id := uuid.New()
		err := fakes.Caller(id)
		log.StackError(ctx, "error %v with id %s", err, id)
	})

	if !strings.Contains(out, "[ERROR]") {
		t.Errorf("expected ERROR prefix, got: %s", out)
	}

	if !strings.Contains(out, "runtime/debug.Stack()") {
		t.Errorf("missing message, got START\n%s\nEND", out)
	}
}

func TestLoggerHandlesNilContext(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logger should never panic, got: %v", r)
		}
	}()
	_ = fakes.CaptureLogs(func() {
		log.Info(context.Background(), "ok")
	})
}
