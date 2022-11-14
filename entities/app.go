package entities

type HomePageReq struct {
	UserID int `json:"userID" binding:"required"` //FIXME add validate
}

type HomePageRes struct {
}
