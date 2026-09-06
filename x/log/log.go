package log

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/google/uuid"

	"github.com/grasp-labs/ds-go-echo-middleware/v3/middleware/claims"
	"github.com/grasp-labs/ds-go-echo-middleware/v3/middleware/requestctx"
)

// defaultContext is used when user context is missing.
// Prevents nil dereference and provides default traceability fields.
var defaultContext = &claims.Context{
	Rsc: "00000000-0000-0000-0000-000000000000:unknown",
	Sub: "unknown",
}

// Logger is the interfaces.Logger implementation handed to the shared
// middleware chain.
type Logger struct{}

func (Logger) Info(ctx context.Context, format string, args ...any) {
	logAt("INFO", ctx, false, format, args...)
}

func (Logger) Warning(ctx context.Context, format string, args ...any) {
	logAt("WARN", ctx, false, format, args...)
}

func (Logger) Error(ctx context.Context, format string, args ...any) {
	logAt("ERROR", ctx, false, format, args...)
}

func (Logger) StackError(ctx context.Context, format string, args ...any) {
	logAt("ERROR", ctx, true, format, args...)
}

func Info(c context.Context, format string, args ...any) {
	logAt("INFO", c, false, format, args...)
}

func Warning(c context.Context, format string, args ...any) {
	logAt("WARN", c, false, format, args...)
}

func Error(c context.Context, format string, args ...any) {
	logAt("ERROR", c, false, format, args...)
}

// StackError logs at ERROR and appends the goroutine's stack. Use it where the
// caller location alone is not actionable — notably from shared helpers, whose
// own line is identical for every failure routed through them.
func StackError(c context.Context, format string, args ...any) {
	logAt("ERROR", c, true, format, args...)
}

// logAt is the single emit path for every exported entry point above. The
// caller is resolved here so the frame arithmetic lives in one place: each
// exported function sits exactly one frame above this one. Every entry point
// must therefore call logAt directly and never route through another exported
// logging function, or the reported location shifts onto this file.
func logAt(level string, ctx context.Context, withStack bool, format string, args ...any) {
	// Frames: 0 = logAt, 1 = the exported wrapper, 2 = the code being logged.
	caller := "unknown"
	if _, file, line, ok := runtime.Caller(2); ok {
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	msg := fmt.Sprintf(format, args...)
	prefix := buildLogPrefix(level, caller, ctx)

	if withStack {
		log.Printf("%s %s\n%s", prefix, msg, debug.Stack())
		return
	}
	log.Printf("%s %s", prefix, msg)
}

// buildLogPrefix tags a line with the request ID, tenant ID, subject and caller
// so logs are searchable on a user and tenant level.
func buildLogPrefix(level, caller string, ctx context.Context) string {
	// Parse (or generate) request ID set by the RequestID middleware.
	requestID, err := uuid.Parse(requestctx.GetRequestID(ctx))
	if err != nil {
		requestID = uuid.New()
	}

	userCtx := requestctx.GetUserContext(ctx)
	if userCtx == nil {
		userCtx = defaultContext
	}

	tenantID, err := userCtx.GetTenantId()
	if err != nil {
		tenantID = uuid.Nil
	}

	return fmt.Sprintf("[%s][%s][%s][%s][%s]", level, requestID, tenantID, userCtx.Sub, caller)
}
