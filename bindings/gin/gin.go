package bindings

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/oalexander-dev/golang-architecture/domain"
)

type GinBindingConfig struct {
	GinMode   string
	SecretKey string
}

func NewGinBinding(ops domain.Ops, config GinBindingConfig) *gin.Engine {
	gin.SetMode(config.GinMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	ginJwtConfig := jwtConfig{
		IdentityKey: "username",
		SecretKey:   config.SecretKey,
		Timeout:     time.Minute * 30,
		MaxRefresh:  time.Hour * 12,
	}

	authMiddleware, err := newGinJwtMiddleware(ginJwtConfig, ops)
	if err != nil {
		panic("failed to initialize JWT middleware")
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

	apiV1Group := r.Group("/api/v1")
	{
		authGroup := apiV1Group.Group("/auth")
		{
			authGroup.POST("/login", authMiddleware.LoginHandler)

			authGroup.POST("/logout", authMiddleware.LogoutHandler)

			authGroup.POST("", func(c *gin.Context) {
				var user domain.UserInput
				err := c.ShouldBindJSON(&user)
				if err != nil {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				savedUser, err := ops.User.Create(user)
				if err != nil {
					fmt.Println(err.Error())
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				c.JSON(http.StatusCreated, gin.H{
					"id":       savedUser.ID,
					"username": savedUser.Username,
					"fullName": user.FullName,
				})
			})
		}

		apiV1Group.Use(authMiddleware.MiddlewareFunc())

		userGroup := apiV1Group.Group("/users")
		{
			userGroup.GET("/me", func(c *gin.Context) {
				claims := jwt.ExtractClaims(c)
				username := claims[ginJwtConfig.IdentityKey].(string)

				if username == "" {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}

				user, err := ops.User.GetByUsername(username)
				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"username": username,
					"id":       user.ID,
					"fullName": user.FullName,
				})
			})

			userGroup.GET("/:username", func(c *gin.Context) {
				username := c.Param("username")
				if err != nil || username == "" {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				user, err := ops.User.GetByUsername(username)
				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				c.JSON(http.StatusOK, user)
			})
		}
	}

	return r
}
