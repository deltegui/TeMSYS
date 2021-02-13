package api

import (
	"context"
	"net/http"
	"strings"
	"temsys"
	"time"

	"github.com/deltegui/phoenix"
)

type ctxKey string

const UserCtxKey ctxKey = "user"

type Cors struct {
	allowOrigin string
}

func NewCors(allowOrigin string) Cors {
	return Cors{allowOrigin}
}

// EnableCors middleware to enable cors in API endpoints
func (cors Cors) EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cors.allowOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		next.ServeHTTP(w, req)
	})
}

// OptionsCors creates a middleware that handles all request using
// Options method and returns 204. This is used to return all
// CORS headers.
func OptionsCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, req)
	})
}

// JWTAuth middleware to check using a JWT Bearer token if user is authorized or is admin.
type JWTAuth struct {
	tokenizer temsys.Tokenizer
}

// NewJWTAuth create JWTAuth middleware.
func NewJWTAuth(tokenizer temsys.Tokenizer) JWTAuth {
	return JWTAuth{tokenizer}
}

// Authorize middleware that checks if exists the header 'Authorization' with
// valid JWT bearer token.
func (authMiddle JWTAuth) Authorize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authMiddle.handleAndCheckToken(w, req, next, func(role temsys.Role) error { return nil })
	})
}

// Admin middleware that checks if exists the header 'Authorization' with
// valid JWT bearer token with admin role.
func (authMiddle JWTAuth) Admin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authMiddle.handleAndCheckToken(w, req, next, func(role temsys.Role) error {
			if role != temsys.AdminRole {
				return temsys.OnlyAdminErr
			}
			return nil
		})
	})
}

func (authMiddle JWTAuth) handleAndCheckToken(w http.ResponseWriter, req *http.Request, next http.Handler, verifyRole func(temsys.Role) error) {
	presenter := phoenix.NewJSONPresenter(w)
	token, err := authMiddle.getToken(req)
	if err != nil {
		presenter.PresentError(err)
		return
	}
	if token.Expires.Before(time.Now()) {
		presenter.PresentError(temsys.ExpiredTokenErr)
		return
	}
	if err := verifyRole(token.Role); err != nil {
		presenter.PresentError(err)
		return
	}
	ctx := context.WithValue(req.Context(), UserCtxKey, token.Owner)
	next.ServeHTTP(w, req.WithContext(ctx))
}

func (authMiddle JWTAuth) getToken(req *http.Request) (temsys.Token, error) {
	const bearerPrefix string = "Bearer "
	bearerToken := req.Header.Get("Authorization")
	if len(bearerToken) == 0 || !strings.HasPrefix(bearerToken, bearerPrefix) {
		return temsys.Token{}, temsys.NotAuthErr
	}
	rawToken := strings.Replace(bearerToken, bearerPrefix, "", 1)
	token, err := authMiddle.tokenizer.Decode(rawToken)
	if err != nil {
		return temsys.Token{}, temsys.InvalidTokenErr
	}
	return token, nil
}
