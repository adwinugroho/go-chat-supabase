package entity

type (
	Message struct {
		ID          string `json:"id"` // uuid
		Content     string `json:"content"`
		SKU         string `json:"sku,omitempty"` // uuid
		Description string `json:"description,omitempty"`
	}
)
