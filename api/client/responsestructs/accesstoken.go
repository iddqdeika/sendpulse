package responsestructs

type AccessToken struct {
	AccessToken	string	`json:"access_token"`
	TokenType	string	`json:"token_type"`
	ExpiresIn	int32	`json:"expires_in"`
}