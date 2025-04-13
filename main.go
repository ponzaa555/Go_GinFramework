package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Gin_Intro/middleware"
)

func main() {

	router := gin.Default()
	/*
		ถ้าใช้ gin.New() อยากให้ Log จะต้องใช้ router.Use(gin.Logger())
		router := gin.New()
		router.Use(gin.Logger())
	*/
	// router.Use(middleware.Authenticate) // applay all route

	router.POST("/CreateUser", middleware.Authenticate, middleware.Addheder, createUser)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.GET("/GetData", getData)
	router.GET("/getQuueryString", getQueryString)

	// Appaly middleware to group
	/*
		admin := router.Group("/admin", middleware.Authenticate )
		{
			//path : /admin/GetData
			admin.GET("/GetData", getData)
		}
	*/
	/* Group Route
	// set auth before access this page
	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})
	admin := router.Group("/admin", auth)
	{
		//path : /admin/GetData
		admin.GET("/GetData", getData)
	}
	client := router.Group("client")
	{
		//path :/client/getQuueryString
		client.GET("/getQuueryString", getQueryString)
	}
	*/
	server := &http.Server{
		// config router
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hi I am gin-framework",
	})
}

func createUser(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"data":     "HI I am POST method from GIN framework",
		"bodyData": string(value),
	})
}

// query request http://localhost:8080/getQuueryString?name=Mark&age=30
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "HI I am getQueryString",
		"name": name,
		"age":  age,
	})
}

// get URLData http://localhost:8080/getUrlData/Mark/30
// http://localhost:8080/getUrlData/:name/:age
func getUrlData(c *gin.Context) {
	param := c.Params
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data":   "HI I am getUrlData",
		"name":   name,
		"age":    age,
		"params": param,
	})
}
