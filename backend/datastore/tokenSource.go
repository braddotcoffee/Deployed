package datastore

import "golang.org/x/oauth2"

// FirebaseTokenSource provides new OAuth token for firebase
type FirebaseTokenSource struct {
	token string
}

// Token returns new OAuth2 token
func (ts FirebaseTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: ts.token,
	}, nil
}
