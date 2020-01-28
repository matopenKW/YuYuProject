package dto

type Team struct {
	Name      string `json:",omitempty"`
	ClassName string `json:",omitempty"`
	Rate      int    `json:",omitempty"`
	Uid       string `json:",omitempty"`
}
