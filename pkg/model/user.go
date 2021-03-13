package model

import "time"

type User struct {
	BaseModel
	UserID   string     `json:"user_id" grom:"index:idx_user_id" name:"用户编号"`
	Name     string     `json:"name" name:"昵称"`
	Status   string     `json:"status" name:"状态:2(封号), 1(正常)"`
	Birthday *time.Time `json:"birthday" name:"出生日期"`
	PicUrl   string     `json:"pic_url" name:"图片地址"`
}
