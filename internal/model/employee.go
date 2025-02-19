package model

import (
	"fmt"
	"gorm.io/gorm"
	"hmshop/common/enum"
	"time"
)

type Employee struct {
	Id         uint64    `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IdNumber   string    `json:"idNumber"`
	Status     int       `json:"status" gorm:"default:1"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

func (e *Employee) TableName() string {
	return "employee"
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("before")
	//tx.Statement.SetColumn("createTime", time.Now())
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	value := tx.Statement.Context.Value("currentId")
	if uid, ok := value.(uint64); ok {
		e.UpdateUser = uid
	}
	fmt.Println("after")
	return nil
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdateTime = time.Now()
	CurrentId := tx.Statement.Context.Value(enum.CurrentId)
	fmt.Println("before update", CurrentId)
	if uid, ok := CurrentId.(uint64); ok {
		e.UpdateUser = uid
	}
	return nil
}
