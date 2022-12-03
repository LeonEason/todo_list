package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index:not null"`
	Status    int    `gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64  //开始记录的时间
	EndTime   int64  //结束记录的时间
}
