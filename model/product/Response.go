package product

type Response struct {
	Id          int64  `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int64  `json:"price,omitempty"`
	CategoryId  int64  `json:"category_id,omitempty"`
	BrandId     int64  `json:"brand_id"`

	Image []Image `json:"image"`
	Type  []Type  `json:"type"`
}
