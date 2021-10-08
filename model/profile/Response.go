package profile

type Response struct {
	PhoneNumber float64 `json:"phone_number"`
	Name        string  `json:"name,omitempty"`
	Address     string  `json:"address,omitempty"`
	Nik         float64 `json:"nik,omitempty"`
	SessionId   string  `json:"session_id,omitempty"`
}
