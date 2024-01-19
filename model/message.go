package model

type (
	NewSendMessageRequest struct {
		Content string `json:"content"`
		SKU     string `json:"sku"`
	}
)
