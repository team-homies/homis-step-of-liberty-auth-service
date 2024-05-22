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
	FindUserInfo(userId uint) (user *entity.User, err error)
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

func (g *gormAuthRepository) CreateUser(email, provider string) (err error) {
	// 	INSERT
	//   INTO "user"(nickname, profile, provider, refresh_token, is_used, email)
	// VALUES('', '', 'google', '', true, 'suhy427@gmail.com');
	tx := g.db.Begin()
	err = tx.Save(&entity.User{
		Nickname:     "",
		Profile:      "",
		Provider:     provider,
		RefreshToken: "",
		IsUsed:       true,
		Email:        email,
	}).Error

	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()

	return
}

func (g *gormAuthRepository) UpdateRefreshToken(userId uint64, refreshToken string) (err error) {
	// update "user"
	// set refresh_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMj'
	// where id = 2;
	tx := g.db.Begin()
	err = tx.Model(&entity.User{}).Where("id = ?", userId).Update("refresh_token", refreshToken).Error

	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()

	return
}

func (g *gormAuthRepository) FindRefreshToken(refreshToken string) (result *entity.User, err error) {
	// SELECT id, refresh_token
	//   FROM "user" u
	//  WHERE refresh_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMj'

	tx := g.db.Select("id", "refresh_token").Where("refresh_token = ?", refreshToken).Find(&result)

	return result, tx.Error
}

func (g *gormAuthRepository) FindUserInfo(userId uint) (user *entity.User, err error) {
	// 	select *
	//    from "user"
	//   where id = 4
	//     AND is_used = true;
	tx := g.db.Where("id = ? AND is_used = true", userId).Find(&user)

	return user, tx.Error

}
