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
