package userservice

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

const (
	FailedToGetUser = "failed to get logged in user"
)

func (us UserService) GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject, user models.User) (oapiprivate.GetLoggedInUserResponseObject, error) {

	currentUserResponse := oapiprivate.CurrentUserResponseSchemaJSONResponse{
		User: oapiprivate.User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return oapiprivate.GetLoggedInUser200JSONResponse{
		CurrentUserResponseSchemaJSONResponse: currentUserResponse,
	}, nil
}
