package forms

type RoleInsertForm struct {
	Name       string   `json:"name" binding:"required"`
	Nickname   string   `json:"nickname"`
	RouterAuth []string `json:"routerAuth"`
	ButtonAuth []string `json:"buttonAuth"`
}
