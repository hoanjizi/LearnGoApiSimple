package loginctl

import (
	"fmt"
	"learngoapisimple/databases"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	name := c.DefaultQuery("name", "hh")
	c.JSON(200, gin.H{"mess": name})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "user created"})
}

func GetListUser(c *gin.Context) {
	err, client, ctx := databases.ConnectToDatabase()
	collection := client.Database("huskydb").Collection("Huskys")

	if err != nil {
		c.AbortWithStatus(500)
	}
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.AbortWithStatus(500)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	c.JSON(200, gin.H{"users": ""})
}
