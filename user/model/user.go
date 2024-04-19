package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Name            string `gorm:"type:varchar(36) not null unique; comment:'昵称'"   json:"name,omitempty"`
	Password        string `gorm:"type:varchar(36) not null unique; comment:'密码'"  json:"password,omitempty"`
	Phone           string `gorm:"type:varchar(11) not null unique; comment:'电话'"  json:"phone,omitempty"`
	Slat            string `gorm:"type:varchar(36) not null unique; comment:'盐值'"  json:"slat,omitempty"`
	Sex             int32  `gorm:"type:int not null; comment:'性别'"  json:"sex,omitempty"`
	Avatar          string `gorm:"type:varchar(36) not null unique; comment:'头像'"  json:"avatar,omitempty"`
	BackgroundImage string `gorm:"type:varchar(36) not null ; comment:'背景图'"   json:"background_image"`
	Uid             string ` gorm:"type:int not null unique; comment:'用户唯一uid'" json:"uid"`
	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}

type UserMode struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserMode {
	//err := db.AutoMigrate(&User{})
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	return &UserMode{db: db}
}

// FindPhone 检测电话号码是否存在
func (u *UserMode) FindPhone(ctx context.Context, phone string) (*User, error) {
	user := new(User)
	err := u.db.WithContext(ctx).Where("phone=?", phone).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindName 名称唯一
func (u *UserMode) FindName(ctx context.Context, name string) (*User, error) {
	user := new(User)
	err := u.db.WithContext(ctx).Where("name=?", name).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// InsertUser 注册用户
func (u *UserMode) InsertUser(ctx context.Context, user *User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *UserMode) FindUser(ctx context.Context, name string) (*User, error) {
	user := new(User)
	err := u.db.WithContext(ctx).Where("name=?", name).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
