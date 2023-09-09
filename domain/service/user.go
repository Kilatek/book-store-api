package service

import (
	"context"
	"time"

	"bookstore.com/domain/entity"
	portError "bookstore.com/port/error"
	"bookstore.com/port/payload"
	"bookstore.com/repository"
	"bookstore.com/tools/mapper"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, req *payload.RegisterRequest) error {
	user := &entity.User{}

	if err := mapper.MapStructsWithJSONTags(req, user); err != nil {
		return err
	}

	userTmp, err := s.userRepo.Find(ctx, user.Username)
	if err != nil && err.Error() != "user not found" {
		return err
	}

	if userTmp != nil {
		return portError.NewBadRequestError("user does exist", nil)
	}

	user.Password, err = Hash(user.Password)
	if err != nil {
		return err
	}

	return s.userRepo.Store(ctx, user)
}

func (s *userService) Login(ctx context.Context, req *payload.LoginRequest) (*payload.LoginResponse, error) {
	var jwtKey = []byte("my_secret_key")
	res := &payload.LoginResponse{}

	user_tmp, err := s.userRepo.Find(ctx, req.Username)
	if err != nil {
		return res, err
	}

	err = CheckPasswordHash(user_tmp.Password, req.Password)
	if err != nil {
		return res, err
	}
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &entity.Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return res, err
	}

	res.Token = tokenString

	return res, nil
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
