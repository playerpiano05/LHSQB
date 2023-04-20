package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func loginRoute(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//
	//
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

}

//alright goodnight!!

/*
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
*/
