package snippets

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	log ctxKey = iota
	pref
)

func checkContext() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, log, "10")
	fmt.Printf("log :: %+v\n", ctx.Value(log))
	fmt.Printf("context :: %+v\n", ctx)

	ctx = context.WithValue(ctx, log, "18")
	fmt.Printf("log :: %+v\n", ctx.Value(log))
	fmt.Printf("context :: %+v\n", ctx)

	ctx = context.WithValue(ctx, log, "1")
	ctx = context.WithValue(ctx, log, "2")
	ctx = context.WithValue(ctx, pref, "3")
	fmt.Printf("log :: %+v\n", ctx.Value(log))
	fmt.Printf("pref :: %+v\n", ctx.Value(pref))
	fmt.Printf("context :: %+v\n", ctx)
}
