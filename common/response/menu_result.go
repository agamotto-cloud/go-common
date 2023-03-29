package response

type MenuResult struct {
	Id       int          `json:"id,key" `  //
	Name     string       `json:"name"   `  //
	Icon     string       `json:"icon"   `  //
	Path     string       `json:"path"`     //
	ParentId int          `json:"parentId"` //
	Child    []MenuResult `json:"child"`
}
