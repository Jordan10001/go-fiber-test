package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Config adalah konfigurasi OAuth untuk Google
var GoogleConfig *oauth2.Config

// InitGoogleOAuth menginisialisasi konfigurasi Google OAuth
func InitGoogleOAuth(clientID, clientSecret string) {
	GoogleConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GetGoogleOAuthURL mengembalikan URL otorisasi
func GetGoogleOAuthURL(state string) string {
	return GoogleConfig.AuthCodeURL(state)
}

// GetGoogleTokens menukar kode otorisasi dengan token
func GetGoogleTokens(ctx context.Context, code string) (*oauth2.Token, error) {
	return GoogleConfig.Exchange(ctx, code)
}