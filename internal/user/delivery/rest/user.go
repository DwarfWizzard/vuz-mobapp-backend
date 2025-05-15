package rest

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	userdto "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserInfoRequest struct {
	UserId uint32 `query:"id"`
}

type UserInfoResponse struct {
	Id     uint32                  `json:"id"`
	Role   string                  `json:"role"`
	Groups []userdto.UserGroupInfo `json:"groups"`
}

// UserInfo - GET /v1/user/:id
func (h *UserHandler) UserInfo(c echo.Context) error {
	ctx := c.Request().Context()

	var rq UserInfoRequest
	if err := c.Bind(&rq); err != nil {
		return response.ErrFormatError
	}

	userInfo, err := h.uc.GetUserInfo(ctx, rq.UserId)
	if err != nil {
		h.logger.Error("Get user info error", zap.Error(err))
		if repository.ErrorIsNoRows(err) {
			return response.ErrUserNotFound
		}

		return response.ErrInternalServerError
	}

	data := &UserInfoResponse{
		Id:     userInfo.User.ID,
		Role:   userInfo.User.Role.Name,
		Groups: userInfo.Groups,
	}

	return response.Success(c, data)
}
