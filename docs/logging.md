# ✨ x/log

Log is a minimal context aware log wrapper to be used in any service, that produce tagged and standardized output.

## Usage

### Import

```bash
import "github.com/grasp-labs/ds-go-kit/x/log"
```

### usage

```go
log.Info(ctx, "upload ok")
log.Error(ctx, "failed: %v", err)
```

For prefixed logging with full error trace, call

```bash
log.StackError(ctx, "Message")
```

Example response (Mac OS)

```bash
2025/11/25 09:29:57 [ERROR][c166e250-dea0-431c-9d23-4156bf318afc][00000000-0000-0000-0000-000000000000][super@user.com][log_test.go:72] error New err with id: 5b92ec60-0a6e-4323-9bd1-ec5464a825d3 with id 5b92ec60-0a6e-4323-9bd1-ec5464a825d3
goroutine 11 [running]:
runtime/debug.Stack()
        /opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/debug/stack.go:26 +0x64
ds-go-kit/x/log.StackError({0x10074e158, 0x1400010cb70}, {0x1006c3d48?, 0x1008fed30?}, {0x14000105e58?, 0x30?, 0x14000105e28?})
        .../ds-go-kit/x/log/log.go:52 +0x50
ds-go-kit/x/log_test.TestStackErrorLogger.func1()
        .../ds-go-kit/x/log/log_test.go:72 +0xb4
ds-go-kit/internal/fakes.CaptureLogs(0x10074ba80)
       .../ds-go-kit/internal/fakes/logs.go:14 +0xa0
ds-go-kit/x/log_test.TestStackErrorLogger(0x1400014c1c0)
        .../ds-go-kit/x/log/log_test.go:67 +0x28
testing.tRunner(0x1400014c1c0, 0x10074b9d8)
        /opt/homebrew/Cellar/go/1.25.1/libexec/src/testing/testing.go:1934 +0xc8
created by testing.(*T).Run in goroutine 1
        /opt/homebrew/Cellar/go/1.25.1/libexec/src/testing/testing.go:1997 +0x364
```

## ✔ Context-aware logging from day one

- request ID
- tenant ID
- subject/user
- file:line
- full stack trace on error

```go
log.Warning(ctx, err)
```

```text
2025/11/21 11:40:37 [WARN][c97319b9-b6dc-4c94-859f-7b4b0baa646a][aee4b2bf-f399-425e-88b0-eee33227dccc][user@email.com][handler.go:30] Invalid JSON in request body
```

## ✔ Future-proof (can swap log.Printf → zap later)

No call sites need to change.
