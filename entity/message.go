package entity

type (
	Message struct {
		ID          string
		Content     string `json:"content"`
		Description string `json:"description,omitempty"`
		CreatedAt   string `json:"createdAt"` // yyyy-mm-dd hh:mm:ss
	}
)
