package models

import "gorm.io/gorm"

type (
	Lokasi struct {
		ID uint 				`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NamaKab string 			`json:"nama_kabupaten" gorm:"column:nama_kabupaten;"`
		NamaNagari string 		`json:"nama_nagari" gorm:"column:nama_nagari;"`
		NamaKecamatan string 	`json:"nama_kecamatan" gorm:"column:nama_kecamatan;"`
		IsActive int64 			`json:"is_active" gorm:"column:is_active;"`
	}
)


func (Lokasi) TableName() string{
	return "data_jorong" //nama table di database
}

type Lokasis []Lokasi

func (p *Lokasis) All(db *gorm.DB) error {
	return db.Model(Lokasi{}).Find(p).Error
}





