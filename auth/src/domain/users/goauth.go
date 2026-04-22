package users

import (
	"amarolio-auth/src/customerrors"
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauthUtil struct {
	Config *oauth2.Config
	UrlAPI string
}

func (gou *GoogleOauthUtil) GetAuthCodeURL(state string) string {
	return gou.Config.AuthCodeURL(state)
}

func (gou *GoogleOauthUtil) Exchange(ctx context.Context, code string) (string, error) {
	token, err := gou.Config.Exchange(ctx, code)
	if err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return token.AccessToken, nil
}

func (gou *GoogleOauthUtil) GetUrlAPI() string {
	return gou.UrlAPI
}

func CreateGoogleOauthUtil() *GoogleOauthUtil {
	redirectUrl := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
	return &GoogleOauthUtil{
		Config: &oauth2.Config{
			RedirectURL:  redirectUrl,
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
		UrlAPI: "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
	}
}
