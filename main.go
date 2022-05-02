package main

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"log"
	"we-tools/internal/apps/memes"
	"we-tools/internal/common/persistence"
	"we-tools/internal/common/storage"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mysqlDB, err := persistence.NewDB("mysql", "root:q123q123@tcp(127.0.0.1:5080)/we_tools?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect to mysql failed, err:%v", err)
		return
	}
	memeRepo := memes.NewRepo(mysqlDB)
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("new snowflake node failed, err:%v", err)
		return
	}
	localStorage, err := storage.NewLocalStorage("storage", "http://localhost:9999/file")
	if err != nil {
		log.Fatalf("new local storage failed, err:%v", err)
		return
	}
	memeUsecase := memes.NewUsecase(memeRepo, node, localStorage)
	memesApi := memes.NewApi(memeUsecase)
	r.GET("/file/*key", localStorage.Handle)
	r.POST("/memes", memesApi.UploadMeme)
	r.GET("/memes/tags", memesApi.GetTags)
	r.GET("/memes", memesApi.ListMemes)

	r.Run("localhost:9999")
}
