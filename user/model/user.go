package model

type User struct {
}

func (u *User) TableName() string {
	return "user"
}
