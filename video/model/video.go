package model

import "zero/pkg/orm"

type Video struct {
}

func (v *Video) TableName() string {
	return "video"
}

type VideoModel struct {
	db *orm.DB
}

func NewVideoModel(db *orm.DB) *VideoModel {
	//初始化

	return &VideoModel{db: db}
}
