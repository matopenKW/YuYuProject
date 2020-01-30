package dto

type Team struct {
	Id        string `json:",omitempty"`
	ClassName string `json:",omitempty"`
	Uid       string `json:",omitempty"`
}

type TeamRate struct {
	Team
	All  int `json:",omitempty"`
	West int `json:",omitempty"`
	East int `json:",omitempty"`
}
