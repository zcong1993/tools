package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/tools/md5"
	"github.com/zcong1993/tools/resolver"
	"github.com/zcong1993/tools/ulid"
	"github.com/zcong1993/tools/utils"
)

var dnsHandler = func(c *gin.Context) {
	host := c.Query("d")
	if host == "" {
		c.JSON(200, gin.H{
			"error": "url should end with '?d=example.com'",
		})
		return
	}
	res, err := resolver.Resolve(host)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(200, res)
}

var ulidHandler = func(c *gin.Context) {
	n := c.Query("n")
	nums, err := utils.Str2int64(n)
	if err != nil {
		nums = 1
	}
	if nums > 100 {
		nums = 100
	}
	c.JSON(200, ulid.GenUlids(nums))
}

var md5Handler = func(c *gin.Context) {
	s := c.Query("s")
	if s == "" {
		c.JSON(200, gin.H{
			"error": "url should end with '?s=string4md5'",
		})
		return
	}
	hash := md5.GetMd5([]byte(s))
	c.JSON(200, gin.H{
		"string": s,
		"md5":    hash,
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/dns", dnsHandler)
	r.GET("/ulid", ulidHandler)
	r.GET("/md5", md5Handler)
	r.Run(fmt.Sprintf(":%s", utils.GetEnvOrDefault("PORT", "8888")))
}
