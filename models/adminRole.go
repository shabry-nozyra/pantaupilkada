package models

import (
	"gorm.io/gorm"
)

type (
	AdminRole struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		RoleAdmin string `json:"role" gorm:"column:role;"`
	}
)
func (AdminRole) TableName() string{
	return "adminrole" //nama table di database
}
func (p *AdminRole) Get(db *gorm.DB, ID int) error {
	return db.Model(AdminRole{}).Where("id = ?", ID).First(p).Error
}


func (p *AdminRole) Create(db *gorm.DB) error {
	return db.Model(AdminRole{}).Create(p).Error
}

func (p *AdminRole) Update(db *gorm.DB) error {
	return db.Model(AdminRole{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *AdminRole) Delete(db *gorm.DB, ID int) error {
	return db.Model(AdminRole{}).Where("id = ?", ID).Delete(p).Error
}

type AdminRoles []AdminRole

func (p *AdminRoles) All(db *gorm.DB) error {
	return db.Model(AdminRole{}).Find(p).Error
}






