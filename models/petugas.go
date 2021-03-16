package models

import (
	"gorm.io/gorm"
)

type (
	Petugas struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NoTPS uint8 `json:"no_tps" gorm:"column:no_tps;"`
		Kecamatan string `json:"kecamatan" gorm:"column:kecamatan;size:100;"`
		Nagari string `json:"nagari" gorm:"column:nagari;size:100;"`
		Jorong string `json:"jorong" gorm:"column:jorong;size:100;"`
		Foto uint8 `json:"foto" gorm:"column:foto;"`
		NoHp uint8 `json:"no_hp" gorm:"column:no_hp;"`
		Email string `json:"email" gorm:"column:email;"`
		Username string `json:"username" gorm:"column:username;size:50;"`
		NamaLengkap string `json:"nama_lengkap" gorm:"column:nama_lengkap;size:50;"`
		Password string `json:"password" gorm:"column:password;"`
		RolePetugas uint8 `json:"role_petugas" gorm:"column:role_petugas;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
		TimeKirim uint8 `json:"timekirim" gorm:"column:timekirim;"`
	}
)
func (Petugas) TableName() string{
	return "petugas" //nama table di database
}

func (p *Petugas) Get(db *gorm.DB, ID int) error {
	return db.Model(Petugas{}).Where("id = ?", ID).First(p).Error
}


func (p *Petugas) Create(db *gorm.DB) error {
	return db.Model(Petugas{}).Create(p).Error
}

func (p *Petugas) Update(db *gorm.DB) error {
	return db.Model(Petugas{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Petugas) Delete(db *gorm.DB, ID int) error {
	return db.Model(Petugas{}).Where("id = ?", ID).Delete(p).Error
}

type Petugass []Petugas

func (p *Petugass) All(db *gorm.DB) error {
	return db.Model(Petugas{}).Find(p).Error
}






