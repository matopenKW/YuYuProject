package dto

type Team struct {
	Id        string `json:",omitempty"`
	ClassName string `json:",omitempty"`
	Uid       string `json:",omitempty"`
	All       int    `json:",omitempty"`
	West      int    `json:",omitempty"`
	East      int    `json:",omitempty"`
	Score     int    `json:",omitempty"`
}

type TeamProduct struct {
	*Team
	TenantId    string
	ProductList []*Product
	ProductNum  int
	Rate        int
}
