package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

	// 1. Validate the provided key
	if req.Key != svc.WriteAccessKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	payload := "write-access"
	h := hmac.New(sha256.New, []byte(svc.ServerSecret))
	h.Write([]byte(payload))
	signature := hex.EncodeToString(h.Sum(nil))

	// 3. Create the combined value (payload.signature)
	cookieValue := fmt.Sprintf("%s.%s", payload, signature)

	// 4. Set the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    cookieValue,
		Path:     "/api",
		HttpOnly: true,
		Secure:   true, // Set to true for Production (HTTPS)
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access granted"))
}

// ValidateWriteAccess checks the cookie and returns nil if valid
func (svc *NBService) ValidateWriteAccess(r *http.Request) error {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return errors.New("missing auth cookie")
	}

	// 1. Split payload and signature
	parts := strings.Split(cookie.Value, ".")
	if len(parts) != 2 {
		return errors.New("invalid cookie format")
	}

	payload := parts[0]
	providedSignature, err := hex.DecodeString(parts[1])
	if err != nil {
		return errors.New("invalid signature encoding")
	}

	// 2. Re-calculate the HMAC
	h := hmac.New(sha256.New, []byte(svc.ServerSecret))
	h.Write([]byte(payload))
	expectedSignature := h.Sum(nil)

	// 3. Constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(providedSignature, expectedSignature) != 1 {
		return errors.New("invalid signature")
	}

	// 4. Optional: Verify the payload content
	if payload != "write-access" {
		return errors.New("invalid permissions")
	}

	return nil
}