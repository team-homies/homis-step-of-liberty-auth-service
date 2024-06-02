package service

import (
	"main/app/api/auth/resource"
	"main/config"
	"main/constant/common"
	"main/database/entity"
	"main/database/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type AuthService interface {
	CreateToken(req *resource.CreateTokenRequest) (res *resource.CreateTokenResponse, err error)
	UpdateRefreshToken(req *resource.UpdateTokenRequest) (res *resource.UpdateTokenResponse, err error)
	UserInfo(userId uint) (res *resource.UserInfoResponse, err error)
	UpdateUserInfo(req *resource.UpdateUserInfoRequest) (err error)
	FindVisual(userId uint) (res *resource.FindVisualResponse, err error)
	FindVisualCode(userId uint) (res *resource.FindVisualCodeResponse, err error)
}

func NewAuthService() AuthService {
	return &authService{
		AuthService: &authService{},
	}
}

type authService struct {
	AuthService
}

func (as *authService) CreateToken(req *resource.CreateTokenRequest) (res *resource.CreateTokenResponse, err error) {
	var userId uint64
	// 1. FindUserByUserInfo 사용해서 db에  id, provider 조회
	user, err := repository.NewRepository().FindUserByUserInfo(req.Id, req.Provider)
	if err != nil {
		return
	}

	// 1-1. 없으면 CreateUser 사용해서 저장 후 로직 진행(2번)
	if !user.IsUsed {
		err = repository.NewRepository().CreateUser(req.Id, req.Provider)
		if err != nil {
			return
		}
		// TODO: 리팩토링 예정(중복 로직)
		user, err = repository.NewRepository().FindUserByUserInfo(req.Id, req.Provider)
		if err != nil {
			return
		}
	}

	if user != nil {
		userId = uint64(user.ID)
	}
	// 1-2. 있으면 res 반환 후 로직 진행(2번)

	// 2. AccessToken, RefreshToken 만들기
	exp := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := CreateAccessToken(userId, exp)
	if err != nil {
		return
	}

	refreshToken, err := CreateRefreshToken(userId)
	if err != nil {
		return
	}

	// 3. RefreshToken db에 저장
	err = repository.NewRepository().UpdateRefreshToken(userId, refreshToken)
	if err != nil {
		return
	}

	res = &resource.CreateTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expired:      exp,
	}
	return
}

func CreateAccessToken(userId uint64, exp int64) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(viper.GetString(config.JWT_SECRET)))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateRefreshToken(userId uint64) (string, error) {
	var err error
	//Creating Refresh Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString(config.JWT_SECRET)))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *authService) UpdateRefreshToken(req *resource.UpdateTokenRequest) (res *resource.UpdateTokenResponse, err error) {
	var userId uint64
	// 1.FindRefreshToken 으로 db에 refresh토큰 조회
	token, err := repository.NewRepository().FindRefreshToken(req.RefreshToken)
	if err != nil {
		return
	}
	// 1-1. refresh 값이 없으면 에러
	// if token == nil {
	//	return
	// }

	// 1-2. refresh 값이 있으면 access token 발급
	if token == nil {
		return
	}
	userId = uint64(token.ID)

	exp := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := CreateAccessToken(userId, exp)
	if err != nil {
		return
	}

	res = &resource.UpdateTokenResponse{
		AccessToken: accessToken,
		Expired:     exp,
	}

	return res, err
}

func (as *authService) UserInfo(userId uint) (res *resource.UserInfoResponse, err error) {
	user, err := repository.NewRepository().FindUserInfo(userId)
	if err != nil {
		return
	}
	res = new(resource.UserInfoResponse)

	// if user != nil {
	// 	res = &resource.UserInfoResponse{
	// 		ID:       user.ID,
	// 		Email:    user.Email,
	// 		Nickname: user.Nickname,
	// 		Profile:  user.Profile,
	// 	}
	// } else {
	// 	return
	// }

	if user == nil {
		return
	}

	res = &resource.UserInfoResponse{
		Id:       user.ID,
		Email:    user.Email,
		Nickname: user.Nickname,
		Profile:  user.Profile,
	}

	return res, err
}

// 사용자 본인 정보 수정 body : Nickname, Propile
func (as *authService) UpdateUserInfo(req *resource.UpdateUserInfoRequest) (err error) {
	authRepository := repository.NewRepository()

	// 1. 검증한 userid
	var userInfo entity.User
	userInfo.Model.ID = req.Id
	if req.Nickname != "" {
		userInfo.Nickname = req.Nickname
	}

	if req.Profile != "" {
		userInfo.Profile = req.Profile
	}

	// 2. 만들어놓은 레포지토리를 사용해서 데이터를 수정
	err = authRepository.UpdateUserInfo(&userInfo)
	if err != nil {
		return
	}

	// 3. 리턴
	return

}

// 시각적 성취도 조회
func (as *authService) FindVisual(userId uint) (res *resource.FindVisualResponse, err error) {
	userRepository := repository.NewRepository()
	res = new(resource.FindVisualResponse)
	// 1. userId를 이용해서 user의 수집률을 구하고 변수에 담는다
	var collectRate uint

	// 2. 수집률로 조건식을 사용하여 코드분류
	code := common.PercentCal(collectRate)

	// 3.  수집률을 담아 만들어놓은 레포지토리를 사용해서 데이터를 가져온다
	visualFind, err := userRepository.FindVisual(code)
	if err != nil {
		return nil, err
	}

	// 4. 가져온 데이터를 하나의 객체(res)에 합친다
	res = &resource.FindVisualResponse{
		Name:     visualFind.Name,
		Code:     visualFind.Code,
		Percent:  int(collectRate),
		ImageUrl: visualFind.ImageUrl,
	}

	// 5. 리턴
	return

}

// 시각적 성취도 코드 조회
func (as *authService) FindVisualCode(userId uint) (res *resource.FindVisualCodeResponse, err error) {
	userRepository := repository.NewRepository()
	res = new(resource.FindVisualCodeResponse)
	// 1. userId를 이용해서 user의 수집률을 구하고 변수에 담는다
	var collectRate uint

	// 2. 수집률로 조건식을 사용하여 코드분류
	code := common.PercentCal(collectRate)

	// 3.  수집률을 담아 만들어놓은 레포지토리를 사용해서 데이터를 가져온다
	visualFind, err := userRepository.FindVisualCode(code)
	if err != nil {
		return nil, err
	}

	// 4. 가져온 데이터를 하나의 객체(res)에 합친다
	res = &resource.FindVisualCodeResponse{
		Name:         visualFind.Name,
		Code:         visualFind.Code,
		DisplayLevel: visualFind.DisplayLevel,
		Description:  visualFind.Description,
	}

	// 5. 리턴
	return

}

// 수집률 grpc
