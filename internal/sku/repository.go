package sku

import (
	"shop/models"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.Sku) error
	FindAll() ([]*models.Sku, error)
}

type SkuRepository struct {
	orm *gorm.DB
}

func NewSkuRepository(orm *gorm.DB) Repository {
	return &SkuRepository{orm: orm}
}

func (u *SkuRepository) Insert(sku *models.Sku) error {
	if err := u.orm.Create(&sku).Error; err != nil {
		return err
	}
	return nil
}
func (u *SkuRepository) FindAll() ([]*models.Sku, error) {
	skus := make([]*models.Sku, 0)
	if err := u.orm.Find(&skus).Error; err != nil {
		return nil, err
	}
	return skus, nil
}
