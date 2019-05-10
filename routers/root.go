package api

import(
	"github.com/gin-gonic/gin"
	"learngoapisimple/routers/login"
)

func Apllyroot(g *gin.Engine) {
	api := g.Group("/api")
	{
		login.Apllygrouplogin(api)
	}
}