package param

type TagPost struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}

type TagPut struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}

type TagList struct {
	Status int  `form:"status" json:"status" xml:"status" binding:"oneof=0 1"`
	Name string `form:"name" json:"name" xml:"name"  binding:"max=100"`
}


