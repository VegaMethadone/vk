package authentication

import (
	"encoding/base64"
	"log"
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
		Expires:  time.Now().Add(30 * time.Minute),
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func CookieCheker(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Fatalf("Cookie is damaged %v", err)
		return nil, err
	}

	token, err := base64Decode(c.Value)
	if err != nil {
		log.Fatalf("Cookie is damaged %v", err)
		return nil, err
	}

	return token, nil
}
