package product

type Available struct {
	Id     int64    `json:"id,omitempty"`
	TypeId int64    `json:"type_id,omitempty"`
	Price  int64    `json:"price,omitempty"`
	Cities []string `json:"cities,omitempty"`
	Stock  []int64  `json:"stock,omitempty"`
}
