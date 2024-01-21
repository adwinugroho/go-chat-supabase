package model

type (
	ListAllMessageRequest struct {
		Filters map[string]interface{} `json:"filters,omitempty"`
	}

	NewSendMessageRequest struct {
		Content string `json:"content"`
	}

	NewRoomRequest struct {
		RoomID string `json:"roomId"`
		Name   string `json:"name"`
	}
)
