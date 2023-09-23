package bindings

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oalexander-dev/golang-architecture/domain"
)

func NewGinBinding(ops domain.Ops) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
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

	r.POST("/users", func(c *gin.Context) {
		var user domain.User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err = ops.User.Create(user)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	return r
}