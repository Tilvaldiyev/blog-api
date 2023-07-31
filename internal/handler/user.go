package handler

import (
	"github.com/Tilvaldiyev/blog-api/api"
	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// createUser registration new user
//
//	@Summary      Create user
//	@Description  Create new user
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param req body api.RegisterRequest true "req body"
//
//	@Success      201
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /user/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	u := &entity.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	err = h.srvs.CreateUser(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}

func (h *Handler) userPosts(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	// logic

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    userID,
	})
}
