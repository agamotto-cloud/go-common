package param

type MenuSortParam struct {
	PosId   int  `json:"posId"`
	MenuId  int  `json:"menuId"`
	IsChild bool `json:"isChild"`
	IsPre   bool `json:"isPre"`
}
