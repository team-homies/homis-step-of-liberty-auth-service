package resource

type CreateTokenRequest struct {
	Id            uint64 `json:"id"`
	Provider      string `json:"provider"`
	FirebaseToken string `json:"firebase_token"`
}
