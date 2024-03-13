package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/go-graph-ent/ent"
	"github.com/rohanshrestha09/go-graph-ent/frameworks/primary/controllers"
	"github.com/rohanshrestha09/go-graph-ent/frameworks/secondary/repositories"
)

func InitUserModule(r *gin.RouterGroup, uc *ent.UserClient) {
	ur := repositories.NewUserRepository(uc)

	controllers.NewUserController(r, ur)
}
