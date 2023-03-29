package param

type MenuParam struct {
	Id       int    `json:"id" binding:"omitempty,min=1"`         // 菜单ID，更新和删除时必填
	Name     string `json:"name" binding:"required,min=2,max=50"` // 菜单名称，必填
	Icon     string `json:"icon"`                                 //
	Path     string `json:"path" binding:"required,max=255"`      // 菜单链接，必填
	Sort     int    `json:"sort" binding:"omitempty,min=1"`       // 菜单排序，选填
	ParentID int    `json:"parent_id" binding:"omitempty,min=1"`  // 父级菜单ID，选填
}
