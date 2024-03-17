package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	docs "github.com/rohanshrestha09/go-graph-ent/docs"
	"github.com/rohanshrestha09/go-graph-ent/ent"
	"github.com/rohanshrestha09/go-graph-ent/infrastructure/configs"
	"github.com/rohanshrestha09/go-graph-ent/modules"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {
	databaseConfig := configs.GetDatabaseConfig()

	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name)

	client, err := ent.Open(os.Getenv("DATABASE_DIALECT"), databaseUrl)

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	defer client.Close()

	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/api/v1")

	modules.InitUserModule(v1.Group("user"), client.User)

	modules.InitBlogModule(v1.Group("blog"), client.Blog)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + os.Getenv("APP_PORT"))
}
