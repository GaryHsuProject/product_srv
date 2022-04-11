package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

type Sku struct {
	BaseModel
	Title    string  `gorm:"column:title" json:"title"`
	SpuId    int     `gorm:"column:spu_id,index:idx_spu_id" json:"spu_id"`
	Image    string  `gorm:"column:image" json:"image"`
	Price    float32 `gorm:"column:price" json:"price"`
	Param    string  `gorm:"column:param" json:"param"`
	Saleable bool    `gorm:"column:saleable,index:idx_saleable" json:"saleable"`
	Valid    bool    `gorm:"column:valid,index:idx_valid" json:"valid"`
}
