package product

type Image struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id,omitempty"`
	Base64    string `json:"base_64,omitempty"`
}
