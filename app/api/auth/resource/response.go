package resource

type CreateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Expired      int    `json:"expired"`
	RefreshToken string `json:"refresh_token"`
}
