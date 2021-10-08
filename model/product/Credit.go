package product

type Credit struct {
	Id          int64  `json:"id"`
	ProductId   int64  `json:"product_id,omitempty"`
	DownPayment int64  `json:"down_payment"`
	Tenor       string `json:"tenor,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
}
