package db

import (
	"github.com/niqinge/test-simple/pkg/model"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (db *UserDB) Create(u *model.User) error {
	return db.db.Create(u).Error
}

func (db *UserDB) Update(fields map[string]interface{}) error {
	return db.db.Updates(fields).Error
}

func (db *UserDB) Delete(userId string) error {
	return db.db.Where("user_id = ?", userId).Delete(&model.User{}).Error
}

func (db *UserDB) QueryByUserId(userId string) (user *model.User, err error) {
	err = db.db.Where("user_id = ?", userId).Last(&user).Error
	return
}

func (db *UserDB) QueryPage(pageNo, size int, fields map[string]interface{}) (users []*model.User, err error) {
	err = db.db.Where(fields).Order("id desc").Limit(size).Offset(pageNo).Find(&users).Error
	return
}
