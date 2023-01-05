package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Avtoelon/api/auth"
	"github.com/Avtoelon/config"
	"github.com/Avtoelon/pkg/logger"
	"github.com/Hatsker01/Docker_implemintation/api/api/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type handlerV1 struct {
	db         *sqlx.DB
	log        logger.Logger
	cfg        config.Config
	JwtHandler auth.JwtHandler
}

type HandlerV1Config struct {
	Db         *sqlx.DB
	Logger     logger.Logger
	Cfg        config.Config
	jwtHandler auth.JwtHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		db:         c.Db,
		log:        c.Logger,
		cfg:        c.Cfg,
		JwtHandler: c.jwtHandler,
	}
}

func CheckClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   model.JwtRequestModel
		claims          jwt.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("Unauthorized request:", logger.Error(err))
		return nil

	}
	h.JwtHandler.Token = authorization.Token
	claims, err = h.JwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("token is invalid:", logger.Error(err))
		return nil
	}
	return claims
}
func GetClaims(h handlerV1, c *gin.Context) (*CustomClaims, error) {

	var claims CustomClaims

	strToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(h.JwtHandler.SigninKey), nil
	})

	if err != nil {
		h.log.Error("invalid access token")
		return nil, err
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Role = rawClaims["role"].(string)
	claims.Exp = rawClaims["exp"].(float64)
	fmt.Printf("%T type of value in map %v\n", rawClaims["exp"], rawClaims["exp"])
	fmt.Printf("%T type of value in map %v\n", rawClaims["iat"], rawClaims["iat"])

	claims.Iat = rawClaims["iat"].(float64)

	var aud = make([]string, len(rawClaims["aud"].([]interface{})))

	for i, v := range rawClaims["aud"].([]interface{}) {
		aud[i] = v.(string)
	}

	claims.Aud = aud
	claims.Iss = rawClaims["iss"].(string)

	return &claims, nil

}

type CustomClaims struct {
	*jwt.Token
	Sub  string   `json:"sub"`
	Iss  string   `json:"iss"`
	Exp  float64  `json:"exp"`
	Iat  float64  `json:"iat"`
	Aud  []string `json:"aud"`
	Role string   `json:"role"`
	// Token string `json:"token"`
}
