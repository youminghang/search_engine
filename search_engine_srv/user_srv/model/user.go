package model

type User struct {
	BaseModel
	Mobile   string `gorm:"index:idx_modile;unique;type:varchar(11)"`
	NickName string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(100);not null"`
}
