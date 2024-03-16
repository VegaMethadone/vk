package authentication

import (
	"encoding/base64"
	"net/http"
	"time"
)

func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func CookieSeter(w http.ResponseWriter, r *http.Request, token string) {
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(10 * time.Minute),
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}
