package login

import(
	"github.com/gin-gonic/gin"
	"learngoapisimple/controllers/loginctl"
)

func Apllygrouplogin(g *gin.RouterGroup){
	v1 := g.Group("/login") 
	{
		//api/login/ping
		v1.GET("/pings",loginctl.GetPing)
		//api/login/create
		v1.POST("/create",loginctl.CreateUser)
		//api/login/list
		v1.GET("/list",loginctl.GetListUser)
	}
}