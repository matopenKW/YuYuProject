package dto

type Tenanto struct {
	Seq         int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Acquisition string `json:",omitempty"`
	ClassName   string `json:",omitempty"`
	Score       int    `json:",omitempty"`
}
