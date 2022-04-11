package models

type Spu struct {
	BaseModel
	Title      string `gorm:"title" json:"title"`
	SubTitle   string `gorm:"sub_title" json:"sub_title"`
	CategoryId int    `gorm:"category_id,index:idx_category_id" json:"category_id"`
	SpgId      int    `gorm:"spg_id,index:idx_spg_id" json:"spg_id"`
	Saleable   bool   `gorm:"saleable,index:idx_saleable" json:"saleable"`
	Valid      string `gorm:"valid,index:idx_valid" json:"valid"`
}
