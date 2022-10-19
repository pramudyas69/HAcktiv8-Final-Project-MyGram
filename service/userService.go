package service

import (
	"MyGramHacktiv8/dto"
	"MyGramHacktiv8/entity"
	"MyGramHacktiv8/pkg/errs"
	"MyGramHacktiv8/pkg/helpers"
	"MyGramHacktiv8/repository/userRepository"
	"fmt"
)

type UserService interface {
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	UpdateUserData(userId uint, userPayload *dto.UpdateUserDataRequest) (*dto.UpdateUserDataResponse, errs.MessageErr)
	DeleteUser(userId uint) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo userRepository.UserRepository
}

func NewUserService(userRepo userRepository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Email: userPayload.Email,
	}

	user, err := u.userRepo.Login(payload)
	if err != nil {
		return nil, err
	}

	validPassword := user.VerifyPassword(userPayload.Password)
	if !validPassword {
		return nil, errs.NewUnAuthorized("Invalid email or password")
	}

	token := user.GenerateToken()

	response := &dto.LoginResponse{
		AccessToken: token,
	}

	return response, nil
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Age:      userPayload.Age,
	}

	err = user.HashPass()
	if err != nil {
		return nil, err
	}

	user, err = u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return response, nil
}

func (u *userService) UpdateUserData(userId uint, userPayload *dto.UpdateUserDataRequest) (*dto.UpdateUserDataResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	user, err = u.userRepo.UpdateUserData(userId, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data user", user.Email)

	response := &dto.UpdateUserDataResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(userId uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "User data has been successfully deleted",
	}

	return response, nil
}
