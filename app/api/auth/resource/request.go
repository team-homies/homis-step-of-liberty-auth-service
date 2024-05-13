package resource

type CreateTokenRequest struct {
	Id            string `json:"id"`
	Provider      string `json:"provider"`
	FirebaseToken string `json:"firebase_token"`
}

type UpdateTokenRequest struct {
	Id           string `json:"id"`
	RefreshToken string `json:"refresh_token"`
}
