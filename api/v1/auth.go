package v1

import (
	"context"
	"net/http"

	"github.com/SaidovZohid/medium_api_gateway/api/models"
	"github.com/SaidovZohid/medium_api_gateway/genproto/user_service"
	"github.com/gin-gonic/gin"
)

// @Router /auth/register [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		req models.RegisterRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, _ := h.grpcClient.UserService().GetByEmail(context.Background(), &user_service.GetByEmailRequest{
		Email: req.Email,
	})
	if user != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(ErrEmailExists))
		return
	}

	_, err = h.grpcClient.AuthService().Register(context.Background(), &user_service.RegisterRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "success",
	})
}
