package model

type Model struct {
	ID  uint32 `gorm:"primary_ke" json:"id"`
	CreatedOn uint32 `gorm:"default:NULl" json:"created_on"`
	DeletedOn uint32 `gorm:"default:NUll" json:"deleted_on"`
}
