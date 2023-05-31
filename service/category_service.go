package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func(service CategoryService) Get(id string) (*entity.Category,error) {
	category := service.repository.FindById(id)
	if category == nil {
		return category,errors.New("category not found!!!")
	}else {
		return category,nil
	}
}
