package model

import "gorm.io/gorm"

type User struct {
	Username string `gorm:"type:varchar(36) not null unique; comment:'用户名'"   json:"username,omitempty"`
	Password string `gorm:"type:varchar(36) not null unique; comment:'密码'"  json:"password,omitempty"`
	Phone    string `gorm:"type:varchar(11) not null unique; comment:'电话'"  json:"phone,omitempty"`
	Slat     string `gorm:"type:varchar(36) not null unique; comment:'盐值'"  json:"slat,omitempty"`
	Url      string `gorm:"type:varchar(36) not null unique; comment:'头像"   json:"url"`
	Uid      uint   ` gorm:"type:int not null unique; comment:'用户唯一id'" json:"uid"`
	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}

type UserMode struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserMode {
	return &UserMode{db: db}
}
