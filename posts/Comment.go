package posts

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/repository"
)

type CreateCommentRequest struct {
	UserId  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `validate:"required" json:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `validate:"required" json:"message" binding:"required"`
}

type ResponCommnet struct {
	UserID  uint                  `json:"userID"`
	User    repository.UserRespon `json:"user"`
	PhotoID uint                  `json:"photoID"`
	Photo   repository.PhotoRespo `json:"photo"`
	Message string                `json:"message"`
}

type CommentService struct {
	CommentRepo RepositoryComment
	Validate    *validator.Validate
}

func (cs *CommentService) UpdateCommentById(commentId uint, request UpdateCommentRequest, userId uint) error {
	err := cs.Validate.Struct(request)
	if err != nil {
		return err
	}
	commentData, err := cs.CommentRepo.FindByIdComment(commentId)
	if err != nil {
		return err
	}
	if commentData.UserID != userId {
		return errors.New("forbidden")
	}
	commentData.Message = request.Message
	cs.CommentRepo.UpdateComment(commentId, commentData)
	return nil
}

func (cs *CommentService) DeleteCommentById(commentId uint, userId uint) error {
	result, err := cs.CommentRepo.FindByIdComment(commentId)
	if err != nil {
		return err
	}
	if result.UserID != userId {
		return errors.New("Forbiden")
	}
	cs.CommentRepo.DeleteComment(commentId)
	return nil
}

func (cs *CommentService) FindByIdComment(commentId uint) (ResponCommnet, error) {
	var finalResult ResponCommnet
	commentByIdResult, err := cs.CommentRepo.FindByIdComment(commentId)
	if err != nil {
		return finalResult, err
	}
	finalResult.UserID = commentByIdResult.UserID
	finalResult.User = commentByIdResult.User
	finalResult.PhotoID = commentByIdResult.PhotoID
	finalResult.Photo = commentByIdResult.Photo
	finalResult.Message = commentByIdResult.Message

	return finalResult, nil
}

func (cs *CommentService) FindAllComment() []ResponCommnet {
	var finalResult []ResponCommnet
	result := cs.CommentRepo.FindAllComment()
	for _, value := range result {
		finalResult = append(
			finalResult,
			ResponCommnet{
				UserID:  value.UserID,
				User:    value.User,
				PhotoID: value.PhotoID,
				Photo:   value.Photo,
				Message: value.Message,
			},
		)
	}
	return finalResult
}

func (cs *CommentService) CreateComment(request CreateCommentRequest) error {
	if err := cs.Validate.Struct(request); err != nil {
		return err
	}
	comment := repository.Comment{
		UserID:  request.UserId,
		PhotoID: request.PhotoID,
		Message: request.Message,
	}
	cs.CommentRepo.SaveComment(comment)
	return nil
}
