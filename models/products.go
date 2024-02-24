package models

type Product struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"product-name"`
	Description string `gorm:"text" json:"description"`
}
