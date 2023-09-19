package usecase

import (
	"context"
	"time"

	"github.com/pws-backend/domain"
	"github.com/pws-backend/internal/tokenutil"
	"gorm.io/gorm"
)

type loginUsecase struct {
	db             *gorm.DB
	contextTimeout time.Duration
}

func NewLoginUsecase(db *gorm.DB, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	_, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	var result = domain.User{Email: email}
	err := lu.db.First(&result).Error
	return result, err
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
