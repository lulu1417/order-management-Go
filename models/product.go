package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	v := validator.New()
	// 在這個初始化函數中註冊自定義驗證函數
	err := v.RegisterValidation("customUniqueName", customUniqueName)
	if err != nil {
		return
	}
}

func customUniqueName(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	if !ok {
		// 如果無法將欄位轉換為字串，返回 false
		return false
	}

	var db *gorm.DB
	if IsNameUnique(db, name) {
		return true
	}
	return false
}

func IsNameUnique(db *gorm.DB, name string) bool {
	var count int
	db.Model(&Product{}).Where("name = ?", name).Count(&count)
	return count == 0
}
