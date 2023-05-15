package main

import (
	"github.com/gin-gonic/gin"
	"go-logger-test/middlerware/logger"
	"go-logger-test/model"
	"log"
)

func main() {
	// logger init
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		logger.Info("Request received") // 记录请求信息
		c.Next()
	})

	r.GET("/questions", func(c *gin.Context) {
		logger.Info("Get API Start")
		// 値を取得する
		questins := model.GetAll()
		logger.Error("This is an Error")
		c.JSON(200, questins)
		logger.Info("Get API End")
	})
	r.GET("/:tag/:num", func(c *gin.Context) {
		tag := c.Param("tag")
		num := c.Param("num")
		question := model.GetBy(tag, num)
		c.JSON(200, question)
	})

	r.Run()
}
