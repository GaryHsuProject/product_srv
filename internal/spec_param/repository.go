package spec_param

import (
	"shop/models"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.SpecParam) error
	FindAll() ([]*models.SpecParam, error)
}

type SpecParamRepository struct {
	orm *gorm.DB
}

func NewCategoryRepository(orm *gorm.DB) Repository {
	return &SpecParamRepository{orm: orm}
}

func (u *SpecParamRepository) Insert(specParam *models.SpecParam) error {
	if err := u.orm.Create(&specParam).Error; err != nil {
		return err
	}
	return nil
}
func (u *SpecParamRepository) FindAll() ([]*models.SpecParam, error) {
	specParam := make([]*models.SpecParam, 0)
	if err := u.orm.Find(&specParam).Error; err != nil {
		return nil, err
	}
	return specParam, nil
}
