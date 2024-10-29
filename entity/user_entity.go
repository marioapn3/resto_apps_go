package entity

import "time"

type User struct {
	ID             int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Name           string `gorm:"type:varchar(255)" json:"name"`
	Occupation     string `gorm:"type:varchar(255)" json:"occupation"`
	Email          string `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password       string `gorm:"type:varchar(255)" json:"password"`
	AvatarFileName string `gorm:"type:varchar(255)" json:"avatar_file_name"`
	Role           string `gorm:"type:varchar(255)" json:"role"`
	Token          string `gorm:"-" json:"token"`

	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}
