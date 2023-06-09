package controller

type CreateCommentRequest struct {
	UserId  uint   `json:"user_id" binding:"required"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" binding:"required"`
}
