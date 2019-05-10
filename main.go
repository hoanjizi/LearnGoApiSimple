package main

import(
	"github.com/gin-gonic/gin"
	"learngoapisimple/routers"
	"learngoapisimple/databases"
)

func main() {
	databases.CreateAllTable()
	app := gin.Default()
	api.Apllyroot(app)
	app.Run(":" + "3000")
}

