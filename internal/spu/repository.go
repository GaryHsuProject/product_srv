package spu

import (
	"shop/models"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.Spu) error
	FindAll() ([]*models.Spu, error)
}

type SpuRepository struct {
	orm *gorm.DB
}

func NewCategoryRepository(orm *gorm.DB) Repository {
	return &SpuRepository{orm: orm}
}

func (u *SpuRepository) Insert(spu *models.Spu) error {
	if err := u.orm.Create(&spu).Error; err != nil {
		return err
	}
	return nil
}
func (u *SpuRepository) FindAll() ([]*models.Spu, error) {
	spus := make([]*models.Spu, 0)
	if err := u.orm.Find(&spus).Error; err != nil {
		return nil, err
	}
	return spus, nil
}
