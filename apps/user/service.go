package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
)

type userService struct {
	userRepository *userRepository
}

func NewUserService(userRepository *userRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (service *userService) CreateUser(ctx context.Context, user User) (User, error) {
	user.Password, user.Salt = service.hashWithSalt(user.Password)
	return service.userRepository.CreateUser(ctx, user)
}

func (service *userService) GetUser(ctx context.Context, id int64) (User, error) {
	return service.userRepository.GetUser(ctx, id)
}

func (service *userService) UpdateUser(ctx context.Context, user User) (User, error) {
	return service.userRepository.UpdateUser(ctx, user)
}

func (service *userService) getNewSalt(saltLen int) string {
	buffer := bytes.NewBuffer(make([]byte, saltLen))
	for i := 0; i < saltLen; i++ {
		buffer.WriteByte(SaltAlphabet[rand.Intn(len(SaltAlphabet))])
	}
	return strings.ReplaceAll(buffer.String(), "\u0000", "")
}

func (service *userService) hashWithSalt(password string) (hashedPassword, salt string) {
	salt = service.getNewSalt(SaltLen)
	hashedArray := sha256.Sum256([]byte(password + salt))
	hashedPassword = hex.EncodeToString(hashedArray[:])
	return
}
