package product

type Color struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Hex       string `json:"hex,omitempty"`
}
