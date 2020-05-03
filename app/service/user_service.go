package service

import (
	"context"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/app/repository"
	"github.com/dwibedis/chat-service/app/util"
	"github.com/dwibedis/chat-service/app/validator"
	"log"
)

type User struct {
	userRepo *repository.User
	userOtpRepo *repository.UserOtp
	userValidator *validator.User
}

func NewUser(userRepo *repository.User, userValidator *validator.User, userOtpRepo *repository.UserOtp) *User {
	return &User{
		userRepo:      userRepo,
		userValidator: userValidator,
		userOtpRepo:   userOtpRepo,
	}
}

func (u *User) RegisterUser(ctx context.Context, user *entities.User) *entities.User {
	isValid := u.userValidator.ValidateUser(ctx, user)
	if !isValid {
		return nil
	}
	existingUser := u.userRepo.GetUserByPhone(ctx, user.Phone)
	if existingUser != nil {
		return existingUser
	}
	newUser, err := u.userRepo.AddUser(ctx, user)
	if err != nil {
		return nil
	}
	return newUser
}

func (u *User) GenerateAndSendUserOtp(ctx context.Context, user *entities.User) error  {
	isValid := u.userValidator.ValidateUser(ctx, user)
	if !isValid {
		return nil
	}
	otp := util.GenerateRandomNumber()
	err := u.userOtpRepo.SaveUserPhoneAndOtp(ctx, user, otp)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) VerifyOtp(ctx context.Context, userOtp *entities.UserOtp) bool  {
	savedOtp := u.userOtpRepo.GetOtpByPhone(ctx, userOtp.UserPhone)
	inputOtp := userOtp.Otp
	log.Println("Saved otp ", savedOtp, " input Otp ", inputOtp)
	if savedOtp == inputOtp {
		u.userOtpRepo.MarkPhoneOtpInvalid(ctx, userOtp.UserPhone)
		return true
	}
	return false
}