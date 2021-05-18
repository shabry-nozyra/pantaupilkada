package models

import (
	"gorm.io/gorm"
)

type (
	Petugas struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NoTPS string `json:"no_tps" gorm:"column:no_tps;"`
		Kecamatan string `json:"kecamatan" gorm:"column:kecamatan;size:100;"`
		Nagari string `json:"nagari" gorm:"column:nagari;size:100;"`
		Jorong string `json:"jorong" gorm:"column:jorong;size:100;"`
		Foto string `json:"foto" gorm:"column:foto;"`
		NoHp string `json:"no_hp" gorm:"column:no_hp;"`
		Email string `json:"email" gorm:"column:gmail;"`
		Username string `json:"username" gorm:"column:username;size:50;"`
		NamaLengkap string `json:"nama_lengkap" gorm:"column:nama_lengkap;size:50;"`
		Password string `json:"password" gorm:"column:password;"`
		RolePetugas int64 `json:"role_petugas" gorm:"column:role_petugas;"`
		IsActive int64 `json:"is_active" gorm:"column:is_active;"`
		Lokasi string `json:"lokasi" gorm:"column:lokasi;size:100;"`
		TimeKirim int64 `json:"timekirim" gorm:"column:timekirim;"`
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

func (p *Petugass) AllTPS(db *gorm.DB) error {
	return db.Model(Petugas{}).Where("role_petugas = 1").Find(p).Error
}

func (p *Petugass) AllKec(db *gorm.DB) error {
	return db.Model(Petugas{}).Where("role_petugas = 2").Find(p).Error
}

func (p *Petugass) AllKab(db *gorm.DB) error {
	return db.Model(Petugas{}).Where("role_petugas = 3").Find(p).Error
}








