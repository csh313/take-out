package model

import (
	"fmt"
	"gorm.io/gorm"
	"hmshop/common/enum"
	"time"
)

type Category struct {
	Id         uint64    `json:"id"`
	Type       int       `json:"type"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

func (Category) TableName() string {
	return "category"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("before")
	//tx.Statement.SetColumn("createTime", time.Now())
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	value := tx.Statement.Context.Value("currentId")
	if uid, ok := value.(uint64); ok {
		c.UpdateUser = uid
	}
	fmt.Println("after")
	return nil
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdateTime = time.Now()
	CurrentId := tx.Statement.Context.Value(enum.CurrentId)
	fmt.Println("before update", CurrentId)
	if uid, ok := CurrentId.(uint64); ok {
		c.UpdateUser = uid
	}
	return nil
}
