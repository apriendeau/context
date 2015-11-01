package context_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/apriendeau/context"
	"github.com/stretchr/testify/assert"
)

func TestGetCtx(t *testing.T) {
	req := &http.Request{}
	ctx := context.Get(req)
	assert.NotNil(t, ctx.Created)
	copyCtx := context.Get(req)
	assert.Equal(t, ctx.Created, copyCtx.Created, "context should point to the same ctx")
	context.Clear(req)
}

func TestCtxClear(t *testing.T) {
	req := &http.Request{}
	ctx := context.Get(req)
	d, _ := time.ParseDuration("2ns")
	context.Clear(req)
	time.Sleep(d)
	cCtx := context.Get(req)
	assert.NotEqual(t, ctx.Created, cCtx.Created, "clearing, should cause a new ctx to be created")
	context.Clear(req)
}
