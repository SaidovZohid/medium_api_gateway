package v1

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/SaidovZohid/medium_api_gateway/genproto/user_service"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	Id        string `json:"id"`
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`
	IssuedAt  string `json:"issued_at"`
	ExpiredAt string `json:"expired_at"`
}

func (h *handlerV1) AuthMiddleWare(ctx *gin.Context) {
	accessToken := ctx.GetHeader(os.Getenv("AUTHORIZATION_HEADER_KEY"))

	if len(accessToken) == 0 {
		err := errors.New("authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
		return
	}
	payload, err := h.grpcClient.AuthService().VerifyToken(context.Background(), &user_service.VerifyTokenRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	ctx.Set(os.Getenv("AUTHORIZATION_PAYLOAD_KEY"), Payload{
		Id:        payload.Id,
		UserID:    payload.UserId,
		Email:     payload.Email,
		UserType:  payload.UserType,
		IssuedAt:  payload.IssuedAt,
		ExpiredAt: payload.ExpiredAt,
	})
	ctx.Next()
}

func (h *handlerV1) GetAuthPayload(ctx *gin.Context) (*Payload, error) {
	i, exist := ctx.Get(os.Getenv("AUTHORIZATION_PAYLOAD_KEY"))
	if !exist {
		return nil, errors.New("not found payload")
	}
	payload, ok := i.(Payload) // error is here
	if !ok {
		return nil, errors.New("unknown user")
	}
	return &payload, nil
}
