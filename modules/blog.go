package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/go-graph-ent/ent"
	"github.com/rohanshrestha09/go-graph-ent/frameworks/primary/controllers"
	"github.com/rohanshrestha09/go-graph-ent/frameworks/secondary/repositories"
)

func InitBlogModule(r *gin.RouterGroup, bc *ent.BlogClient) {
	br := repositories.NewBlogRepository(bc)

	controllers.NewBlogController(r, br)
}
