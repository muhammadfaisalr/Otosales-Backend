package product

type Type struct {
	Id              int64  `json:"id"`
	ProductId       int64  `json:"product_id"`
	Name            string `json:"name,omitempty"`
	Transmission    string `json:"transmission,omitempty"`
	CentimeterCubic int64  `json:"centimeter_cubic,omitempty"`
	EngineType      string `json:"engine_type,omitempty"`
	MaxTorque       string `json:"max_torque,omitempty"`
	Type            string `json:"type"`

	Available []Available `json:"available"`
	Color     []Color     `json:"color"`
	Credit    []Credit    `json:"credit"`
}
