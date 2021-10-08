package article

type Request struct {
	Title    string  `json:"title,omitempty"`
	Content  string  `json:"content,omitempty"`
	AuthorId float64 `json:"author_id,omitempty"`
}
