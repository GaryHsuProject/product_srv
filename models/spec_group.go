package models

type SpecGroup struct {
	Id    int    `gorm:"column:id,primary" json:"id"`
	SpgId int    `gorm:"column:spg_id,index:idx_spg_id,unique" json:"spg_id"`
	Name  string `gorm:"column:name,unique" json:"name"`
}
