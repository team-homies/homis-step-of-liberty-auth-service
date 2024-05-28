package resource

type CreateTokenRequest struct {
	Id            string `json:"id"`
	Provider      string `json:"provider"`
	FirebaseToken string `json:"firebase_token"`
}

type UpdateTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// type UserInfoRequest struct {
// 	UserId uint64 `json:"id"`
// }

type FindVisualRequest struct {
	UserId uint `json:"user_id"`
}
