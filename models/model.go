package models

import "gorm.io/gorm"

func MigrateModel(db *gorm.DB) error  {
	return db.AutoMigrate(&Paslon{},&TPS{}, &Admin{}, &AdminRole{}, &Jorong{},
	&Kecurangan{}, &Konten{}, &Paslon{}, &Pesan{}, &Petugas{}, &TPS{})
}
