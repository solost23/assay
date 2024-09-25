package forms

import "assay/infra/constant"

type UserInsertForm struct {
	Username string `json:"username" binding:"required`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Email    string `json:"email" binding:"required"`
	RoleId   uint   `json:"roleId" binding:"required"`
}

type UserListForm struct {
	constant.PageForm
	Username string `form:"username"`
}

type UserList struct {
	constant.PageList
	Records []UserListRecord `json:"records"`
}

type UserListRecord struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	UpdatedAt string `json:"updatedAt"`
}
