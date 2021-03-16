package models

import (
	"gorm.io/gorm"
)

type Jorong struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	NamaJorong string `json:"nama_jorong" gorm:"column:nama_jorong;"`
	NamaNagari string `json:"nama_nagari" gorm:"column:nama_nagari;"`
	NamaKecamatan string `json:"nama_kecamatan" gorm:"column:nama_kecamatan;"`
	IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
}

func (Jorong) TableName() string{
	return "data_jorong" //nama table di database
}
func (p *Jorong) Get(db *gorm.DB, ID int) error {
	return db.Model(Jorong{}).Where("id = ?", ID).First(p).Error
}


func (p *Jorong) Create(db *gorm.DB) error {
	return db.Model(Jorong{}).Create(p).Error
}

func (p *Jorong) Update(db *gorm.DB) error {
	return db.Model(Jorong{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Jorong) Delete(db *gorm.DB, ID int) error {
	return db.Model(Jorong{}).Where("id = ?", ID).Delete(p).Error
}

type Jorongs []Jorong

func (p *Jorongs) All(db *gorm.DB) error {
	return db.Model(Jorong{}).Find(p).Error
}





