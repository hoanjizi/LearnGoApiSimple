package loginctl

import (
	"learngoapisimple/common"
	"learngoapisimple/databases"
	models "learngoapisimple/models/login"
	"log"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	name := c.DefaultQuery("name", "hh")
	c.JSON(200, gin.H{"mess": name})
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	err, session, _ := databases.ConnectToDatabase()
	defer session.Close()
	if err != nil {
		c.AbortWithStatus(500)
	}
	//create table

	if err := session.Query(`INSERT INTO users (id, username, password, email, birthdate, phonenumber, tokenuser) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		common.RandStringBytes(20), name, pass, "", "", "", "").Exec(); err != nil {
		log.Fatalln(err)
		c.JSON(400, gin.H{"message": err.Error})
	} else {
		c.JSON(200, gin.H{"message": "user created"})
	}
}

func GetListUser(c *gin.Context) {
	err, session, _ := databases.ConnectToDatabase()
	defer session.Close()
	if err != nil {
		c.AbortWithStatus(500)
	}
	var userList []models.User
	m := map[string]interface{}{}
	iter := session.Query("SELECT username, password, email, birthdate, phonenumber, tokenuser FROM users").Iter()
	for iter.Scan(m) {
		userList = append(userList, models.User{
			Username:    m["username"].(string),
			Password:    m["password"].(string),
			Birthdate:   m["birthdate"].(string),
			Email:       m["email"].(string),
			Phonenumber: m["phonenumber"].(string),
			Tokenuser:   m["tokenuser"].(string),
		})
		m = map[string]interface{}{}
	}
	c.JSON(200,gin.H{"users":userList})
}
