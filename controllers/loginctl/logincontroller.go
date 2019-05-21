package loginctl

import (
	"context"
	"encoding/json"
	"learngoapisimple/databases"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetPing(c *gin.Context) {
	name := c.DefaultQuery("name", "hh")
	c.JSON(200, gin.H{"mess": name})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "user created"})
}

func LoginUser(c *gin.Context) {
	userName := c.PostForm("username")
	passWord := c.PostForm("password")
	if userName != "" && passWord != "" {
		err, client := databases.ConnectToDatabase()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer client.Disconnect(ctx)
		collection := client.Database("huskydb").Collection("huskies")
		if err != nil {
			c.AbortWithStatus(500)
		}
		var result bson.M
		filter := bson.M{"idImg": passWord}
		errCol := collection.FindOne(ctx, filter).Decode(&result)
		if errCol != nil {
			c.JSON(400, "bad request 1")
		} else {
			if result != nil {
				c.JSON(200, "ok")
			} else {
				c.JSON(400, "bad request 2")
			}
		}

	} else {
		c.JSON(400, "bad request 3")
	}

}

func GetListUser(c *gin.Context) {
	err, client := databases.ConnectToDatabase()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database("huskydb").Collection("huskies")

	if err != nil {
		c.AbortWithStatus(500)
	}

	var jsonDocuments []map[string]interface{}
	var jsonDocument map[string]interface{}
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		var result bson.M
		var temporaryBytes []byte
		errcur := cur.Decode(&result)
		if errcur != nil {
			log.Fatal(errcur)
		}
		temporaryBytes, _ = bson.MarshalJSON(result)
		json.Unmarshal(temporaryBytes, &jsonDocument)
		jsonDocuments = append(jsonDocuments, jsonDocument)
	}
	c.JSON(200, gin.H{"users": jsonDocuments})
}
