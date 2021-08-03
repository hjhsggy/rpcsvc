package interceptor

import (
	"context"
	"fmt"
)

type handler func(ctx context.Context)

type interceptor func(ctx context.Context, h handler, ivk invoker) error

type invoker func(ctx context.Context, ceps []interceptor, h handler) error

func getChainInterceptor(ctx context.Context, ceps []interceptor, ivk invoker) interceptor {

	if len(ceps) == 0 {
		return nil
	}
	if len(ceps) == 1 {
		return ceps[0]
	}
	return func(ctx context.Context, h handler, ivk invoker) error {
		return ceps[0](ctx, h, getInvoker(ctx, ceps, 0, ivk))
	}

}

func getInvoker(ctx context.Context, ceps []interceptor, cur int, ivk invoker) invoker {

	if cur == len(ceps)-1 {
		return ivk
	}

	return func(ctx context.Context, ceps []interceptor, h handler) error {
		return ceps[cur+1](ctx, h, getInvoker(ctx, ceps, cur+1, ivk))
	}
}

func Interceptor() {

	var ctx context.Context
	var ceps []interceptor
	var h = func(ctx context.Context) {
		fmt.Println("do something")
	}

	var inter1 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	var inter2 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}

	var inter3 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}

	ceps = append(ceps, inter1, inter2, inter3)
	var ivk = func(ctx context.Context, interceptors []interceptor, h handler) error {
		fmt.Println("invoker start")
		return nil
	}

	cep := getChainInterceptor(ctx, ceps, ivk)
	cep(ctx, h, ivk)

}
