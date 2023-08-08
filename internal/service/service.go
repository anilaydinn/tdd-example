package service

import "github.com/anilaydinn/tdd-example/internal/model"

type Service interface {
	CreateStudent(request model.CreateStudentRequest) (*model.CreateStudentResponse, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) CreateStudent(request model.CreateStudentRequest) (*model.CreateStudentResponse, error) {
	return nil, nil
}
