package appctx

import "context"

type ctxKey string

const (
	RequestIDKey ctxKey = "request_id"
	UserIDKey    ctxKey = "user_id"
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

func RequestID(ctx context.Context) string {
	if v := ctx.Value(RequestIDKey); v != nil {
		if rid, ok := v.(string); ok {
			return rid
		}
	}
	return ""
}
