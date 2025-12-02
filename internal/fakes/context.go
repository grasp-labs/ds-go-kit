package fakes

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/grasp-labs/ds-go-echo-middleware/v2/middleware/claims"
	"github.com/grasp-labs/ds-go-echo-middleware/v2/middleware/requestctx"
)

type Ctx struct {
	RequestID   uuid.UUID
	UserContext *claims.Context
}

func (c *Ctx) New() context.Context {
	ctx := context.Background()
	c.RequestID = uuid.New()
	c.UserContext = &claims.Context{
		Sub: "super@user.com",
		Rsc: fmt.Sprintf("%s:%s", "customer x", uuid.New().String()),
	}
	ctx = requestctx.SetRequestID(ctx, c.RequestID.String())
	ctx = requestctx.SetUserContext(ctx, c.UserContext)
	return ctx
}
