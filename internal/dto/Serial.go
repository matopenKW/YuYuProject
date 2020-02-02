package dto

type Serial struct {
	FloorId     string `json:",omitempty"`
	Seq         int    `json:",omitempty"`
	Acquired    bool   `json:",omitempty"`
	Acquisition string `json:",omitempty"`
}
