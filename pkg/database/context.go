package database

import (
	"context"
	"fmt"
	"regexp"
)

type contextKey string

const (
	tenantKey contextKey = "tenant"
	userKey   contextKey = "user_id"
	roleKey   contextKey = "role"
)

var schemaRegex = regexp.MustCompile(`^[a-z0-9_]+$`)

func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(ctx, tenantKey, tenant)
}

func GetTenant(ctx context.Context) (string, error) {
	tenant, ok := ctx.Value(tenantKey).(string)
	if !ok || tenant == "" {
		return "", fmt.Errorf("tenant not found in context")
	}
	if !schemaRegex.MatchString(tenant) {
		return "", fmt.Errorf("invalid tenant format")
	}
	return tenant, nil
}

func WithUser(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userKey, userID)
}

func GetUserID(ctx context.Context) (int64, error) {
	id, ok := ctx.Value(userKey).(int64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in context")
	}
	return id, nil
}

func WithRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetRole(ctx context.Context) (string, error) {
	role, ok := ctx.Value(roleKey).(string)
	if !ok || role == "" {
		return "", fmt.Errorf("role not found in context")
	}
	return role, nil
}
