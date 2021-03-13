package dao

import (
	"github.com/niqinge/test-simple/common"
	"github.com/niqinge/test-simple/pkg/model"
	"github.com/pkg/errors"
)

type UserDB interface {
	Create(u *model.User) error
	Update(fields map[string]interface{}) error
	Delete(userId string) error
	QueryByUserId(userId string) (user *model.User, err error)
	QueryPage(pageNo, size int, fields map[string]interface{}) (users []*model.User, err error)
}

type UserDAO struct {
	db UserDB
}

func (dao *UserDAO) Create(u *model.User) error {
	return dao.db.Create(u)
}

func (dao *UserDAO) UpdateByUserId(fields map[string]interface{}) error {
	if len(fields) == 0 || fields["user_id"] == nil {
		return errors.Errorf("update bu user id, user_id not found, fields:%+v", fields)
	}
	return dao.db.Update(fields)
}

func (dao *UserDAO) DelByUserId (userId string) error {
	return dao.db.Delete(userId)
}

func (dao *UserDAO) QueryByUserId(userId string) (*model.User, error) {
	return dao.db.QueryByUserId(userId)
}

func (dao *UserDAO) QueryList (pageNo, size int, fields map[string]interface{}) ([]*model.User, error) {
	if size == 0 {
		size = common.PageSize
	}
	
	return dao.db.QueryPage(pageNo, size, fields)
}