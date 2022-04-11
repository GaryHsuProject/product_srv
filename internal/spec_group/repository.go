package spec_group

import (
	"shop/models"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.SpecGroup) error
	FindAll() ([]*models.SpecGroup, error)
}

type SpecGroupRepository struct {
	orm *gorm.DB
}

func NewCategoryRepository(orm *gorm.DB) Repository {
	return &SpecGroupRepository{orm: orm}
}

func (u *SpecGroupRepository) Insert(specGroup *models.SpecGroup) error {
	if err := u.orm.Create(&specGroup).Error; err != nil {
		return err
	}
	return nil
}
func (u *SpecGroupRepository) FindAll() ([]*models.SpecGroup, error) {
	specGroup := make([]*models.SpecGroup, 0)
	if err := u.orm.Find(&specGroup).Error; err != nil {
		return nil, err
	}
	return specGroup, nil
}
