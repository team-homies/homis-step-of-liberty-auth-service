package resource

type CreateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Expired      int64  `json:"expired"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expired     int64  `json:"expired"`
}

type UserInfoResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

// 시각적 성취도 조회
type FindVisualResponse struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Percent  int    `json:"percent"`
	ImageUrl string `json:"image_url"`
}

// 시각적 성취도 코드 조회
type FindVisualCodeResponse struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	DisplayLevel int    `json:"display_level"`
	Description  string `json:"description"`
}
