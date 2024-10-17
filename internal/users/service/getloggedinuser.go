package users

import (
	"common/models"
	"common/oapiprivate"
	"context"
)

const (
	FailedToGetUser = "failed to get logged in user"
)

func (us UserService) GetLoggedInUser(ctx context.Context, input oapiprivate.GetLoggedInUserRequestObject, user models.User) (oapiprivate.GetLoggedInUserResponseObject, error) {

	var updatedAt = ""

	if user.UpdatedAt != nil {
		updatedAt = user.UpdatedAt.String()
	}

	currentUserResponse := oapiprivate.CurrentUserResponseSchemaJSONResponse{
		User: oapiprivate.User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: updatedAt,
		},
	}

	return oapiprivate.GetLoggedInUser200JSONResponse{
		CurrentUserResponseSchemaJSONResponse: currentUserResponse,
	}, nil
}
