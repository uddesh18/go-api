package authentication

import "time"

type UserLoginDetail struct {
	ID            int       `gorm:"type:int;primary_key"`
	Username      string    `gorm:"type:varchar(255)"`
	Password      string    `gorm:"type:varchar(255)"`
	LoginDateTime time.Time `gorm:"type:datetime"`
}
