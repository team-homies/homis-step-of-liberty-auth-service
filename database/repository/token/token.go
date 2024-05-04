package token

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
	FindUserByUserInfo(email, provider string) (user *entity.User, err error)
	CreateUser(email, provider string) error
}

type gormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &gormAuthRepository{db}
}

func (g *gormAuthRepository) Create(user *entity.User) error {

	panic("")
}

func (g *gormAuthRepository) Delete(id uint) error {
	panic("")
}

func (g *gormAuthRepository) GetAll() ([]entity.User, error) {
	panic("")
}

func (g *gormAuthRepository) GetByID(id uint) (*entity.User, error) {
	panic("")
}

func (g *gormAuthRepository) Update(user *entity.User) error {
	panic("")
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
