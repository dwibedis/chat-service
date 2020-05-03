package validator

import (
	"context"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/app/repository"
	"github.com/dwibedis/chat-service/app/util"
)

type User struct {
	userRepo *repository.User
}

func NewUser(userRepo *repository.User) *User {
	return &User{userRepo:userRepo};
}

func (u *User) ValidateUser(ctx context.Context, user *entities.User) bool {
	if u.isEmptyUserInput(user) {
		return false
	}

	if !util.IsEmpty(user.Id) {
		return u.ValidateUserById(ctx, user)
	}

	if user.Phone != 0 {
		return u.ValidateUserByPhone(ctx, user)
	}

	return true
}

func (u *User) isEmptyUserInput(user *entities.User) bool {
	if util.IsEmpty(user.Name) {
		return true
	}
	if user.Phone == 0 {
		return true
	}
	return false
}

func (u *User) ValidateUserById(ctx context.Context,user *entities.User) bool {
	existingUser := u.userRepo.GetUserById(ctx, user.Id)
	if existingUser.Phone == user.Phone &&
			existingUser.Name == user.Name {
		return true
	}
	return false
}

func (u *User) ValidateUserByPhone(ctx context.Context,user *entities.User) bool {
	existingUser := u.userRepo.GetUserByPhone(ctx, user.Phone)
	if existingUser.Id == user.Id &&
		existingUser.Name == user.Name {
		return true
	}
	return false
}