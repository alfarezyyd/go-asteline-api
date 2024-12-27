package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleIdentityProvider struct {
	GoogleLoginConfig oauth2.Config
}

var IdentityProviderHolder GoogleIdentityProvider

func GoogleConfig(viperConfig *viper.Viper) oauth2.Config {
	IdentityProviderHolder.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/redirect",
		ClientID:     viperConfig.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: viperConfig.GetString("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return IdentityProviderHolder.GoogleLoginConfig
}
