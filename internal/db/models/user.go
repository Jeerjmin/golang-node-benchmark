package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Password string `gorm:"size:255" json:"-"`
}
