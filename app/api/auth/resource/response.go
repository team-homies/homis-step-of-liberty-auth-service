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

type FindVisualResponse struct {
	Id       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Percent  int    `json:"percent"`
	ImageUrl string `json:"image_url"`
}

type FindVisualCodeResponse struct {
	Id           uint   `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	DisplayLevel int    `json:"display_level"`
	Description  string `json:"description"`
}
