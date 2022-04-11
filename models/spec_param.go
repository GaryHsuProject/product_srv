package models

type SpecParam struct {
	Id        int    `gorm:"column:id,primary" json:"id"`
	SpgId     int    `gorm:"column:spg_id,index:idx_spg_id" json:"spg_id"`
	Name      string `gorm:"column:name" json:"name"`
	Numeric   bool   `gorm:"column:numeric" json:"numberic"`
	Unit      string `gorm:"unit" json:"unit"`
	Generic   bool   `gorm:"generic" json:"generic"`
	Searching bool   `gorm:"searching" json:"searching"`
	Segements string `gorm:"segements" json:"segements"`
}
