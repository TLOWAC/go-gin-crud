package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// 신기한점 : 별도의 body 인자를 전달받지 않고 ctx.BindJson 으로 body 데이터를 json 파싱한다.
func postAlbums(c *gin.Context) {
	var newAlbum album // 전달 받을 body

	// context 로 넘어온 body 데이터와 newAlbum 을 json 으로 파싱
	if err := c.BindJSON(&newAlbum); err != nil {
			return
	}

	// 메모리 디비로 사용중인 albums 에 post 요청으로 들어온 newAlbum 데이터 추가 (메모리 DB 업데이트)
	albums = append(albums, newAlbum)

	// response 반환. statusCode 와 newAlbum 데이터 반환
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // ctx 에서 param("id") 에 접근한다. post 에서 body 를 가져올때랑은 다른 방식.

	for _, a := range albums { // idx, albums[] 를 순회한다. 여기서 idx 는 불필요하여 _ 로 처리했다. 
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// albums 엔드포인트
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}