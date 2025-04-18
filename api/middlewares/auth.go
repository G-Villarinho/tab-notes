package middlewares

import (
	"net/http"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/handlers"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	Authenticated(next http.HandlerFunc) http.HandlerFunc
}

type authMiddleware struct {
	kp pkgs.EcdsaKeyPair
	rc pkgs.RequestContext
	ss services.SessionService
}

func NewAuthMiddleware(
	ecdsaKeyPair pkgs.EcdsaKeyPair,
	requestContext pkgs.RequestContext,
	sessionService services.SessionService) AuthMiddleware {
	return &authMiddleware{
		kp: ecdsaKeyPair,
		rc: requestContext,
		ss: sessionService,
	}
}

func (a *authMiddleware) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := handlers.GetTokenCookie(r)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		if tokenStr == "" {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		publicKey, err := a.kp.ParseECDSAPublicKey(configs.Env.Key.PublicKey)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusInternalServerError)
			handlers.DeleteTokenCookie(w)
			return
		}

		var claims models.AuthTokenClaims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (any, error) {
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			handlers.DeleteTokenCookie(w)
			return
		}

		revoked, err := a.ss.IsSessionRevoked(r.Context(), claims.SessionID)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusInternalServerError)
			return
		}

		if revoked {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			handlers.DeleteTokenCookie(w)
			return
		}

		ctx := a.rc.SetToken(r.Context(), tokenStr)
		ctx = a.rc.SetSessionID(ctx, claims.SessionID)
		ctx = a.rc.SetUserID(ctx, claims.Subject)

		r = r.WithContext(ctx)

		next(w, r)
	}
}
