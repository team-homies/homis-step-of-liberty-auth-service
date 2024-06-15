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
	UpdateUserInfo(user *entity.User) (err error)
	FindVisual(Code string) (user *entity.Visual, err error)
	FindVisualCode(Code string) (res *entity.Visual, err error)
	FindUserList(userId uint) (res *entity.User, err error)
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

// 사용자 본인 정보 수정 body : Nickname, Propile
func (g *gormAuthRepository) UpdateUserInfo(user *entity.User) (err error) {
	// 	update "user"
	//    set nickname  = 'woorim', profile  = 'image.url'
	//  where "user".id = 8 and "user".is_used is not null;

	// 1. gorm 적용
	tx := g.db.Begin()
	err = tx.Model(&entity.User{}).Select("nickname", "profile").Where("id = ? AND is_used = true", user.ID).Updates(&user).Error

	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()

	return

}

// 시각적 성취도 조회
func (g *gormAuthRepository) FindVisual(Code string) (res *entity.Visual, err error) {
	// 	select "code", "name", "percent", "image_url"
	// 	from visual v
	//    where code = "AM";

	// 1. gorm 적용
	err = g.db.Model(&entity.Visual{}).Select("code", "name", "percent", "image_url").Where("code = ?", Code).First(&res).Error

	if err != nil {
		return
	}
	return
}

// 시각적 성취도 코드 조회
func (g *gormAuthRepository) FindVisualCode(Code string) (res *entity.Visual, err error) {
	// 	select "code", "name", "display_level", "description"
	// 	from visual v
	//    where code = "AM";

	// 1. gorm 적용
	err = g.db.Model(&entity.Visual{}).Select("code", "name", "display_level", "description").Where("code = ?", Code).First(&res).Error

	if err != nil {
		return
	}
	return
}

// 유저 리스트 조회
func (g *gormAuthRepository) FindUserList(userId uint) (res *entity.User, err error) {
	// 	select u.id, u.email, u.nickname, u.profile, u.created_at
	//   from "user" u
	//  where u.id  = 8
	// code 추가

	// 1. gorm 적용
	tx := g.db
	err = tx.Model(&entity.User{}).Select("id", "email", "nickname", "profile", "created_at").Where("id = ? AND is_used = true", userId).Find(&res).Error
	if err != nil {
		return
	}
	return
}
