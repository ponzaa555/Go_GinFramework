package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Gin_Intro/log"
	"github.com/ponzaa555/Gin_Intro/middleware"
)

func main() {

	router := gin.Default()

	//เก็บ log ไปที่ file แทน
	// os.Stdout write log file as terminal
	f, _ := os.Create("ginLoggin.log")

	/*
		LogFormat
		{ ::1 - [Sun, 13 Apr 2025 18:04:39 +07] "POST /CreateUser HTTP/1.1 255625 PostmanRuntime/7.43.3 "" %!s(MISSING)"}
			[GIN] 2025/04/13 - 18:04:39 | 200 |     922.959µs |             ::1 | POST     "/CreateUser"
	*/
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	/*
		//ถ้าใช้ gin.New() อยากให้ Log จะต้องใช้ router.Use(gin.Logger())
		router := gin.New()
		//router.Use(gin.Logger())
	*/

	// gin will logging with this format output  log.FormatsLogs
	router.Use(gin.LoggerWithFormatter(log.FormatLogsJson))

	// router.Use(middleware.Authenticate) // applay all route
	/*
		GET    /getUrlData/:name/:age    --> main.getUrlData (3 handlers)
		config Termianl log
		2025/04/13 17:52:47 endpoint formatted imformation is 5 /CreateUser POST main.
		[GIN] 2025/04/13 - 17:53:01 | 200 |      87.167µs |             ::1 | POST     "/CreateUser"

		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			log.Printf("endpoint formatted imformation is %v %v %v %v\n", nuHandlers, absolutePath, httpMethod, handlerName)
		}
	*/

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
