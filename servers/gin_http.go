package servers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func g_test1(c *gin.Context) {
	c.String(200, "test1")
}

func g_test2(c *gin.Context) {
	sample := make(map[string]interface{})
	sample["key1"] = "value1"
	sample["key2"] = "value2"

	// json_byte, _ := json.Marshal(sample)
	c.JSON(200, sample)
}

func ReadReqBody(body io.ReadCloser) string {
	var body_byte []byte
	body_byte, e := io.ReadAll(body)
	if e == nil {
		return string(body_byte)
	}
	return e.Error()
}

func g_test_req(c *gin.Context) {

	res := fmt.Sprintln("query=", c.Query("name"))
	// b, _ := c.GetRawData()
	// res += fmt.Sprintln("body=", string(b))
	res += fmt.Sprintln("post form=", c.PostForm("name"))
	res += fmt.Sprintln("param=", c.Param("name"))

	c.String(200, fmt.Sprintf("%s:\n%s", "g_test_req", res))
}

type NameStruct struct {
	Name string `form:"name" json:"name" uri:"name" binding:"required"`
}

func g_test_req_bind(c *gin.Context) {
	var data NameStruct
	if c.ShouldBind(&data) != nil {
		if c.ShouldBindUri(&data) != nil {
			c.String(500, "Name is not provided")
			return
		}
	}

	c.String(200, fmt.Sprintf("%s:\nname: %s", "g_test_req_bind", data.Name))
}

func Gin_server() {
	main_router := gin.Default()

	// main_router.StaticFS("/", gin.Dir("./static", true))
	main_router.Use(static.Serve("/", static.LocalFile("./static", true)))

	middleware := func(c *gin.Context) {
		fmt.Println("in middleware", c.Request.RequestURI)
		c.Next()
	}

	main_router.Use(middleware)

	{
		router := main_router.Group("/group/")
		router.GET("/test1", g_test1)
		router.GET("/test2", g_test2)
	}

	main_router.GET("/test1/", g_test1)
	main_router.GET("/test2/", g_test2)
	main_router.GET("/req/:name", g_test_req)
	main_router.Any("/req/", g_test_req)
	main_router.Any("/req2/", g_test_req_bind)
	main_router.Any("/req2/:name", g_test_req_bind)

	http.ListenAndServe(":8899", main_router)

}
