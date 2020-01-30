package dto

type Tenanto struct {
	Seq         int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Acquisition string `json:",omitempty"`
}

type TenantoView struct {
	Tenanto
	ClassName string `json:",omitempty"`
}
