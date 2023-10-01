package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/pws-backend/domain"
	"github.com/pws-backend/internal/tokenutil"
)

type authInitUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func AuthInitUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &authInitUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *authInitUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *authInitUsecase) CheckUserExists(c context.Context) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	users, err := su.userRepository.Fetch(ctx, 1)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("No user was found!")
	}
	return &users[0], nil
}

func (su *authInitUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *authInitUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
