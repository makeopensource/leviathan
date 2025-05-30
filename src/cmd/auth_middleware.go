package cmd

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
)

const AuthHeader = "Authorization"

type authInterceptor struct {
	authKey string
}

func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		token := conn.RequestHeader().Get(AuthHeader)
		if token == "" {
			return connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("no auth header found"))
		}

		err := i.verifyAuthHeader(token)
		if err != nil {
			return err
		}

		return next(ctx, conn)
	}
}

func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		clientToken := req.Header().Get(AuthHeader)

		err := i.verifyAuthHeader(clientToken)
		if err != nil {
			return nil, err
		}

		return next(ctx, req)
	}
}

func (*authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	}
}

func (i *authInterceptor) verifyAuthHeader(clientToken string) error {
	if i.authKey != clientToken {
		return connect.NewError(
			connect.CodeUnauthenticated,
			fmt.Errorf("invalid auth key"),
		)
	}
	return nil
}
