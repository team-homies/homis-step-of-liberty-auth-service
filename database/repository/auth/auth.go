package auth

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUserByUserInfo(email, provider string) (user *entity.User, err error)
	CreateUser(email, provider string) error
	UpdateRefreshToken(userId uint64, refreshToken string) error
	FindRefreshToken(refreshToken string) (result *entity.User, err error)
}

type gormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &gormAuthRepository{db}
}

func (g *gormAuthRepository) FindUserByUserInfo(email, provider string) (user *entity.User, err error) {
	// 	SELECT *
	//  FROM user
	// WHERE email = '해리이메일'
	//     AND provider = '구글'
	//     AND is_used = true;
	tx := g.db.Where("email = ? AND provider = ? AND is_used = true", email, provider).Find(&user)

	return user, tx.Error
}

func (g *gormAuthRepository) CreateUser(email, provider string) error {
	// TODO: 트랜잭션 고려
	// 	INSERT
	//   INTO "user"(nickname, profile, provider, refresh_token, is_used, email)
	// VALUES('', '', 'google', '', true, 'suhy427@gmail.com');
	tx := g.db.Save(&entity.User{
		Nickname:     "",
		Profile:      "",
		Provider:     provider,
		RefreshToken: "",
		IsUsed:       true,
		Email:        email,
	})

	return tx.Error
}

func (g *gormAuthRepository) UpdateRefreshToken(userId uint64, refreshToken string) error {
	// TODO: 트랜잭션 고려
	// update "user"
	// set refresh_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMj'
	// where id = 2;
	tx := g.db.Model(&entity.User{}).Where("id = ?", userId).Update("refresh_token", refreshToken)

	return tx.Error
}

func (g *gormAuthRepository) FindRefreshToken(refreshToken string) (result *entity.User, err error) {
	// SELECT refresh_token
	//   FROM "user" u
	//  WHERE refresh_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMj'

	tx := g.db.Select("refresh_token").Where("refresh_token = ?", refreshToken).Find(&result)

	return result, tx.Error
}