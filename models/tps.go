package models

import (
	"gorm.io/gorm"
)
type (
	TPS struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NoTPS uint8 `json:"no_tps" gorm:"column:no_tps;"`
		Lokasi string `json:"lokasi" gorm:"column:lokasi;size:100;"`
		Kecamatan string `json:"kecamatan" gorm:"column:kecamatan;size:50;"`
		Nagari string `json:"nagari" gorm:"column:nagari;size:50;"`
		Jorong string `json:"jorong" gorm:"column:jorong;size:50;"`
		JPL uint8 `json:"jpl" gorm:"column:jpl;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
	}
)
func (TPS) TableName() string{
	return "tps" //nama table di database
}

func (p *TPS) Get(db *gorm.DB, ID int) error {
	return db.Model(TPS{}).Where("id = ?", ID).First(p).Error
}


func (p *TPS) Create(db *gorm.DB) error {
	return db.Model(TPS{}).Create(p).Error
}

func (p *TPS) Update(db *gorm.DB) error {
	return db.Model(TPS{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *TPS) Delete(db *gorm.DB, ID int) error {
	return db.Model(TPS{}).Where("id = ?", ID).Delete(p).Error
}

type TPSs []TPS

func (p *TPSs) All(db *gorm.DB) error {
	return db.Model(TPS{}).Find(p).Error
}






