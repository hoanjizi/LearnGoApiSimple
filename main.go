package main

import(
	"github.com/gin-gonic/gin"
	"learngoapisimple/routers"
)

func main() {
	app := gin.Default()
	api.Apllyroot(app)
	app.Run(":" + "3000")
}

