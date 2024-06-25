package chain

func Chain[T any](ctx T, middlewares []func(ctx T, next func()), final func(ctx T)) {
	if len(middlewares) == 0 {
		final(ctx)
		return
	}

	first := middlewares[0]
	rest := middlewares[1:]
	next := func() {
		Chain(ctx, rest, final)
	}

	first(ctx, next)
}
