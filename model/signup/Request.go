package signup

type Request struct {
	PhoneNumber float64 `json:"phone_number,omitempty"`
	Name        string  `json:"name,omitempty"`
	Address     string  `json:"address,omitempty"`
	Nik         float64 `json:"nik,omitempty"`
	SessionId   string  `json:"session_id,omitempty"`
}
