package services

import (
	"context"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo interfaces.IUserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, req *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashPassword)

	req.Role = string(models.RoleUser)

	err = s.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}

func (s *UserService) RegisterAdmin(ctx context.Context, req *models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashPassword)

	req.Role = string(models.RoleAdmin)

	err = s.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}