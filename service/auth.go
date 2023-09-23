package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/repository"
)

type authService struct {
	authRepo  repository.AuthRepository
	SecretKey []byte
}

type claims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (a authService) Register(ctx context.Context, user model.User) error {
	err := a.authRepo.Register(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a authService) Login(ctx context.Context, role string, email string, password string) (string, error) {
	user, err := a.authRepo.Login(ctx, role, email, password)
	if err != nil {
		return "", err
	}

	claim := claims{
		ID:             user.ID,
		Name:           user.Name,
		Role:           user.Role,
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	tokenString, err := token.SignedString(a.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewAuthService(authRepo repository.AuthRepository, secretKey []byte) AuthService {
	return authService{
		authRepo:  authRepo,
		SecretKey: secretKey,
	}
}
