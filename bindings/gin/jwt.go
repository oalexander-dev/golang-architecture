package bindings

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/oalexander-dev/golang-architecture/domain"
)

type jwtConfig struct {
	IdentityKey string
	SecretKey   string
	Timeout     time.Duration
	MaxRefresh  time.Duration
	Realm       string
}

func newGinJwtMiddleware(config jwtConfig, ops domain.Ops) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(config.SecretKey),
		Timeout:     config.Timeout,
		MaxRefresh:  config.MaxRefresh,
		IdentityKey: config.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					config.IdentityKey: user.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				Username: claims[config.IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals domain.UserLoginInput
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := ops.User.CheckPassword(loginVals.Username, loginVals.Password)
			if err == nil {
				return &domain.User{
					ID:       user.ID,
					Username: user.Username,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.Status(code)
		},
		TokenLookup:       "cookie: jwt",
		TokenHeadName:     "Bearer",
		TimeFunc:          time.Now,
		SendCookie:        true,
		SendAuthorization: false,
		CookieHTTPOnly:    true,
		SecureCookie:      true,
	})

	if err != nil {
		return &jwt.GinJWTMiddleware{}, nil
	}

	return authMiddleware, nil
}
