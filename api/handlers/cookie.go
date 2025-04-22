package handlers

import (
	"net/http"
	"time"
)

func SetTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "tabnotes_id",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((time.Hour * 24 * 7).Seconds()),
	})
}

func GetTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("tabnotes_id")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func DeleteTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "tabnotes_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
}
