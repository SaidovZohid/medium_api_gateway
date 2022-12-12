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

// @Router /auth/verify [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Verify(ctx *gin.Context) {
	var (
		req models.VerifyRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := h.grpcClient.AuthService().Verify(context.Background(), &user_service.VerifyRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: user.AccessToken,
	})

}

// @Router /auth/login [post]
// @Summary Login User
// @Description Login User
// @Tags register
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Login(ctx *gin.Context) {
	var (
		req models.LoginRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := h.grpcClient.AuthService().Login(context.Background(), &user_service.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: user.AccessToken,
	})

}

// @Router /auth/forgot-password [post]
// @Summary Forgot  password
// @Description Forgot  password
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.ForgotPasswordRequest true "Data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ForgotPassword(ctx *gin.Context) {
	var (
		req models.ForgotPasswordRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err := h.grpcClient.AuthService().ForgotPassword(context.Background(), &user_service.ForgotPasswordRequest{
		Email: req.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Success: "Validation code has been sent",
	})
}

// @Router /auth/verify-forgot-password [post]
// @Summary Verify forgot password
// @Description Verify forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) VerifyForgotPassword(ctx *gin.Context) {
	var (
		req *models.VerifyRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.grpcClient.AuthService().VerifyForgotPassword(context.Background(), &user_service.VerifyRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.AuthResponse{
		Id:          result.Id,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: result.AccessToken,
	})
}
