package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type AuthWriteAccessRequeset struct {
	Key string `json:"key"`
}

func (svc *NBService) AuthWriteAccess(w http.ResponseWriter, r *http.Request) {
	var req AuthWriteAccessRequeset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Key != svc.WriteAccessKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	payload := "write-access"
	h := hmac.New(sha256.New, []byte(svc.ServerSecret))
	h.Write([]byte(payload))
	signature := hex.EncodeToString(h.Sum(nil))

	cookieValue := fmt.Sprintf("%s.%s", payload, signature)
	host, _, _ := net.SplitHostPort(r.Host)

	if host == "localhost" {
		http.SetCookie(w, &http.Cookie{
			Name:        "auth_token",
			Value:       cookieValue,
			Path:        "/",
			HttpOnly:    true,
			Secure:      false,
			SameSite:    http.SameSiteNoneMode,
			Expires:     time.Now().Add(24 * time.Hour),
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    cookieValue,
			Path:     "/api",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(24 * time.Hour),
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access granted"))
}

// ValidateWriteAccess checks the cookie and returns nil if valid
func (svc *NBService) ValidateWriteAccess(r *http.Request) error {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		log.Print("failed auth request, missing cookie")
		return errors.New("missing auth cookie")
	}

	// 1. Split payload and signature
	parts := strings.Split(cookie.Value, ".")
	if len(parts) != 2 {
		log.Print("failed auth request, invalid cookie")
		return errors.New("invalid cookie format")
	}

	payload := parts[0]
	providedSignature, err := hex.DecodeString(parts[1])
	if err != nil {
		log.Print("failed auth request, invalid signature")
		return errors.New("invalid signature encoding")
	}

	// 2. Re-calculate the HMAC
	h := hmac.New(sha256.New, []byte(svc.ServerSecret))
	h.Write([]byte(payload))
	expectedSignature := h.Sum(nil)

	// 3. Constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(providedSignature, expectedSignature) != 1 {
		log.Print("failed auth request, invalid signature")
		return errors.New("invalid signature")
	}

	// 4. Optional: Verify the payload content
	if payload != "write-access" {
		log.Print("failed auth request, invalid permissions")
		return errors.New("invalid permissions")
	}

	return nil
}
