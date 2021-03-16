package models

import (
	"gorm.io/gorm"
)

type Paslon struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Bupati string `json:"bupati" gorm:"column:bupati; size:200"`
	Wakil string `json:"wakil" gorm:"column:wakil; size:200"`
	NoUrut uint8 `json:"no_urut" gorm:"column:no_urut;"`
	Foto string `json:"foto" gorm:"column:foto"`
}
func (Paslon) TableName() string{
	return "paslon" //nama table di database
}

func (p *Paslon) Get(db *gorm.DB, ID int) error {
	return db.Model(Paslon{}).Where("id = ?", ID).First(p).Error
}


func (p *Paslon) Create(db *gorm.DB) error {
	return db.Model(Paslon{}).Create(p).Error
}

func (p *Paslon) Update(db *gorm.DB) error {
	return db.Model(Paslon{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Paslon) Delete(db *gorm.DB, ID int) error {
	return db.Model(Paslon{}).Where("id = ?", ID).Delete(p).Error
}

type Paslons []Paslon

func (p *Paslons) All(db *gorm.DB) error {
	return db.Model(Paslon{}).Find(p).Error
}





