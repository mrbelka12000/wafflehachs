package request

type CreateRoomRequest struct {
	Username string `json:"username"`
}

func (crr *CreateRoomRequest) IsValid() bool {
	return crr.Username != ""
}
