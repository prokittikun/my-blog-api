package routes

import (
	"my-blog-api/controllers"
	"my-blog-api/middleware"
	"my-blog-api/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")

	var loginService services.LoginService = services.StaticLoginService()
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)
	v1.POST("login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	grp1 := v1.Group("/user-api")
	// {
	// 	grp1.GET("user", controllers.GetUsers)

	// 	grp1.POST("user", controllers.CreateUser)
	// 	grp1.GET("user/:id", controllers.GetUserByID)
	// 	grp1.PUT("user/:id", controllers.UpdateUser)
	// 	grp1.DELETE("user/:id", controllers.DeleteUser)
	// }
	grp1.Use(middleware.AuthorizeJWT())
	{
		grp1.GET("user", controllers.GetUsers)

		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserByID)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	grp2 := v1.Group("/blog-api")
	{
		grp2.GET("blog", controllers.GetBlogs)
		grp2.POST("blog", controllers.CreateBlog)
		grp2.GET("blog/:id", controllers.GetBlogByID)
		grp2.PUT("blog/:id", controllers.UpdateBlog)
		grp2.DELETE("blog/:id", controllers.DeleteBlog)
	}

	return r
}
