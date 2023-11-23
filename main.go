package main

import (
	"final_satu/internal/controller"
	"final_satu/internal/repository"
	"final_satu/internal/service"
	"fmt"

	_ "final_satu/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https
// @host final1todos-production.up.railway.app
// @BasePath /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func route(r *gin.Engine, ctr controller.Controller) {
	r.GET("/todos", ctr.GetAll)
	r.GET("/todos/:id", ctr.GetOne)
	r.POST("/todos/", ctr.Create)
	r.DELETE("/todos/:id", ctr.Delete)
	r.PUT("/todos/:id", ctr.Update)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func startRouter(r *gin.Engine, ctr controller.Controller) {
	route(r, ctr)
	r.Run(":8080")
}
func main() {
	// var err error
	fmt.Println("Hello, world!")
	db := repository.New()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	s := service.New(db)
	ctr := controller.New(s)

	startRouter(r, ctr)
}
