package tags

type Tags struct {
	Id       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(255)"`
	LastName string `gorm:"type:varchar(255)"`
}
