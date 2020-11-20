package model

type Model struct {
	ID  uint32 `gorm:"primary_ke" json:"id"`
	CreatedOn string `json:"created_on"`
	DeleteOn string `json:"delete_on"`
}
