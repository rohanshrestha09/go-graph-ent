package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/rohanshrestha09/go-graph-ent/common"
	. "github.com/rohanshrestha09/go-graph-ent/core/domains"
	"github.com/rohanshrestha09/go-graph-ent/core/ports"
	"github.com/rohanshrestha09/go-graph-ent/frameworks/primary/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserRepository ports.UserRepository
}

func NewUserController(r *gin.RouterGroup, ur ports.UserRepository) {
	uc := &UserController{
		UserRepository: ur,
	}

	r.GET(":id", uc.GetUser)
	r.GET("", uc.GetUsers)
	r.POST("", uc.CreateUser)
	r.PATCH(":id", uc.UpdateUser)
}

// Get User godoc
//
//	@Summary	Get a user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"id"
//	@Success	200	{object}	dto.Response[User]
//	@Router		/user/{id}/ [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	user, err := uc.UserRepository.FindUser(ctx, User{ID: id})

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	ctx.JSON(http.StatusOK, dto.NewResponse("User fetched", user))
}

// Get All User godoc
//
//	@Summary	Get all user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		pagination	query		common.Query	false	"Query"
//	@Success	200			{object}	dto.PaginatedResponse[User]
//	@Router		/user/ [get]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	var query common.Query

	err := ctx.BindQuery(&query)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	users, count, err := uc.UserRepository.FindUsers(ctx, User{}, query)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	ctx.JSON(http.StatusOK, dto.NewPaginatedResponse(
		dto.PaginatedResponse[*User]{
			Message: "Users Fetched",
			Result:  users,
			Count:   count,
		},
		query.Page,
		query.Size,
	))
}

// Create User godoc
//
//	@Summary	Create user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.CreateUserDto	true	"Request Body"
//	@Success	201		{object}	dto.Response[User]
//	@Router		/user/ [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var createUserDto dto.CreateUserDto

	if err := ctx.BindJSON(&createUserDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,
			dto.NewErrorResponse(
				err.Error(), http.StatusUnprocessableEntity,
			))
		return
	}

	if err := validator.New().Struct(createUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadGateway,
			))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,
			dto.NewErrorResponse(
				err.Error(), http.StatusUnprocessableEntity,
			))
		return
	}

	createUserDto.Password = string(hash)

	user, err := uc.UserRepository.CreateUser(ctx, &User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Active:   createUserDto.Active,
		Image:    createUserDto.Image,
		Password: createUserDto.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			dto.NewErrorResponse(
				err.Error(), http.StatusInternalServerError,
			))
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewResponse("User updated", user))
}

// Update User godoc
//
//	@Summary	Update user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"id"
//	@Param		body	body		dto.UpdateUserDto	true	"Request Body"
//	@Success	201		{object}	dto.Response[User]
//	@Router		/user/{id}/ [patch]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	var updateUserDto dto.UpdateUserDto

	if err := ctx.BindJSON(&updateUserDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,
			dto.NewErrorResponse(
				err.Error(), http.StatusUnprocessableEntity,
			))
		return
	}

	if err := validator.New().Struct(updateUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest,
			dto.NewErrorResponse(
				err.Error(), http.StatusBadRequest,
			))
		return
	}

	if updateUserDto.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(updateUserDto.Password), bcrypt.DefaultCost)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity,
				dto.NewErrorResponse(
					err.Error(), http.StatusUnprocessableEntity,
				))
			return
		}

		updateUserDto.Password = string(hash)
	}

	user, err := uc.UserRepository.UpdateUser(
		ctx,
		User{
			ID: id,
		},
		&User{
			Name:     updateUserDto.Name,
			Active:   updateUserDto.Active,
			Image:    updateUserDto.Image,
			Password: updateUserDto.Password,
		})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			dto.NewErrorResponse(
				err.Error(), http.StatusInternalServerError,
			))
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewResponse("User updated", user))
}
