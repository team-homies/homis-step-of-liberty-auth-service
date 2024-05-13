package service

import (
	"main/app/api/auth/resource"
	"main/config"
	"main/database/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type AuthService interface {
	CreateToken(req *resource.CreateTokenRequest) (res *resource.CreateTokenResponse, err error)
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

func CreateAccessToken(userid uint64, exp int64) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(viper.GetString(config.JWT_SECRET)))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateRefreshToken(userid uint64) (string, error) {
	var err error
	//Creating Refresh Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString(config.JWT_SECRET)))
	if err != nil {
		return "", err
	}
	return token, nil
}