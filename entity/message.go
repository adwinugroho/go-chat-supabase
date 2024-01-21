package entity

type (
	Message struct {
		MessageID   string   `json:"message_id"`
		Content     []string `json:"content"`
		UserID      []string `json:"user_id"`
		Description string   `json:"description,omitempty"`
		CreatedAt   string   `json:"created_at"` // yyyy-mm-dd hh:mm:ss
	}
)
