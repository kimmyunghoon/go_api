package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunGinExample() {

	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()
	// Blank Gin without middleware by default
	// r := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	//Multipart/Urlencoded Form
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	/**
	Another example: query + post form

	POST /post?id=1234&page=1 HTTP/1.1
	Content-Type: application/x-www-form-urlencoded

	name=manu&message=this_is_great

	return
	id: 1234; page: 1; name: manu; message: this_is_great
	*/
	router.POST("/post1", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	/**
	Map as querystring or postform parameters

	POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	Content-Type: application/x-www-form-urlencoded

	names[first]=thinkerou&names[second]=tianou

	return
	ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]

	*/
	router.POST("/post2", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	//Grouping routes
	var fn = func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	}

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", fn)
		v1.POST("/submit", fn)
		v1.POST("/read", fn)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", fn)
		v2.POST("/submit", fn)
		v2.POST("/read", fn)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

//
//func RunUsingMiddleware() {
//	// Creates a router without any middleware by default
//	r := gin.New()
//
//	// Global middleware
//	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
//	// By default gin.DefaultWriter = os.Stdout
//	r.Use(gin.Logger())
//
//	// Recovery middleware recovers from any panics and writes a 500 if there was one.
//	r.Use(gin.Recovery())
//
//	// Per route middleware, you can add as many as you desire.
//	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
//	var fn = func(c *gin.Context) {
//		c.String(http.StatusOK, "test")
//	}
//
//	// Authorization group
//	// authorized := r.Group("/", AuthRequired())
//	// exactly the same as:
//	authorized := r.Group("/")
//	// per group middleware! in this case we use the custom created
//	// AuthRequired() middleware just in the "authorized" group.
//	authorized.Use(AuthRequired())
//	{
//		authorized.POST("/login", fn)
//		authorized.POST("/submit", fn)
//		authorized.POST("/read", fn)
//
//		// nested group
//		testing := authorized.Group("testing")
//		// visit 0.0.0.0:8080/testing/analytics
//		testing.GET("/analytics", analyticsEndpoint)
//	}
//
//	// Listen and serve on 0.0.0.0:8080
//	r.Run(":8080")
//}
