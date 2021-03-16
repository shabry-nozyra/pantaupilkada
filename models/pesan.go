package models

import (
	"gorm.io/gorm"
)

type (
	Pesan struct {
		ID uint 		`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Name string 	`json:"name" gorm:"column:name;size:40"`
		Phone string 	`json:"phone" gorm:"column:phone; size:15"`
		Email string 	`json:"email" gorm:"column:email;size:40"`
		Perihal string `json:"perihal" gorm:"column:perihal;"`
		Message string `json:"message" gorm:"column:message;"`
		TimeCreated uint8 `json:"timecreate" gorm:"column:timecreate;"`
	}
)


func (Pesan) TableName() string{
	return "pesan" //nama table di database
}

func (p *Pesan) Get(db *gorm.DB, ID int) error {
	return db.Model(Pesan{}).Where("id = ?", ID).First(p).Error
}


func (p *Pesan) Create(db *gorm.DB) error {
	return db.Model(Pesan{}).Create(p).Error
}

func (p *Pesan) Update(db *gorm.DB) error {
	return db.Model(Pesan{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Pesan) Delete(db *gorm.DB, ID int) error {
	return db.Model(Pesan{}).Where("id = ?", ID).Delete(p).Error
}

type Pesans []Pesan

func (p *Pesans) All(db *gorm.DB) error {
	return db.Model(Pesan{}).Find(p).Error
}






