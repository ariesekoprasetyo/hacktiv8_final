package posts

import (
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type CommentService struct {
	CommentRepo RepositoryComment
	Validate    *validator.Validate
}

func (cs *CommentService) UpdateComment(request controller.UpdateCommentRequest) {
	err := cs.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	commentData, err := cs.CommentRepo.FindByIdComment(request.ID)
	if err != nil {
		panic(err)
	}
	commentData.Message = request.Message
	cs.CommentRepo.UpdateComment(commentData)
}

func (cs *CommentService) DeleteComment(commentId uint) {
	cs.CommentRepo.DeleteComment(commentId)
}

func (cs *CommentService) FindByIdComment(commentId uint) (repository.Comment, error) {
	commentByIdResult, err := cs.CommentRepo.FindByIdComment(commentId)
	if err != nil {
		return commentByIdResult, err
	}
	return commentByIdResult, nil
}

func (cs *CommentService) FindAllComment() []repository.Comment {
	allComment := cs.CommentRepo.FindAllComment()
	return allComment
}

func (cs *CommentService) CreateComment(request controller.CreateCommentRequest) {
	if err := cs.Validate.Struct(request); err != nil {
		panic(err)
	}
	comment := repository.Comment{
		UserId:  request.UserId,
		PhotoID: request.UserId,
		Message: request.Message,
	}
	cs.CommentRepo.SaveComment(comment)
}
