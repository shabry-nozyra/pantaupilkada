package models

import "gorm.io/gorm"

type (
	Konten struct {
		ID uint 			`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Halaman string 		`json:"halaman" gorm:"column:halaman;size:100"`
		Konten string 		`json:"konten" gorm:"column:konten; size:100"`
		IsiKonten string 	`json:"isikonten" gorm:"column:isikonten;"`
		URL string 			`json:"url" gorm:"column:url;"`
		IsActive uint8 		`json:"isactive" gorm:"column:isactive;"`
	}
)


func (Konten) TableName() string{
	return "konten" //nama table di database
}
func (p *Konten) Get(db *gorm.DB, ID int) error {
	return db.Model(Konten{}).Where("id = ?", ID).First(p).Error
}


func (p *Konten) Create(db *gorm.DB) error {
	return db.Model(Konten{}).Create(p).Error
}

func (p *Konten) Update(db *gorm.DB) error {
	return db.Model(Konten{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Konten) Delete(db *gorm.DB, ID int) error {
	return db.Model(Konten{}).Where("id = ?", ID).Delete(p).Error
}

type Kontens []Konten

func (p *Kontens) All(db *gorm.DB) error {
	return db.Model(Konten{}).Find(p).Error
}





