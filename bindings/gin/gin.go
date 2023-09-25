package bindings

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oalexander-dev/golang-architecture/domain"
)

func NewGinBinding(ops domain.Ops) *gin.Engine {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode("release")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	authMiddleware, err := newGinJwtMiddleware(jwtConfig{
		IdentityKey: "username",
		SecretKey:   "12342213412asdfajs89epurhna98chcb",
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 12,
		Realm:       "test",
	}, ops)
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

				c.JSON(http.StatusCreated, savedUser)
			})
		}

		apiV1Group.Use(authMiddleware.MiddlewareFunc())

		userGroup := apiV1Group.Group("/users")
		{
			userGroup.GET("/me", func(c *gin.Context) {
				claims, exists := c.Get("JWT_PAYLOAD")
				if exists {
					c.JSON(http.StatusOK, gin.H{
						"claims": claims,
					})
					return
				}

				c.JSON(http.StatusUnauthorized, gin.H{
					"claims": claims,
				})
			})

			userGroup.GET("/:id", func(c *gin.Context) {
				id, err := strconv.ParseInt(c.Param("id"), 10, 64)
				if err != nil {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				user, err := ops.User.GetByID(id)
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
