package platform

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Account struct {
	InternalId interface{}
	Token      string `json:"auth_token"`
	PlatformId string `json:"id"`
	FirstName  string `json:"localizedFirstName"`
	LastName   string `json:"localizedLastName"`
}

const (
	LOGIN_URL_PATTERN = "https://www.linkedin.com/oauth/v2/authorization?response_type=code&state=%d&scope=w_member_social&client_id=%s&redirect_uri=%s"
	OAUTH_URL         = "https://www.linkedin.com/oauth/v2/accessToken"
	PROFILE_SELF_URL  = "https://www.linkedin.com/oauth/v2/me"
)

// Returns an OAuth access token, it's expiration time and possibly an error
func OAuth(loginToken string, redirectURL string, clientId string, clientSecret string) (string, int, error) {
	oAuthPayload := map[string]string{
		"grant_type":    "authorization_code",
		"code":          loginToken,
		"redirect_uri":  url.QueryEscape(redirectURL),
		"client_id":     clientId,
		"client_secret": clientSecret,
	}

	serializedOAuthPayload, err := json.Marshal(oAuthPayload)
	if err != nil {
		return "", 0, err
	}

	res, err := http.Post(OAUTH_URL, "application/json", bytes.NewBuffer(serializedOAuthPayload))
	if err != nil {
		return "", 0, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	var oAuthResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &oAuthResponse); err != nil {
		return "", 0, err
	}

	return oAuthResponse.AccessToken, oAuthResponse.ExpiresIn, nil
}

func GetProfile(token string) (*Account, error) {
	profileRequest, err := http.NewRequest(http.MethodGet, PROFILE_SELF_URL, http.NoBody)
	if err != nil {
		return nil, err
	}

	profileRequest.Header.Set("Authorization", "Bearer: "+token)

	res, err := (&http.Client{}).Do(profileRequest)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var account Account
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	account.Token = token

	return &account, nil
}

// In the future, contact my college's administration so they allow access to the LinkedIn API via their company page.
// For now, I'll be using a mock system
