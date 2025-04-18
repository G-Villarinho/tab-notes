package pkgs

import "context"

type contextKey string

const (
	UserIDKey    contextKey = "user_id"
	TokenKey     contextKey = "user_token"
	SessionIDKey contextKey = "session_id"
)

type RequestContext interface {
	SetUserID(ctx context.Context, userID string) context.Context
	SetToken(ctx context.Context, token string) context.Context
	SetSessionID(ctx context.Context, sessionID string) context.Context
	GetUserID(ctx context.Context) (string, bool)
	GetToken(ctx context.Context) (string, bool)
	GetSessionID(ctx context.Context) (string, bool)
}

type requestContext struct {
	userIDKey    contextKey
	tokenKey     contextKey
	sessionIDKey contextKey
}

func NewRequestContext() RequestContext {
	return &requestContext{
		userIDKey:    UserIDKey,
		tokenKey:     TokenKey,
		sessionIDKey: SessionIDKey,
	}
}

func (r *requestContext) SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, r.userIDKey, userID)
}

func (r *requestContext) SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, r.sessionIDKey, sessionID)
}

func (r *requestContext) SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, r.tokenKey, token)
}

func (r *requestContext) GetUserID(ctx context.Context) (string, bool) {
	if userID, ok := ctx.Value(r.userIDKey).(string); ok {
		return userID, true
	}

	return "", false
}

func (r *requestContext) GetToken(ctx context.Context) (string, bool) {
	if token, ok := ctx.Value(r.tokenKey).(string); ok {
		return token, true
	}

	return "", false
}

func (r *requestContext) GetSessionID(ctx context.Context) (string, bool) {
	if sessionID, ok := ctx.Value(r.sessionIDKey).(string); ok {
		return sessionID, true
	}

	return "", false
}
