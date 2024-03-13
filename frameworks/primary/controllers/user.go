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
}

// Get User godoc
//
//	@Summary	Get a user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"id"
//	@Router		/user/{id} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	user, err := uc.UserRepository.FindUser(ctx, User{ID: id})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{"data": user})
}

// Get All User godoc
//
//	@Summary	Get all user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		pagination	query	common.Pagination	false	"Query"
//	@Router		/user/ [get]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	var pagination common.Pagination

	err := ctx.BindQuery(&pagination)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	users, count, err := uc.UserRepository.FindUsers(ctx, User{}, pagination)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users, "count": count})
}

// Create User godoc
//
//	@Summary	Create user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		body	body	dto.CreateUserDto	true	"Requesy Body"
//	@Router		/user/ [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var createUserDto dto.CreateUserDto

	if err := ctx.BindJSON(&createUserDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := validator.New().Struct(createUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := uc.UserRepository.CreateUser(ctx, &User{
		Name:   createUserDto.Name,
		Age:    createUserDto.Age,
		Active: createUserDto.Active,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
