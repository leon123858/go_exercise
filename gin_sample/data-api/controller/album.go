package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leon123858/committee-meeting-assistan/data-api/model"
	"github.com/leon123858/committee-meeting-assistan/data-api/utils"
)

// @Summary		albums example
// @Description	get list of albums
// @Tags			album
// @Success		200	{array}	model.Album
// @Failure		400	{array}	model.Album
// @Router			/albums [get]
// @Security		Bearer
func GetAlbums(c *gin.Context) {
	// token := utils.GetUserToken(c)
	// println(token.UID)
	albums := []model.Album{}
	err := utils.DB.Find(&albums).Error
	if err != nil {
		log.Println("Get albums failed: ", err)
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// @Summary		postAlbums adds an album from JSON received in the request body.
// @Description	post list of albums
// @Param			request	body	model.Album	true	"album 實體結構, id 不用"
// @Tags			album
// @Success		200	{object}	utils.Config
// @Failure		400	{object}	utils.Config
// @Router			/albums [post]
// @Security		Bearer
func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	err := utils.DB.Create(&newAlbum).Error
	if err != nil {
		log.Println("create albums failed: ", err)
		panic(err)
	}
	// fetch new
	albums := []model.Album{}
	err = utils.DB.Find(&albums).Error
	if err != nil {
		log.Println("Get albums failed: ", err)
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// @Summary		getAlbumByID locates the album whose ID value matches the id
// @Description	parameter sent by the client, then returns that album as a response.
// @Param			id	path	string	true	"要抓的目標的 id ,唯一編號"
// @Tags			album
// @Success		200	{object}	model.Album
// @Router			/albums/{id} [get]
// @Security		Bearer
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album model.Album
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	err := utils.DB.Where("id = ?", id).First(&album).Error
	if err != nil {
		log.Println("Get albums failed: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
