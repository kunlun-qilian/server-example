package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID string `gorm:"index:I_user_id,unique" column:"f_user_id" type:"varchar" size:64 not null`
}
