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
	"github.com/rohanshrestha09/go-graph-ent/utils"
)

type BlogController struct {
	BlogRepository ports.BlogRepository
}

func NewBlogController(r *gin.RouterGroup, br ports.BlogRepository) {
	bc := &BlogController{
		BlogRepository: br,
	}

	r.GET(":slug", bc.GetBlog)
	r.GET("", bc.GetBlogs)
	r.POST("", bc.CreateBlog)
	r.PATCH(":slug", bc.UpdateBlog)
}

// Get Blog godoc
//
//	@Summary	Get a blog
//	@Tags		Blog
//	@Accept		json
//	@Produce	json
//	@Param		slug	path	string	true	"slug"
//	@Router		/blog/{slug}/ [get]
func (uc *BlogController) GetBlog(ctx *gin.Context) {
	slug := ctx.Param("slug")

	blog, err := uc.BlogRepository.FindBlog(ctx, Blog{Slug: slug})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{"data": blog})
}

// Get All Blog godoc
//
//	@Summary	Get all blog
//	@Tags		Blog
//	@Accept		json
//	@Produce	json
//	@Param		query	query	dto.QueryBlogDto	false	"Query"
//	@Router		/blog/ [get]
func (bc *BlogController) GetBlogs(ctx *gin.Context) {
	var query dto.QueryBlogDto

	err := ctx.BindQuery(&query)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	var userId uuid.UUID

	if query.UserID != "" {
		userId, err = uuid.Parse(query.UserID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			return
		}
	}

	blogs, count, err := bc.BlogRepository.FindBlogs(
		ctx,
		Blog{
			UserID: userId,
		},
		common.Query{
			Page:   query.Page,
			Size:   query.Size,
			Sort:   query.Sort,
			Order:  query.Order,
			Search: query.Search,
		},
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": blogs, "count": count})
}

// Create Blog godoc
//
//	@Summary	Create blog
//	@Tags		Blog
//	@Accept		json
//	@Produce	json
//	@Param		body	body	dto.CreateBlogDto	true	"Request Body"
//	@Router		/blog/ [post]
func (bc *BlogController) CreateBlog(ctx *gin.Context) {
	var createBlogDto dto.CreateBlogDto

	if err := ctx.BindJSON(&createBlogDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := validator.New().Struct(createBlogDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := uuid.Parse(createBlogDto.UserID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	blog, err := bc.BlogRepository.CreateBlog(ctx, &Blog{
		Slug:    utils.ToKebabCase(createBlogDto.Title),
		Title:   createBlogDto.Title,
		Content: createBlogDto.Content,
		Image:   createBlogDto.Image,
		Status:  createBlogDto.Status,
		UserID:  userId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": blog})
}

// Update Blog godoc
//
//	@Summary	Update blog
//	@Tags		Blog
//	@Accept		json
//	@Produce	json
//	@Param		slug	path	string				true	"slug"
//	@Param		body	body	dto.UpdateBlogDto	true	"Request Body"
//	@Router		/blog/{slug}/ [patch]
func (bc *BlogController) UpdateBlog(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var updateBlogDto dto.UpdateBlogDto

	if err := ctx.BindJSON(&updateBlogDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := validator.New().Struct(updateBlogDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	blog, err := bc.BlogRepository.UpdateBlog(
		ctx,
		Blog{
			Slug: slug,
		},
		&Blog{
			Title:   updateBlogDto.Title,
			Content: updateBlogDto.Content,
			Image:   updateBlogDto.Image,
			Status:  updateBlogDto.Status,
		})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": blog})
}
