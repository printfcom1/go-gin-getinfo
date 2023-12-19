package informations

type Informations struct {
	ID         int    `json:"id" db:"id"`
	ParentID   int    `json:"parent_id" db:"parent_id"`
	Title      string `json:"title" db:"title"`
	Name       string `json:"name" db:"name"`
	Route      string `json:"route" db:"route"`
	Icon       string `json:"icon" db:"icon"`
	IsChildren bool   `json:"is_children" db:"is_children"`
}

type InformationsChild struct {
	ID         int                 `json:"id" db:"id"`
	ParentID   int                 `json:"parent_id" db:"parent_id"`
	Title      string              `json:"title" db:"title"`
	Name       string              `json:"name" db:"name"`
	Route      string              `json:"route" db:"route"`
	Icon       string              `json:"icon" db:"icon"`
	IsChildren bool                `json:"is_children" db:"is_children"`
	Children   []InformationsChild `json:"children"`
}

type InformationsRes struct {
	ID         int                 `json:"id"`
	ParentID   int                 `json:"parent_id" `
	Title      string              `json:"title"`
	Name       string              `json:"name" `
	Route      string              `json:"route" `
	Icon       string              `json:"icon" `
	IsChildren bool                `json:"is_children" `
	Children   []InformationsChild `json:"children"`
}
