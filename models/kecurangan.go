package models

import (
	"gorm.io/gorm"
)

type (
	Kecurangan struct {
		ID uint 			`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Detail string 		`json:"halaman" gorm:"column:halaman;size:100"`
		WaktuKejadian string `json:"waktu" gorm:"column:waktu;"`
		Jam uint 			`json:"jam" gorm:"column:jam;"`
		Tempat string 		`json:"tempat" gorm:"column:tempat;"`
		Bukti1 string 		`json:"bukti1" gorm:"column:bukti1;"`
		Bukti2 string 		`json:"bukti2" gorm:"column:bukti2;"`
		Bukti3 string 		`json:"bukti3" gorm:"column:bukti3;"`
	}
)


func (Kecurangan) TableName() string{
	return "kecurangan" //nama table di database
}
func (p *Kecurangan) Get(db *gorm.DB, ID int) error {
	return db.Model(Kecurangan{}).Where("id = ?", ID).First(p).Error
}


func (p *Kecurangan) Create(db *gorm.DB) error {
	return db.Model(Kecurangan{}).Create(p).Error
}

func (p *Kecurangan) Update(db *gorm.DB) error {
	return db.Model(Kecurangan{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Kecurangan) Delete(db *gorm.DB, ID int) error {
	return db.Model(Kecurangan{}).Where("id = ?", ID).Delete(p).Error
}

type Kecurangans []Kecurangan

func (p *Kecurangans) All(db *gorm.DB) error {
	return db.Model(Kecurangan{}).Find(p).Error
}






