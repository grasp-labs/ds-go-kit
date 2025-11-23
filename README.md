# DS Go Kit

![Build](https://github.com/grasp-labs/ds-go-kit/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/grasp-labs/ds-go-kit)](https://goreportcard.com/report/github.com/grasp-labs/ds-go-kit)
[![codecov](https://codecov.io/gh/grasp-labs/ds-go-kit/branch/main/graph/badge.svg)](https://codecov.io/gh/grasp-labs/ds-go-kit)
[![GitHub release](https://img.shields.io/github/v/release/grasp-labs/ds-go-kit)](https://github.com/grasp-labs/ds-go-kit/releases)
![License](https://img.shields.io/github/license/grasp-labs/ds-go-kit?cacheSeconds=60)


A shared, reusable Go library across all your Go services.

## 📦 Dependencies

- DS Go Echo Middleware [Github](https://github.com/grasp-labs/ds-go-echo-middleware/tree/main)

## ✨ x/log/

## ✔ Beautiful imports

```bash
import "github.com/grasp-labs/ds-go-kit/x/log"
```

## ✔ Beautiful usage

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

## ✔ Future-proof (can swap log.Printf → zap later)

No call sites need to change.
