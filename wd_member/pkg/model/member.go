package model

import (
	"context"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Member struct {
	MemberID     string    `gorm:"column:member_id;type:varchar(40);primaryKey:MEMBER_ID_PK"`
	Name         string    `gorm:"column:username;type:varchar(255)"`
	Email        string    `gorm:"column:email;type:varchar(255);unique"`
	PhoneNumber  string    `gorm:"column:phone_number;type:varchar(255)"`
	Nickname     string    `gorm:"column:nickname;type:varchar(255)"`
	Password     string    `gorm:"column:password;type:varchar(255)"`
	ProfileIMG   string    `gorm:"column:profile_img;type:text"`
	Status       string    `gorm:"column:status;type:varchar(100)"`
	MemberStatus string    `gorm:"column:member_status;type:varchar(50)"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(6)"`
	UpdatedAt    time.Time `gorm:"column:created_at;type:datetime(6)"`
	DeletedAt    time.Time `gorm:"column:created_at;type:datetime(6)"`
}

func (m *Member) TableName() string {
	return "member"
}

func (m *Member) create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(m).Error
}

func (m *Member) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(m.Password), 10)
	if err != nil {
		return errors.New("fail hash password")
	}
	m.Password = hex.EncodeToString(hash)
	return nil
}
