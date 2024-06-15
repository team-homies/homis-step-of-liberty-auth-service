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

// 사용자 본인 정보 수정
type UpdateUserInfoRequest struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}
