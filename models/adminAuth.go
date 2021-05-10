package models

import (
	"gorm.io/gorm"
)

type (
	Admin struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement;unique"`
		Name string `json:"name" gorm:"column:name;size:100;"`
		Email string `json:"email" gorm:"column:email;size:50;unique"`
		Password string `json:"password" gorm:"column:password;"`
		Image string `json:"image" gorm:"column:image;"`
		RoleAdmin uint `json:"role_id" gorm:"column:role_id;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
		TimeCreated int64 `json:"timecreated" gorm:"column:time_created;null"`
	}
)
func (Admin) TableName() string{
	return "admin_auth" //nama table di database
}
func (p *Admin) Get(db *gorm.DB, ID int) error {
	return db.Model(Admin{}).Where("id = ?", ID).First(p).Error
}

func (p *Admin) GetByEmail(db *gorm.DB, Email string) error {
	return db.Model(Admin{}).Where("email = ?", Email).First(p).Error
}


func (p *Admin) Create(db *gorm.DB) error {
	return db.Model(Admin{}).Create(p).Error
}

func (p *Admin) Update(db *gorm.DB) error {
	return db.Model(Admin{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Admin) Delete(db *gorm.DB, ID int) error {
	return db.Model(Admin{}).Where("id = ?", ID).Delete(p).Error
}


func MaxId(db *gorm.DB) uint{
	var max uint
	db.Table("admin_auth").Select("max(id)").Row().Scan(&max)
	return max
}

type Admins []Admin

func (a *Admins) All(db *gorm.DB) error {
	return db.Model(Admin{}).Find(a).Error
}





