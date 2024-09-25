package constant

type PageForm struct {
	Page int `form:"page" comment:"当前页码" binding:"omitempty,min=1"`
	Size int `form:"size" comment:"每页显示记录数" binding:"omitempty,min=1"`
}

type PageList struct {
	Size    int   `json:"size"`
	Pages   int64 `json:"pages"`
	Total   int64 `json:"total"`
	Current int   `json:"current"`
}

type UIdForm struct {
	Id uint `uri:"id" comment:"id" binding:"min=1"`
}
